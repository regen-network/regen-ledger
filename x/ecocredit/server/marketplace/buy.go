package marketplace

import (
	"context"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	v1 "github.com/regen-network/regen-ledger/x/ecocredit/marketplace"
)

func (k Keeper) Buy(ctx context.Context, req *v1.MsgBuy) (*v1.MsgBuyResponse, error) {
	// setup
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	buyerAcc, err := sdk.AccAddressFromBech32(req.Buyer)
	if err != nil {
		return nil, err
	}

	// range over orders, keeping a slice of order ID's created.
	buyOrderIds := make([]uint64, len(req.Orders))
	for i, order := range req.Orders {
		// verify expiration is in the future
		if order.Expiration != nil && order.Expiration.Before(sdkCtx.BlockTime()) {
			return nil, sdkerrors.ErrInvalidRequest.Wrapf("expiration must be in the future: %s", order.Expiration)
		}

		if !k.bankKeeper.HasBalance(sdkCtx, buyerAcc, *order.BidPrice) {
			return nil, sdkerrors.ErrInsufficientFunds
		}

		switch selection := order.Selection.Sum.(type) {
		case *v1.MsgBuy_Order_Selection_SellOrderId:
			// get the sell order
			sellOrder, err := k.stateStore.SellOrderStore().Get(ctx, selection.SellOrderId)
			if err != nil {
				return nil, fmt.Errorf("sell order %d: %w", selection.SellOrderId, err)
			}
			market, err := k.stateStore.MarketStore().Get(ctx, sellOrder.MarketId)
			if err != nil {
				return nil, fmt.Errorf("market id %d: %w", sellOrder.MarketId, err)
			}
			// check that the bid and ask denoms are the same.
			if order.BidPrice.Denom != market.BankDenom {
				return nil, sdkerrors.ErrInvalidRequest.Wrapf("bid price denom does not match ask price denom: "+
					" %s, expected %s", order.BidPrice.Denom, market.BankDenom)
			}
			// check that bid price is at least equal to ask price
			askAmount, ok := sdk.NewIntFromString(sellOrder.AskPrice)
			if !ok {
				return nil, fmt.Errorf("could not convert ask price to sdk.Int: %s", sellOrder.AskPrice)
			}
			if order.BidPrice.Amount.LT(askAmount) {
				return nil, sdkerrors.ErrInsufficientFunds.Wrapf("bid price too low: got %s, needed at least %s",
					order.BidPrice.Amount.String(), sellOrder.AskPrice)
			}

		}

	}
	return &v1.MsgBuyResponse{BuyOrderIds: buyOrderIds}, nil
}
