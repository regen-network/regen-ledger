package keeper

import (
	"context"
	"time"

	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	api "github.com/regen-network/regen-ledger/api/v2/regen/ecocredit/basket/v1"
	baseapi "github.com/regen-network/regen-ledger/api/v2/regen/ecocredit/v1"
	regenmath "github.com/regen-network/regen-ledger/types/v2/math"
	"github.com/regen-network/regen-ledger/x/ecocredit/v3"
	"github.com/regen-network/regen-ledger/x/ecocredit/v3/base"
	basetypes "github.com/regen-network/regen-ledger/x/ecocredit/v3/base/types/v1"
	basketsub "github.com/regen-network/regen-ledger/x/ecocredit/v3/basket"
	types "github.com/regen-network/regen-ledger/x/ecocredit/v3/basket/types/v1"
)

// Put deposits ecocredits into a basket, returning fungible coins to the depositor.
// NOTE: the credits MUST adhere to the following specifications set by the basket: credit type, class, and date criteria.
func (k Keeper) Put(ctx context.Context, req *types.MsgPut) (*types.MsgPutResponse, error) {
	ownerAddr, err := sdk.AccAddressFromBech32(req.Owner)
	if err != nil {
		return nil, err
	}

	// get the basket
	basket, err := k.stateStore.BasketTable().GetByBasketDenom(ctx, req.BasketDenom)
	if err != nil {
		if ormerrors.IsNotFound(err) {
			return nil, sdkerrors.ErrNotFound.Wrapf("basket %s not found", req.BasketDenom)
		}
		return nil, err
	}

	// get the credit type
	creditType, err := k.baseStore.CreditTypeTable().Get(ctx, basket.CreditTypeAbbrev)
	if err != nil {
		return nil, err
	}

	// keep track of the total amount of tokens to give to the depositor
	amountReceived := sdk.NewInt(0)
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	ownerString := ownerAddr.String()
	moduleAddrString := k.moduleAddress.String()
	for _, credit := range req.Credits {
		// get credit batch
		batch, err := k.baseStore.BatchTable().GetByDenom(ctx, credit.BatchDenom)
		if err != nil {
			return nil, sdkerrors.ErrInvalidRequest.Wrapf("could not get batch %s: %s", credit.BatchDenom, err.Error())
		}

		// validate that the credit batch adheres to the basket's specifications
		if err := k.canBasketAcceptCredit(ctx, basket, batch); err != nil {
			return nil, err
		}
		// get the amount of credits in dec
		amt, err := regenmath.NewPositiveFixedDecFromString(credit.Amount, creditType.Precision)
		if err != nil {
			return nil, err
		}
		// update the user and basket balances
		if err = k.transferToBasket(ctx, ownerAddr, amt, basket.Id, batch, creditType.Precision); err != nil {
			if sdkerrors.ErrInsufficientFunds.Is(err) {
				return nil, ecocredit.ErrInsufficientCredits
			}
			return nil, err
		}
		// get the amount of basket tokens to give to the depositor
		tokens, err := creditAmountToBasketCoin(amt, creditType.Precision, basket.BasketDenom)
		if err != nil {
			return nil, err
		}
		// update the total amount received so far
		amountReceived = amountReceived.Add(tokens.Amount)

		if err = sdkCtx.EventManager().EmitTypedEvent(&basetypes.EventTransfer{
			Sender:         ownerString,
			Recipient:      moduleAddrString, // basket submodule
			BatchDenom:     credit.BatchDenom,
			TradableAmount: credit.Amount,
		}); err != nil {
			return nil, err
		}

		sdkCtx.GasMeter().ConsumeGas(ecocredit.GasCostPerIteration, "ecocredit/basket/MsgPut credit iteration")
	}

	// mint and send tokens to depositor
	coinsToSend := sdk.Coins{sdk.NewCoin(basket.BasketDenom, amountReceived)}
	if err = k.bankKeeper.MintCoins(sdkCtx, basketsub.BasketSubModuleName, coinsToSend); err != nil {
		return nil, err
	}
	if err = k.bankKeeper.SendCoinsFromModuleToAccount(sdkCtx, basketsub.BasketSubModuleName, ownerAddr, coinsToSend); err != nil {
		return nil, err
	}

	amountReceivedString := amountReceived.String()

	if err = sdkCtx.EventManager().EmitTypedEvent(&types.EventPut{
		Owner:       ownerString,
		BasketDenom: basket.BasketDenom,
		Credits:     req.Credits,          // deprecated
		Amount:      amountReceivedString, // deprecated
	}); err != nil {
		return nil, err
	}

	return &types.MsgPutResponse{AmountReceived: amountReceivedString}, nil
}

