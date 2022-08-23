package marketplace

import (
	"context"

	sdkmath "cosmossdk.io/math"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/marketplace/v1"
	"github.com/regen-network/regen-ledger/types/math"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
	"github.com/regen-network/regen-ledger/x/ecocredit/server/utils"
)

// isDenomAllowed checks if the denom is allowed to be used in orders.
func isDenomAllowed(ctx context.Context, bankDenom string, table api.AllowedDenomTable) (bool, error) {
	return table.Has(ctx, bankDenom)
}

type orderOptions struct {
	autoRetire   bool
	batchDenom   string
	jurisdiction string
}

// fillOrder moves credits and coins according to the order. It will:
// - update a sell order, removing it if quantity becomes 0 as a result of this purchase.
// - remove the purchaseQty from the seller's escrowed balance.
// - add credits to the buyer's tradable/retired address (based on the DisableAutoRetire field).
// - update the supply accordingly.
// - send the coins specified in the bid to the seller.
func (k Keeper) fillOrder(ctx context.Context, orderIndex string, sellOrder *api.SellOrder, buyerAcc sdk.AccAddress, purchaseQty math.Dec,
	cost sdk.Coin, opts orderOptions) error {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	sellOrderQty, err := math.NewDecFromString(sellOrder.Quantity)
	if err != nil {
		return err
	}

	switch sellOrderQty.Cmp(purchaseQty) {
	case math.LessThan:
		return sdkerrors.ErrInvalidRequest.Wrapf(
			"%s: requested quantity: %v, sell order quantity %s",
			orderIndex, purchaseQty, sellOrder.Quantity,
		)
	case math.EqualTo:
		if err := k.stateStore.SellOrderTable().Delete(ctx, sellOrder); err != nil {
			return err
		}
	case math.GreaterThan:
		newSellOrderQty, err := sellOrderQty.Sub(purchaseQty)
		if err != nil {
			return err
		}
		sellOrder.Quantity = newSellOrderQty.String()
		if err = k.stateStore.SellOrderTable().Update(ctx, sellOrder); err != nil {
			return err
		}
	}

	// remove the credits from the seller's escrowed balance
	sellerBal, err := k.coreStore.BatchBalanceTable().Get(ctx, sellOrder.Seller, sellOrder.BatchKey)
	if err != nil {
		return err
	}
	escrowBal, err := math.NewDecFromString(sellerBal.EscrowedAmount)
	if err != nil {
		return err
	}
	escrowBal, err = math.SafeSubBalance(escrowBal, purchaseQty)
	if err != nil {
		return err
	}
	sellerBal.EscrowedAmount = escrowBal.String()
	if err = k.coreStore.BatchBalanceTable().Update(ctx, sellerBal); err != nil {
		return err
	}

	// update the buyers balance and the batch supply
	supply, err := k.coreStore.BatchSupplyTable().Get(ctx, sellOrder.BatchKey)
	if err != nil {
		return err
	}
	buyerBal, err := utils.GetBalance(ctx, k.coreStore.BatchBalanceTable(), buyerAcc, sellOrder.BatchKey)
	if err != nil {
		return err
	}
	// if auto retire is disabled, we move the credits into the buyer's tradable balance.
	// supply is not updated because supply does not distinguish between tradable and escrowed credits.
	if !opts.autoRetire {
		tradableBalance, err := math.NewDecFromString(buyerBal.TradableAmount)
		if err != nil {
			return err
		}
		tradableBalance, err = math.SafeAddBalance(tradableBalance, purchaseQty)
		if err != nil {
			return err
		}
		buyerBal.TradableAmount = tradableBalance.String()
		if err = sdkCtx.EventManager().EmitTypedEvent(&core.EventTransfer{
			Sender:         sdk.AccAddress(sellOrder.Seller).String(),
			Recipient:      buyerAcc.String(),
			BatchDenom:     opts.batchDenom,
			TradableAmount: purchaseQty.String(),
			RetiredAmount:  "0",
		}); err != nil {
			return err
		}
	} else {
		retiredBalance, err := math.NewDecFromString(buyerBal.RetiredAmount)
		if err != nil {
			return err
		}
		retiredBalance, err = math.SafeAddBalance(retiredBalance, purchaseQty)
		if err != nil {
			return err
		}
		buyerBal.RetiredAmount = retiredBalance.String()

		supplyTradable, err := math.NewDecFromString(supply.TradableAmount)
		if err != nil {
			return err
		}
		supplyTradable, err = math.SafeSubBalance(supplyTradable, purchaseQty)
		if err != nil {
			return err
		}
		supply.TradableAmount = supplyTradable.String()

		supplyRetired, err := math.NewDecFromString(supply.RetiredAmount)
		if err != nil {
			return err
		}
		supplyRetired, err = math.SafeAddBalance(supplyRetired, purchaseQty)
		if err != nil {
			return err
		}
		supply.RetiredAmount = supplyRetired.String()
		if err = k.coreStore.BatchSupplyTable().Update(ctx, supply); err != nil {
			return err
		}
		if err = sdkCtx.EventManager().EmitTypedEvent(&core.EventRetire{
			Owner:        buyerAcc.String(),
			BatchDenom:   opts.batchDenom,
			Amount:       purchaseQty.String(),
			Jurisdiction: opts.jurisdiction,
		}); err != nil {
			return err
		}
	}
	if err = k.coreStore.BatchBalanceTable().Save(ctx, buyerBal); err != nil {
		return err
	}

	return k.bankKeeper.SendCoins(sdkCtx, buyerAcc, sellOrder.Seller, sdk.NewCoins(cost))
}

// getTotalCost calculates the cost of the order by multiplying the price per credit, and the amount of credits
// desired in the order.
func getTotalCost(pricePerCredit sdkmath.Int, amtCredits math.Dec) (sdkmath.Int, error) {
	unitPrice, err := math.NewPositiveFixedDecFromString(pricePerCredit.String(), amtCredits.NumDecimalPlaces())
	if err != nil {
		return sdkmath.Int{}, err
	}
	cost, err := amtCredits.Mul(unitPrice)
	if err != nil {
		return sdkmath.Int{}, err
	}
	return cost.SdkIntTrim(), nil
}
