package marketplace

import (
	"context"

	"github.com/regen-network/regen-ledger/x/ecocredit/marketplace"
)

func (k Keeper) SellOrders(context.Context, *marketplace.QuerySellOrdersRequest) (*marketplace.QuerySellOrdersResponse, error) {
	panic("not implemented")
}

func (k Keeper) SellOrdersByAddress(context.Context, *marketplace.QuerySellOrdersByAddressRequest) (*marketplace.QuerySellOrdersByAddressResponse, error) {
	panic("not implemented")
}

func (k Keeper) SellOrdersByBatchDenom(context.Context, *marketplace.QuerySellOrdersByBatchDenomRequest) (*marketplace.QuerySellOrdersByBatchDenomResponse, error) {
	panic("not implemented")
}
