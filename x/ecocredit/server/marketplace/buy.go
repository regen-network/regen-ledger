package marketplace

import (
	"context"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	marketplacev1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/marketplace/v1"
	v1 "github.com/regen-network/regen-ledger/x/ecocredit/marketplace"
)

func (k Keeper) Buy(ctx context.Context, req *v1.MsgBuy) (*v1.MsgBuyResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	buyerAcc, err := sdk.AccAddressFromBech32(req.Buyer)
	if err != nil {
		return nil, err
	}

	buyOrderIds := make([]uint64, len(req.Orders))
	for i, order := range req.Orders {
		// verify expiration is in the future
		if order.Expiration != nil && order.Expiration.Before(sdkCtx.BlockTime()) {
			return nil, sdkerrors.ErrInvalidRequest.Wrapf("expiration must be in the future: %s", order.Expiration)
		}
		// TODO: is bid price the price for 1 credit? or the price they're willing to pay
		// for `Quantity`?
		if !k.bankKeeper.HasBalance(sdkCtx, buyerAcc, *order.BidPrice) {
			return nil, sdkerrors.ErrInsufficientFunds
		}

		switch selection := order.Selection.Sum.(type) {
		case *v1.MsgBuy_Order_Selection_SellOrderId:
			sellOrder, err := k.stateStore.SellOrderStore().Get(ctx, selection.SellOrderId)
			if err != nil {
				return nil, fmt.Errorf("sell order %d: %w", selection.SellOrderId, err)
			}
			market, err := k.stateStore.MarketStore().Get(ctx, sellOrder.MarketId)
			if err != nil {
				return nil, fmt.Errorf("market id %d: %w", sellOrder.MarketId, err)
			}
			if order.BidPrice.Denom != market.BankDenom {
				return nil, sdkerrors.ErrInvalidRequest.Wrapf("bid price denom does not match ask price denom: "+
					" %s, expected %s", order.BidPrice.Denom, market.BankDenom)
			}
			askAmount, ok := sdk.NewIntFromString(sellOrder.AskPrice)
			if !ok {
				return nil, fmt.Errorf("could not convert ask price to sdk.Int: %s", sellOrder.AskPrice)
			}
			if order.BidPrice.Amount.LT(askAmount) {
				return nil, sdkerrors.ErrInsufficientFunds.Wrapf("bid price too low: got %s, needed at least %s",
					order.BidPrice.Amount.String(), sellOrder.AskPrice)
			}
			// TODO: we're doing escrow right? usually we check if seller still has credits, but wont need to with escrow model.

		}

	}
}
