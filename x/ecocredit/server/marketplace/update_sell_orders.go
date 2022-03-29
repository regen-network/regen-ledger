package marketplace

import (
	"context"

	"github.com/regen-network/regen-ledger/x/ecocredit/server/utils"
	"google.golang.org/protobuf/types/known/timestamppb"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/marketplace/v1"
	"github.com/regen-network/regen-ledger/types/math"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/regen-network/regen-ledger/x/ecocredit/marketplace"
)

// UpdateSellOrders updates the sellOrder with the provided values.
// Note: only the DisableAutoRetire lacks field presence, so if the existing value
// is true and you do not want to change that, you MUST provide a value of true in the update.
// Otherwise, the sell order will be changed to false.
func (k Keeper) UpdateSellOrders(ctx context.Context, req *marketplace.MsgUpdateSellOrders) (*marketplace.MsgUpdateSellOrdersResponse, error) {
	seller, err := sdk.AccAddressFromBech32(req.Owner)
	if err != nil {
		return nil, err
	}

	for _, update := range req.Updates {
		sellOrder, err := k.stateStore.SellOrderTable().Get(ctx, update.SellOrderId)
		if err != nil {
			return nil, sdkerrors.ErrInvalidRequest.Wrapf("could not get sell order with id %d: %s", update.SellOrderId, err.Error())
		}
		sellOrderAddr := sdk.AccAddress(sellOrder.Seller)
		if !seller.Equals(sellOrderAddr) {
			return nil, sdkerrors.ErrUnauthorized.Wrapf("unable to update sell order: got: %s, want: %s", req.Owner, sellOrderAddr.String())
		}
		if err = k.applySellOrderUpdates(ctx, sellOrder, update); err != nil {
			return nil, err
		}
	}
	return &marketplace.MsgUpdateSellOrdersResponse{}, nil
}

// applySellOrderUpdates applies the updates to the order.
func (k Keeper) applySellOrderUpdates(ctx context.Context, order *api.SellOrder, update *marketplace.MsgUpdateSellOrders_Update) error {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	var creditType *ecocredit.CreditType
	event := marketplace.EventUpdateSellOrder{}

	order.DisableAutoRetire = update.DisableAutoRetire
	event.DisableAutoRetire = update.DisableAutoRetire

	if update.NewAskPrice != nil {
		market, err := k.stateStore.MarketTable().Get(ctx, order.MarketId)
		if err != nil {
			return err
		}
		// TODO: pending param refactor https://github.com/regen-network/regen-ledger/issues/624
		//has, err := isDenomAllowed(ctx, k.stateStore, update.AskPrice.Denom)
		//if err != nil {
		//	return nil, err
		//}
		//if !has {
		//	return nil, ecocredit.ErrInvalidRequest.Wrapf("cannot use coin with denom %s in sell orders", order.AskPrice.Denom)
		//}
		if market.BankDenom != update.NewAskPrice.Denom {
			creditType, err = k.getCreditTypeFromBatchId(ctx, order.BatchId)
			if err != nil {
				return err
			}
			marketId, err := k.getOrCreateMarketId(ctx, creditType.Abbreviation, update.NewAskPrice.Denom)
			if err != nil {
				return err
			}
			order.MarketId = marketId
		}
		order.AskPrice = update.NewAskPrice.Amount.String()
		event.NewAskPrice = update.NewAskPrice
	}
	if update.NewExpiration != nil {
		// verify expiration is in the future
		if update.NewExpiration.Before(sdkCtx.BlockTime()) {
			return sdkerrors.ErrInvalidRequest.Wrapf("expiration must be in the future: %s", update.NewExpiration)
		}
		order.Expiration = timestamppb.New(*update.NewExpiration)
		event.NewExpiration = update.NewExpiration
	}
	if update.NewQuantity != "" {
		if creditType == nil {
			var err error
			creditType, err = k.getCreditTypeFromBatchId(ctx, order.BatchId)
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
			if err = k.escrowCredits(ctx, order.Seller, order.BatchId, amtToEscrow); err != nil {
				return err
			}
			order.Quantity = update.NewQuantity
			event.NewQuantity = update.NewQuantity
		case math.LessThan:
			amtToUnescrow, err := existingQty.Sub(newQty)
			if err != nil {
				return err
			}
			if err = k.unescrowCredits(ctx, order.Seller, order.BatchId, amtToUnescrow.String()); err != nil {
				return err
			}
			order.Quantity = update.NewQuantity
			event.NewQuantity = update.NewQuantity
		}
	}

	if err := k.stateStore.SellOrderTable().Update(ctx, order); err != nil {
		return err
	}

	return sdkCtx.EventManager().EmitTypedEvent(&event)
}

// getCreditTypeFromBatchId gets the credit type given a batch id.
func (k Keeper) getCreditTypeFromBatchId(ctx context.Context, id uint64) (*ecocredit.CreditType, error) {
	batch, err := k.coreStore.BatchInfoTable().Get(ctx, id)
	if err != nil {
		return nil, err
	}
	creditType, err := utils.GetCreditTypeFromBatchDenom(ctx, k.coreStore, k.paramsKeeper, batch.BatchDenom)
	if err != nil {
		return nil, err
	}
	return &creditType, nil
}
