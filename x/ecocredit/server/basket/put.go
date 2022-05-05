package basket

import (
	"context"
	"time"

	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/basket/v1"
	ecoApi "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	regenmath "github.com/regen-network/regen-ledger/types/math"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	baskettypes "github.com/regen-network/regen-ledger/x/ecocredit/basket"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

// Put deposits ecocredits into a basket, returning fungible coins to the depositor.
// NOTE: the credits MUST adhere to the following specifications set by the basket: credit type, class, and date criteria.
func (k Keeper) Put(ctx context.Context, req *baskettypes.MsgPut) (*baskettypes.MsgPutResponse, error) {
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

	// keep track of the total amount of tokens to give to the depositor
	amountReceived := sdk.NewInt(0)
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	for _, credit := range req.Credits {
		// get credit batch info
		batchInfo, err := k.coreStore.BatchTable().GetByDenom(ctx, credit.BatchDenom)
		if err != nil {
			return nil, sdkerrors.ErrInvalidRequest.Wrapf("could not get batch %s: %s", credit.BatchDenom, err.Error())
		}

		// validate that the credit batch adheres to the basket's specifications
		if err := k.canBasketAcceptCredit(ctx, basket, batchInfo); err != nil {
			return nil, err
		}
		// get the amount of credits in dec
		amt, err := regenmath.NewPositiveFixedDecFromString(credit.Amount, basket.Exponent)
		if err != nil {
			return nil, err
		}
		// update the user and basket balances
		if err = k.transferToBasket(ctx, ownerAddr, amt, basket, batchInfo); err != nil {
			if sdkerrors.ErrInsufficientFunds.Is(err) {
				return nil, ecocredit.ErrInsufficientCredits
			}
			return nil, err
		}
		// get the amount of basket tokens to give to the depositor
		tokens, err := creditAmountToBasketCoins(amt, basket.Exponent, basket.BasketDenom)
		if err != nil {
			return nil, err
		}
		// update the total amount received so far
		amountReceived = amountReceived.Add(tokens[0].Amount)

		sdkCtx.GasMeter().ConsumeGas(ecocredit.GasCostPerIteration, "ecocredit/basket/MsgPut credit iteration")
	}

	// mint and send tokens to depositor
	coinsToSend := sdk.Coins{sdk.NewCoin(basket.BasketDenom, amountReceived)}
	if err = k.bankKeeper.MintCoins(sdkCtx, baskettypes.BasketSubModuleName, coinsToSend); err != nil {
		return nil, err
	}
	if err = k.bankKeeper.SendCoinsFromModuleToAccount(sdkCtx, baskettypes.BasketSubModuleName, ownerAddr, coinsToSend); err != nil {
		return nil, err
	}

	if err = sdkCtx.EventManager().EmitTypedEvent(&baskettypes.EventPut{
		Owner:       ownerAddr.String(),
		BasketDenom: basket.BasketDenom,
		Amount:      amountReceived.String(),
	}); err != nil {
		return nil, err
	}

	return &baskettypes.MsgPutResponse{AmountReceived: amountReceived.String()}, nil
}

// canBasketAcceptCredit checks that a credit adheres to the specifications of a basket. Specifically, it checks:
//  - batch's start time is within the basket's specified time window or min start date
//  - class is in the basket's allowed class store
//  - type matches the baskets specified credit type.
func (k Keeper) canBasketAcceptCredit(ctx context.Context, basket *api.Basket, batchInfo *ecoApi.Batch) error {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	blockTime := sdkCtx.BlockTime()
	errInvalidReq := sdkerrors.ErrInvalidRequest

	if basket.DateCriteria != nil {
		// check time window match
		var minStartDate time.Time
		var criteria = basket.DateCriteria
		if criteria.MinStartDate != nil {
			minStartDate = criteria.MinStartDate.AsTime()
		} else if criteria.StartDateWindow != nil {
			window := criteria.StartDateWindow.AsDuration()
			minStartDate = blockTime.Add(-window)
		} else if criteria.YearsInThePast != 0 {
			year := blockTime.Year() - int(criteria.YearsInThePast)
			minStartDate = time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)
		}

		startDate := batchInfo.StartDate.AsTime()
		if startDate.Before(minStartDate) {
			return errInvalidReq.Wrapf("cannot put a credit from a batch with start date %s "+
				"into a basket that requires an earliest start date of %s", batchInfo.StartDate.AsTime().String(), minStartDate.String())
		}

	}

	classId := core.GetClassIdFromBatchDenom(batchInfo.Denom)

	// check credit class match
	found, err := k.stateStore.BasketClassTable().Has(ctx, basket.Id, classId)
	if err != nil {
		return err
	}
	if !found {
		return errInvalidReq.Wrapf("credit class %s is not allowed in this basket", classId)
	}

	// check credit type match
	class, err := k.coreStore.ClassTable().GetById(ctx, classId)
	if err != nil {
		return err
	}
	if class.CreditTypeAbbrev != basket.CreditTypeAbbrev {
		return errInvalidReq.Wrapf("basket requires credit type %s but a credit with type %s was given", basket.CreditTypeAbbrev, class.CreditTypeAbbrev)
	}

	return nil
}

// transferToBasket moves credits from the user's tradable balance, into the basket's balance
func (k Keeper) transferToBasket(ctx context.Context, sender sdk.AccAddress, amt regenmath.Dec, basket *api.Basket, batchInfo *ecoApi.Batch) error {
	// update user balance, subtracting from their tradable balance
	userBal, err := k.coreStore.BatchBalanceTable().Get(ctx, sender, batchInfo.Key)
	if err != nil {
		return ecocredit.ErrInsufficientCredits.Wrapf("could not get batch %s balance for %s", batchInfo.Denom, sender.String())
	}
	tradable, err := regenmath.NewPositiveDecFromString(userBal.Tradable)
	if err != nil {
		return err
	}
	newTradable, err := regenmath.SafeSubBalance(tradable, amt)
	if err != nil {
		return ecocredit.ErrInsufficientCredits.Wrapf("cannot put %v credits into the basket with a balance of %v: %s", amt, tradable, err.Error())
	}
	userBal.Tradable = newTradable.String()
	if err = k.coreStore.BatchBalanceTable().Update(ctx, userBal); err != nil {
		return err
	}

	// update basket balance with amount sent, adding to the basket's balance.
	var bal *api.BasketBalance
	bal, err = k.stateStore.BasketBalanceTable().Get(ctx, basket.Id, batchInfo.Denom)
	if err != nil {
		if ormerrors.IsNotFound(err) {
			bal = &api.BasketBalance{
				BasketId:       basket.Id,
				BatchDenom:     batchInfo.Denom,
				Balance:        amt.String(),
				BatchStartDate: batchInfo.StartDate,
			}
		} else {
			return err
		}
	} else {
		newBalance, err := regenmath.NewPositiveFixedDecFromString(bal.Balance, basket.Exponent)
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

// creditAmountToBasketCoins calculates the tokens to award to the depositor
func creditAmountToBasketCoins(creditAmt regenmath.Dec, exp uint32, denom string) (sdk.Coins, error) {
	var coins sdk.Coins
	multiplier := regenmath.NewDecFinite(1, int32(exp))
	tokenAmt, err := multiplier.MulExact(creditAmt)
	if err != nil {
		return coins, err
	}

	amtInt, err := tokenAmt.BigInt()
	if err != nil {
		return coins, err
	}

	return sdk.Coins{sdk.NewCoin(denom, sdk.NewIntFromBigInt(amtInt))}, nil
}
