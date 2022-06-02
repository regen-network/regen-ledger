package marketplace

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	v1 "github.com/regen-network/regen-ledger/x/ecocredit/marketplace"
)

// CancelSellOrder cancels a sell order and returns the escrowed credits to the seller.
func (k Keeper) CancelSellOrder(ctx context.Context, req *v1.MsgCancelSellOrder) (*v1.MsgCancelSellOrderResponse, error) {
	sellerAcc, err := sdk.AccAddressFromBech32(req.Seller)
	if err != nil {
		return nil, err
	}

	sellOrder, err := k.stateStore.SellOrderTable().Get(ctx, req.SellOrderId)
	if err != nil {
		return nil, fmt.Errorf("sell order with id %d: %w", req.SellOrderId, err)
	}

	if !sellerAcc.Equals(sdk.AccAddress(sellOrder.Seller)) {
		return nil, sdkerrors.ErrUnauthorized.Wrapf("seller must be the owner of the sell order")
	}

	err = k.unescrowCredits(ctx, sellerAcc, sellOrder.BatchKey, sellOrder.Quantity)
	if err != nil {
		return nil, err
	}

	return &v1.MsgCancelSellOrderResponse{}, k.stateStore.SellOrderTable().Delete(ctx, sellOrder)
}
