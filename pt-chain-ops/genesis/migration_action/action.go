package migration_action

import (
	"context"
	"math/big"
	"path/filepath"

	"github.com/patex-ecosystem/patex-network/pt-chain-ops/crossdomain"

	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/patex-ecosystem/patex-network/pt-chain-ops/genesis"
)

type Config struct {
	DeployConfig          *genesis.DeployConfig
	PVMAddressesPath      string
	EVMAddressesPath      string
	PVMAllowancesPath     string
	PVMMessagesPath       string
	EVMMessagesPath       string
	Network               string
	HardhatDeployments    []string
	L1URL                 string
	StartingL1BlockNumber uint64
	L2DBPath              string
	DryRun                bool
	NoCheck               bool
}

func Migrate(cfg *Config) (*genesis.MigrationResult, error) {
	deployConfig := cfg.DeployConfig

	pvmAddresses, err := crossdomain.NewAddresses(cfg.PVMAddressesPath)
	if err != nil {
		return nil, err
	}
	evmAddresess, err := crossdomain.NewAddresses(cfg.EVMAddressesPath)
	if err != nil {
		return nil, err
	}
	pvmAllowances, err := crossdomain.NewAllowances(cfg.PVMAllowancesPath)
	if err != nil {
		return nil, err
	}
	pvmMessages, err := crossdomain.NewSentMessageFromJSON(cfg.PVMMessagesPath)
	if err != nil {
		return nil, err
	}
	evmMessages, err := crossdomain.NewSentMessageFromJSON(cfg.EVMMessagesPath)
	if err != nil {
		return nil, err
	}

	migrationData := crossdomain.MigrationData{
		OvmAddresses:  pvmAddresses,
		EvmAddresses:  evmAddresess,
		OvmAllowances: pvmAllowances,
		OvmMessages:   pvmMessages,
		EvmMessages:   evmMessages,
	}

	l1Client, err := ethclient.Dial(cfg.L1URL)
	if err != nil {
		return nil, err
	}
	var blockNumber *big.Int
	bnum := cfg.StartingL1BlockNumber
	if bnum != 0 {
		blockNumber = new(big.Int).SetUint64(bnum)
	}

	block, err := l1Client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		return nil, err
	}

	chaindataPath := filepath.Join(cfg.L2DBPath, "geth", "chaindata")
	ancientPath := filepath.Join(chaindataPath, "ancient")
	ldb, err := rawdb.Open(
		rawdb.OpenOptions{
			Type:              "leveldb",
			Directory:         chaindataPath,
			Cache:             4096,
			Handles:           120,
			AncientsDirectory: ancientPath,
			Namespace:         "",
			ReadOnly:          false,
		})
	if err != nil {
		return nil, err
	}
	defer ldb.Close()

	return genesis.MigrateDB(ldb, deployConfig, block, &migrationData, !cfg.DryRun, cfg.NoCheck)
}
