package proposer

import (
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/urfave/cli"

	"github.com/patex-ecosystem/patex-network/pt-node/sources"
	"github.com/patex-ecosystem/patex-network/pt-proposer/flags"

	ptlog "github.com/patex-ecosystem/patex-network/pt-service/log"
	ptmetrics "github.com/patex-ecosystem/patex-network/pt-service/metrics"
	ptpprof "github.com/patex-ecosystem/patex-network/pt-service/pprof"
	ptrpc "github.com/patex-ecosystem/patex-network/pt-service/rpc"
	"github.com/patex-ecosystem/patex-network/pt-service/txmgr"
)

// Config contains the well typed fields that are used to initialize the output submitter.
// It is intended for programmatic use.
type Config struct {
	L2OutputOracleAddr common.Address
	PollInterval       time.Duration
	NetworkTimeout     time.Duration
	TxManager          txmgr.TxManager
	L1Client           *ethclient.Client
	RollupClient       *sources.RollupClient
	AllowNonFinalized  bool
}

// CLIConfig is a well typed config that is parsed from the CLI params.
// This also contains config options for auxiliary services.
// It is transformed into a `Config` before the L2 output submitter is started.
type CLIConfig struct {
	/* Required Params */

	// L1EthRpc is the HTTP provider URL for L1.
	L1EthRpc string

	// RollupRpc is the HTTP provider URL for the rollup node.
	RollupRpc string

	// L2OOAddress is the L2OutputOracle contract address.
	L2OOAddress string

	// PollInterval is the delay between querying L2 for more transaction
	// and creating a new batch.
	PollInterval time.Duration

	// AllowNonFinalized can be set to true to propose outputs
	// for L2 blocks derived from non-finalized L1 data.
	AllowNonFinalized bool

	TxMgrConfig txmgr.CLIConfig

	RPCConfig ptrpc.CLIConfig

	LogConfig ptlog.CLIConfig

	MetricsConfig ptmetrics.CLIConfig

	PprofConfig ptpprof.CLIConfig
}

func (c CLIConfig) Check() error {
	if err := c.RPCConfig.Check(); err != nil {
		return err
	}
	if err := c.LogConfig.Check(); err != nil {
		return err
	}
	if err := c.MetricsConfig.Check(); err != nil {
		return err
	}
	if err := c.PprofConfig.Check(); err != nil {
		return err
	}
	if err := c.TxMgrConfig.Check(); err != nil {
		return err
	}
	return nil
}

// NewConfig parses the Config from the provided flags or environment variables.
func NewConfig(ctx *cli.Context) CLIConfig {
	return CLIConfig{
		// Required Flags
		L1EthRpc:     ctx.GlobalString(flags.L1EthRpcFlag.Name),
		RollupRpc:    ctx.GlobalString(flags.RollupRpcFlag.Name),
		L2OOAddress:  ctx.GlobalString(flags.L2OOAddressFlag.Name),
		PollInterval: ctx.GlobalDuration(flags.PollIntervalFlag.Name),
		TxMgrConfig:  txmgr.ReadCLIConfig(ctx),
		// Optional Flags
		AllowNonFinalized: ctx.GlobalBool(flags.AllowNonFinalizedFlag.Name),
		RPCConfig:         ptrpc.ReadCLIConfig(ctx),
		LogConfig:         ptlog.ReadCLIConfig(ctx),
		MetricsConfig:     ptmetrics.ReadCLIConfig(ctx),
		PprofConfig:       ptpprof.ReadCLIConfig(ctx),
	}
}
