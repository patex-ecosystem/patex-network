package immutables_test

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/patex-ecosystem/patex-network/pt-chain-ops/immutables"
	"github.com/stretchr/testify/require"
)

func TestBuildPatex(t *testing.T) {
	results, err := immutables.BuildPatex(immutables.ImmutableConfig{
		"L2StandardBridge": {
			"otherBridge": common.HexToAddress("0x1234567890123456789012345678901234567890"),
		},
		"L2CrossDomainMessenger": {
			"otherMessenger": common.HexToAddress("0x1234567890123456789012345678901234567890"),
		},
		"L2ERC721Bridge": {
			"otherBridge": common.HexToAddress("0x1234567890123456789012345678901234567890"),
			"messenger":   common.HexToAddress("0x1234567890123456789012345678901234567890"),
		},
		"PatexMintableERC721Factory": {
			"remoteChainId": big.NewInt(1),
			"bridge":        common.HexToAddress("0x1234567890123456789012345678901234567890"),
		},
		"SequencerFeeVault": {
			"recipient": common.HexToAddress("0x1234567890123456789012345678901234567890"),
		},
		"L1FeeVault": {
			"recipient": common.HexToAddress("0x1234567890123456789012345678901234567890"),
		},
		"BaseFeeVault": {
			"recipient": common.HexToAddress("0x1234567890123456789012345678901234567890"),
		},
	})
	require.Nil(t, err)
	require.NotNil(t, results)

	contracts := map[string]bool{
		"GasPriceOracle":             true,
		"L1Block":                    true,
		"L2CrossDomainMessenger":     true,
		"L2StandardBridge":           true,
		"L2ToL1MessagePasser":        true,
		"SequencerFeeVault":          true,
		"BaseFeeVault":               true,
		"L1FeeVault":                 true,
		"PatexMintableERC20Factory":  true,
		"DeployerWhitelist":          true,
		"LegacyMessagePasser":        true,
		"L1BlockNumber":              true,
		"L2ERC721Bridge":             true,
		"PatexMintableERC721Factory": true,
		"LegacyERC20ETH":             true,
	}

	// Only the exact contracts that we care about are being
	// modified
	require.Equal(t, len(results), len(contracts))

	for name, bytecode := range results {
		// There is bytecode there
		require.Greater(t, len(bytecode), 0)
		// It is in the set of contracts that we care about
		require.True(t, contracts[name])
	}
}
