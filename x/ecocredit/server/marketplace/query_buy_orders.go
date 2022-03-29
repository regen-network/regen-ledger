package marketplace

import (
	"context"

	"github.com/cosmos/cosmos-sdk/orm/model/ormlist"
	sdk "github.com/cosmos/cosmos-sdk/types"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/marketplace/v1"
	"github.com/regen-network/regen-ledger/types/ormutil"
	"github.com/regen-network/regen-ledger/x/ecocredit/marketplace"
)

// BuyOrders queries all buy orders with optional pagination
func (k Keeper) BuyOrders(ctx context.Context, request *marketplace.QueryBuyOrdersRequest) (*marketplace.QueryBuyOrdersResponse, error) {
	pg, err := ormutil.GogoPageReqToPulsarPageReq(request.Pagination)
	if err != nil {
		return nil, err
	}

	it, err := k.stateStore.BuyOrderTable().List(ctx, api.BuyOrderPrimaryKey{}, ormlist.Paginate(pg))
	if err != nil {
		return nil, err
	}
	defer it.Close()

	orders := make([]*marketplace.BuyOrder, 0, 10)
	for it.Next() {
		v, err := it.Value()
		if err != nil {
			return nil, err
		}
		var order marketplace.BuyOrder
		if err = ormutil.PulsarToGogoSlow(v, &order); err != nil {
			return nil, err
		}
		orders = append(orders, &order)
	}

	pr, err := ormutil.PulsarPageResToGogoPageRes(it.PageResponse())
	if err != nil {
		return nil, err
	}

	return &marketplace.QueryBuyOrdersResponse{BuyOrders: orders, Pagination: pr}, nil
}

// BuyOrdersByAddress queries all buy orders created by the given address with optional pagination
func (k Keeper) BuyOrdersByAddress(ctx context.Context, request *marketplace.QueryBuyOrdersByAddressRequest) (*marketplace.QueryBuyOrdersByAddressResponse, error) {
	pg, err := ormutil.GogoPageReqToPulsarPageReq(request.Pagination)
	if err != nil {
		return nil, err
	}

	buyer, err := sdk.AccAddressFromBech32(request.Address)
	if err != nil {
		return nil, err
	}

	it, err := k.stateStore.BuyOrderTable().List(ctx, api.BuyOrderBuyerIndexKey{}.WithBuyer(buyer), ormlist.Paginate(pg))
	if err != nil {
		return nil, err
	}
	defer it.Close()

	orders := make([]*marketplace.BuyOrder, 0, 10)
	for it.Next() {
		v, err := it.Value()
		if err != nil {
			return nil, err
		}
		var order marketplace.BuyOrder
		if err = ormutil.PulsarToGogoSlow(v, &order); err != nil {
			return nil, err
		}
		orders = append(orders, &order)
	}

	pr, err := ormutil.PulsarPageResToGogoPageRes(it.PageResponse())
	if err != nil {
		return nil, err
	}

	return &marketplace.QueryBuyOrdersByAddressResponse{BuyOrders: orders, Pagination: pr}, nil
}
