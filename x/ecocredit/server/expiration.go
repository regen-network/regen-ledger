package server

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// PruneOrders checks if there are any expired sell or buy orders and removes them from state.
func (s serverImpl) PruneOrders(ctx sdk.Context) error {
	return s.MarketplaceKeeper.PruneSellOrders(sdk.WrapSDKContext(ctx))
}
