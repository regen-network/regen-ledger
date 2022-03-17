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

func (k Keeper) SellOrders(ctx context.Context, req *v1.QuerySellOrdersRequest) (*v1.QuerySellOrdersResponse, error) {
	var seller sdk.AccAddress
	if req.Address != "" {
		addr, err := sdk.AccAddressFromBech32(req.Address)
		if err != nil {
			return nil, err
		}
		seller = addr
	}

	pg, err := ormutil.GogoPageReqToPulsarPageReq(req.Pagination)
	if err != nil {
		return nil, err
	}

	var batchId uint64
	batch, err := k.coreStore.BatchInfoTable().GetByBatchDenom(ctx, req.BatchDenom)
	if err != nil {
		return nil, sdkerrors.ErrInvalidRequest.Wrapf("could not get batch with denom %s: %s", req.BatchDenom, err.Error())
	}

	it, err := k.stateStore.SellOrderTable().List(ctx, marketplacev1.SellOrderBatchIdSellerIndexKey{}.WithBatchIdSeller(batch.Id, seller), ormlist.Paginate(pg))
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
	return &v1.QuerySellOrdersResponse{SellOrders: orders}, nil
}
