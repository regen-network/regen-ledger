package marketplace

import (
	"context"
	"fmt"

	"google.golang.org/protobuf/types/known/timestamppb"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/marketplace/v1"
	ecoApi "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/types/math"
	"github.com/regen-network/regen-ledger/x/ecocredit/marketplace"
	"github.com/regen-network/regen-ledger/x/ecocredit/server/utils"
)

// UpdateSellOrders updates the sellOrder with the provided values.
// Note: only the DisableAutoRetire lacks field presence, so if the existing value
// is true, and you do not want to change that, you MUST provide a value of true in the update.
// Otherwise, the sell order will be changed to false.
func (k Keeper) UpdateSellOrders(ctx context.Context, req *marketplace.MsgUpdateSellOrders) (*marketplace.MsgUpdateSellOrdersResponse, error) {
	seller, err := sdk.AccAddressFromBech32(req.Owner)
	if err != nil {
		return nil, err
	}

	for i, update := range req.Updates {
		// orderIndex is passed to helper functions for more granular error messages
		// when an individual order in a list of orders fails to process
		orderIndex := fmt.Sprintf("order[%d]", i)

		sellOrder, err := k.stateStore.SellOrderTable().Get(ctx, update.SellOrderId)
		if err != nil {
			return nil, sdkerrors.ErrInvalidRequest.Wrapf("could not get sell order with id %d: %s", update.SellOrderId, err.Error())
		}
		sellOrderAddr := sdk.AccAddress(sellOrder.Seller)
		if !seller.Equals(sellOrderAddr) {
			return nil, sdkerrors.ErrUnauthorized.Wrapf("unable to update sell order: got: %s, want: %s", req.Owner, sellOrderAddr.String())
		}
		if err = k.applySellOrderUpdates(ctx, orderIndex, sellOrder, update); err != nil {
			return nil, err
		}
	}
	return &marketplace.MsgUpdateSellOrdersResponse{}, nil
}

// applySellOrderUpdates applies the updates to the order.
func (k Keeper) applySellOrderUpdates(ctx context.Context, orderIndex string, order *api.SellOrder, update *marketplace.MsgUpdateSellOrders_Update) error {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	var creditType *ecoApi.CreditType
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
			return sdkerrors.ErrInvalidRequest.Wrapf("%s is not allowed to be used in sell orders", update.NewAskPrice.Denom)
		}

		if market.BankDenom != update.NewAskPrice.Denom {
			creditType, err = k.getCreditTypeFromBatchKey(ctx, order.BatchKey)
			if err != nil {
				return err
			}
			marketId, err := k.getOrCreateMarketId(ctx, creditType.Abbreviation, update.NewAskPrice.Denom)
			if err != nil {
				return err
			}
			order.MarketId = marketId
		}
		order.AskAmount = update.NewAskPrice.Amount.String()
	}
	if update.NewExpiration != nil {
		// verify expiration is in the future
		if update.NewExpiration.Before(sdkCtx.BlockTime()) {
			return sdkerrors.ErrInvalidRequest.Wrapf("expiration must be in the future: %s", update.NewExpiration)
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
			if err = k.escrowCredits(ctx, orderIndex, order.Seller, order.BatchKey, amtToEscrow); err != nil {
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

	if err := k.stateStore.SellOrderTable().Update(ctx, order); err != nil {
		return err
	}

	return sdkCtx.EventManager().EmitTypedEvent(&marketplace.EventUpdateSellOrder{
		SellOrderId: order.Id,
	})
}

// getCreditTypeFromBatchKey gets the credit type given a batch key.
func (k Keeper) getCreditTypeFromBatchKey(ctx context.Context, key uint64) (*ecoApi.CreditType, error) {
	batch, err := k.coreStore.BatchTable().Get(ctx, key)
	if err != nil {
		return nil, err
	}
	creditType, err := utils.GetCreditTypeFromBatchDenom(ctx, k.coreStore, batch.Denom)
	if err != nil {
		return nil, err
	}
	return creditType, nil
}
