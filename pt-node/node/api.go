package node

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/log"

	"github.com/patex-ecosystem/patex-network/pt-bindings/predeploys"
	"github.com/patex-ecosystem/patex-network/pt-node/eth"
	"github.com/patex-ecosystem/patex-network/pt-node/rollup"
	"github.com/patex-ecosystem/patex-network/pt-node/version"
)

type l2EthClient interface {
	InfoByHash(ctx context.Context, hash common.Hash) (eth.BlockInfo, error)
	// GetProof returns a proof of the account, it may return a nil result without error if the address was not found.
	// Optionally keys of the account storage trie can be specified to include with corresponding values in the proof.
	GetProof(ctx context.Context, address common.Address, storage []common.Hash, blockTag string) (*eth.AccountResult, error)
}

type driverClient interface {
	SyncStatus(ctx context.Context) (*eth.SyncStatus, error)
	BlockRefWithStatus(ctx context.Context, num uint64) (eth.L2BlockRef, *eth.SyncStatus, error)
	ResetDerivationPipeline(context.Context) error
	StartSequencer(ctx context.Context, blockHash common.Hash) error
	StopSequencer(context.Context) (common.Hash, error)
	CalibrateCurrentL1Origin(context.Context)
}

type rpcMetrics interface {
	// RecordRPCServerRequest returns a function that records the duration of serving the given RPC method
	RecordRPCServerRequest(method string) func()
}

type adminAPI struct {
	dr driverClient
	m  rpcMetrics
}

func NewAdminAPI(dr driverClient, m rpcMetrics) *adminAPI {
	return &adminAPI{
		dr: dr,
		m:  m,
	}
}

func (n *adminAPI) ResetDerivationPipeline(ctx context.Context) error {
	recordDur := n.m.RecordRPCServerRequest("admin_resetDerivationPipeline")
	defer recordDur()
	return n.dr.ResetDerivationPipeline(ctx)
}

func (n *adminAPI) StartSequencer(ctx context.Context, blockHash common.Hash) error {
	recordDur := n.m.RecordRPCServerRequest("admin_startSequencer")
	defer recordDur()
	return n.dr.StartSequencer(ctx, blockHash)
}

func (n *adminAPI) StopSequencer(ctx context.Context) (common.Hash, error) {
	recordDur := n.m.RecordRPCServerRequest("admin_stopSequencer")
	defer recordDur()
	return n.dr.StopSequencer(ctx)
}

type nodeAPI struct {
	config *rollup.Config
	client l2EthClient
	dr     driverClient
	log    log.Logger
	m      rpcMetrics
}

func NewNodeAPI(config *rollup.Config, l2Client l2EthClient, dr driverClient, log log.Logger, m rpcMetrics) *nodeAPI {
	return &nodeAPI{
		config: config,
		client: l2Client,
		dr:     dr,
		log:    log,
		m:      m,
	}
}

func (n *nodeAPI) OutputAtBlock(ctx context.Context, number hexutil.Uint64) (*eth.OutputResponse, error) {
	recordDur := n.m.RecordRPCServerRequest("patex_outputAtBlock")
	defer recordDur()

	ref, status, err := n.dr.BlockRefWithStatus(ctx, uint64(number))
	if err != nil {
		return nil, fmt.Errorf("failed to get L2 block ref with sync status: %w", err)
	}

	head, err := n.client.InfoByHash(ctx, ref.Hash)
	if err != nil {
		return nil, fmt.Errorf("failed to get L2 block by hash %s: %w", ref, err)
	}
	if head == nil {
		return nil, ethereum.NotFound
	}

	proof, err := n.client.GetProof(ctx, predeploys.L2ToL1MessagePasserAddr, []common.Hash{}, ref.Hash.String())
	if err != nil {
		return nil, fmt.Errorf("failed to get contract proof at block %s: %w", ref, err)
	}
	if proof == nil {
		return nil, fmt.Errorf("proof %w", ethereum.NotFound)
	}
	// make sure that the proof (including storage hash) that we retrieved is correct by verifying it against the state-root
	if err := proof.Verify(head.Root()); err != nil {
		n.log.Error("invalid withdrawal root detected in block", "stateRoot", head.Root(), "blocknum", number, "msg", err)
		return nil, fmt.Errorf("invalid withdrawal root hash, state root was %s: %w", head.Root(), err)
	}

	var l2OutputRootVersion eth.Bytes32 // it's zero for now
	l2OutputRoot, err := rollup.ComputeL2OutputRootV0(head, proof.StorageHash)
	if err != nil {
		n.log.Error("Error computing L2 output root, nil ptr passed to hashing function")
		return nil, err
	}

	return &eth.OutputResponse{
		Version:               l2OutputRootVersion,
		OutputRoot:            l2OutputRoot,
		BlockRef:              ref,
		WithdrawalStorageRoot: proof.StorageHash,
		StateRoot:             head.Root(),
		Status:                status,
	}, nil
}

func (n *nodeAPI) SyncStatus(ctx context.Context) (*eth.SyncStatus, error) {
	recordDur := n.m.RecordRPCServerRequest("patex_syncStatus")
	defer recordDur()
	return n.dr.SyncStatus(ctx)
}

func (n *nodeAPI) RollupConfig(_ context.Context) (*rollup.Config, error) {
	recordDur := n.m.RecordRPCServerRequest("patex_rollupConfig")
	defer recordDur()
	return n.config, nil
}

func (n *nodeAPI) Version(ctx context.Context) (string, error) {
	recordDur := n.m.RecordRPCServerRequest("patex_version")
	defer recordDur()
	return version.Version + "-" + version.Meta, nil
}

func (n *nodeAPI) CalibrateCurrentL1Origin(ctx context.Context) (string, error) {
	recordDur := n.m.RecordRPCServerRequest("admin_calibrateL1Origin")
	defer recordDur()
	n.dr.CalibrateCurrentL1Origin(ctx)
	return "calibrated", nil
}

//complete *atomic.Bool
