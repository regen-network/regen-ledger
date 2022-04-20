package marketplace

import (
	"context"
	"fmt"

	"github.com/regen-network/regen-ledger/types/math"
	"github.com/regen-network/regen-ledger/x/ecocredit/marketplace"
	"github.com/regen-network/regen-ledger/x/ecocredit/server/utils"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// BuyDirect allows for the purchase of credits directly from sell orders.
func (k Keeper) BuyDirect(ctx context.Context, req *marketplace.MsgBuyDirect) (*marketplace.MsgBuyDirectResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	buyerAcc, err := sdk.AccAddressFromBech32(req.Buyer)
	if err != nil {
		return nil, err
	}
	for _, order := range req.Orders {
		sellOrder, err := k.stateStore.SellOrderTable().Get(ctx, order.SellOrderId)
		if err != nil {
			return nil, fmt.Errorf("sell order %d: %w", order.SellOrderId, err)
		}
		if order.DisableAutoRetire && !sellOrder.DisableAutoRetire {
			return nil, sdkerrors.ErrInvalidRequest.Wrapf("cannot disable auto retire when purchasing credits " +
				"from a sell order that does not have auto retire disabled")
		}
		batch, err := k.coreStore.BatchInfoTable().Get(ctx, sellOrder.BatchId)
		if err != nil {
			return nil, sdkerrors.ErrIO.Wrapf("error getting batch id %d: %s", sellOrder.BatchId, err.Error())
		}
		ct, err := utils.GetCreditTypeFromBatchDenom(ctx, k.coreStore, k.paramsKeeper, batch.BatchDenom)
		if err != nil {
			return nil, err
		}
		creditOrderQty, err := math.NewPositiveFixedDecFromString(order.Quantity, ct.Precision)
		if err != nil {
			return nil, err
		}

		// check that bid price and ask price denoms match
		market, err := k.stateStore.MarketTable().Get(ctx, sellOrder.MarketId)
		if err != nil {
			return nil, fmt.Errorf("market id %d: %w", sellOrder.MarketId, err)
		}
		if order.BidPrice.Denom != market.BankDenom {
			return nil, sdkerrors.ErrInvalidRequest.Wrapf("bid price denom does not match ask price denom: "+
				"%s, expected %s", order.BidPrice.Denom, market.BankDenom)
		}
		// check that bid price >= sell price
		sellOrderAskPrice, ok := sdk.NewIntFromString(sellOrder.AskPrice)
		if !ok {
			return nil, fmt.Errorf("could not convert %s to %T", sellOrder.AskPrice, sdk.Int{})
		}
		sellOrderPriceCoin := sdk.Coin{Denom: market.BankDenom, Amount: sellOrderAskPrice}
		if sellOrderAskPrice.GT(order.BidPrice.Amount) {
			return nil, sdkerrors.ErrInvalidRequest.Wrapf("price per credit too low: sell order ask per credit: %v, request: %v", sellOrderPriceCoin, order.BidPrice)
		}

		// check address has the total cost (price per * order quantity)
		bal := k.bankKeeper.GetBalance(sdkCtx, buyerAcc, order.BidPrice.Denom)
		cost, err := getTotalCost(sellOrderAskPrice, creditOrderQty)
		if err != nil {
			return nil, err
		}
		coinCost := sdk.Coin{Amount: cost, Denom: market.BankDenom}
		if bal.IsLT(coinCost) {
			return nil, sdkerrors.ErrInsufficientFunds.Wrapf("requested to purchase %s credits @ %s%s per "+
				"credit (total %v) with a balance of %v", order.Quantity, sellOrder.AskPrice, market.BankDenom, coinCost, bal)
		}

		// fill the order, updating balances and the sell order in state
		if err = k.fillOrder(ctx, sellOrder, buyerAcc, creditOrderQty, coinCost, orderOptions{
			autoRetire:         !order.DisableAutoRetire,
			canPartialFill:     false,
			batchDenom:         batch.BatchDenom,
			retirementLocation: order.RetirementJurisdiction,
		}); err != nil {
			return nil, fmt.Errorf("error filling order: %w", err)
		}
	}

	return &marketplace.MsgBuyDirectResponse{}, nil
}