// canBasketAcceptCredit checks that a credit adheres to the specifications of a basket. Specifically, it checks:
// - batch's start time is within the basket's specified time window or min start date
// - class is in the basket's allowed class store
// - type matches the baskets specified credit type.
func (k Keeper) canBasketAcceptCredit(ctx context.Context, basket *api.Basket, batch *baseapi.Batch) error {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	blockTime := sdkCtx.BlockTime()
	errInvalidReq := sdkerrors.ErrInvalidRequest

	if basket.DateCriteria != nil {
		// check time window match
		var minStartDate time.Time
		var criteria = basket.DateCriteria
		switch {
		case criteria.MinStartDate != nil:
			minStartDate = criteria.MinStartDate.AsTime()
		case criteria.StartDateWindow != nil:
			window := criteria.StartDateWindow.AsDuration()
			minStartDate = blockTime.Add(-window)
		case criteria.YearsInThePast != 0:
			year := blockTime.Year() - int(criteria.YearsInThePast)
			minStartDate = time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)
		}

		startDate := batch.StartDate.AsTime()
		if startDate.Before(minStartDate) {
			return errInvalidReq.Wrapf("cannot put a credit from a batch with start date %s "+
				"into a basket that requires an earliest start date of %s", batch.StartDate.AsTime().String(), minStartDate.String())
		}

	}

	classID := base.GetClassIDFromBatchDenom(batch.Denom)

	// check credit class match
	found, err := k.stateStore.BasketClassTable().Has(ctx, basket.Id, classID)
	if err != nil {
		return err
	}
	if !found {
		return errInvalidReq.Wrapf("credit class %s is not allowed in this basket", classID)
	}

	// check credit type match
	class, err := k.baseStore.ClassTable().GetById(ctx, classID)
	if err != nil {
		return err
	}
	if class.CreditTypeAbbrev != basket.CreditTypeAbbrev {
		return errInvalidReq.Wrapf("basket requires credit type %s but a credit with type %s was given", basket.CreditTypeAbbrev, class.CreditTypeAbbrev)
	}

	return nil
}

// transferToBasket moves credits from the user's tradable balance, into the basket's balance
func (k Keeper) transferToBasket(ctx context.Context, sender sdk.AccAddress, amt regenmath.Dec, basketID uint64, batch *baseapi.Batch, exponent uint32) error {
	// update user balance, subtracting from their tradable balance
	userBal, err := k.baseStore.BatchBalanceTable().Get(ctx, sender, batch.Key)
	if err != nil {
		return ecocredit.ErrInsufficientCredits.Wrapf("could not get batch %s balance for %s", batch.Denom, sender.String())
	}
	tradable, err := regenmath.NewPositiveDecFromString(userBal.TradableAmount)
	if err != nil {
		return err
	}
	newTradable, err := regenmath.SafeSubBalance(tradable, amt)
	if err != nil {
		return ecocredit.ErrInsufficientCredits.Wrapf("cannot put %v credits into the basket with a balance of %v: %s", amt, tradable, err.Error())
	}
	userBal.TradableAmount = newTradable.String()
	if err = k.baseStore.BatchBalanceTable().Update(ctx, userBal); err != nil {
		return err
	}

	// update basket balance with amount sent, adding to the basket's balance.
	var bal *api.BasketBalance
	bal, err = k.stateStore.BasketBalanceTable().Get(ctx, basketID, batch.Denom)
	if err != nil {
		if ormerrors.IsNotFound(err) {
			bal = &api.BasketBalance{
				BasketId:       basketID,
				BatchDenom:     batch.Denom,
				Balance:        amt.String(),
				BatchStartDate: batch.StartDate,
			}
		} else {
			return err
		}
	} else {
		newBalance, err := regenmath.NewPositiveFixedDecFromString(bal.Balance, exponent)
		if err != nil {
			return err
		}
		newBalance, err = newBalance.Add(amt)
		if err != nil {
			return err
		}
		bal.Balance = newBalance.String()
	}
	if err = k.stateStore.BasketBalanceTable().Save(ctx, bal); err != nil {
		return err
	}
	return nil
}
