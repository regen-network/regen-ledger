package marketplace

import (
	"context"

	"github.com/cosmos/cosmos-sdk/orm/model/ormlist"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/marketplace/v1"
	"github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/types/ormutil"
	"github.com/regen-network/regen-ledger/x/ecocredit/marketplace"
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

// SellOrdersByBatchDenom queries all sell orders under a specific batch denom with optional pagination
func (k Keeper) SellOrdersByBatchDenom(ctx context.Context, req *marketplace.QuerySellOrdersByBatchDenomRequest) (*marketplace.QuerySellOrdersByBatchDenomResponse, error) {
	pg, err := ormutil.GogoPageReqToPulsarPageReq(req.Pagination)
	if err != nil {
		return nil, err
	}

	batch, err := k.coreStore.BatchTable().GetByDenom(ctx, req.BatchDenom)
	if err != nil {
		return nil, sdkerrors.ErrInvalidRequest.Wrapf("could not get batch with denom %s: %s", req.BatchDenom, err.Error())
	}

	it, err := k.stateStore.SellOrderTable().List(ctx, api.SellOrderBatchKeyIndexKey{}.WithBatchKey(batch.Key), ormlist.Paginate(pg))
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

	return &marketplace.QuerySellOrdersByBatchDenomResponse{SellOrders: orders, Pagination: pr}, nil
}

// SellOrdersByAddress queries all sell orders created by the given address with optional pagination
func (k Keeper) SellOrdersByAddress(ctx context.Context, req *marketplace.QuerySellOrdersByAddressRequest) (*marketplace.QuerySellOrdersByAddressResponse, error) {
	pg, err := ormutil.GogoPageReqToPulsarPageReq(req.Pagination)
	if err != nil {
		return nil, err
	}

	seller, err := sdk.AccAddressFromBech32(req.Address)
	if err != nil {
		return nil, err
	}

	it, err := k.stateStore.SellOrderTable().List(ctx, api.SellOrderSellerIndexKey{}.WithSeller(seller), ormlist.Paginate(pg))
	if err != nil {
		return nil, err
	}
	defer it.Close()

	sellerString := seller.String()
	orders := make([]*marketplace.SellOrderInfo, 0, 10)
	for it.Next() {
		order, err := it.Value()
		if err != nil {
			return nil, err
		}

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
			Seller:            sellerString,
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

	return &marketplace.QuerySellOrdersByAddressResponse{SellOrders: orders, Pagination: pr}, nil
}
