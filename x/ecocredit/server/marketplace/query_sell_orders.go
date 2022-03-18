package marketplace

import (
	"context"

	"github.com/cosmos/cosmos-sdk/orm/model/ormlist"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	marketplacev1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/marketplace/v1"
	ecocreditv1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/types/ormutil"
	v1 "github.com/regen-network/regen-ledger/x/ecocredit/marketplace"
)

func (k Keeper) SellOrders(ctx context.Context, req *v1.QuerySellOrdersRequest) (*v1.QuerySellOrdersResponse, error) {

	index, err := k.extractIndexFromRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	pg, err := ormutil.GogoPageReqToPulsarPageReq(req.Pagination)
	if err != nil {
		return nil, err
	}

	it, err := k.stateStore.SellOrderTable().List(ctx, index, ormlist.Paginate(pg))
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
		var sellOrder v1.SellOrder
		if err = ormutil.PulsarToGogoSlow(v, &sellOrder); err != nil {
			return nil, err
		}
		orders = append(orders, &sellOrder)
	}

	pr, err := ormutil.PulsarPageResToGogoPageRes(it.PageResponse())
	if err != nil {
		return nil, err
	}

	return &v1.QuerySellOrdersResponse{SellOrders: orders, Pagination: pr}, nil
}

func (k Keeper) extractIndexFromRequest(ctx context.Context, req *v1.QuerySellOrdersRequest) (marketplacev1.SellOrderIndexKey, error) {
	var index marketplacev1.SellOrderIndexKey

	if req.Address == "" && req.BatchDenom == "" {
		return marketplacev1.SellOrderSellerIndexKey{}, nil
	}

	// using a closure here to avoid copy pasting the wordy error
	getBatch := func(denom string) (*ecocreditv1.BatchInfo, error) {
		batch, err := k.coreStore.BatchInfoTable().GetByBatchDenom(ctx, req.BatchDenom)
		if err != nil {
			return nil, sdkerrors.ErrInvalidRequest.Wrapf("could not get batch with denom %s: %s", req.BatchDenom, err.Error())
		}
		return batch, nil
	}

	if req.Address != "" {
		addr, err := sdk.AccAddressFromBech32(req.Address)
		if err != nil {
			return nil, err
		}
		if req.BatchDenom != "" {
			batch, err := getBatch(req.BatchDenom)
			if err != nil {
				return nil, err
			}
			index = marketplacev1.SellOrderSellerBatchIdIndexKey{}.WithSellerBatchId(addr, batch.Id)
		} else {
			index = marketplacev1.SellOrderSellerIndexKey{}.WithSeller(addr)
		}
	} else {
		if req.BatchDenom != "" {
			batch, err := getBatch(req.BatchDenom)
			if err != nil {
				return nil, err
			}
			index = marketplacev1.SellOrderBatchIdIndexKey{}.WithBatchId(batch.Id)
		}
	}
	return index, nil
}
