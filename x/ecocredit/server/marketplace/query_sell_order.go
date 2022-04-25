package marketplace

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/x/ecocredit/marketplace"
)

func (k Keeper) SellOrder(ctx context.Context, req *marketplace.QuerySellOrderRequest) (*marketplace.QuerySellOrderResponse, error) {
	order, err := k.stateStore.SellOrderTable().Get(ctx, req.SellOrderId)
	if err != nil {
		return nil, sdkerrors.ErrInvalidRequest.Wrapf("could not get sell order with id %d: %s", req.SellOrderId, err.Error())
	}

	seller := sdk.AccAddress(order.Seller)

	batch, err := k.coreStore.BatchTable().Get(ctx, order.BatchId)
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
		AskPrice:          order.AskPrice,
		DisableAutoRetire: order.DisableAutoRetire,
		Expiration:        types.ProtobufToGogoTimestamp(order.Expiration),
	}

	return &marketplace.QuerySellOrderResponse{SellOrder: &info}, nil
}
