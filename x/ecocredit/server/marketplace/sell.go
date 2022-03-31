package marketplace

import (
	"context"

	"github.com/regen-network/regen-ledger/x/ecocredit/server/utils"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	marketApi "github.com/regen-network/regen-ledger/api/regen/ecocredit/marketplace/v1"
	"github.com/regen-network/regen-ledger/types/math"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	marketplacev1 "github.com/regen-network/regen-ledger/x/ecocredit/marketplace"
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
		batch, err := k.coreStore.BatchInfoTable().GetByBatchDenom(ctx, order.BatchDenom)
		if err != nil {
			return nil, sdkerrors.ErrInvalidRequest.Wrapf("batch denom %s: %s", order.BatchDenom, err.Error())
		}
		ct, err := utils.GetCreditTypeFromBatchDenom(ctx, k.coreStore, k.paramsKeeper, batch.BatchDenom)
		if err != nil {
			return nil, err
		}
		marketId, err := k.getOrCreateMarketId(ctx, ct.Abbreviation, order.AskPrice.Denom)
		if err != nil {
			return nil, err
		}

		// verify expiration is in the future
		if order.Expiration != nil && order.Expiration.Before(sdkCtx.BlockTime()) {
			return nil, sdkerrors.ErrInvalidRequest.Wrapf("expiration must be in the future: %s", order.Expiration)
		}
		sellQty, err := math.NewPositiveFixedDecFromString(order.Quantity, ct.Precision)
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

		var expiration *timestamppb.Timestamp
		if order.Expiration != nil {
			expiration = timestamppb.New(*order.Expiration)
		}

		id, err := k.stateStore.SellOrderTable().InsertReturningID(ctx, &marketApi.SellOrder{
			Seller:            ownerAcc,
			BatchId:           batch.Id,
			Quantity:          order.Quantity,
			MarketId:          marketId,
			AskPrice:          order.AskPrice.Amount.String(),
			DisableAutoRetire: order.DisableAutoRetire,
			Expiration:        expiration,
			Maker:             true, // maker is always true for sell orders
		})
		if err != nil {
			return nil, err
		}

		sellOrderIds[i] = id
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

		sdkCtx.GasMeter().ConsumeGas(ecocredit.GasCostPerIteration, "ecocredit/core/MsgSell order iteration")
	}
	return &marketplacev1.MsgSellResponse{SellOrderIds: sellOrderIds}, nil
}

// getOrCreateMarketId attempts to get a market, creating one otherwise, and return the Id.
func (k Keeper) getOrCreateMarketId(ctx context.Context, creditTypeAbbrev, bankDenom string) (uint64, error) {
	market, err := k.stateStore.MarketTable().GetByCreditTypeBankDenom(ctx, creditTypeAbbrev, bankDenom)
	switch err {
	case nil:
		return market.Id, nil
	case ormerrors.NotFound:
		return k.stateStore.MarketTable().InsertReturningID(ctx, &marketApi.Market{
			CreditType:        creditTypeAbbrev,
			BankDenom:         bankDenom,
			PrecisionModifier: 0,
		})
	default:
		return 0, err
	}
}

func (k Keeper) escrowCredits(ctx context.Context, ownerAcc sdk.AccAddress, batchId uint64, sellQty math.Dec) error {
	bal, err := k.coreStore.BatchBalanceTable().Get(ctx, ownerAcc, batchId)
	if err != nil {
		return err
	}
	tradable, err := math.NewDecFromString(bal.Tradable)
	if err != nil {
		return err
	}
	newTradable, err := math.SafeSubBalance(tradable, sellQty)
	if err != nil {
		return sdkerrors.ErrInvalidRequest.Wrapf("tradable balance: %v, sell order request: %v - %s", tradable, sellQty, err.Error())
	}

	escrowed, err := math.NewDecFromString(bal.Escrowed)
	if err != nil {
		return err
	}
	newEscrowed, err := math.SafeAddBalance(escrowed, sellQty)
	if err != nil {
		return sdkerrors.ErrInvalidRequest.Wrapf("escrowed balance: %v, sell order request: %v - %s", escrowed, sellQty, err.Error())
	}
	bal.Tradable = newTradable.String()
	bal.Escrowed = newEscrowed.String()

	if err = k.coreStore.BatchBalanceTable().Update(ctx, bal); err != nil {
		return err
	}

	supply, err := k.coreStore.BatchSupplyTable().Get(ctx, batchId)
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
	return k.coreStore.BatchSupplyTable().Save(ctx, supply)
}
