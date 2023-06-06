package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"

	"github.com/ethereum/go-ethereum/log"
	"github.com/patex-ecosystem/patex-network/pt-batcher/batcher"
	"github.com/patex-ecosystem/patex-network/pt-batcher/cmd/doc"
	"github.com/patex-ecosystem/patex-network/pt-batcher/flags"
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
	app.Name = "pt-batcher"
	app.Usage = "Batch Submitter Service"
	app.Description = "Service for generating and submitting L2 tx batches to L1"
	app.Action = curryMain(Version)
	app.Commands = []cli.Command{
		{
			Name:        "doc",
			Subcommands: doc.Subcommands,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Crit("Application failed", "message", err)
	}
}

// curryMain transforms the batcher.Main function into an app.Action
// This is done to capture the Version of the batcher.
func curryMain(version string) func(ctx *cli.Context) error {
	return func(ctx *cli.Context) error {
		return batcher.Main(version, ctx)
	}
}
