package flags

import (
	"github.com/urfave/cli"

	ptservice "github.com/patex-ecosystem/patex-network/pt-service"
	ptlog "github.com/patex-ecosystem/patex-network/pt-service/log"
	ptmetrics "github.com/patex-ecosystem/patex-network/pt-service/metrics"
	ptpprof "github.com/patex-ecosystem/patex-network/pt-service/pprof"
	ptrpc "github.com/patex-ecosystem/patex-network/pt-service/rpc"
	"github.com/patex-ecosystem/patex-network/pt-service/txmgr"
)

const envVarPrefix = "PT_PROPOSER"

var (
	// Required Flags
	L1EthRpcFlag = cli.StringFlag{
		Name:     "l1-eth-rpc",
		Usage:    "HTTP provider URL for L1",
		Required: true,
		EnvVar:   ptservice.PrefixEnvVar(envVarPrefix, "L1_ETH_RPC"),
	}
	RollupRpcFlag = cli.StringFlag{
		Name:     "rollup-rpc",
		Usage:    "HTTP provider URL for the rollup node",
		Required: true,
		EnvVar:   ptservice.PrefixEnvVar(envVarPrefix, "ROLLUP_RPC"),
	}
	L2OOAddressFlag = cli.StringFlag{
		Name:     "l2oo-address",
		Usage:    "Address of the L2OutputOracle contract",
		Required: true,
		EnvVar:   ptservice.PrefixEnvVar(envVarPrefix, "L2OO_ADDRESS"),
	}
	PollIntervalFlag = cli.DurationFlag{
		Name: "poll-interval",
		Usage: "Delay between querying L2 for more transactions and " +
			"creating a new batch",
		Required: true,
		EnvVar:   ptservice.PrefixEnvVar(envVarPrefix, "POLL_INTERVAL"),
	}
	// Optional flags
	AllowNonFinalizedFlag = cli.BoolFlag{
		Name:   "allow-non-finalized",
		Usage:  "Allow the proposer to submit proposals for L2 blocks derived from non-finalized L1 blocks.",
		EnvVar: ptservice.PrefixEnvVar(envVarPrefix, "ALLOW_NON_FINALIZED"),
	}
	// Legacy Flags
	L2OutputHDPathFlag = txmgr.L2OutputHDPathFlag
)

var requiredFlags = []cli.Flag{
	L1EthRpcFlag,
	RollupRpcFlag,
	L2OOAddressFlag,
	PollIntervalFlag,
}

var optionalFlags = []cli.Flag{
	AllowNonFinalizedFlag,
}

func init() {
	requiredFlags = append(requiredFlags, ptrpc.CLIFlags(envVarPrefix)...)

	optionalFlags = append(optionalFlags, ptlog.CLIFlags(envVarPrefix)...)
	optionalFlags = append(optionalFlags, ptmetrics.CLIFlags(envVarPrefix)...)
	optionalFlags = append(optionalFlags, ptpprof.CLIFlags(envVarPrefix)...)
	optionalFlags = append(optionalFlags, txmgr.CLIFlags(envVarPrefix)...)

	Flags = append(requiredFlags, optionalFlags...)
}

// Flags contains the list of configuration options available to the binary.
var Flags []cli.Flag
