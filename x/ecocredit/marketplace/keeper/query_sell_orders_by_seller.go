package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	api "github.com/regen-network/regen-ledger/api/v2/regen/ecocredit/marketplace/v1"
	regentypes "github.com/regen-network/regen-ledger/types/v2"
	regenerrors "github.com/regen-network/regen-ledger/types/v2/errors"
	"github.com/regen-network/regen-ledger/types/v2/ormutil"
	types "github.com/regen-network/regen-ledger/x/ecocredit/v3/marketplace/types/v1"
)

// SellOrdersBySeller queries all sell orders created by the given address with optional pagination
func (k Keeper) SellOrdersBySeller(ctx context.Context, req *types.QuerySellOrdersBySellerRequest) (*types.QuerySellOrdersBySellerResponse, error) {
	if req == nil {
		return nil, regenerrors.ErrInvalidArgument.Wrap("empty request")
	}

	seller, err := sdk.AccAddressFromBech32(req.Seller)
	if err != nil {
		return nil, regenerrors.ErrInvalidArgument.Wrapf("seller: %s", err.Error())
	}

	pg := ormutil.PageReqToOrmPaginate(req.Pagination)
	it, err := k.stateStore.SellOrderTable().List(
		ctx, api.SellOrderSellerIndexKey{}.WithSeller(seller), pg)
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

		batch, err := k.baseStore.BatchTable().Get(ctx, order.BatchKey)
		if err != nil {
			return nil, regenerrors.ErrNotFound.Wrapf("could not get batch with key: %d: %s", order.BatchKey, err.Error())
		}

		market, err := k.stateStore.MarketTable().Get(ctx, order.MarketId)
		if err != nil {
			return nil, regenerrors.ErrNotFound.Wrapf("could not get market with id: %d: %s", order.MarketId, err.Error())
		}

		info := types.SellOrderInfo{
			Id:                order.Id,
			Seller:            req.Seller,
			BatchDenom:        batch.Denom,
			Quantity:          order.Quantity,
			AskDenom:          market.BankDenom,
			AskAmount:         order.AskAmount,
			DisableAutoRetire: order.DisableAutoRetire,
			Expiration:        regentypes.ProtobufToGogoTimestamp(order.Expiration),
		}

		orders = append(orders, &info)
	}

	pr := ormutil.PageResToCosmosTypes(it.PageResponse())
	return &types.QuerySellOrdersBySellerResponse{SellOrders: orders, Pagination: pr}, nil
}
