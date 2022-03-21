package marketplace

import (
	"context"

	"github.com/regen-network/regen-ledger/x/ecocredit/marketplace"
)

func (k Keeper) SellOrders(context.Context, *marketplace.QuerySellOrdersRequest) (*marketplace.QuerySellOrdersResponse, error) {
	panic("implement me")
}

func (k Keeper) SellOrdersByAddress(context.Context, *marketplace.QuerySellOrdersByAddressRequest) (*marketplace.QuerySellOrdersByAddressResponse, error) {
	panic("implement me")
}

func (k Keeper) SellOrdersByBatchDenom(context.Context, *marketplace.QuerySellOrdersByBatchDenomRequest) (*marketplace.QuerySellOrdersByBatchDenomResponse, error) {
	panic("implement me")
}
