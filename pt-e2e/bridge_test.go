package pt_e2e

import (
	"math"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
	"github.com/patex-ecosystem/patex-network/pt-bindings/bindings"
	"github.com/patex-ecosystem/patex-network/pt-bindings/predeploys"
	"github.com/patex-ecosystem/patex-network/pt-node/rollup/derive"
	"github.com/patex-ecosystem/patex-network/pt-node/testlog"
	"github.com/stretchr/testify/require"
)

// TestERC20BridgeDeposits tests the the L1StandardBridge bridge ERC20
// functionality.
func TestERC20BridgeDeposits(t *testing.T) {
	parallel(t)
	if !verboseGethNodes {
		log.Root().SetHandler(log.DiscardHandler())
	}

	cfg := DefaultSystemConfig(t)

	sys, err := cfg.Start()
	require.Nil(t, err, "Error starting up system")
	defer sys.Close()

	log := testlog.Logger(t, log.LvlInfo)
	log.Info("genesis", "l2", sys.RollupConfig.Genesis.L2, "l1", sys.RollupConfig.Genesis.L1, "l2_time", sys.RollupConfig.Genesis.L2Time)

	l1Client := sys.Clients["l1"]
	l2Client := sys.Clients["sequencer"]

	opts, err := bind.NewKeyedTransactorWithChainID(sys.cfg.Secrets.Alice, cfg.L1ChainIDBig())
	require.Nil(t, err)

	// Deploy WETH9
	weth9Address, tx, WETH9, err := bindings.DeployWETH9(opts, l1Client)
	require.NoError(t, err)
	_, err = waitForTransaction(tx.Hash(), l1Client, 3*time.Duration(cfg.DeployConfig.L1BlockTime)*time.Second)
	require.NoError(t, err, "Waiting for deposit tx on L1")

	// Get some WETH
	opts.Value = big.NewInt(params.Ether)
	tx, err = WETH9.Fallback(opts, []byte{})
	require.NoError(t, err)
	_, err = waitForTransaction(tx.Hash(), l1Client, 3*time.Duration(cfg.DeployConfig.L1BlockTime)*time.Second)
	require.NoError(t, err)
	opts.Value = nil
	wethBalance, err := WETH9.BalanceOf(&bind.CallOpts{}, opts.From)
	require.NoError(t, err)
	require.Equal(t, big.NewInt(params.Ether), wethBalance)

	// Deploy L2 WETH9
	l2Opts, err := bind.NewKeyedTransactorWithChainID(sys.cfg.Secrets.Alice, cfg.L2ChainIDBig())
	require.NoError(t, err)
	patexMintableTokenFactory, err := bindings.NewPatexMintableERC20Factory(predeploys.PatexMintableERC20FactoryAddr, l2Client)
	require.NoError(t, err)
	tx, err = patexMintableTokenFactory.CreatePatexMintableERC20(l2Opts, weth9Address, "L2-WETH", "L2-WETH")
	require.NoError(t, err)
	_, err = waitForTransaction(tx.Hash(), l2Client, 3*time.Duration(cfg.DeployConfig.L2BlockTime)*time.Second)
	require.NoError(t, err)

	// Get the deployment event to have access to the L2 WETH9 address
	it, err := patexMintableTokenFactory.FilterPatexMintableERC20Created(&bind.FilterOpts{Start: 0}, nil, nil)
	require.NoError(t, err)
	var event *bindings.PatexMintableERC20FactoryPatexMintableERC20Created
	for it.Next() {
		event = it.Event
	}
	require.NotNil(t, event)

	// Approve WETH9 with the bridge
	tx, err = WETH9.Approve(opts, predeploys.DevL1StandardBridgeAddr, new(big.Int).SetUint64(math.MaxUint64))
	require.NoError(t, err)
	_, err = waitForTransaction(tx.Hash(), l1Client, 3*time.Duration(cfg.DeployConfig.L1BlockTime)*time.Second)
	require.NoError(t, err)

	// Bridge the WETH9
	l1StandardBridge, err := bindings.NewL1StandardBridge(predeploys.DevL1StandardBridgeAddr, l1Client)
	require.NoError(t, err)
	tx, err = l1StandardBridge.BridgeERC20(opts, weth9Address, event.LocalToken, big.NewInt(100), 100000, []byte{})
	require.NoError(t, err)
	depositReceipt, err := waitForTransaction(tx.Hash(), l1Client, 3*time.Duration(cfg.DeployConfig.L1BlockTime)*time.Second)
	require.NoError(t, err)

	t.Log("Deposit through L1StandardBridge", "gas used", depositReceipt.GasUsed)

	// compute the deposit transaction hash + poll for it
	portal, err := bindings.NewPatexPortal(predeploys.DevPatexPortalAddr, l1Client)
	require.NoError(t, err)

	depIt, err := portal.FilterTransactionDeposited(&bind.FilterOpts{Start: 0}, nil, nil, nil)
	require.NoError(t, err)
	var depositEvent *bindings.PatexPortalTransactionDeposited
	for depIt.Next() {
		depositEvent = depIt.Event
	}
	require.NotNil(t, depositEvent)

	depositTx, err := derive.UnmarshalDepositLogEvent(&depositEvent.Raw)
	require.NoError(t, err)
	_, err = waitForTransaction(types.NewTx(depositTx).Hash(), l2Client, 3*time.Duration(cfg.DeployConfig.L2BlockTime)*time.Second)
	require.NoError(t, err)

	// Ensure that the deposit went through
	patexMintableToken, err := bindings.NewPatexMintableERC20(event.LocalToken, l2Client)
	require.NoError(t, err)

	// Should have balance on L2
	l2Balance, err := patexMintableToken.BalanceOf(&bind.CallOpts{}, opts.From)
	require.NoError(t, err)
	require.Equal(t, l2Balance, big.NewInt(100))
}
