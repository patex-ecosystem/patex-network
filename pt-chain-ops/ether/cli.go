package ether

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	"github.com/ethereum/go-ethereum/core/state"
)

// getPVMETHTotalSupply returns PVM ETH's total supply by reading
// the appropriate storage slot.
func getPVMETHTotalSupply(db *state.StateDB) *big.Int {
	key := getPVMETHTotalSupplySlot()
	return db.GetState(PVMETHAddress, key).Big()
}

func getPVMETHTotalSupplySlot() common.Hash {
	position := common.Big2
	key := common.BytesToHash(common.LeftPadBytes(position.Bytes(), 32))
	return key
}

func GetPVMETHTotalSupplySlot() common.Hash {
	return getPVMETHTotalSupplySlot()
}

// GetPVMETHBalance gets a user's PVM ETH balance from state by querying the
// appropriate storage slot directly.
func GetPVMETHBalance(db *state.StateDB, addr common.Address) *big.Int {
	return db.GetState(PVMETHAddress, CalcPVMETHStorageKey(addr)).Big()
}
