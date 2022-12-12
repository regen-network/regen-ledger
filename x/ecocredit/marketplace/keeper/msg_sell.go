package keeper

import (
	"context"
	"fmt"

	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	marketApi "github.com/regen-network/regen-ledger/api/v2/regen/ecocredit/marketplace/v1"
	"github.com/regen-network/regen-ledger/types/v2/math"
	"github.com/regen-network/regen-ledger/x/ecocredit/v3"
	types "github.com/regen-network/regen-ledger/x/ecocredit/v3/marketplace/types/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/v3/server/utils"
)

// Sell creates new sell orders for credits
func (k Keeper) Sell(ctx context.Context, req *types.MsgSell) (*types.MsgSellResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	sellerAcc, err := sdk.AccAddressFromBech32(req.Seller)
	if err != nil {
		return nil, err
	}

	sellOrderIDs := make([]uint64, len(req.Orders))

	for i, order := range req.Orders {
		// orderIndex is used for more granular error messages when
		// an individual order in a list of orders fails to process
		orderIndex := fmt.Sprintf("orders[%d]", i)

		batch, err := k.baseStore.BatchTable().GetByDenom(ctx, order.BatchDenom)
		if err != nil {
			return nil, sdkerrors.ErrInvalidRequest.Wrapf(
				"%s: batch denom %s: %s", orderIndex, order.BatchDenom, err,
			)
		}

		creditType, err := utils.GetCreditTypeFromBatchDenom(ctx, k.baseStore, batch.Denom)
		if err != nil {
			return nil, err
		}

		marketID, err := k.getOrCreateMarketID(ctx, creditType.Abbreviation, order.AskPrice.Denom)
		if err != nil {
			return nil, err
		}

		// verify expiration is in the future
		if order.Expiration != nil && !order.Expiration.UTC().After(sdkCtx.BlockTime()) {
			return nil, sdkerrors.ErrInvalidRequest.Wrapf(
				"%s: expiration must be in the future: %s", orderIndex, order.Expiration,
			)
		}

		sellQty, err := math.NewPositiveFixedDecFromString(order.Quantity, creditType.Precision)
		if err != nil {
			return nil, err
		}

		// convert seller balance tradable credits to escrowed credits
		if err = k.escrowCredits(ctx, orderIndex, sellerAcc, batch.Key, sellQty); err != nil {
			return nil, err
		}

		// check if ask price denom is an allowed denom
		allowed, err := isDenomAllowed(ctx, order.AskPrice.Denom, k.stateStore.AllowedDenomTable())
		if err != nil {
			return nil, err
		}
		if !allowed {
			return nil, sdkerrors.ErrInvalidRequest.Wrapf(
				"%s: %s is not allowed to be used in sell orders", orderIndex, order.AskPrice.Denom,
			)
		}

		var expiration *timestamppb.Timestamp
		if order.Expiration != nil {
			expiration = timestamppb.New(*order.Expiration)
		}

		id, err := k.stateStore.SellOrderTable().InsertReturningID(ctx, &marketApi.SellOrder{
			Seller:            sellerAcc,
			BatchKey:          batch.Key,
			Quantity:          order.Quantity,
			MarketId:          marketID,
			AskAmount:         order.AskPrice.Amount.String(),
			DisableAutoRetire: order.DisableAutoRetire,
			Expiration:        expiration,
			Maker:             true, // maker is always true for sell orders
		})
		if err != nil {
			return nil, err
		}

		sellOrderIDs[i] = id

		if err = sdkCtx.EventManager().EmitTypedEvent(&types.EventSell{
			SellOrderId: id,
		}); err != nil {
			return nil, err
		}

		sdkCtx.GasMeter().ConsumeGas(ecocredit.GasCostPerIteration, "ecocredit/marketplace/MsgSell order iteration")
	}

	return &types.MsgSellResponse{SellOrderIds: sellOrderIDs}, nil
}

// getOrCreateMarketID attempts to get a market, otherwise creating a market, and returns the ID.
func (k Keeper) getOrCreateMarketID(ctx context.Context, creditTypeAbbrev, bankDenom string) (uint64, error) {
	market, err := k.stateStore.MarketTable().GetByCreditTypeAbbrevBankDenom(ctx, creditTypeAbbrev, bankDenom)
	switch err {
	case nil:
		return market.Id, nil
	case ormerrors.NotFound:
		return k.stateStore.MarketTable().InsertReturningID(ctx, &marketApi.Market{
			CreditTypeAbbrev:  creditTypeAbbrev,
			BankDenom:         bankDenom,
			PrecisionModifier: 0,
		})
	default:
		return 0, err
	}
}

// escrowCredits updates seller balance, subtracting the provided quantity from tradable amount
// and adding it to escrowed amount.
func (k Keeper) escrowCredits(ctx context.Context, orderIndex string, sellerAcc sdk.AccAddress, batchKey uint64, quantity math.Dec) error {

	// get seller balance to be updated
	bal, err := k.baseStore.BatchBalanceTable().Get(ctx, sellerAcc, batchKey)
	if err != nil {
		return ecocredit.ErrInsufficientCredits.Wrapf(
			"%s: credit quantity: %v, tradable balance: 0", orderIndex, quantity,
		)
	}

	// calculate and set seller balance tradable amount (subtract credits)
	tradable, err := math.NewDecFromString(bal.TradableAmount)
	if err != nil {
		return err
	}
	newTradable, err := math.SafeSubBalance(tradable, quantity)
	if err != nil {
		return ecocredit.ErrInsufficientCredits.Wrapf(
			"%s: credit quantity: %v, tradable balance: %v", orderIndex, quantity, tradable,
		)
	}
	bal.TradableAmount = newTradable.String()

	// calculate and set seller balance escrowed amount (add credits)
	escrowed, err := math.NewDecFromString(bal.EscrowedAmount)
	if err != nil {
		return err
	}
	newEscrowed, err := math.SafeAddBalance(escrowed, quantity)
	if err != nil {
		return err
	}
	bal.EscrowedAmount = newEscrowed.String()

	// update seller balance with new tradable and escrowed amounts
	return k.baseStore.BatchBalanceTable().Update(ctx, bal)
}
