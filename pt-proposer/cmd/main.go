package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"

	"github.com/ethereum/go-ethereum/log"
	"github.com/patex-ecosystem/patex-network/pt-proposer/flags"
	"github.com/patex-ecosystem/patex-network/pt-proposer/proposer"
	ptlog "github.com/patex-ecosystem/patex-network/pt-service/log"
)

var (
	Version   = "v0.10.14"
	GitCommit = ""
	GitDate   = ""
)

func main() {
	ptlog.SetupDefaults()

	app := cli.NewApp()
	app.Flags = flags.Flags
	app.Version = fmt.Sprintf("%s-%s-%s", Version, GitCommit, GitDate)
	app.Name = "pt-proposer"
	app.Usage = "L2Output Submitter"
	app.Description = "Service for generating and submitting L2 Output checkpoints to the L2OutputOracle contract"

	app.Action = curryMain(Version)
	err := app.Run(os.Args)
	if err != nil {
		log.Crit("Application failed", "message", err)
	}
}

// curryMain transforms the proposer.Main function into an app.Action
// This is done to capture the Version of the proposer.
func curryMain(version string) func(ctx *cli.Context) error {
	return func(ctx *cli.Context) error {
		return proposer.Main(version, ctx)
	}
}
