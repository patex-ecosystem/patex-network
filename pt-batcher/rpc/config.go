package rpc

import (
	"github.com/urfave/cli"

	ptservice "github.com/patex-ecosystem/patex-network/pt-service"
	ptrpc "github.com/patex-ecosystem/patex-network/pt-service/rpc"
)

const (
	EnableAdminFlagName = "rpc.enable-admin"
)

func CLIFlags(envPrefix string) []cli.Flag {
	return []cli.Flag{
		cli.BoolFlag{
			Name:   EnableAdminFlagName,
			Usage:  "Enable the admin API (experimental)",
			EnvVar: ptservice.PrefixEnvVar(envPrefix, "RPC_ENABLE_ADMIN"),
		},
	}
}

type CLIConfig struct {
	ptrpc.CLIConfig
	EnableAdmin bool
}

func ReadCLIConfig(ctx *cli.Context) CLIConfig {
	return CLIConfig{
		CLIConfig:   ptrpc.ReadCLIConfig(ctx),
		EnableAdmin: ctx.GlobalBool(EnableAdminFlagName),
	}
}
