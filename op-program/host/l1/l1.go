package l1

import (
	"context"

	"github.com/ethereum-optimism/pepe/op-node/client"
	"github.com/ethereum-optimism/pepe/op-node/rollup/derive"
	"github.com/ethereum-optimism/pepe/op-node/sources"
	cll1 "github.com/ethereum-optimism/pepe/op-program/client/l1"
	"github.com/ethereum-optimism/pepe/op-program/host/config"
	"github.com/ethereum-optimism/pepe/op-program/preimage"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
)

func NewFetchingOracle(ctx context.Context, logger log.Logger, cfg *config.Config) (cll1.Oracle, error) {
	rpc, err := client.NewRPC(ctx, logger, cfg.L1URL)
	if err != nil {
		return nil, err
	}

	source, err := sources.NewL1Client(rpc, logger, nil, sources.L1ClientDefaultConfig(cfg.Rollup, cfg.L1TrustRPC, cfg.L1RPCKind))
	if err != nil {
		return nil, err
	}
	return NewFetchingL1Oracle(ctx, logger, source), nil
}

func NewSource(logger log.Logger, oracle preimage.Oracle, hint preimage.Hinter, l1Head common.Hash) derive.L1Fetcher {
	l1Oracle := cll1.NewCachingOracle(cll1.NewPreimageOracle(oracle, hint))
	return cll1.NewOracleL1Client(logger, l1Oracle, l1Head)
}
