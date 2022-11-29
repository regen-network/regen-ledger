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

// SellOrders queries all sell orders in state with optional pagination
func (k Keeper) SellOrders(ctx context.Context, req *types.QuerySellOrdersRequest) (*types.QuerySellOrdersResponse, error) {
	if req == nil {
		return nil, regenerrors.ErrInvalidArgument.Wrap("empty request")
	}

	pg, err := ormutil.GogoPageReqToPulsarPageReq(req.Pagination)
	if err != nil {
		return nil, regenerrors.ErrInvalidArgument.Wrap(err.Error())
	}

	it, err := k.stateStore.SellOrderTable().List(ctx, api.SellOrderSellerIndexKey{}, ormlist.Paginate(pg))
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

		batch, err := k.baseStore.BatchTable().Get(ctx, order.BatchKey)
		if err != nil {
			return nil, regenerrors.ErrNotFound.Wrapf("unable to get batch with key: %d: %s", order.BatchKey, err.Error())
		}

		market, err := k.stateStore.MarketTable().Get(ctx, order.MarketId)
		if err != nil {
			return nil, regenerrors.ErrNotFound.Wrapf("unable to get market with id: %d: %s", order.MarketId, err.Error())
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

	return &types.QuerySellOrdersResponse{SellOrders: orders, Pagination: pr}, nil
}
