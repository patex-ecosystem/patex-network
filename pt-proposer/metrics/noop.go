package metrics

import (
	"github.com/patex-ecosystem/patex-network/pt-node/eth"
	ptmetrics "github.com/patex-ecosystem/patex-network/pt-service/metrics"
	txmetrics "github.com/patex-ecosystem/patex-network/pt-service/txmgr/metrics"
)

type noptmetrics struct {
	ptmetrics.NoopRefMetrics
	txmetrics.NoopTxMetrics
}

var Noptmetrics Metricer = new(noptmetrics)

func (*noptmetrics) RecordInfo(version string) {}
func (*noptmetrics) RecordUp()                 {}

func (*noptmetrics) RecordL2BlocksProposed(l2ref eth.L2BlockRef) {}
