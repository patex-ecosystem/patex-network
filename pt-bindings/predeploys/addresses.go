package predeploys

import "github.com/ethereum/go-ethereum/common"

const (
	L2ToL1MessagePasser        = "0x4200000000000000000000000000000000000016"
	DeployerWhitelist          = "0x4200000000000000000000000000000000000002"
	LegacyERC20ETH             = "0xDeadDeAddeAddEAddeadDEaDDEAdDeaDDeAD0000"
	WETH9                      = "0x4200000000000000000000000000000000000006"
	L2CrossDomainMessenger     = "0x4200000000000000000000000000000000000007"
	L2StandardBridge           = "0x4200000000000000000000000000000000000010"
	SequencerFeeVault          = "0x4200000000000000000000000000000000000011"
	PatexMintableERC20Factory  = "0x4200000000000000000000000000000000000012"
	L1BlockNumber              = "0x4200000000000000000000000000000000000013"
	GasPriceOracle             = "0x420000000000000000000000000000000000000F"
	L1Block                    = "0x4200000000000000000000000000000000000015"
	GovernanceToken            = "0x4200000000000000000000000000000000000042"
	LegacyMessagePasser        = "0x4200000000000000000000000000000000000000"
	L2ERC721Bridge             = "0x4200000000000000000000000000000000000014"
	PatexMintableERC721Factory = "0x4200000000000000000000000000000000000017"
	ProxyAdmin                 = "0x4200000000000000000000000000000000000018"
	BaseFeeVault               = "0x4200000000000000000000000000000000000019"
	L1FeeVault                 = "0x420000000000000000000000000000000000001a"
)

var (
	L2ToL1MessagePasserAddr        = common.HexToAddress(L2ToL1MessagePasser)
	DeployerWhitelistAddr          = common.HexToAddress(DeployerWhitelist)
	LegacyERC20ETHAddr             = common.HexToAddress(LegacyERC20ETH)
	WETH9Addr                      = common.HexToAddress(WETH9)
	L2CrossDomainMessengerAddr     = common.HexToAddress(L2CrossDomainMessenger)
	L2StandardBridgeAddr           = common.HexToAddress(L2StandardBridge)
	SequencerFeeVaultAddr          = common.HexToAddress(SequencerFeeVault)
	PatexMintableERC20FactoryAddr  = common.HexToAddress(PatexMintableERC20Factory)
	L1BlockNumberAddr              = common.HexToAddress(L1BlockNumber)
	GasPriceOracleAddr             = common.HexToAddress(GasPriceOracle)
	L1BlockAddr                    = common.HexToAddress(L1Block)
	GovernanceTokenAddr            = common.HexToAddress(GovernanceToken)
	LegacyMessagePasserAddr        = common.HexToAddress(LegacyMessagePasser)
	L2ERC721BridgeAddr             = common.HexToAddress(L2ERC721Bridge)
	PatexMintableERC721FactoryAddr = common.HexToAddress(PatexMintableERC721Factory)
	ProxyAdminAddr                 = common.HexToAddress(ProxyAdmin)
	BaseFeeVaultAddr               = common.HexToAddress(BaseFeeVault)
	L1FeeVaultAddr                 = common.HexToAddress(L1FeeVault)

	Predeploys = make(map[string]*common.Address)
)

func init() {
	Predeploys["L2ToL1MessagePasser"] = &L2ToL1MessagePasserAddr
	Predeploys["DeployerWhitelist"] = &DeployerWhitelistAddr
	Predeploys["LegacyERC20ETH"] = &LegacyERC20ETHAddr
	Predeploys["WETH9"] = &WETH9Addr
	Predeploys["L2CrossDomainMessenger"] = &L2CrossDomainMessengerAddr
	Predeploys["L2StandardBridge"] = &L2StandardBridgeAddr
	Predeploys["SequencerFeeVault"] = &SequencerFeeVaultAddr
	Predeploys["PatexMintableERC20Factory"] = &PatexMintableERC20FactoryAddr
	Predeploys["L1BlockNumber"] = &L1BlockNumberAddr
	Predeploys["GasPriceOracle"] = &GasPriceOracleAddr
	Predeploys["L1Block"] = &L1BlockAddr
	Predeploys["GovernanceToken"] = &GovernanceTokenAddr
	Predeploys["LegacyMessagePasser"] = &LegacyMessagePasserAddr
	Predeploys["L2ERC721Bridge"] = &L2ERC721BridgeAddr
	Predeploys["PatexMintableERC721Factory"] = &PatexMintableERC721FactoryAddr
	Predeploys["ProxyAdmin"] = &ProxyAdminAddr
	Predeploys["BaseFeeVault"] = &BaseFeeVaultAddr
	Predeploys["L1FeeVault"] = &L1FeeVaultAddr
}
