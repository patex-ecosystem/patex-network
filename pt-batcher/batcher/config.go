package batcher

import (
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/urfave/cli"

	"github.com/patex-ecosystem/patex-network/pt-batcher/compressor"
	"github.com/patex-ecosystem/patex-network/pt-batcher/flags"
	"github.com/patex-ecosystem/patex-network/pt-batcher/metrics"
	"github.com/patex-ecosystem/patex-network/pt-batcher/rpc"
	"github.com/patex-ecosystem/patex-network/pt-node/rollup"
	"github.com/patex-ecosystem/patex-network/pt-node/sources"
	ptlog "github.com/patex-ecosystem/patex-network/pt-service/log"
	ptmetrics "github.com/patex-ecosystem/patex-network/pt-service/metrics"
	ptpprof "github.com/patex-ecosystem/patex-network/pt-service/pprof"
	"github.com/patex-ecosystem/patex-network/pt-service/txmgr"
)

type Config struct {
	log        log.Logger
	metr       metrics.Metricer
	L1Client   *ethclient.Client
	L2Client   *ethclient.Client
	RollupNode *sources.RollupClient
	TxManager  txmgr.TxManager

	NetworkTimeout         time.Duration
	PollInterval           time.Duration
	MaxPendingTransactions uint64

	// RollupConfig is queried at startup
	Rollup *rollup.Config

	// Channel builder parameters
	Channel ChannelConfig
}

// Check ensures that the [Config] is valid.
func (c *Config) Check() error {
	if err := c.Rollup.Check(); err != nil {
		return err
	}
	if err := c.Channel.Check(); err != nil {
		return err
	}
	return nil
}

type CLIConfig struct {
	// L1EthRpc is the HTTP provider URL for L1.
	L1EthRpc string

	// L2EthRpc is the HTTP provider URL for the L2 execution engine.
	L2EthRpc string

	// RollupRpc is the HTTP provider URL for the L2 rollup node.
	RollupRpc string

	// MaxChannelDuration is the maximum duration (in #L1-blocks) to keep a
	// channel open. This allows to more eagerly send batcher transactions
	// during times of low L2 transaction volume. Note that the effective
	// L1-block distance between batcher transactions is then MaxChannelDuration
	// + NumConfirmations because the batcher waits for NumConfirmations blocks
	// after sending a batcher tx and only then starts a new channel.
	//
	// If 0, duration checks are disabled.
	MaxChannelDuration uint64

	// The batcher tx submission safety margin (in #L1-blocks) to subtract from
	// a channel's timeout and sequencing window, to guarantee safe inclusion of
	// a channel on L1.
	SubSafetyMargin uint64

	// PollInterval is the delay between querying L2 for more transaction
	// and creating a new batch.
	PollInterval time.Duration

	// MaxPendingTransactions is the maximum number of concurrent pending
	// transactions sent to the transaction manager (0 == no limit).
	MaxPendingTransactions uint64

	// MaxL1TxSize is the maximum size of a batch tx submitted to L1.
	MaxL1TxSize uint64

	Stopped bool

	TxMgrConfig      txmgr.CLIConfig
	RPCConfig        rpc.CLIConfig
	LogConfig        ptlog.CLIConfig
	MetricsConfig    ptmetrics.CLIConfig
	PprofConfig      ptpprof.CLIConfig
	CompressorConfig compressor.CLIConfig
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
		/* Required Flags */
		L1EthRpc:        ctx.GlobalString(flags.L1EthRpcFlag.Name),
		L2EthRpc:        ctx.GlobalString(flags.L2EthRpcFlag.Name),
		RollupRpc:       ctx.GlobalString(flags.RollupRpcFlag.Name),
		SubSafetyMargin: ctx.GlobalUint64(flags.SubSafetyMarginFlag.Name),
		PollInterval:    ctx.GlobalDuration(flags.PollIntervalFlag.Name),

		/* Optional Flags */
		MaxPendingTransactions: ctx.GlobalUint64(flags.MaxPendingTransactionsFlag.Name),
		MaxChannelDuration:     ctx.GlobalUint64(flags.MaxChannelDurationFlag.Name),
		MaxL1TxSize:            ctx.GlobalUint64(flags.MaxL1TxSizeBytesFlag.Name),
		Stopped:                ctx.GlobalBool(flags.StoppedFlag.Name),
		TxMgrConfig:            txmgr.ReadCLIConfig(ctx),
		RPCConfig:              rpc.ReadCLIConfig(ctx),
		LogConfig:              ptlog.ReadCLIConfig(ctx),
		MetricsConfig:          ptmetrics.ReadCLIConfig(ctx),
		PprofConfig:            ptpprof.ReadCLIConfig(ctx),
		CompressorConfig:       compressor.ReadCLIConfig(ctx),
	}
}
