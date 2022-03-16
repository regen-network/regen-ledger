package marketplace

import (
	"context"

	"google.golang.org/protobuf/types/known/timestamppb"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/marketplace/v1"
	v1 "github.com/regen-network/regen-ledger/x/ecocredit/marketplace"
	"github.com/regen-network/regen-ledger/x/ecocredit/server"
)

func (k Keeper) UpdateSellOrders(ctx context.Context, req *v1.MsgUpdateSellOrders) (*v1.MsgUpdateSellOrdersResponse, error) {
	seller, err := sdk.AccAddressFromBech32(req.Owner)
	if err != nil {
		return nil, err
	}

	for _, update := range req.Updates {
		// verify expiration is in the future
		if update.NewExpiration != nil && update.NewExpiration.Before(ctx.BlockTime()) {
			return nil, sdkerrors.ErrInvalidRequest.Wrapf("expiration must be in the future: %s", update.NewExpiration)
		}
		sellOrder, err := k.stateStore.SellOrderTable().Get(ctx, update.SellOrderId)
		if err != nil {
			return nil, err
		}
		if err = k.applyOrderUpdates(ctx, sellOrder, update); err != nil {
			return nil, err
		}
	}
	return &v1.MsgUpdateSellOrdersResponse{}, nil
}

func (k Keeper) applyOrderUpdates(ctx context.Context, order *api.SellOrder, update *v1.MsgUpdateSellOrders_Update) error {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	event := v1.EventUpdateSellOrder{}
	order.DisableAutoRetire = update.DisableAutoRetire
	event.DisableAutoRetire = update.DisableAutoRetire

	if update.NewAskPrice != nil {
		// TODO: this should probably be updated to check that the denom is allowed
		// once https://github.com/regen-network/regen-ledger/issues/624 is resolved.
		market, err := k.stateStore.MarketTable().Get(ctx, order.MarketId)
		if err != nil {
			return err
		}
		if market.BankDenom != update.NewAskPrice.Denom {
			batch, err := k.coreStore.BatchInfoTable().Get(ctx, order.BatchId)
			if err != nil {
				return err
			}
			ct, err := server.GetCreditTypeFromBatchDenom(ctx, k.coreStore, k.params, batch.BatchDenom)
			if err != nil {
				return err
			}
			marketId, err := k.getOrCreateMarketId(ctx, ct.Abbreviation, update.NewAskPrice.Denom)
			if err != nil {
				return err
			}
			order.MarketId = marketId
		}
		order.AskPrice = update.NewAskPrice.Amount.String()
		event.NewAskPrice = update.NewAskPrice
	}
	if update.NewExpiration != nil {
		order.Expiration = timestamppb.New(*update.NewExpiration)
		event.NewExpiration = update.NewExpiration
	}
	if update.NewQuantity != "" {
		order.Quantity = update.NewQuantity
		event.NewQuantity = update.NewQuantity
	}
}
