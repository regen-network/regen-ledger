package basket

import (
	"context"
	"fmt"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/regen-network/regen-ledger/types/math"

	basketv1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/basket/v1"
	baskettypes "github.com/regen-network/regen-ledger/x/ecocredit/basket"
)

func (k Keeper) Take(ctx context.Context, msg *baskettypes.MsgTake) (*baskettypes.MsgTakeResponse, error) {
	basket, err := k.stateStore.BasketStore().GetByBasketDenom(ctx, msg.BasketDenom)
	if err != nil {
		return nil, err
	}

	retire := msg.RetireOnTake
	if !basket.DisableAutoRetire && !retire {
		return nil, ErrCantDisableRetire
	}

	amountBasketTokens, ok := sdk.NewIntFromString(msg.Amount)
	if !ok {
		return nil, sdkerrors.ErrInvalidRequest.Wrapf("bad integer %s", msg.Amount)
	}

	acct, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return nil, err
	}

	sdkContext := sdk.UnwrapSDKContext(ctx)
	basketCoins := sdk.NewCoins(sdk.NewCoin(basket.BasketDenom, amountBasketTokens))
	err = k.bankKeeper.SendCoinsFromAccountToModule(sdkContext, acct, k.moduleAccountName, basketCoins)
	if err != nil {
		return nil, err
	}

	err = k.bankKeeper.BurnCoins(sdkContext, k.moduleAccountName, basketCoins)
	if err != nil {
		return nil, err
	}

	amountBasketTokensDec, err := math.NewDecFromString(msg.Amount)
	if err != nil {
		return nil, err
	}

	multiplier := math.NewDecFinite(1, int32(basket.Exponent))
	amountCreditsNeeded, err := amountBasketTokensDec.QuoExact(multiplier)
	if err != nil {
		return nil, err
	}

	var credits []*baskettypes.BasketCredit
	for {
		it, err := k.stateStore.BasketBalanceStore().List(ctx,
			basketv1.BasketBalanceBasketIdBatchStartDateIndexKey{}.WithBasketId(basket.Id),
		)
		if err != nil {
			return nil, err
		}

		if !it.Next() {
			return nil, fmt.Errorf("unexpected failure - balance invariant broken")
		}

		basketBalance, err := it.Value()
		if err != nil {
			return nil, err
		}
		it.Close()

		balance, err := math.NewDecFromString(basketBalance.Balance)
		if err != nil {
			return nil, err
		}

		cmp := balance.Cmp(amountCreditsNeeded)
		if cmp > 0 {
			credits = append(credits, &baskettypes.BasketCredit{
				BatchDenom: basketBalance.BatchDenom,
				Amount:     amountCreditsNeeded.String(),
			})

			err = k.ecocreditKeeper.AddCreditBalance(
				ctx,
				acct,
				basketBalance.BatchDenom,
				amountCreditsNeeded,
				retire,
				msg.RetirementLocation,
			)
			if err != nil {
				return nil, err
			}

			newBalance, err := balance.Sub(amountCreditsNeeded)
			if err != nil {
				return nil, err
			}

			basketBalance.Balance = newBalance.String()
			err = k.stateStore.BasketBalanceStore().Update(ctx, basketBalance)
			if err != nil {
				return nil, err
			}

			break
		} else {
			credits = append(credits, &baskettypes.BasketCredit{
				BatchDenom: basketBalance.BatchDenom,
				Amount:     balance.String(),
			})

			err = k.ecocreditKeeper.AddCreditBalance(
				ctx,
				acct,
				basketBalance.BatchDenom,
				balance,
				retire,
				msg.RetirementLocation,
			)
			if err != nil {
				return nil, err
			}

			err = k.stateStore.BasketBalanceStore().Delete(ctx, basketBalance)
			if err != nil {
				return nil, err
			}

			// basket balance == credits needed
			if cmp == 0 {
				break
			}

			amountCreditsNeeded, err = amountCreditsNeeded.Sub(balance)
			if err != nil {
				return nil, err
			}
		}
	}

	return &baskettypes.MsgTakeResponse{
		Credits: credits,
	}, nil
}
