package marketplace

import (
	"context"
	"fmt"
	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/marketplace/v1"
	ecocreditv1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/types/math"
	v1 "github.com/regen-network/regen-ledger/x/ecocredit/marketplace"
)

func (k Keeper) Buy(ctx context.Context, req *v1.MsgBuy) (*v1.MsgBuyResponse, error) {
	// setup
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	buyerAcc, err := sdk.AccAddressFromBech32(req.Buyer)
	if err != nil {
		return nil, err
	}

	for _, order := range req.Orders {
		// verify expiration is in the future
		if order.Expiration != nil && order.Expiration.Before(sdkCtx.BlockTime()) {
			return nil, sdkerrors.ErrInvalidRequest.Wrapf("expiration must be in the future: %s", order.Expiration)
		}

		// assert they have the balance they're  bidding with
		if !k.bankKeeper.HasBalance(sdkCtx, buyerAcc, *order.BidPrice) {
			return nil, sdkerrors.ErrInsufficientFunds
		}

		switch selection := order.Selection.Sum.(type) {
		case *v1.MsgBuy_Order_Selection_SellOrderId:
			orderAmt, err := math.NewDecFromString(order.Quantity)
			if err != nil {
				return nil, err
			}
			// get the sell order
			sellOrder, err := k.stateStore.SellOrderStore().Get(ctx, selection.SellOrderId)
			if err != nil {
				return nil, fmt.Errorf("sell order %d: %w", selection.SellOrderId, err)
			}
			if sellOrder.DisableAutoRetire != order.DisableAutoRetire {
				return nil, sdkerrors.ErrInvalidRequest.Wrapf("auto-retire mismatch: sell order set to %t, buy "+
					"order set to %t", sellOrder.DisableAutoRetire, order.DisableAutoRetire)
			}
			// get the market
			market, err := k.stateStore.MarketStore().Get(ctx, sellOrder.MarketId)
			if err != nil {
				return nil, fmt.Errorf("market id %d: %w", sellOrder.MarketId, err)
			}
			// check that the order denoms match
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
			if err = k.updateBalances(ctx, sellOrder, buyerAcc, orderAmt, *order.BidPrice); err != nil {
				return nil, fmt.Errorf("error updating balances: %w", err)
			}
		default:
			return nil, sdkerrors.ErrInvalidRequest.Wrap("only direct buy orders are enabled at this time")
		}
	}
	return &v1.MsgBuyResponse{}, nil
}

func (k Keeper) updateBalances(ctx context.Context, sellOrder *api.SellOrder, buyerAcc sdk.AccAddress, purchaseQty math.Dec, bidCoin sdk.Coin) error {
	// update the sell order
	sellOrderQty, err := math.NewDecFromString(sellOrder.Quantity)
	if err != nil {
		return err
	}
	sellOrderQty, err = math.SafeSubBalance(sellOrderQty, purchaseQty)
	if err != nil {
		return err
	}
	if sellOrderQty.IsZero() { // remove the sell order
		if err := k.stateStore.SellOrderStore().Delete(ctx, sellOrder); err != nil {
			return err
		}
	} else { // update the sell order
		sellOrder.Quantity = sellOrderQty.String()
		if err = k.stateStore.SellOrderStore().Update(ctx, sellOrder); err != nil {
			return err
		}
	}

	// update the sellers balance
	sellerBal, err := k.coreStore.BatchBalanceStore().Get(ctx, sellOrder.Seller, sellOrder.BatchId)
	if err != nil {
		return err
	}
	escrowBal, err := math.NewDecFromString(sellerBal.Escrowed)
	if err != nil {
		return err
	}
	escrowBal, err = math.SafeSubBalance(escrowBal, purchaseQty)
	if err != nil {
		return err
	}
	sellerBal.Escrowed = escrowBal.String()
	if err = k.coreStore.BatchBalanceStore().Update(ctx, sellerBal); err != nil {
		return err
	}

	// update the buyers balance and supply
	supply, err := k.coreStore.BatchSupplyStore().Get(ctx, sellOrder.BatchId)
	if err != nil {
		return err
	}
	buyerBal, err := k.coreStore.BatchBalanceStore().Get(ctx, buyerAcc, sellOrder.BatchId)
	if err != nil {
		if ormerrors.IsNotFound(err) {
			buyerBal = &ecocreditv1.BatchBalance{
				Address:  buyerAcc,
				BatchId:  sellOrder.BatchId,
				Tradable: "0",
				Retired:  "0",
				Escrowed: "0",
			}
		} else {
			return err
		}
	}
	if sellOrder.DisableAutoRetire {
		tradableBalance, err := math.NewDecFromString(buyerBal.Tradable)
		if err != nil {
			return err
		}
		tradableBalance, err = math.SafeAddBalance(tradableBalance, purchaseQty)
		if err != nil {
			return err
		}
		buyerBal.Tradable = tradableBalance.String()

		supplyTradable, err := math.NewDecFromString(supply.TradableAmount)
		if err != nil {
			return err
		}
		supplyTradable, err = math.SafeAddBalance(supplyTradable, purchaseQty)
		if err != nil {
			return err
		}
		supply.TradableAmount = supplyTradable.String()
	} else {
		retiredBalance, err := math.NewDecFromString(buyerBal.Retired)
		if err != nil {
			return err
		}
		retiredBalance, err = math.SafeAddBalance(retiredBalance, purchaseQty)
		if err != nil {
			return err
		}
		buyerBal.Retired = retiredBalance.String()

		supplyRetired, err := math.NewDecFromString(supply.RetiredAmount)
		if err != nil {
			return err
		}
		supplyRetired, err = math.SafeAddBalance(supplyRetired, purchaseQty)
		if err != nil {
			return err
		}
		supply.RetiredAmount = supplyRetired.String()
	}
	// we can update the escrowed supply outside the condition since its the same either way
	supplyEscrowed, err := math.NewDecFromString(supply.EscrowedAmount)
	supplyEscrowed, err = math.SafeSubBalance(supplyEscrowed, purchaseQty)
	if err != nil {
		return err
	}
	supply.EscrowedAmount = supplyEscrowed.String()
	if err = k.coreStore.BatchSupplyStore().Update(ctx, supply); err != nil {
		return err
	}
	if err = k.coreStore.BatchBalanceStore().Save(ctx, buyerBal); err != nil {
		return err
	}

	// send the coins to the seller
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	return k.bankKeeper.SendCoins(sdkCtx, buyerAcc, sellOrder.Seller, sdk.NewCoins(bidCoin))
}
