package metrics

import (
	"github.com/patex-ecosystem/patex-network/pt-node/eth"
	"github.com/patex-ecosystem/patex-network/pt-node/rollup/derive"
	ptmetrics "github.com/patex-ecosystem/patex-network/pt-service/metrics"
	txmetrics "github.com/patex-ecosystem/patex-network/pt-service/txmgr/metrics"
)

type noptmetrics struct {
	ptmetrics.NoopRefMetrics
	txmetrics.NoopTxMetrics
}

var Noptmetrics Metricer = new(noptmetrics)

func (*noptmetrics) Document() []ptmetrics.DocumentedMetric { return nil }

func (*noptmetrics) RecordInfo(version string) {}
func (*noptmetrics) RecordUp()                 {}

func (*noptmetrics) RecordLatestL1Block(l1ref eth.L1BlockRef)               {}
func (*noptmetrics) RecordL2BlocksLoaded(eth.L2BlockRef)                    {}
func (*noptmetrics) RecordChannelOpened(derive.ChannelID, int)              {}
func (*noptmetrics) RecordL2BlocksAdded(eth.L2BlockRef, int, int, int, int) {}

func (*noptmetrics) RecordChannelClosed(derive.ChannelID, int, int, int, int, error) {}

func (*noptmetrics) RecordChannelFullySubmitted(derive.ChannelID) {}
func (*noptmetrics) RecordChannelTimedOut(derive.ChannelID)       {}

func (*noptmetrics) RecordBatchTxSubmitted() {}
func (*noptmetrics) RecordBatchTxSuccess()   {}
func (*noptmetrics) RecordBatchTxFailed()    {}
