package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	regentypes "github.com/regen-network/regen-ledger/types"
	types "github.com/regen-network/regen-ledger/x/ecocredit/marketplace/types/v1"
)

func (k Keeper) SellOrder(ctx context.Context, req *types.QuerySellOrderRequest) (*types.QuerySellOrderResponse, error) {
	order, err := k.stateStore.SellOrderTable().Get(ctx, req.SellOrderId)
	if err != nil {
		return nil, sdkerrors.ErrInvalidRequest.Wrapf("could not get sell order with id %d: %s", req.SellOrderId, err.Error())
	}

	seller := sdk.AccAddress(order.Seller)

	batch, err := k.baseStore.BatchTable().Get(ctx, order.BatchKey)
	if err != nil {
		return nil, err
	}

	market, err := k.stateStore.MarketTable().Get(ctx, order.MarketId)
	if err != nil {
		return nil, err
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

	return &types.QuerySellOrderResponse{SellOrder: &info}, nil
}
