package marketplace

import (
	"context"
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
		return nil, err
	}
	if !sellerAcc.Equals(sdk.AccAddress(sellOrder.Seller)) {
		return nil, sdkerrors.ErrUnauthorized.Wrapf("sell order was created by %s", sdk.AccAddress(sellOrder.Seller).String())
	}

	return &v1.MsgCancelSellOrderResponse{}, k.unescrowCredits(ctx, sellerAcc, sellOrder.BatchId, sellOrder.Quantity)
}
