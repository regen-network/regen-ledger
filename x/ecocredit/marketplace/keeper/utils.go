package keeper

import (
	"context"

	sdkmath "cosmossdk.io/math"
	gogoproto "github.com/gogo/protobuf/proto"
	protov2 "google.golang.org/protobuf/proto"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	api "github.com/regen-network/regen-ledger/api/v2/regen/ecocredit/marketplace/v1"
	"github.com/regen-network/regen-ledger/types/v2/math"
	basetypes "github.com/regen-network/regen-ledger/x/ecocredit/v3/base/types/v1"
	types "github.com/regen-network/regen-ledger/x/ecocredit/v3/marketplace/types/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/v3/server/utils"
)

// isDenomAllowed checks if the denom is allowed to be used in sell orders.
func isDenomAllowed(ctx context.Context, bankDenom string, table api.AllowedDenomTable) (bool, error) {
	return table.Has(ctx, bankDenom)
}

type fillOrderParams struct {
	orderIndex   string
	sellOrder    *api.SellOrder
	buyerAcc     sdk.AccAddress
	buyQuantity  math.Dec
	buyerFee     math.Dec
	subTotalCost math.Dec
	totalCost    math.Dec
	autoRetire   bool
	batchDenom   string
	bankDenom    string
	jurisdiction string
	reason       string
	feeParams    *api.FeeParams
}

