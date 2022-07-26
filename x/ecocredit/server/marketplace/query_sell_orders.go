package marketplace

import (
	"context"

	"github.com/cosmos/cosmos-sdk/orm/model/ormlist"
	sdk "github.com/cosmos/cosmos-sdk/types"
	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/marketplace/v1"
	"github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/types/ormutil"
	"github.com/regen-network/regen-ledger/x/ecocredit/v2/marketplace"
)

// SellOrders queries all sell orders in state with optional pagination
func (k Keeper) SellOrders(ctx context.Context, req *marketplace.QuerySellOrdersRequest) (*marketplace.QuerySellOrdersResponse, error) {
	pg, err := ormutil.GogoPageReqToPulsarPageReq(req.Pagination)
	if err != nil {
		return nil, err
	}

	it, err := k.stateStore.SellOrderTable().List(ctx, api.SellOrderSellerIndexKey{}, ormlist.Paginate(pg))
	if err != nil {
		return nil, err
	}
	defer it.Close()

	orders := make([]*marketplace.SellOrderInfo, 0, 10)
	for it.Next() {
		order, err := it.Value()
		if err != nil {
			return nil, err
		}

		seller := sdk.AccAddress(order.Seller)

		batch, err := k.coreStore.BatchTable().Get(ctx, order.BatchKey)
		if err != nil {
			return nil, err
		}

		market, err := k.stateStore.MarketTable().Get(ctx, order.MarketId)
		if err != nil {
			return nil, err
		}

		info := marketplace.SellOrderInfo{
			Id:                order.Id,
			Seller:            seller.String(),
			BatchDenom:        batch.Denom,
			Quantity:          order.Quantity,
			AskDenom:          market.BankDenom,
			AskAmount:         order.AskAmount,
			DisableAutoRetire: order.DisableAutoRetire,
			Expiration:        types.ProtobufToGogoTimestamp(order.Expiration),
		}

		orders = append(orders, &info)
	}

	pr, err := ormutil.PulsarPageResToGogoPageRes(it.PageResponse())
	if err != nil {
		return nil, err
	}

	return &marketplace.QuerySellOrdersResponse{SellOrders: orders, Pagination: pr}, nil
}
