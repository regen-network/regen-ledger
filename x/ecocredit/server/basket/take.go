package basket

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/basket/v1"
	"github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/types/math"
	"github.com/regen-network/regen-ledger/x/ecocredit"
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
	err = k.bankKeeper.SendCoinsFromAccountToModule(sdkContext, acct, baskettypes.BasketSubModuleName, basketCoins)
	if err != nil {
		return nil, err
	}

	err = k.bankKeeper.BurnCoins(sdkContext, baskettypes.BasketSubModuleName, basketCoins)
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
			api.BasketBalanceBasketIdBatchStartDateIndexKey{}.WithBasketId(basket.Id),
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
		sdkContext.GasMeter().ConsumeGas(ecocredit.GasCostPerIteration, "ecocredit/basket/MsgTake iteration")

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

			err = k.addCreditBalance(
				ctx,
				acct,
				basketBalance.BatchDenom,
				amountCreditsNeeded,
				basket.BasketDenom,
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

			err = k.addCreditBalance(
				ctx,
				acct,
				basketBalance.BatchDenom,
				balance,
				basket.BasketDenom,
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

	err = sdk.UnwrapSDKContext(ctx).EventManager().EmitTypedEvent(&baskettypes.EventTake{
		Owner:       msg.Owner,
		BasketDenom: msg.BasketDenom,
		Credits:     credits,
		Amount:      msg.Amount,
	})
	return &baskettypes.MsgTakeResponse{
		Credits: credits,
	}, err
}

func (k Keeper) addCreditBalance(ctx context.Context, owner sdk.AccAddress, batchDenom string, amount math.Dec, basketDenom string, retire bool, retirementLocation string) error {
	sdkCtx := types.UnwrapSDKContext(ctx)
	store := sdkCtx.KVStore(k.storeKey)
	if !retire {
		err := ecocredit.AddAndSetDecimal(store, ecocredit.TradableBalanceKey(owner, ecocredit.BatchDenomT(batchDenom)), amount)
		if err != nil {
			return err
		}

		return sdkCtx.EventManager().EmitTypedEvent(&ecocredit.EventReceive{
			Recipient:      owner.String(),
			BatchDenom:     batchDenom,
			TradableAmount: amount.String(),
			BasketDenom:    basketDenom,
		})
	} else {
		err := ecocredit.AddAndSetDecimal(store, ecocredit.RetiredBalanceKey(owner, ecocredit.BatchDenomT(batchDenom)), amount)
		if err != nil {
			return err
		}

		err = ecocredit.AddAndSetDecimal(store, ecocredit.RetiredSupplyKey(ecocredit.BatchDenomT(batchDenom)), amount)
		if err != nil {
			return err
		}

		err = ecocredit.SubAndSetDecimal(store, ecocredit.TradableSupplyKey(ecocredit.BatchDenomT(batchDenom)), amount)
		if err != nil {
			return err
		}

		err = sdkCtx.EventManager().EmitTypedEvent(&ecocredit.EventReceive{
			Recipient:     owner.String(),
			BatchDenom:    batchDenom,
			RetiredAmount: amount.String(),
			BasketDenom:   basketDenom,
		})
		if err != nil {
			return err
		}

		return sdkCtx.EventManager().EmitTypedEvent(&ecocredit.EventRetire{
			Retirer:    owner.String(),
			BatchDenom: batchDenom,
			Amount:     amount.String(),
			Location:   retirementLocation,
		})
	}
}
