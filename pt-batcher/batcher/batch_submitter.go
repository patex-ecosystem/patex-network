package batcher

import (
	"context"
	"fmt"
	_ "net/http/pprof"

	gethrpc "github.com/ethereum/go-ethereum/rpc"
	"github.com/urfave/cli"

	"github.com/patex-ecosystem/patex-network/pt-batcher/flags"
	"github.com/patex-ecosystem/patex-network/pt-batcher/metrics"
	"github.com/patex-ecosystem/patex-network/pt-batcher/rpc"
	ptservice "github.com/patex-ecosystem/patex-network/pt-service"
	ptlog "github.com/patex-ecosystem/patex-network/pt-service/log"
	"github.com/patex-ecosystem/patex-network/pt-service/opio"
	ptpprof "github.com/patex-ecosystem/patex-network/pt-service/pprof"
	ptrpc "github.com/patex-ecosystem/patex-network/pt-service/rpc"
)

// Main is the entrypoint into the Batch Submitter. This method returns a
// closure that executes the service and blocks until the service exits. The use
// of a closure allows the parameters bound to the tpt-level main package, e.g.
// GitVersion, to be captured and used once the function is executed.
func Main(version string, cliCtx *cli.Context) error {
	if err := flags.CheckRequired(cliCtx); err != nil {
		return err
	}
	cfg := NewConfig(cliCtx)
	if err := cfg.Check(); err != nil {
		return fmt.Errorf("invalid CLI flags: %w", err)
	}

	l := ptlog.NewLogger(cfg.LogConfig)
	ptservice.ValidateEnvVars(flags.EnvVarPrefix, flags.Flags, l)
	m := metrics.NewMetrics("default")
	l.Info("Initializing Batch Submitter")

	batchSubmitter, err := NewBatchSubmitterFromCLIConfig(cfg, l, m)
	if err != nil {
		l.Error("Unable to create Batch Submitter", "error", err)
		return err
	}

	if !cfg.Stopped {
		if err := batchSubmitter.Start(); err != nil {
			l.Error("Unable to start Batch Submitter", "error", err)
			return err
		}
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // Stop pprof and metrics only after main loop returns
	defer batchSubmitter.StopIfRunning(context.Background())

	pprofConfig := cfg.PprofConfig
	if pprofConfig.Enabled {
		l.Info("starting pprof", "addr", pprofConfig.ListenAddr, "port", pprofConfig.ListenPort)
		go func() {
			if err := ptpprof.ListenAndServe(ctx, pprofConfig.ListenAddr, pprofConfig.ListenPort); err != nil {
				l.Error("error starting pprof", "err", err)
			}
		}()
	}

	metricsCfg := cfg.MetricsConfig
	if metricsCfg.Enabled {
		l.Info("starting metrics server", "addr", metricsCfg.ListenAddr, "port", metricsCfg.ListenPort)
		go func() {
			if err := m.Serve(ctx, metricsCfg.ListenAddr, metricsCfg.ListenPort); err != nil {
				l.Error("error starting metrics server", err)
			}
		}()
		m.StartBalanceMetrics(ctx, l, batchSubmitter.L1Client, batchSubmitter.TxManager.From())
	}

	rpcCfg := cfg.RPCConfig
	server := ptrpc.NewServer(
		rpcCfg.ListenAddr,
		rpcCfg.ListenPort,
		version,
		ptrpc.WithLogger(l),
	)
	if rpcCfg.EnableAdmin {
		server.AddAPI(gethrpc.API{
			Namespace: "admin",
			Service:   rpc.NewAdminAPI(batchSubmitter),
		})
		l.Info("Admin RPC enabled")
	}
	if err := server.Start(); err != nil {
		cancel()
		return fmt.Errorf("error starting RPC server: %w", err)
	}

	m.RecordInfo(version)
	m.RecordUp()

	opio.BlockOnInterrupts()
	if err := server.Stop(); err != nil {
		l.Error("Error shutting down http server: %w", err)
	}
	return nil
}
