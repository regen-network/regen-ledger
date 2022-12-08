package keeper

import (
	"context"
	"fmt"

	"google.golang.org/protobuf/types/known/timestamppb"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	api "github.com/regen-network/regen-ledger/api/v2/regen/ecocredit/marketplace/v1"
	baseapi "github.com/regen-network/regen-ledger/api/v2/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/types/v2/math"
	types "github.com/regen-network/regen-ledger/x/ecocredit/v3/marketplace/types/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/v3/server/utils"
)

// UpdateSellOrders updates the sellOrder with the provided values.
// Note: only the DisableAutoRetire lacks field presence, so if the existing value
// is true, and you do not want to change that, you MUST provide a value of true in the update.
// Otherwise, the sell order will be changed to false.
func (k Keeper) UpdateSellOrders(ctx context.Context, req *types.MsgUpdateSellOrders) (*types.MsgUpdateSellOrdersResponse, error) {
	seller, err := sdk.AccAddressFromBech32(req.Seller)
	if err != nil {
		return nil, err
	}

	for i, update := range req.Updates {
		// updateIndex is used for more granular error messages when
		// an individual update in a list of updates fails to process
		updateIndex := fmt.Sprintf("updates[%d]", i)

		sellOrder, err := k.stateStore.SellOrderTable().Get(ctx, update.SellOrderId)
		if err != nil {
			return nil, sdkerrors.ErrInvalidRequest.Wrapf("%s: sell order with id %d: %s", updateIndex, update.SellOrderId, err)
		}

		sellOrderAddr := sdk.AccAddress(sellOrder.Seller)
		if !seller.Equals(sellOrderAddr) {
			return nil, sdkerrors.ErrUnauthorized.Wrapf("%s: seller must be the seller of the sell order", updateIndex)
		}

		// apply the updates to the sell order
		if err = k.applySellOrderUpdates(ctx, updateIndex, sellOrder, update); err != nil {
			return nil, err
		}
	}

	return &types.MsgUpdateSellOrdersResponse{}, nil
}

// applySellOrderUpdates applies the updates to the order.
func (k Keeper) applySellOrderUpdates(ctx context.Context, updateIndex string, order *api.SellOrder, update *types.MsgUpdateSellOrders_Update) error {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	order.Maker = true // maker is always true for sell orders

	// set order disable auto-retire based on update, note that if the update does
	// include disable auto-retire, disable auto-retire will be updated to false
	order.DisableAutoRetire = update.DisableAutoRetire

	// get credit type from batch key to get/create market and check precision
	creditType, err := k.getCreditTypeFromBatchKey(ctx, order.BatchKey)
	if err != nil {
		return err
	}

	if update.NewAskPrice != nil {
		// get market to check if new ask price denom is an allowed denom
		market, err := k.stateStore.MarketTable().Get(ctx, order.MarketId)
		if err != nil {
			return err
		}

		// check if new ask price denom is an allowed denom
		allowed, err := isDenomAllowed(ctx, update.NewAskPrice.Denom, k.stateStore.AllowedDenomTable())
		if err != nil {
			return err
		}
		if !allowed {
			return sdkerrors.ErrInvalidRequest.Wrapf(
				"%s: %s is not allowed to be used in sell orders",
				updateIndex, update.NewAskPrice.Denom,
			)
		}

		if market.BankDenom != update.NewAskPrice.Denom {
			marketID, err := k.getOrCreateMarketID(ctx, creditType.Abbreviation, update.NewAskPrice.Denom)
			if err != nil {
				return err
			}
			order.MarketId = marketID
		}

		// set order ask amount to new ask price amount
		order.AskAmount = update.NewAskPrice.Amount.String()
	}

	if update.NewExpiration != nil {
		// force to UTC
		expirationUTC := update.NewExpiration.UTC()
		update.NewExpiration = &expirationUTC

		// verify expiration is in the future
		if !update.NewExpiration.After(sdkCtx.BlockTime()) {
			return sdkerrors.ErrInvalidRequest.Wrapf(
				"%s: expiration must be in the future: %s",
				updateIndex, update.NewExpiration,
			)
		}

		// set order expiration to new expiration
		order.Expiration = timestamppb.New(*update.NewExpiration)
	}

	if update.NewQuantity != "" {

		// get decimal of new quantity
		newQuantity, err := math.NewPositiveFixedDecFromString(update.NewQuantity, creditType.Precision)
		if err != nil {
			return err
		}

		// get decimal of current quantity
		currentQuantity, err := math.NewDecFromString(order.Quantity)
		if err != nil {
			return err
		}

		// compare newQuantity and currentQuantity
		// if newQuantity > currentQuantity, we need to increase escrowed amount by the difference.
		// if newQuantity < currentQuantity, we need to decrease escrowed amount by the difference.
		switch newQuantity.Cmp(currentQuantity) {
		case math.GreaterThan:
			// calculate quantity of credits to escrow
			escrowQuantity, err := newQuantity.Sub(currentQuantity)
			if err != nil {
				return err
			}

			// convert seller balance tradable credits to escrowed credits
			if err = k.escrowCredits(ctx, updateIndex, order.Seller, order.BatchKey, escrowQuantity); err != nil {
				return err
			}

			// set order quantity to new quantity
			order.Quantity = update.NewQuantity
		case math.LessThan:
			unescrowQuantity, err := currentQuantity.Sub(newQuantity)
			if err != nil {
				return err
			}

			// convert seller balance escrowed credits to tradable credits
			if err = k.unescrowCredits(ctx, order.Seller, order.BatchKey, unescrowQuantity.String()); err != nil {
				return err
			}

			// set order quantity to new quantity
			order.Quantity = update.NewQuantity
		}
	}

	// update sell order with new sell order properties
	if err := k.stateStore.SellOrderTable().Update(ctx, order); err != nil {
		return err
	}

	return sdkCtx.EventManager().EmitTypedEvent(&types.EventUpdateSellOrder{
		SellOrderId: order.Id,
	})
}

// getCreditTypeFromBatchKey gets the credit type given a batch key.
func (k Keeper) getCreditTypeFromBatchKey(ctx context.Context, key uint64) (*baseapi.CreditType, error) {
	batch, err := k.baseStore.BatchTable().Get(ctx, key)
	if err != nil {
		return nil, err
	}
	creditType, err := utils.GetCreditTypeFromBatchDenom(ctx, k.baseStore, batch.Denom)
	if err != nil {
		return nil, err
	}
	return creditType, nil
}
