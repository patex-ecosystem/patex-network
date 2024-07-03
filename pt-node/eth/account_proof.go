package eth

import (
	"bytes"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethdb/memorydb"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/trie"
)

type StorageProofEntry struct {
	Key   common.Hash     `json:"key"`
	Value hexutil.Big     `json:"value"`
	Proof []hexutil.Bytes `json:"proof"`
}

type AccountResult struct {
	AccountProof []hexutil.Bytes `json:"accountProof"`

	Address     common.Address `json:"address"`
	Balance     *hexutil.Big   `json:"balance"`
	Fixed       *hexutil.Big   `json:"fixed"`
	Shares      *hexutil.Big   `json:"shares"`
	Remainder   *hexutil.Big   `json:"remainder"`
	CodeHash    common.Hash    `json:"codeHash"`
	Nonce       hexutil.Uint64 `json:"nonce"`
	Flags       hexutil.Uint64 `json:"flags"`
	StorageHash common.Hash    `json:"storageHash"`

	// Optional
	StorageProof []StorageProofEntry `json:"storageProof,omitempty"`
}

// Verify an account (and optionally storage) proof from the getProof RPC. See https://eips.ethereum.org/EIPS/eip-1186
func (res *AccountResult) verify(stateRoot common.Hash, accountClaimed []any) error {
	// verify storage proof values, if any, against the storage trie root hash of the account
	for i, entry := range res.StorageProof {
		// load all MPT nodes into a DB
		db := memorydb.New()
		for j, encodedNode := range entry.Proof {
			nodeKey := encodedNode
			if len(encodedNode) >= 32 { // small MPT nodes are not hashed
				nodeKey = crypto.Keccak256(encodedNode)
			}
			if err := db.Put(nodeKey, encodedNode); err != nil {
				return fmt.Errorf("failed to load storage proof node %d of storage value %d into mem db: %w", j, i, err)
			}
		}
		path := crypto.Keccak256(entry.Key[:])
		val, err := trie.VerifyProof(res.StorageHash, path, db)
		if err != nil {
			return fmt.Errorf("failed to verify storage value %d with key %s (path %x) in storage trie %s: %w", i, entry.Key, path, res.StorageHash, err)
		}
		if val == nil && entry.Value.ToInt().Cmp(common.Big0) == 0 { // empty storage is zero by default
			continue
		}
		comparison, err := rlp.EncodeToBytes(entry.Value.ToInt().Bytes())
		if err != nil {
			return fmt.Errorf("failed to encode storage value %d with key %s (path %x) in storage trie %s: %w", i, entry.Key, path, res.StorageHash, err)
		}
		if !bytes.Equal(val, comparison) {
			return fmt.Errorf("value %d in storage proof does not match proven value at key %s (path %x)", i, entry.Key, path)
		}
	}

	accountClaimedValue, err := rlp.EncodeToBytes(accountClaimed)
	if err != nil {
		return fmt.Errorf("failed to encode account from retrieved values: %w", err)
	}

	// create a db with all account trie nodes
	db := memorydb.New()
	for i, encodedNode := range res.AccountProof {
		nodeKey := encodedNode
		if len(encodedNode) >= 32 { // small MPT nodes are not hashed
			nodeKey = crypto.Keccak256(encodedNode)
		}
		if err := db.Put(nodeKey, encodedNode); err != nil {
			return fmt.Errorf("failed to load account proof node %d into mem db: %w", i, err)
		}
	}
	path := crypto.Keccak256(res.Address[:])
	accountProofValue, err := trie.VerifyProof(stateRoot, path, db)
	if err != nil {
		return fmt.Errorf("failed to verify account value with key %s (path %x) in account trie %s: %w", res.Address, path, stateRoot, err)
	}

	if !bytes.Equal(accountClaimedValue, accountProofValue) {
		return fmt.Errorf("L1 RPC is tricking us, account proof does not match provided deserialized values:\n"+
			"  claimed: %x\n"+
			"  proof:   %x", accountClaimedValue, accountProofValue)
	}
	return err
}

func (res *AccountResult) Verify(stateRoot common.Hash) error {
	accountClaimed := []any{uint64(res.Nonce), res.Balance.ToInt().Bytes(), res.StorageHash, res.CodeHash}
	return res.verify(stateRoot, accountClaimed)
}

func (res *AccountResult) VerifyL2(stateRoot common.Hash) error {
	accountClaimed := []any{uint64(res.Nonce), uint8(res.Flags), res.Fixed.ToInt().Bytes(), res.Shares.ToInt().Bytes(), res.Remainder.ToInt().Bytes(), res.StorageHash, res.CodeHash}
	return res.verify(stateRoot, accountClaimed)
}
