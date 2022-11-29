package keeper

import (
	"context"

	api "github.com/regen-network/regen-ledger/api/v2/regen/ecocredit/v1"
	types "github.com/regen-network/regen-ledger/x/ecocredit/v3/base/types/v1"
)

func (k Keeper) AllowedBridgeChains(ctx context.Context, _ *types.QueryAllowedBridgeChainsRequest) (*types.QueryAllowedBridgeChainsResponse, error) {
	it, err := k.stateStore.AllowedBridgeChainTable().List(ctx, api.AllowedBridgeChainPrimaryKey{})
	if err != nil {
		return nil, err
	}
	defer it.Close()

	var chains []string
	for it.Next() {
		entry, err := it.Value()
		if err != nil {
			return nil, err
		}
		chains = append(chains, entry.ChainName)
	}

	return &types.QueryAllowedBridgeChainsResponse{AllowedBridgeChains: chains}, nil
}
