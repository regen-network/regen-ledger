package marketplace

import (
	"context"
	"fmt"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	marketApi "github.com/regen-network/regen-ledger/api/regen/ecocredit/marketplace/v1"
	"github.com/regen-network/regen-ledger/types/math"
	marketplacev1 "github.com/regen-network/regen-ledger/x/ecocredit/marketplace"
	"github.com/regen-network/regen-ledger/x/ecocredit/server"
)

// Sell creates new sell orders for credits
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
		ct, err := server.GetCreditTypeFromBatchDenom(ctx, k.coreStore, k.params, batch.BatchDenom)
		if err != nil {
			return nil, err
		}
		marketId, err := k.getOrCreateMarketId(ctx, ct.Abbreviation, order.AskPrice.Denom)
		// verify expiration is in the future
		if order.Expiration != nil && order.Expiration.Before(sdkCtx.BlockTime()) {
			return nil, sdkerrors.ErrInvalidRequest.Wrapf("expiration must be in the future: %s", order.Expiration)
		}
		sellQty, err := math.NewDecFromString(order.Quantity)
		if err != nil {
			return nil, err
		}
		if err = k.escrowCredits(ctx, ownerAcc, batch.Id, sellQty); err != nil {
			return nil, err
		}
		// TODO: pending param refactor https://github.com/regen-network/regen-ledger/issues/624
		//has, err := isDenomAllowed(ctx, k.stateStore, order.AskPrice.Denom)
		//if err != nil {
		//	return nil, err
		//}
		//if !has {
		//	return nil, ecocredit.ErrInvalidSellOrder.Wrapf("cannot use coin with denom %s in sell orders", order.AskPrice.Denom)
		//}

		id, err := k.stateStore.SellOrderStore().InsertReturningID(ctx, &marketApi.SellOrder{
			Seller:            ownerAcc,
			BatchId:           batch.Id,
			Quantity:          order.Quantity,
			MarketId:          marketId,
			AskPrice:          order.AskPrice.Amount.String(),
			DisableAutoRetire: order.DisableAutoRetire,
			Expiration:        timestamppb.New(*order.Expiration),
			Maker:             true, // maker is always true for sell orders
		})
		if err != nil {
			return nil, err
		}
		sellOrderIds[i] = id
		sdkCtx.GasMeter().ConsumeGas(gasCostPerIteration, "create sell order")
		if err = sdkCtx.EventManager().EmitTypedEvent(&marketplacev1.EventSell{
			OrderId:           id,
			BatchDenom:        batch.BatchDenom,
			Quantity:          order.Quantity,
			AskPrice:          order.AskPrice,
			DisableAutoRetire: order.DisableAutoRetire,
			Expiration:        order.Expiration,
		}); err != nil {
			return nil, err
		}
	}
	return &marketplacev1.MsgSellResponse{SellOrderIds: sellOrderIds}, nil
}

// getOrCreateMarketId attempts to get a market, creating one otherwise, and return the Id.
func (k Keeper) getOrCreateMarketId(ctx context.Context, creditTypeAbbrev, bankDenom string) (uint64, error) {
	market, err := k.stateStore.MarketStore().GetByCreditTypeBankDenom(ctx, creditTypeAbbrev, bankDenom)
	switch err {
	case nil:
		return market.Id, nil
	case ormerrors.NotFound:
		return k.stateStore.MarketStore().InsertReturningID(ctx, &marketApi.Market{
			CreditType:        creditTypeAbbrev,
			BankDenom:         bankDenom,
			PrecisionModifier: 0,
		})
	default:
		return 0, err
	}
}

func (k Keeper) escrowCredits(ctx context.Context, ownerAcc sdk.AccAddress, batchId uint64, sellQty math.Dec) error {
	// get the seller's balance
	bal, err := k.coreStore.BatchBalanceStore().Get(ctx, ownerAcc, batchId)
	if err != nil {
		return err
	}
	tradable, err := math.NewDecFromString(bal.Tradable)
	if err != nil {
		return err
	}
	// subtract the desired sell amount from tradable balance
	newTradable, err := math.SafeSubBalance(tradable, sellQty)
	if err != nil {
		return fmt.Errorf("tradable balance: %s, sell order request: %s - %w", tradable.String(), sellQty.String(), err)
	}

	escrowed, err := math.NewDecFromString(bal.Escrowed)
	if err != nil {
		return err
	}
	newEscrowed, err := math.SafeAddBalance(escrowed, sellQty)
	if err != nil {
		return fmt.Errorf("escrowed balance: %s, sell order request: %s - %w", escrowed.String(), sellQty.String(), err)
	}
	// set the new balance
	bal.Tradable = newTradable.String()
	bal.Escrowed = newEscrowed.String()

	// save
	if err = k.coreStore.BatchBalanceStore().Update(ctx, bal); err != nil {
		return err
	}

	// update the batch supply
	supply, err := k.coreStore.BatchSupplyStore().Get(ctx, batchId)
	if err != nil {
		return err
	}
	supTradable, err := math.NewDecFromString(supply.TradableAmount)
	if err != nil {
		return err
	}
	supEscrow, err := math.NewDecFromString(supply.EscrowedAmount)
	if err != nil {
		return err
	}
	supTradable, err = math.SafeSubBalance(supTradable, sellQty)
	if err != nil {
		return err
	}
	supEscrow, err = math.SafeAddBalance(supEscrow, sellQty)
	if err != nil {
		return err
	}
	supply.EscrowedAmount = supEscrow.String()
	supply.TradableAmount = supTradable.String()
	return k.coreStore.BatchSupplyStore().Save(ctx, supply)
}
