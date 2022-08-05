package marketplace

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/regen-network/regen-ledger/x/ecocredit/marketplace"
)

// CancelSellOrder cancels a sell order and returns the escrowed credits to the seller.
func (k Keeper) CancelSellOrder(ctx context.Context, req *marketplace.MsgCancelSellOrder) (*marketplace.MsgCancelSellOrderResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	sellerAcc, err := sdk.AccAddressFromBech32(req.Seller)
	if err != nil {
		return nil, err
	}

	sellOrder, err := k.stateStore.SellOrderTable().Get(ctx, req.SellOrderId)
	if err != nil {
		return nil, sdkerrors.ErrInvalidRequest.Wrapf("sell order with id %d: %s", req.SellOrderId, err.Error())
	}

	if !sellerAcc.Equals(sdk.AccAddress(sellOrder.Seller)) {
		return nil, sdkerrors.ErrUnauthorized.Wrapf("seller must be the owner of the sell order")
	}

	err = k.unescrowCredits(ctx, sellerAcc, sellOrder.BatchKey, sellOrder.Quantity)
	if err != nil {
		return nil, err
	}

	if err = sdkCtx.EventManager().EmitTypedEvent(&marketplace.EventCancelSellOrder{
		SellOrderId: sellOrder.Id,
	}); err != nil {
		return nil, err
	}

	return &marketplace.MsgCancelSellOrderResponse{}, k.stateStore.SellOrderTable().Delete(ctx, sellOrder)
}
