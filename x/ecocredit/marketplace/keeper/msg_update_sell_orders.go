package keeper

import (
	"context"
	"fmt"

	"google.golang.org/protobuf/types/known/timestamppb"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/marketplace/v1"
	baseapi "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/types/math"
	types "github.com/regen-network/regen-ledger/x/ecocredit/marketplace/types/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/server/utils"
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
			return nil, sdkerrors.ErrInvalidRequest.Wrapf("%s: sell order with id %d: %s", updateIndex, update.SellOrderId, err.Error())
		}

		sellOrderAddr := sdk.AccAddress(sellOrder.Seller)
		if !seller.Equals(sellOrderAddr) {
			return nil, sdkerrors.ErrUnauthorized.Wrapf("%s: seller must be the seller of the sell order", updateIndex)
		}

		if err = k.applySellOrderUpdates(ctx, updateIndex, sellOrder, update); err != nil {
			return nil, err
		}
	}
	return &types.MsgUpdateSellOrdersResponse{}, nil
}

// applySellOrderUpdates applies the updates to the order.
func (k Keeper) applySellOrderUpdates(ctx context.Context, updateIndex string, order *api.SellOrder, update *types.MsgUpdateSellOrders_Update) error {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	var creditType *baseapi.CreditType

	order.DisableAutoRetire = update.DisableAutoRetire

	if update.NewAskPrice != nil {
		market, err := k.stateStore.MarketTable().Get(ctx, order.MarketId)
		if err != nil {
			return err
		}

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
			creditType, err = k.getCreditTypeFromBatchKey(ctx, order.BatchKey)
			if err != nil {
				return err
			}
			marketID, err := k.getOrCreateMarketID(ctx, creditType.Abbreviation, update.NewAskPrice.Denom)
			if err != nil {
				return err
			}
			order.MarketId = marketID
		}

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
		order.Expiration = timestamppb.New(*update.NewExpiration)
	}

	if update.NewQuantity != "" {
		if creditType == nil {
			var err error
			creditType, err = k.getCreditTypeFromBatchKey(ctx, order.BatchKey)
			if err != nil {
				return err
			}
		}
		newQty, err := math.NewPositiveFixedDecFromString(update.NewQuantity, creditType.Precision)
		if err != nil {
			return err
		}
		existingQty, err := math.NewDecFromString(order.Quantity)
		if err != nil {
			return err
		}
		// compare newQty and the existingQty
		// if newQty > existingQty, we need to increase our amount escrowed by the difference of new - existing.
		// if newQty < existingQty we need to decrease our amount escrowed by the difference of existing - new.
		switch newQty.Cmp(existingQty) {
		case math.GreaterThan:
			amtToEscrow, err := newQty.Sub(existingQty)
			if err != nil {
				return err
			}
			if err = k.escrowCredits(ctx, updateIndex, order.Seller, order.BatchKey, amtToEscrow); err != nil {
				return err
			}
			order.Quantity = update.NewQuantity
		case math.LessThan:
			amtToUnescrow, err := existingQty.Sub(newQty)
			if err != nil {
				return err
			}
			if err = k.unescrowCredits(ctx, order.Seller, order.BatchKey, amtToUnescrow.String()); err != nil {
				return err
			}
			order.Quantity = update.NewQuantity
		}
	}

	order.Maker = true // maker is always true for sell orders

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
