package marketplace

import (
	"context"

	"github.com/cosmos/cosmos-sdk/orm/model/ormlist"
	sdk "github.com/cosmos/cosmos-sdk/types"

	marketplacev1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/marketplace/v1"
	"github.com/regen-network/regen-ledger/types/ormutil"
	v1 "github.com/regen-network/regen-ledger/x/ecocredit/marketplace"
)

// BuyOrders queries all buy orders with optional pagination
func (k Keeper) BuyOrders(ctx context.Context, request *v1.QueryBuyOrdersRequest) (*v1.QueryBuyOrdersResponse, error) {
	pg, err := ormutil.GogoPageReqToPulsarPageReq(request.Pagination)
	if err != nil {
		return nil, err
	}

	it, err := k.stateStore.BuyOrderTable().List(ctx, marketplacev1.BuyOrderPrimaryKey{}, ormlist.Paginate(pg))
	if err != nil {
		return nil, err
	}
	defer it.Close()

	orders := make([]*v1.BuyOrder, 0, 10)
	for it.Next() {
		v, err := it.Value()
		if err != nil {
			return nil, err
		}
		var order v1.BuyOrder
		if err = ormutil.PulsarToGogoSlow(v, &order); err != nil {
			return nil, err
		}
		orders = append(orders, &order)
	}
	return &v1.QueryBuyOrdersResponse{BuyOrders: orders}, nil
}

// BuyOrdersByAddress queries all buy orders created by the given address with optional pagination
func (k Keeper) BuyOrdersByAddress(ctx context.Context, request *v1.QueryBuyOrdersByAddressRequest) (*v1.QueryBuyOrdersByAddressResponse, error) {
	pg, err := ormutil.GogoPageReqToPulsarPageReq(request.Pagination)
	if err != nil {
		return nil, err
	}

	buyer, err := sdk.AccAddressFromBech32(request.Address)
	if err != nil {
		return nil, err
	}

	it, err := k.stateStore.BuyOrderTable().List(ctx, marketplacev1.BuyOrderBuyerIndexKey{}.WithBuyer(buyer), ormlist.Paginate(pg))
	if err != nil {
		return nil, err
	}
	defer it.Close()

	orders := make([]*v1.BuyOrder, 0, 10)
	for it.Next() {
		v, err := it.Value()
		if err != nil {
			return nil, err
		}
		var order v1.BuyOrder
		if err = ormutil.PulsarToGogoSlow(v, &order); err != nil {
			return nil, err
		}
		orders = append(orders, &order)
	}
	return &v1.QueryBuyOrdersByAddressResponse{BuyOrders: orders}, nil
}