// fillOrder updates seller balance, buyer balance, batch supply, and transfers calculated total cost
// from buyer account to seller account.
func (k Keeper) fillOrder(ctx context.Context, params fillOrderParams) error {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	// get sell order quantity to be checked and/or updated
	sellOrderQty, err := math.NewDecFromString(params.sellOrder.Quantity)
	if err != nil {
		return err
	}

	// If the sell order quantity is less than the purchase quantity, return an error.
	// If the sell order quantity is equal to the purchase quantity, remove the sell order.
	// If the sell order quantity is greater than the buy quantity, subtract the buy quantity
	// from the sell order quantity and update the sell order.
	switch sellOrderQty.Cmp(params.buyQuantity) {
	case math.LessThan:
		return sdkerrors.ErrInvalidRequest.Wrapf(
			"%s: requested quantity: %v, sell order quantity %s",
			params.orderIndex, params.buyQuantity, params.sellOrder.Quantity,
		)
	case math.EqualTo:
		if err := k.stateStore.SellOrderTable().Delete(ctx, params.sellOrder); err != nil {
			return err
		}
	case math.GreaterThan:
		newSellOrderQty, err := sellOrderQty.Sub(params.buyQuantity)
		if err != nil {
			return err
		}
		params.sellOrder.Quantity = newSellOrderQty.String()
		if err = k.stateStore.SellOrderTable().Update(ctx, params.sellOrder); err != nil {
			return err
		}
	}

	// calculate and set seller balance escrowed amount (subtract credits)
	sellerBal, err := k.baseStore.BatchBalanceTable().Get(ctx, params.sellOrder.Seller, params.sellOrder.BatchKey)
	if err != nil {
		return err
	}
	escrowBal, err := math.NewDecFromString(sellerBal.EscrowedAmount)
	if err != nil {
		return err
	}
	escrowBal, err = math.SafeSubBalance(escrowBal, params.buyQuantity)
	if err != nil {
		return err
	}
	sellerBal.EscrowedAmount = escrowBal.String()

	// update seller balance with new escrowed amount
	if err = k.baseStore.BatchBalanceTable().Update(ctx, sellerBal); err != nil {
		return err
	}

	// get buyer balance to be updated
	buyerBal, err := utils.GetBalance(ctx, k.baseStore.BatchBalanceTable(), params.buyerAcc, params.sellOrder.BatchKey)
	if err != nil {
		return err
	}

	// If auto-retire is disabled, we update buyer balance tradable amount. We emit a transfer event with
	// the credit quantity being purchased as the tradable amount. We do not update batch supply because
	// we do not distinguish between tradable and escrowed credits in batch supply.
	//
	// If auto-retire is enabled, we update buyer balance retired amount. We emit a transfer event with the
	// credit quantity being purchased as the retired amount and a retire event with the credit quantity as
	// the amount. We also update batch supply to reflect the retirement of the credits.
	if !params.autoRetire {

		// calculate and set buyer balance tradable amount (add credits)
		tradableBalance, err := math.NewDecFromString(buyerBal.TradableAmount)
		if err != nil {
			return err
		}
		tradableBalance, err = math.SafeAddBalance(tradableBalance, params.buyQuantity)
		if err != nil {
			return err
		}
		buyerBal.TradableAmount = tradableBalance.String()

		// emit transfer event with purchase quantity as tradable amount
		if err = sdkCtx.EventManager().EmitTypedEvent(&basetypes.EventTransfer{
			Sender:         sdk.AccAddress(params.sellOrder.Seller).String(),
			Recipient:      params.buyerAcc.String(),
			BatchDenom:     params.batchDenom,
			TradableAmount: params.buyQuantity.String(),
			RetiredAmount:  "0", // add zero to prevent empty string
		}); err != nil {
			return err
		}
	} else {

		// calculate and set buyer balance retired amount (add credits)
		retiredBalance, err := math.NewDecFromString(buyerBal.RetiredAmount)
		if err != nil {
			return err
		}
		retiredBalance, err = math.SafeAddBalance(retiredBalance, params.buyQuantity)
		if err != nil {
			return err
		}
		buyerBal.RetiredAmount = retiredBalance.String()

		// get batch supply to be updated
		supply, err := k.baseStore.BatchSupplyTable().Get(ctx, params.sellOrder.BatchKey)
		if err != nil {
			return err
		}

		// calculate and set batch supply tradable amount (subtract credits)
		supplyTradable, err := math.NewDecFromString(supply.TradableAmount)
		if err != nil {
			return err
		}
		supplyTradable, err = math.SafeSubBalance(supplyTradable, params.buyQuantity)
		if err != nil {
			return err
		}
		supply.TradableAmount = supplyTradable.String()

		// calculate and set batch supply retired amount (add credits)
		supplyRetired, err := math.NewDecFromString(supply.RetiredAmount)
		if err != nil {
			return err
		}
		supplyRetired, err = math.SafeAddBalance(supplyRetired, params.buyQuantity)
		if err != nil {
			return err
		}
		supply.RetiredAmount = supplyRetired.String()

		// update batch supply with new tradable and retired amount
		if err = k.baseStore.BatchSupplyTable().Update(ctx, supply); err != nil {
			return err
		}

		// emit transfer event with purchase quantity as retired amount
		if err = sdkCtx.EventManager().EmitTypedEvent(&basetypes.EventTransfer{
			Sender:         sdk.AccAddress(params.sellOrder.Seller).String(),
			Recipient:      params.buyerAcc.String(),
			BatchDenom:     params.batchDenom,
			TradableAmount: "0", // add zero to prevent empty string
			RetiredAmount:  params.buyQuantity.String(),
		}); err != nil {
			return err
		}

		// emit retire event with purchase quantity as amount
		if err = sdkCtx.EventManager().EmitTypedEvent(&basetypes.EventRetire{
			Owner:        params.buyerAcc.String(),
			BatchDenom:   params.batchDenom,
			Amount:       params.buyQuantity.String(),
			Jurisdiction: params.jurisdiction,
			Reason:       params.reason,
		}); err != nil {
			return err
		}
	}

	// update buyer credit balance with new tradable or retired amount
	if err = k.baseStore.BatchBalanceTable().Save(ctx, buyerBal); err != nil {
		return err
	}

	// calculate seller fee = subtotal * seller percentage fee
	sellerFee, err := getSellerFee(params.subTotalCost, params.feeParams)
	if err != nil {
		return err
	}

	// calculate total fee = buyer fee + seller fee
	totalFee, err := params.buyerFee.Add(sellerFee)
	if err != nil {
		return err
	}

	// if total fee > 0, then transfer total fee from buyer account to fee pool
	if totalFee.IsPositive() {
		feeCoins := sdk.NewCoins(sdk.NewCoin(params.bankDenom, totalFee.SdkIntTrim()))

		err = k.bankKeeper.SendCoinsFromAccountToModule(
			sdkCtx,
			params.buyerAcc,
			k.feePoolName,
			feeCoins)
		if err != nil {
			return err
		}

		// if bank denom == uregen, then burn the total fee from the fee pool
		if params.bankDenom == "uregen" {
			err = k.bankKeeper.BurnCoins(sdkCtx, k.feePoolName, feeCoins)
			if err != nil {
				return err
			}
		}
	}

	// calculate seller payment = subtotal - seller fee
	sellerPayment, err := params.subTotalCost.Sub(sellerFee)
	if err != nil {
		return err
	}

	seller := sdk.AccAddress(params.sellOrder.Seller)

	// send seller payment from buyer account to seller account
	err = k.bankKeeper.SendCoins(sdkCtx, params.buyerAcc, seller, sdk.NewCoins(sdk.NewCoin(
		params.bankDenom,
		sellerPayment.SdkIntTrim(),
	)))
	if err != nil {
		return err
	}

	return sdkCtx.EventManager().EmitTypedEvent(&types.EventBuyDirect{
		SellOrderId:   params.sellOrder.Id,
		Buyer:         params.buyerAcc.String(),
		BuyerFeePaid:  &sdk.Coin{Denom: params.bankDenom, Amount: params.buyerFee.SdkIntTrim()},
		Seller:        seller.String(),
		SellerFeePaid: &sdk.Coin{Denom: params.bankDenom, Amount: sellerFee.SdkIntTrim()},
	})
}

