package basket

import (
	"context"
	"fmt"
	"math"
	"time"

	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	basketv1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/basket/v1"
	regenmath "github.com/regen-network/regen-ledger/types/math"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	baskettypes "github.com/regen-network/regen-ledger/x/ecocredit/basket"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func (k Keeper) Put(ctx context.Context, req *baskettypes.MsgPut) (*baskettypes.MsgPutResponse, error) {
	ownerAddr, _ := sdk.AccAddressFromBech32(req.Owner)

	basket, err := k.stateStore.BasketStore().GetByBasketDenom(ctx, req.BasketDenom)
	if err != nil {
		return nil, err
	}

	amountReceived := sdk.NewInt(0)
	for _, credit := range req.Credits {
		// get credit batch info
		res, err := k.ecocreditKeeper.BatchInfo(ctx, &ecocredit.QueryBatchInfoRequest{BatchDenom: credit.BatchDenom})
		if err != nil {
			return nil, err
		}
		batchInfo := res.Info
		if err := k.validateCredit(ctx, basket, batchInfo); err != nil {
			return nil, err
		}
		// get the amount of credits in dec
		amt, err := regenmath.NewPositiveFixedDecFromString(credit.Amount, basket.Exponent)
		if err != nil {
			return nil, err
		}
		// update the user and basket balances
		if err = k.updateBalances(ctx, ownerAddr, amt, basket, batchInfo); err != nil {
			return nil, err
		}
		// get the amount of basket tokens to award to the depositor
		tokenAward, err := calculateTokenAward(amt, basket.Exponent, basket.BasketDenom)
		if err != nil {
			return nil, err
		}
		// update the total amount received so far
		amountReceived = amountReceived.Add(tokenAward[0].Amount)
		// mint and send tokens to depositor
		sdkCtx := sdk.UnwrapSDKContext(ctx)
		if err = k.bankKeeper.MintCoins(sdkCtx, ecocredit.ModuleName, tokenAward); err != nil {
			return nil, err
		}
		if err = k.bankKeeper.SendCoinsFromModuleToAccount(sdkCtx, ecocredit.ModuleName, ownerAddr, tokenAward); err != nil {
			return nil, err
		}
	}
	return &baskettypes.MsgPutResponse{AmountReceived: amountReceived.String()}, nil
}

// validateCredit checks that a credit adheres to the specifications of a basket. Specifically, it checks:
//  - batch's start time is within the basket's MinStartTime
//  - class is in the basket's allowed class store
//  - type matches the baskets specified credit type.
func (k Keeper) validateCredit(ctx context.Context, basket *basketv1.Basket, batchInfo *ecocredit.BatchInfo) error {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	blockTime := sdkCtx.BlockTime()
	// check time window match
	var minStartDate time.Time
	switch basket.DateCriteria.Sum.(type) {
	case *basketv1.DateCriteria_MinStartDate:
		min := basket.DateCriteria.Sum.(*basketv1.DateCriteria_MinStartDate).MinStartDate
		minStartDate = min.AsTime()
	case *basketv1.DateCriteria_StartDateWindow:
		window := basket.DateCriteria.Sum.(*basketv1.DateCriteria_StartDateWindow).StartDateWindow.AsDuration()
		minStartDate = blockTime.Add(window)
	default:
		return sdkerrors.ErrInvalidRequest.Wrap("no date criteria was given") // TODO: is date criteria required?
	}

	if batchInfo.StartDate.Before(minStartDate) {
		return sdkerrors.ErrInvalidRequest.Wrapf("cannot put a credit from a batch with start time %s "+
			"into a basket that requires a min start time of %s", batchInfo.StartDate.String(), minStartDate.String())
	}

	// check credit class match
	found, err := k.stateStore.BasketClassStore().Has(ctx, basket.Id, batchInfo.ClassId)
	if err != nil {
		return err
	}
	if !found {
		return sdkerrors.ErrInvalidRequest.Wrapf("credit class %s is not allowed in this basket", batchInfo.ClassId)
	}

	// check credit type match
	requiredCreditType := basket.CreditTypeName
	res2, err := k.ecocreditKeeper.ClassInfo(ctx, &ecocredit.QueryClassInfoRequest{ClassId: batchInfo.ClassId})
	if err != nil {
		return err
	}
	gotCreditType := res2.Info.CreditType.Name
	if requiredCreditType != gotCreditType {
		return sdkerrors.ErrInvalidRequest.Wrapf("cannot use credit of type %s in a basket that requires credit type %s", gotCreditType, requiredCreditType)
	}
	return nil
}

// updateBalances updates the balance of the user in the legacy KVStore as well as the basket's balance in the ORM.
func (k Keeper) updateBalances(ctx context.Context, sender sdk.AccAddress, amt regenmath.Dec, basket *basketv1.Basket, batchInfo *ecocredit.BatchInfo) error {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	store := sdkCtx.KVStore(k.storeKey)

	// update the user balance
	userBalanceKey := ecocredit.TradableBalanceKey(sender, ecocredit.BatchDenomT(batchInfo.BatchDenom))
	userBalance, err := ecocredit.GetDecimal(store, userBalanceKey)
	if err != nil {
		return err
	}
	newUserBalance, err := regenmath.SafeSubBalance(userBalance, amt)
	if err != nil {
		return err
	}
	ecocredit.SetDecimal(store, userBalanceKey, newUserBalance)

	// update basket balance with amount sent
	var bal *basketv1.BasketBalance
	bal, err = k.stateStore.BasketBalanceStore().Get(ctx, basket.Id, batchInfo.BatchDenom)
	if err != nil {
		if ormerrors.IsNotFound(err) {
			bal = &basketv1.BasketBalance{
				BasketId:       basket.Id,
				BatchDenom:     batchInfo.BatchDenom,
				Balance:        amt.String(),
				BatchStartDate: timestamppb.New(*batchInfo.StartDate),
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
	if err = k.stateStore.BasketBalanceStore().Save(ctx, bal); err != nil {
		return err
	}
	return nil
}

// calculateTokenAward calculates the tokens to award to the depositor
func calculateTokenAward(creditAmt regenmath.Dec, exp uint32, denom string) (sdk.Coins, error) {
	multiplier := math.Pow10(int(exp))
	multiStr := fmt.Sprint(multiplier)
	dec, err := regenmath.NewPositiveFixedDecFromString(multiStr, exp)
	if err != nil {
		return sdk.Coins{}, err
	}

	tokens, err := creditAmt.Mul(dec)
	if err != nil {
		return sdk.Coins{}, err
	}

	i64Tokens, err := tokens.Int64()
	if err != nil {
		return sdk.Coins{}, err
	}

	return sdk.Coins{sdk.NewCoin(denom, sdk.NewInt(i64Tokens))}, nil
}
