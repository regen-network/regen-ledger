package basket

import (
	"context"

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

	amount, err := math.NewPositiveFixedDecFromString(msg.Amount, basket.Exponent)
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	it, err := k.stateStore.BasketBalanceStore().List(ctx, basketv1.BasketBalanceBatchStartDateIndexKey{})
	if err != nil {
		return nil, err
	}

	var credits []*baskettypes.BasketCredit
	for it.Next() {
		basketBalance, err := it.Value()
		if err != nil {
			return nil, err
		}

		balance, err := math.NewDecFromString(basketBalance.Balance)
		if err != nil {
			return nil, err
		}

		if balance.Cmp(amount) > 0 {
			credits = append(credits, &baskettypes.BasketCredit{
				BatchDenom: basketBalance.BatchDenom,
				Amount:     amount.String(),
			})

			newBalance, err := balance.Sub(amount)
			if err != nil {
				return nil, err
			}

			basketBalance.Balance = newBalance.String()
			err = it.Update(basketBalance)
			if err != nil {
				return nil, err
			}

			break
		} else {
			err = it.Delete()
			if err != nil {
				return nil, err
			}
		}

	}
}
