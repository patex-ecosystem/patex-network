package predeploys

import "github.com/ethereum/go-ethereum/common"

const (
	DevL2OutputOracle            = "0x6900000000000000000000000000000000000000"
	DevPatexPortal               = "0x6900000000000000000000000000000000000001"
	DevL1CrossDomainMessenger    = "0x6900000000000000000000000000000000000002"
	DevL1StandardBridge          = "0x6900000000000000000000000000000000000003"
	DevPatexMintableERC20Factory = "0x6900000000000000000000000000000000000004"
	DevAddressManager            = "0x6900000000000000000000000000000000000005"
	DevProxyAdmin                = "0x6900000000000000000000000000000000000006"
	DevWETH9                     = "0x6900000000000000000000000000000000000007"
	DevL1ERC721Bridge            = "0x6900000000000000000000000000000000000008"
	DevSystemConfig              = "0x6900000000000000000000000000000000000009"
)

var (
	DevL2OutputOracleAddr            = common.HexToAddress(DevL2OutputOracle)
	DevPatexPortalAddr               = common.HexToAddress(DevPatexPortal)
	DevL1CrossDomainMessengerAddr    = common.HexToAddress(DevL1CrossDomainMessenger)
	DevL1StandardBridgeAddr          = common.HexToAddress(DevL1StandardBridge)
	DevPatexMintableERC20FactoryAddr = common.HexToAddress(DevPatexMintableERC20Factory)
	DevAddressManagerAddr            = common.HexToAddress(DevAddressManager)
	DevProxyAdminAddr                = common.HexToAddress(DevProxyAdmin)
	DevWETH9Addr                     = common.HexToAddress(DevWETH9)
	DevL1ERC721BridgeAddr            = common.HexToAddress(DevL1ERC721Bridge)
	DevSystemConfigAddr              = common.HexToAddress(DevSystemConfig)

	DevPredeploys = make(map[string]*common.Address)
)

func init() {
	DevPredeploys["L2OutputOracle"] = &DevL2OutputOracleAddr
	DevPredeploys["PatexPortal"] = &DevPatexPortalAddr
	DevPredeploys["L1CrossDomainMessenger"] = &DevL1CrossDomainMessengerAddr
	DevPredeploys["L1StandardBridge"] = &DevL1StandardBridgeAddr
	DevPredeploys["PatexMintableERC20Factory"] = &DevPatexMintableERC20FactoryAddr
	DevPredeploys["AddressManager"] = &DevAddressManagerAddr
	DevPredeploys["Admin"] = &DevProxyAdminAddr
	DevPredeploys["WETH9"] = &DevWETH9Addr
	DevPredeploys["L1ERC721Bridge"] = &DevL1ERC721BridgeAddr
	DevPredeploys["SystemConfig"] = &DevSystemConfigAddr
}
