package metrics

import (
	"github.com/ethereum-optimism/pepe/op-node/eth"
	opmetrics "github.com/ethereum-optimism/pepe/op-service/metrics"
	txmetrics "github.com/ethereum-optimism/pepe/op-service/txmgr/metrics"
)

type noopMetrics struct {
	opmetrics.NoopRefMetrics
	txmetrics.NoopTxMetrics
}

var NoopMetrics Metricer = new(noopMetrics)

func (*noopMetrics) RecordInfo(version string) {}
func (*noopMetrics) RecordUp()                 {}

func (*noopMetrics) RecordL2BlocksProposed(l2ref eth.L2BlockRef) {}
