package marketplace

import (
	"context"
	"fmt"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/marketplace/v1"
	ecocreditv1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/types/math"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
	v1 "github.com/regen-network/regen-ledger/x/ecocredit/marketplace"
	"github.com/regen-network/regen-ledger/x/ecocredit/server"

	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// Buy allows users to purchase credits by either directly specifying a sell order, or
// defining a set of filters with attributes to match against.
//
// Currently, only the former is supported. Calls to this function with anything other than
// MsgBuy_Order_Selection_SellOrderId will fail.
func (k Keeper) Buy(ctx context.Context, req *v1.MsgBuy) (*v1.MsgBuyResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	buyerAcc, err := sdk.AccAddressFromBech32(req.Buyer)
	if err != nil {
		return nil, err
	}

	for _, order := range req.Orders {
		// assert they have the balance they're  bidding with
		bal := k.bankKeeper.GetBalance(sdkCtx, buyerAcc, order.BidPrice.Denom)
		if bal.IsLT(*order.BidPrice) {
			return nil, sdkerrors.ErrInsufficientFunds.Wrapf("cannot bid %v with a balance of %v", bal, *order.BidPrice)
		}

		switch selection := order.Selection.Sum.(type) {
		case *v1.MsgBuy_Order_Selection_SellOrderId:
			sellOrder, err := k.stateStore.SellOrderTable().Get(ctx, selection.SellOrderId)
			if err != nil {
				return nil, fmt.Errorf("sell order %d: %w", selection.SellOrderId, err)
			}
			batch, err := k.coreStore.BatchInfoTable().Get(ctx, sellOrder.BatchId)
			if err != nil {

				return nil, sdkerrors.ErrIO.Wrapf("error getting batch id %d: %s", sellOrder.BatchId, err.Error())
			}
			ct, err := server.GetCreditTypeFromBatchDenom(ctx, k.coreStore, k.paramsKeeper, batch.BatchDenom)
			if err != nil {
				return nil, err
			}
			creditOrderQty, err := math.NewPositiveFixedDecFromString(order.Quantity, ct.Precision)
			if err != nil {
				return nil, err
			}

			if sellOrder.DisableAutoRetire != order.DisableAutoRetire {
				return nil, sdkerrors.ErrInvalidRequest.Wrapf("auto-retire mismatch: sell order set to %t, buy "+
					"order set to %t", sellOrder.DisableAutoRetire, order.DisableAutoRetire)
			}

			// check that bid price and ask price denoms match
			market, err := k.stateStore.MarketTable().Get(ctx, sellOrder.MarketId)
			if err != nil {
				return nil, fmt.Errorf("market id %d: %w", sellOrder.MarketId, err)
			}
			if order.BidPrice.Denom != market.BankDenom {
				return nil, sdkerrors.ErrInvalidRequest.Wrapf("bid price denom does not match ask price denom: "+
					" %s, expected %s", order.BidPrice.Denom, market.BankDenom)
			}

			// check that bid price is at least equal to ask price
			askAmount, ok := sdk.NewIntFromString(sellOrder.AskPrice)
			if !ok {
				return nil, fmt.Errorf("could not convert sell order's ask price to %T: %s", sdk.Int{}, sellOrder.AskPrice)
			}
			if order.BidPrice.Amount.LT(askAmount) {
				return nil, sdkerrors.ErrInsufficientFunds.Wrapf("bid price too low: got %s, needed at least %s",
					order.BidPrice.Amount.String(), sellOrder.AskPrice)
			}

			if err = k.updateBalances(ctx, sellOrder, buyerAcc, creditOrderQty, *order.BidPrice, false); err != nil {
				return nil, fmt.Errorf("error updating balances: %w", err)
			}
			if !sellOrder.DisableAutoRetire {
				if err = sdkCtx.EventManager().EmitTypedEvent(&core.EventRetire{
					Retirer:    buyerAcc.String(),
					BatchDenom: batch.BatchDenom,
					Amount:     order.Quantity,
					Location:   order.RetirementLocation,
				}); err != nil {
					return nil, err
				}
			} else {
				if err = sdkCtx.EventManager().EmitTypedEvent(&core.EventReceive{
					Sender:         sdk.AccAddress(sellOrder.Seller).String(),
					Recipient:      buyerAcc.String(),
					BatchDenom:     batch.BatchDenom,
					TradableAmount: order.Quantity,
					RetiredAmount:  "",
					BasketDenom:    "",
				}); err != nil {
					return nil, err
				}
			}
		case *v1.MsgBuy_Order_Selection_Filter:
			return nil, sdkerrors.ErrInvalidRequest.Wrap("only direct buy orders are enabled at this time")
			// verify expiration is in the future
			//if order.Expiration.Before(sdkCtx.BlockTime()) {
			//	return nil, sdkerrors.ErrInvalidRequest.Wrapf("expiration must be in the future: %s", order.Expiration)
			//}
		default:
			return nil, sdkerrors.ErrInvalidRequest.Wrap("only direct buy orders are enabled at this time")
		}
	}
	return &v1.MsgBuyResponse{}, nil
}

// updateBalances moves credits according to the order. it will:
// - update a sell order, removing it if quantity becomes 0 as a result of this purchase.
// - remove the purchaseQty from the seller's escrowed balance.
// - add credits to the buyer's tradable/retired address (based on the DisableAutoRetire field).
// - update the supply accordingly.
// - send the coins specified in the bid to the seller.
func (k Keeper) updateBalances(ctx context.Context, sellOrder *api.SellOrder, buyerAcc sdk.AccAddress, purchaseQty math.Dec, bidCoin sdk.Coin, canPartialFill bool) error {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	sellOrderQty, err := math.NewDecFromString(sellOrder.Quantity)
	if err != nil {
		return err
	}
	// since we only support direct buy orders at this time, we fail when the requested purchase amount is more than
	// the available credits for sale, rather than partial filling the buy order.
	newSellOrderQty, err := sellOrderQty.Sub(purchaseQty)
	if err != nil {
		return err
	}
	if newSellOrderQty.IsNegative() {
		if !canPartialFill {
			return ecocredit.ErrInsufficientCredits.Wrapf("cannot purchase %v credits from a sell order that has %s credits", purchaseQty, sellOrder.Quantity)
		} else {
			// if we can partial fill, we just delete the sellOrder and take whatever
			// credits are left from that order.
			if err := k.stateStore.SellOrderTable().Delete(ctx, sellOrder); err != nil {
				return err
			}
			purchaseQty = sellOrderQty
		}
	}
	if newSellOrderQty.IsZero() { // remove the sell order if no credits are left
		if err := k.stateStore.SellOrderTable().Delete(ctx, sellOrder); err != nil {
			return err
		}
	} else { // update the sell order with the new value otherwise
		sellOrder.Quantity = newSellOrderQty.String()
		if err = k.stateStore.SellOrderTable().Update(ctx, sellOrder); err != nil {
			return err
		}
	}

	// remove the credits from the seller's escrowed balance
	sellerBal, err := k.coreStore.BatchBalanceTable().Get(ctx, sellOrder.Seller, sellOrder.BatchId)
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
	if err = k.coreStore.BatchBalanceTable().Update(ctx, sellerBal); err != nil {
		return err
	}

	// update the buyers balance and the batch supply
	supply, err := k.coreStore.BatchSupplyTable().Get(ctx, sellOrder.BatchId)
	if err != nil {
		return err
	}
	buyerBal, err := k.coreStore.BatchBalanceTable().Get(ctx, buyerAcc, sellOrder.BatchId)
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
	if sellOrder.DisableAutoRetire { // if auto retire is disabled, we move the credits into the buyer's/supply's tradable balance.
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

	supplyEscrowed, err := math.NewDecFromString(supply.EscrowedAmount)
	if err != nil {
		return err
	}
	supplyEscrowed, err = math.SafeSubBalance(supplyEscrowed, purchaseQty)
	if err != nil {
		return err
	}
	supply.EscrowedAmount = supplyEscrowed.String()
	if err = k.coreStore.BatchSupplyTable().Update(ctx, supply); err != nil {
		return err
	}
	if err = k.coreStore.BatchBalanceTable().Save(ctx, buyerBal); err != nil {
		return err
	}

	return k.bankKeeper.SendCoins(sdkCtx, buyerAcc, sellOrder.Seller, sdk.NewCoins(bidCoin))
}
