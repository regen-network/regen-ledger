package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/orm/model/ormlist"
	sdk "github.com/cosmos/cosmos-sdk/types"

	api "github.com/regen-network/regen-ledger/api/v2/regen/ecocredit/marketplace/v1"
	regentypes "github.com/regen-network/regen-ledger/types/v2"
	regenerrors "github.com/regen-network/regen-ledger/types/v2/errors"
	"github.com/regen-network/regen-ledger/types/v2/ormutil"
	types "github.com/regen-network/regen-ledger/x/ecocredit/v3/marketplace/types/v1"
)

// SellOrdersByBatch queries all sell orders under a specific batch denom with optional pagination
func (k Keeper) SellOrdersByBatch(ctx context.Context, req *types.QuerySellOrdersByBatchRequest) (*types.QuerySellOrdersByBatchResponse, error) {
	if req == nil {
		return nil, regenerrors.ErrInvalidArgument.Wrap("empty request")
	}

	pg, err := ormutil.GogoPageReqToPulsarPageReq(req.Pagination)
	if err != nil {
		return nil, regenerrors.ErrInvalidArgument.Wrap(err.Error())
	}

	batch, err := k.baseStore.BatchTable().GetByDenom(ctx, req.BatchDenom)
	if err != nil {
		return nil, regenerrors.ErrNotFound.Wrapf("could not get batch with denom %s: %s", req.BatchDenom, err.Error())
	}

	it, err := k.stateStore.SellOrderTable().List(ctx, api.SellOrderBatchKeyIndexKey{}.WithBatchKey(batch.Key), ormlist.Paginate(pg))
	if err != nil {
		return nil, err
	}
	defer it.Close()

	orders := make([]*types.SellOrderInfo, 0, 10)
	for it.Next() {
		order, err := it.Value()
		if err != nil {
			return nil, err
		}

		seller := sdk.AccAddress(order.Seller)

		market, err := k.stateStore.MarketTable().Get(ctx, order.MarketId)
		if err != nil {
			return nil, regenerrors.ErrNotFound.Wrapf("could not get market with id %d: %s", order.MarketId, err.Error())
		}

		info := types.SellOrderInfo{
			Id:                order.Id,
			Seller:            seller.String(),
			BatchDenom:        batch.Denom,
			Quantity:          order.Quantity,
			AskDenom:          market.BankDenom,
			AskAmount:         order.AskAmount,
			DisableAutoRetire: order.DisableAutoRetire,
			Expiration:        regentypes.ProtobufToGogoTimestamp(order.Expiration),
		}

		orders = append(orders, &info)
	}

	pr, err := ormutil.PulsarPageResToGogoPageRes(it.PageResponse())
	if err != nil {
		return nil, regenerrors.ErrInternal.Wrap(err.Error())
	}

	return &types.QuerySellOrdersByBatchResponse{SellOrders: orders, Pagination: pr}, nil
}
