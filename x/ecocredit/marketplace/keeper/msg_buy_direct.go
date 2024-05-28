package keeper

import (
	"context"
	"fmt"

	sdkmath "cosmossdk.io/math"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/regen-network/regen-ledger/types/v2/math"
	types "github.com/regen-network/regen-ledger/x/ecocredit/v3/marketplace/types/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/v3/server/utils"
)

// BuyDirect allows for the purchase of credits directly from sell orders.
func (k Keeper) BuyDirect(ctx context.Context, req *types.MsgBuyDirect) (*types.MsgBuyDirectResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	buyerAcc, err := sdk.AccAddressFromBech32(req.Buyer)
	if err != nil {
		return nil, err
	}

	for i, order := range req.Orders {
		// orderIndex is used for more granular error messages when
		// an individual order in a list of orders fails to process
		orderIndex := fmt.Sprintf("orders[%d]", i)

		sellOrder, err := k.stateStore.SellOrderTable().Get(ctx, order.SellOrderId)
		if err != nil {
			return nil, sdkerrors.ErrInvalidRequest.Wrapf(
				"%s: sell order with id %d: %s",
				orderIndex, order.SellOrderId, err,
			)
		}

		// check if buyer account is equal to seller account
		if buyerAcc.Equals(sdk.AccAddress(sellOrder.Seller)) {
			return nil, sdkerrors.ErrUnauthorized.Wrapf(
				"%s: buyer account cannot be the same as seller account", orderIndex,
			)
		}

		// check if disable auto-retire is required
		if order.DisableAutoRetire && !sellOrder.DisableAutoRetire {
			return nil, sdkerrors.ErrInvalidRequest.Wrapf(
				"%s: cannot disable auto-retire for a sell order with auto-retire enabled", orderIndex,
			)
		}

		// check decimal places does not exceed credit type precision
		batch, err := k.baseStore.BatchTable().Get(ctx, sellOrder.BatchKey)
		if err != nil {
			return nil, err
		}
		ct, err := utils.GetCreditTypeFromBatch(ctx, k.baseStore, batch)
		if err != nil {
			return nil, err
		}
		buyQuantity, err := math.NewPositiveFixedDecFromString(order.Quantity, ct.Precision)
		if err != nil {
			return nil, sdkerrors.ErrInvalidRequest.Wrapf(
				"%s: decimal places exceeds precision: quantity: %s, credit type precision: %d",
				orderIndex, order.Quantity, ct.Precision,
			)
		}

		// check that bid price and ask price denoms match
		market, err := k.stateStore.MarketTable().Get(ctx, sellOrder.MarketId)
		if err != nil {
			return nil, sdkerrors.ErrInvalidRequest.Wrapf("market id %d: %s", sellOrder.MarketId, err)
		}
		if order.BidPrice.Denom != market.BankDenom {
			return nil, sdkerrors.ErrInvalidRequest.Wrapf(
				"%s: bid price denom: %s, ask price denom: %s",
				orderIndex, order.BidPrice.Denom, market.BankDenom,
			)
		}

		// check that bid price >= sell price
		sellOrderAskAmount, ok := sdk.NewIntFromString(sellOrder.AskAmount)
		if !ok {
			return nil, sdkerrors.ErrInvalidType.Wrapf("could not convert %s to %T", sellOrder.AskAmount, sdkmath.Int{})
		}
		sellOrderPriceCoin := sdk.Coin{Denom: market.BankDenom, Amount: sellOrderAskAmount}
		if sellOrderAskAmount.GT(order.BidPrice.Amount) {
			return nil, sdkerrors.ErrInvalidRequest.Wrapf(
				"%s: ask price: %v, bid price: %v, insufficient bid price",
				orderIndex, sellOrderPriceCoin, order.BidPrice,
			)
		}

		// calc sub-total (price per * order quantity)
		subtotal, err := getSubTotalCost(sellOrderAskAmount, buyQuantity)
		if err != nil {
			return nil, err
		}

		// add buyer fees
		feeParams, err := k.stateStore.FeeParamsTable().Get(ctx)
		if err != nil {
			return nil, err
		}
		total, buyerFee, err := getTotalCostAndBuyerFee(subtotal, feeParams)
		if err != nil {
			return nil, err
		}
		totalCost := sdk.Coin{Amount: total.SdkIntTrim(), Denom: market.BankDenom}

		// check max fee
		maxFee := order.MaxFeeAmount
		if maxFee == nil {
			maxFee = &sdk.Coin{Amount: sdk.NewInt(0), Denom: market.BankDenom}
		}
		buyerFeeCoin := sdk.Coin{Amount: buyerFee.SdkIntTrim(), Denom: market.BankDenom}
		if maxFee.IsLT(buyerFeeCoin) {
			return nil, sdkerrors.ErrInvalidRequest.Wrapf(
				"%s: max fee: %s, required fee: %s",
				orderIndex, maxFee, buyerFeeCoin,
			)
		}

		// check address has the total cost
		buyerBalance := k.bankKeeper.GetBalance(sdkCtx, buyerAcc, order.BidPrice.Denom)
		if buyerBalance.IsLT(totalCost) {
			return nil, sdkerrors.ErrInsufficientFunds.Wrapf(
				"%s: quantity: %s, ask price: %s%s, total price: %v, bank balance: %v",
				orderIndex, order.Quantity, sellOrder.AskAmount, market.BankDenom, totalCost, buyerBalance,
			)
		}

		// fillOrder updates seller balance, buyer balance, batch supply, and transfers calculated
		// total cost from buyer account to seller account.
		if err = k.fillOrder(ctx, fillOrderParams{
			orderIndex:   orderIndex,
			sellOrder:    sellOrder,
			buyerAcc:     buyerAcc,
			buyQuantity:  buyQuantity,
			totalCost:    total,
			subTotalCost: subtotal,
			buyerFee:     buyerFee,
			autoRetire:   !order.DisableAutoRetire,
			batchDenom:   batch.Denom,
			bankDenom:    market.BankDenom,
			jurisdiction: order.RetirementJurisdiction,
			reason:       order.RetirementReason,
			feeParams:    feeParams,
		}); err != nil {
			return nil, err
		}
	}

	return &types.MsgBuyDirectResponse{}, nil
}
