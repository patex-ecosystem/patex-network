package main

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"strings"

	"github.com/patex-ecosystem/patex-network/pt-chain-ops/crossdomain"

	"github.com/mattn/go-isatty"
	"github.com/patex-ecosystem/patex-network/pt-chain-ops/db"

	"github.com/ethereum/go-ethereum/common"
	"github.com/patex-ecosystem/patex-network/pt-node/eth"
	"github.com/patex-ecosystem/patex-network/pt-node/rollup/derive"

	"github.com/ethereum/go-ethereum/core/types"

	"github.com/ethereum/go-ethereum/log"

	"github.com/patex-ecosystem/patex-network/pt-bindings/hardhat"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/patex-ecosystem/patex-network/pt-chain-ops/genesis"

	"github.com/urfave/cli"
)

func main() {
	log.Root().SetHandler(log.StreamHandler(os.Stderr, log.TerminalFormat(isatty.IsTerminal(os.Stderr.Fd()))))

	app := &cli.App{
		Name:  "check-migration",
		Usage: "Run sanity checks on a migrated database",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "l1-rpc-url",
				Value:    "http://127.0.0.1:8545",
				Usage:    "RPC URL for an L1 Node",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "pvm-addresses",
				Usage:    "Path to pvm-addresses.json",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "pvm-allowances",
				Usage:    "Path to pvm-allowances.json",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "pvm-messages",
				Usage:    "Path to pvm-messages.json",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "witness-file",
				Usage:    "Path to witness file",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "db-path",
				Usage:    "Path to database",
				Required: true,
			},
			cli.StringFlag{
				Name:     "deploy-config",
				Usage:    "Path to hardhat deploy config file",
				Required: true,
			},
			cli.StringFlag{
				Name:     "network",
				Usage:    "Name of hardhat deploy network",
				Required: true,
			},
			cli.StringFlag{
				Name:     "hardhat-deployments",
				Usage:    "Comma separated list of hardhat deployment directories",
				Required: true,
			},
			cli.IntFlag{
				Name:  "db-cache",
				Usage: "LevelDB cache size in mb",
				Value: 1024,
			},
			cli.IntFlag{
				Name:  "db-handles",
				Usage: "LevelDB number of handles",
				Value: 60,
			},
		},
		Action: func(ctx *cli.Context) error {
			deployConfig := ctx.String("deploy-config")
			config, err := genesis.NewDeployConfig(deployConfig)
			if err != nil {
				return err
			}

			pvmAddresses, err := crossdomain.NewAddresses(ctx.String("pvm-addresses"))
			if err != nil {
				return err
			}
			pvmAllowances, err := crossdomain.NewAllowances(ctx.String("pvm-allowances"))
			if err != nil {
				return err
			}
			pvmMessages, err := crossdomain.NewSentMessageFromJSON(ctx.String("pvm-messages"))
			if err != nil {
				return err
			}
			evmMessages, evmAddresses, err := crossdomain.ReadWitnessData(ctx.String("witness-file"))
			if err != nil {
				return err
			}

			log.Info(
				"Loaded witness data",
				"pvmAddresses", len(pvmAddresses),
				"evmAddresses", len(evmAddresses),
				"pvmAllowances", len(pvmAllowances),
				"pvmMessages", len(pvmMessages),
				"evmMessages", len(evmMessages),
			)

			migrationData := crossdomain.MigrationData{
				OvmAddresses:  pvmAddresses,
				EvmAddresses:  evmAddresses,
				OvmAllowances: pvmAllowances,
				OvmMessages:   pvmMessages,
				EvmMessages:   evmMessages,
			}

			network := ctx.String("network")
			deployments := strings.Split(ctx.String("hardhat-deployments"), ",")
			hh, err := hardhat.New(network, []string{}, deployments)
			if err != nil {
				return err
			}

			l1RpcURL := ctx.String("l1-rpc-url")
			l1Client, err := ethclient.Dial(l1RpcURL)
			if err != nil {
				return err
			}

			var block *types.Block
			tag := config.L1StartingBlockTag
			if tag.BlockNumber != nil {
				block, err = l1Client.BlockByNumber(context.Background(), big.NewInt(tag.BlockNumber.Int64()))
			} else if tag.BlockHash != nil {
				block, err = l1Client.BlockByHash(context.Background(), *tag.BlockHash)
			} else {
				return fmt.Errorf("invalid l1StartingBlockTag in deploy config: %v", tag)
			}
			if err != nil {
				return err
			}

			dbCache := ctx.Int("db-cache")
			dbHandles := ctx.Int("db-handles")

			// Read the required deployment addresses from disk if required
			if err := config.GetDeployedAddresses(hh); err != nil {
				return err
			}

			if err := config.Check(); err != nil {
				return err
			}

			postLDB, err := db.Open(ctx.String("db-path"), dbCache, dbHandles)
			if err != nil {
				return err
			}

			if err := genesis.PostCheckMigratedDB(
				postLDB,
				migrationData,
				&config.L1CrossDomainMessengerProxy,
				config.L1ChainID,
				config.L2ChainID,
				config.FinalSystemOwner,
				config.ProxyAdminOwner,
				&derive.L1BlockInfo{
					Number:        block.NumberU64(),
					Time:          block.Time(),
					BaseFee:       block.BaseFee(),
					BlockHash:     block.Hash(),
					BatcherAddr:   config.BatchSenderAddress,
					L1FeeOverhead: eth.Bytes32(common.BigToHash(new(big.Int).SetUint64(config.GasPriceOracleOverhead))),
					L1FeeScalar:   eth.Bytes32(common.BigToHash(new(big.Int).SetUint64(config.GasPriceOracleScalar))),
				},
			); err != nil {
				return err
			}

			if err := postLDB.Close(); err != nil {
				return err
			}

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Crit("error in migration", "err", err)
	}
}
