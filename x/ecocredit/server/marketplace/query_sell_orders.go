package marketplace

import (
	"context"

	"github.com/cosmos/cosmos-sdk/orm/model/ormlist"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	marketplacev1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/marketplace/v1"
	"github.com/regen-network/regen-ledger/types/ormutil"
	v1 "github.com/regen-network/regen-ledger/x/ecocredit/marketplace"
)

// SellOrders queries all sell orders in state with optional pagination
func (k Keeper) SellOrders(ctx context.Context, req *v1.QuerySellOrdersRequest) (*v1.QuerySellOrdersResponse, error) {
	pg, err := ormutil.GogoPageReqToPulsarPageReq(req.Pagination)
	if err != nil {
		return nil, err
	}

	it, err := k.stateStore.SellOrderTable().List(ctx, marketplacev1.SellOrderSellerIndexKey{}, ormlist.Paginate(pg))
	if err != nil {
		return nil, err
	}
	defer it.Close()

	orders := make([]*v1.SellOrder, 0, 10)
	for it.Next() {
		v, err := it.Value()
		if err != nil {
			return nil, err
		}
		var order v1.SellOrder
		err = ormutil.PulsarToGogoSlow(v, &order)
		if err != nil {
			return nil, err
		}
		orders = append(orders, &order)
	}
	return &v1.QuerySellOrdersResponse{SellOrders: orders}, nil
}

// SellOrdersByBatchDenom queries all sell orders under a specific batch denom with optional pagination
func (k Keeper) SellOrdersByBatchDenom(ctx context.Context, req *v1.QuerySellOrdersByBatchDenomRequest) (*v1.QuerySellOrdersByBatchDenomResponse, error) {
	pg, err := ormutil.GogoPageReqToPulsarPageReq(req.Pagination)
	if err != nil {
		return nil, err
	}

	batch, err := k.coreStore.BatchInfoTable().GetByBatchDenom(ctx, req.BatchDenom)
	if err != nil {
		return nil, sdkerrors.ErrInvalidRequest.Wrapf("could not get batch with denom %s: %s", req.BatchDenom, err.Error())
	}

	it, err := k.stateStore.SellOrderTable().List(ctx, marketplacev1.SellOrderBatchIdIndexKey{}.WithBatchId(batch.Id), ormlist.Paginate(pg))
	if err != nil {
		return nil, err
	}
	defer it.Close()

	orders := make([]*v1.SellOrder, 0, 10)
	for it.Next() {
		v, err := it.Value()
		if err != nil {
			return nil, err
		}

		var order v1.SellOrder
		err = ormutil.PulsarToGogoSlow(v, &order)
		if err != nil {
			return nil, err
		}

		orders = append(orders, &order)
	}

	return &v1.QuerySellOrdersByBatchDenomResponse{SellOrders: orders}, nil
}

// SellOrdersByAddress queries all sell orders created by the given address with optional pagination
func (k Keeper) SellOrdersByAddress(ctx context.Context, req *v1.QuerySellOrdersByAddressRequest) (*v1.QuerySellOrdersByAddressResponse, error) {
	pg, err := ormutil.GogoPageReqToPulsarPageReq(req.Pagination)
	if err != nil {
		return nil, err
	}

	seller, err := sdk.AccAddressFromBech32(req.Address)
	if err != nil {
		return nil, err
	}

	it, err := k.stateStore.SellOrderTable().List(ctx, marketplacev1.SellOrderSellerIndexKey{}.WithSeller(seller), ormlist.Paginate(pg))
	if err != nil {
		return nil, err
	}
	defer it.Close()

	orders := make([]*v1.SellOrder, 0, 10)
	for it.Next() {
		v, err := it.Value()
		if err != nil {
			return nil, err
		}

		var order v1.SellOrder
		err = ormutil.PulsarToGogoSlow(v, &order)
		if err != nil {
			return nil, err
		}

		orders = append(orders, &order)
	}

	return &v1.QuerySellOrdersByAddressResponse{SellOrders: orders}, nil
}
