package marketplace

import (
	"context"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	marketApi "github.com/regen-network/regen-ledger/api/regen/ecocredit/marketplace/v1"
	"github.com/regen-network/regen-ledger/types/math"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	marketplacev1 "github.com/regen-network/regen-ledger/x/ecocredit/marketplace"
	"github.com/regen-network/regen-ledger/x/ecocredit/server/core"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (k Keeper) Sell(ctx context.Context, req *marketplacev1.MsgSell) (*marketplacev1.MsgSellResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	ownerAcc, err := sdk.AccAddressFromBech32(req.Owner)
	if err != nil {
		return nil, err
	}

	sellOrderIds := make([]uint64, len(req.Orders))

	for i, order := range req.Orders {
		batch, err := k.coreStore.BatchInfoStore().GetByBatchDenom(ctx, order.BatchDenom)
		if err != nil {
			return nil, fmt.Errorf("batch denom %v: %v", order.BatchDenom, err)
		}
		// TODO: if not found, do we create a market??
		market, err := k.getMarket(ctx, order.BatchDenom, order.AskPrice.Denom)
		if err != nil {
			return nil, fmt.Errorf("market: batch denom %v, bank denom %v: %v", order.BatchDenom, order.AskPrice.Denom, err)
		}
		// verify expiration is in the future
		if order.Expiration != nil && order.Expiration.Before(sdkCtx.BlockTime()) {
			return nil, sdkerrors.ErrInvalidRequest.Wrapf("expiration must be in the future: %s", order.Expiration)
		}
		sellQty, err := math.NewDecFromString(order.Quantity)
		if err != nil {
			return nil, err
		}
		if err := assertHasBalance(ctx, k.coreStore, ownerAcc, batch.Id, sellQty); err != nil {
			return nil, err
		}
		has, err := isDenomAllowed(ctx, k.stateStore, order.AskPrice.Denom)
		if err != nil {
			return nil, err
		}
		if !has {
			return nil, ecocredit.ErrInvalidSellOrder.Wrapf("cannot use coin with denom %s in sell orders", order.AskPrice.Denom)
		}
		id, err := k.stateStore.SellOrderStore().InsertReturningID(ctx, &marketApi.SellOrder{
			Seller:            ownerAcc,
			BatchId:           batch.Id,
			Quantity:          order.Quantity,
			MarketId:          market.Id,
			AskPrice:          order.AskPrice.Amount.String(),
			DisableAutoRetire: order.DisableAutoRetire,
			Expiration:        timestamppb.New(*order.Expiration),
			Maker:             false, // TODO: ?????????
		})
		if err != nil {
			return nil, err
		}
		sellOrderIds[i] = id
		sdkCtx.GasMeter().ConsumeGas(gasCostPerIteration, "create sell order")
	}
	return &marketplacev1.MsgSellResponse{SellOrderIds: sellOrderIds}, nil
}

func (k Keeper) getMarket(ctx context.Context, batchDenom, bankDenom string) (*marketApi.Market, error) {
	ct, err := core.GetCreditTypeFromBatchDenom(ctx, k.coreStore, k.params, batchDenom)
	if err != nil {
		return nil, err
	}
	return k.stateStore.MarketStore().GetByCreditTypeBankDenom(ctx, ct.Abbreviation, bankDenom)
}
