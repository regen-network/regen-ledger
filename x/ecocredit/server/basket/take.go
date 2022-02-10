package basket

import (
	"context"

	"github.com/cockroachdb/apd/v3"
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

	amountBasketTokens, _, err := apd.NewFromString(msg.Amount)
	if err != nil {
		return nil, err
	}

	multiplier := apd.New(10, int32(basket.Exponent))
	amountCreditsNeeded := &apd.Decimal{}
	_, err = apd.BaseContext.Quo(amountCreditsNeeded, amountBasketTokens, multiplier)
	if err != nil {
		return nil, err
	}

	var credits []*baskettypes.BasketCredit
	for {
		it, err := k.stateStore.BasketBalanceStore().List(ctx, basketv1.BasketBalanceBatchStartDateIndexKey{})
		if err != nil {
			return nil, err
		}

		basketBalance, err := it.Value()
		if err != nil {
			return nil, err
		}
		it.Close()

		balance, _, err := apd.NewFromString(basketBalance.Balance)
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
				msg.Owner,
				basketBalance.BatchDenom,
				amountCreditsNeeded,
				retire,
				msg.RetirementLocation,
			)
			if err != nil {
				return nil, err
			}

			_, err := apd.BaseContext.Sub(balance, balance, amountCreditsNeeded)
			if err != nil {
				return nil, err
			}

			basketBalance.Balance = balance.Text('f')
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
				msg.Owner,
				basketBalance.BatchDenom,
				balance,
				retire,
				msg.RetirementLocation,
			)
			if err != nil {
				return nil, err
			}

			err = k.stateStore.BasketBalanceStore().Update(ctx, basketBalance)
			if err != nil {
				return nil, err
			}

			// basket balance == credits needed
			if cmp == 0 {
				break
			}

			_, err = apd.BaseContext.Sub(amountCreditsNeeded, amountCreditsNeeded, balance)
			if err != nil {
				return nil, err
			}
		}
	}

	return &baskettypes.MsgTakeResponse{
		Credits: credits,
	}, nil
}