// getSubTotalCost calculates the total cost of the buy order by multiplying the price per credit specified
// in the sell order (i.e. the ask amount) and the quantity of credits specified in the buy order.
func getSubTotalCost(askAmount sdkmath.Int, buyQuantity math.Dec) (math.Dec, error) {
	unitPrice, err := math.NewPositiveFixedDecFromString(askAmount.String(), buyQuantity.NumDecimalPlaces())
	if err != nil {
		return math.Dec{}, err
	}
	subtotal, err := buyQuantity.Mul(unitPrice)
	if err != nil {
		return math.Dec{}, err
	}
	return subtotal, nil
}

// getTotalCostAndBuyerFee calculates the total cost of the buy order by multiplying the subtotal by the buyer percentage fee.
func getTotalCostAndBuyerFee(subtotal math.Dec, feeParams *api.FeeParams) (total math.Dec, buyerFee math.Dec, err error) {
	buyerPercentageFee := math.NewDecFromInt64(0)
	if feeParams != nil && feeParams.BuyerPercentageFee != "" {
		buyerPercentageFee, err = math.NewPositiveDecFromString(feeParams.BuyerPercentageFee)
		if err != nil {
			return
		}
	}

	buyerFee, err = subtotal.Mul(buyerPercentageFee)
	if err != nil {
		return
	}

	total, err = subtotal.Add(buyerFee)
	return
}

func getSellerFee(subtotal math.Dec, feeParams *api.FeeParams) (math.Dec, error) {
	sellerPercentageFee := math.NewDecFromInt64(0)
	if feeParams != nil && feeParams.SellerPercentageFee != "" {
		var err error
		sellerPercentageFee, err = math.NewPositiveDecFromString(feeParams.SellerPercentageFee)
		if err != nil {
			return math.Dec{}, err
		}
	}

	return subtotal.Mul(sellerPercentageFee)
}

func gogoToProtoReflect(from gogoproto.Message, to protov2.Message) error {
	bz, err := gogoproto.Marshal(from)
	if err != nil {
		return err
	}

	return protov2.Unmarshal(bz, to)
}
