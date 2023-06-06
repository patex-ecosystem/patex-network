package pprof

import (
	"errors"
	"math"

	ptservice "github.com/patex-ecosystem/patex-network/pt-service"
	"github.com/urfave/cli"
)

const (
	EnabledFlagName    = "pprof.enabled"
	ListenAddrFlagName = "pprof.addr"
	PortFlagName       = "pprof.port"
)

func CLIFlags(envPrefix string) []cli.Flag {
	return []cli.Flag{
		cli.BoolFlag{
			Name:   EnabledFlagName,
			Usage:  "Enable the pprof server",
			EnvVar: ptservice.PrefixEnvVar(envPrefix, "PPROF_ENABLED"),
		},
		cli.StringFlag{
			Name:   ListenAddrFlagName,
			Usage:  "pprof listening address",
			Value:  "0.0.0.0",
			EnvVar: ptservice.PrefixEnvVar(envPrefix, "PPROF_ADDR"),
		},
		cli.IntFlag{
			Name:   PortFlagName,
			Usage:  "pprof listening port",
			Value:  6060,
			EnvVar: ptservice.PrefixEnvVar(envPrefix, "PPROF_PORT"),
		},
	}
}

type CLIConfig struct {
	Enabled    bool
	ListenAddr string
	ListenPort int
}

func (m CLIConfig) Check() error {
	if !m.Enabled {
		return nil
	}

	if m.ListenPort < 0 || m.ListenPort > math.MaxUint16 {
		return errors.New("invalid pprof port")
	}

	return nil
}

func ReadCLIConfig(ctx *cli.Context) CLIConfig {
	return CLIConfig{
		Enabled:    ctx.GlobalBool(EnabledFlagName),
		ListenAddr: ctx.GlobalString(ListenAddrFlagName),
		ListenPort: ctx.GlobalInt(PortFlagName),
	}
}
