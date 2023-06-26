package chaincfg

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	"github.com/patex-ecosystem/patex-network/pt-node/eth"
	"github.com/patex-ecosystem/patex-network/pt-node/rollup"
)

var Mainnet = rollup.Config{
	Genesis: rollup.Genesis{
		L1: eth.BlockID{
			// moose: Update this during migration
			Hash: common.HexToHash("0x"),
			// moose: Update this during migration
			Number: 0,
		},
		L2: eth.BlockID{
			// moose: Update this during migration
			Hash: common.HexToHash("0x"),
			// moose: Update this during migration
			Number: 0,
		},
		// moose: Update this during migration
		L2Time: 0,
		SystemConfig: eth.SystemConfig{
			BatcherAddr: common.HexToAddress("0x6887246668a3b87f54deb3b94ba47a6f63f32985"),
			Overhead:    eth.Bytes32(common.HexToHash("0x00000000000000000000000000000000000000000000000000000000000000bc")),
			Scalar:      eth.Bytes32(common.HexToHash("0x00000000000000000000000000000000000000000000000000000000000a6fe0")),
			GasLimit:    30_000_000,
		},
	},
	BlockTime:              2,
	MaxSequencerDrift:      600,
	SeqWindowSize:          3600,
	ChannelTimeout:         300,
	L1ChainID:              big.NewInt(1),
	L2ChainID:              big.NewInt(10),
	BatchInboxAddress:      common.HexToAddress("0xff00000000000000000000000000000000000010"),
	DepositContractAddress: common.HexToAddress("0xbEb5Fc579115071764c7423A4f12eDde41f106Ed"),
	L1SystemConfigAddress:  common.HexToAddress("0x229047fed2591dbec1eF1118d64F7aF3dB9EB290"),
	RegolithTime:           u64Ptr(0),
}

var Goerli = rollup.Config{
	Genesis: rollup.Genesis{
		L1: eth.BlockID{
			Hash:   common.HexToHash("0x6ffc1bf3754c01f6bb9fe057c1578b87a8571ce2e9be5ca14bace6eccfd336c7"),
			Number: 8300214,
		},
		L2: eth.BlockID{
			Hash:   common.HexToHash("0x0f783549ea4313b784eadd9b8e8a69913b368b7366363ea814d7707ac505175f"),
			Number: 4061224,
		},
		L2Time: 1673550516,
		SystemConfig: eth.SystemConfig{
			BatcherAddr: common.HexToAddress("0x7431310e026B69BFC676C0013E12A1A11411EEc9"),
			Overhead:    eth.Bytes32(common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000834")),
			Scalar:      eth.Bytes32(common.HexToHash("0x00000000000000000000000000000000000000000000000000000000000f4240")),
			GasLimit:    25_000_000,
		},
	},
	BlockTime:              2,
	MaxSequencerDrift:      600,
	SeqWindowSize:          3600,
	ChannelTimeout:         300,
	L1ChainID:              big.NewInt(5),
	L2ChainID:              big.NewInt(420),
	BatchInboxAddress:      common.HexToAddress("0xff00000000000000000000000000000000000420"),
	DepositContractAddress: common.HexToAddress("0x5b47E1A08Ea6d985D6649300584e6722Ec4B1383"),
	L1SystemConfigAddress:  common.HexToAddress("0xAe851f927Ee40dE99aaBb7461C00f9622ab91d60"),
	RegolithTime:           u64Ptr(1679079600),
}

var PatexSepolia = rollup.Config{
	Genesis: rollup.Genesis{
		L1: eth.BlockID{
			Hash:   common.HexToHash("0x42fb77897fbc2ce4fdbd4c3416da22efe6e0a306b20d9aaf5ed1e7d852a9fdff"),
			Number: 3635475,
		},
		L2: eth.BlockID{
			Hash:   common.HexToHash("0xbb78fd2ea1c0012fe6bdce3f1cdb60de0fc896fa87c9282d6de660a39107dff7"),
			Number: 0,
		},
		L2Time: 1686058464,
		SystemConfig: eth.SystemConfig{
			BatcherAddr: common.HexToAddress("0x87c0da3a03218be364b89b460c1376752c5cec96"),
			Overhead:    eth.Bytes32(common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000834")),
			Scalar:      eth.Bytes32(common.HexToHash("0x00000000000000000000000000000000000000000000000000000000000f4240")),
			GasLimit:    30_000_000,
		},
	},
	BlockTime:              10,
	MaxSequencerDrift:      600,
	SeqWindowSize:          3600,
	ChannelTimeout:         300,
	L1ChainID:              big.NewInt(11155111),
	L2ChainID:              big.NewInt(471100),
	BatchInboxAddress:      common.HexToAddress("0xff00000000000000000000000000000000042069"),
	DepositContractAddress: common.HexToAddress("0xd7400a9e3bd054264be87443939770dcf23e5b95"),
	L1SystemConfigAddress:  common.HexToAddress("0x95eb0167854a4a342ae5b6636afa4015e13f67fd"),
	RegolithTime:           u64Ptr(0),
}

var PatexMainnet = rollup.Config{
	Genesis: rollup.Genesis{
		L1: eth.BlockID{
			Hash:   common.HexToHash("0x44b2078a814136a8af22c3135c88a7e2279c7dd0ff270ba538c776905a7b2594"),
			Number: 17434125,
		},
		L2: eth.BlockID{
			Hash:   common.HexToHash("0x37379affb896facb96f88ac39a1bc91ea312263df9b18dd3ec978dfa59fd48bb"),
			Number: 0,
		},
		L2Time: 1686209495,
		SystemConfig: eth.SystemConfig{
			BatcherAddr: common.HexToAddress("0x6088b06c5a187058434655b71057a9ee93e13d0d"),
			Overhead:    eth.Bytes32(common.HexToHash("0x00000000000000000000000000000000000000000000000000000000000000bc")),
			Scalar:      eth.Bytes32(common.HexToHash("0x00000000000000000000000000000000000000000000000000000000000a6fe0")),
			GasLimit:    30_000_000,
		},
	},
	BlockTime:              10,
	MaxSequencerDrift:      600,
	SeqWindowSize:          3600,
	ChannelTimeout:         300,
	L1ChainID:              big.NewInt(1),
	L2ChainID:              big.NewInt(789),
	BatchInboxAddress:      common.HexToAddress("0xff00000000000000000000000000000000000010"),
	DepositContractAddress: common.HexToAddress("0x0d6e11e2a3b2b3a245bf839c07d775983acb787d"),
	L1SystemConfigAddress:  common.HexToAddress("0x04c9303d5727335d067e5085704976110f9f088e"),
	RegolithTime:           u64Ptr(0),
}

var NetworksByName = map[string]rollup.Config{
	"goerli":        Goerli,
	"patex-sepolia": PatexSepolia,
	"patex-mainnet": PatexMainnet,
	// moose: Update this during migration
	// "mainnet": Mainnet,
}

var L2ChainIDToNetworkName = func() map[string]string {
	out := make(map[string]string)
	for name, netCfg := range NetworksByName {
		out[netCfg.L2ChainID.String()] = name
	}
	return out
}()

func AvailableNetworks() []string {
	var networks []string
	for name := range NetworksByName {
		networks = append(networks, name)
	}
	return networks
}

func GetRollupConfig(name string) (rollup.Config, error) {
	network, ok := NetworksByName[name]
	if !ok {
		return rollup.Config{}, fmt.Errorf("invalid network %s", name)
	}

	return network, nil
}

func u64Ptr(v uint64) *uint64 {
	return &v
}
