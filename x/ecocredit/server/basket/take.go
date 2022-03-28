package basket

import (
	"context"
	"fmt"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/basket/v1"
	ecoApi "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/types/math"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	baskettypes "github.com/regen-network/regen-ledger/x/ecocredit/basket"

	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k Keeper) Take(ctx context.Context, msg *baskettypes.MsgTake) (*baskettypes.MsgTakeResponse, error) {
	basket, err := k.stateStore.BasketTable().GetByBasketDenom(ctx, msg.BasketDenom)
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

	sdkCtx := sdk.UnwrapSDKContext(ctx)
	basketCoins := sdk.NewCoins(sdk.NewCoin(basket.BasketDenom, amountBasketTokens))
	err = k.bankKeeper.SendCoinsFromAccountToModule(sdkCtx, acct, baskettypes.BasketSubModuleName, basketCoins)
	if err != nil {
		return nil, err
	}

	err = k.bankKeeper.BurnCoins(sdkCtx, baskettypes.BasketSubModuleName, basketCoins)
	if err != nil {
		return nil, err
	}

	amountBasketCreditsDec, err := math.NewDecFromString(msg.Amount)
	if err != nil {
		return nil, err
	}

	multiplier := math.NewDecFinite(1, int32(basket.Exponent))
	amountCreditsNeeded, err := amountBasketCreditsDec.QuoExact(multiplier)
	if err != nil {
		return nil, err
	}

	var credits []*baskettypes.BasketCredit
	for {
		it, err := k.stateStore.BasketBalanceTable().List(ctx,
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
			err = k.stateStore.BasketBalanceTable().Update(ctx, basketBalance)
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

			err = k.stateStore.BasketBalanceTable().Delete(ctx, basketBalance)
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

		sdkCtx.GasMeter().ConsumeGas(ecocredit.GasCostPerIteration, "ecocredit/basket/MsgTake balance iteration")
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
	batch, err := k.coreStore.BatchInfoTable().GetByBatchDenom(ctx, batchDenom)
	if err != nil {
		return err
	}
	if !retire {
		if err = k.addAndSaveBalance(ctx, owner, batch.Id, amount); err != nil {
			return err
		}

		return sdkCtx.EventManager().EmitTypedEvent(&ecocredit.EventReceive{
			Recipient:      owner.String(),
			BatchDenom:     batchDenom,
			TradableAmount: amount.String(),
			BasketDenom:    basketDenom,
		})
	} else {
		if err := k.retireAndSaveBalance(ctx, owner, batch.Id, amount); err != nil {
			return err
		}
		if err := k.retireSupply(ctx, batch.Id, amount); err != nil {
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

func (k Keeper) addAndSaveBalance(ctx context.Context, user sdk.AccAddress, batchId uint64, amount math.Dec) error {
	userBal, err := k.coreStore.BatchBalanceTable().Get(ctx, user, batchId)
	if err != nil {
		if ormerrors.IsNotFound(err) {
			userBal = &ecoApi.BatchBalance{
				Address:  user,
				BatchId:  batchId,
				Tradable: "0",
				Retired:  "0",
				Escrowed: "0",
			}
		} else {
			return err
		}
	}
	tradable, err := math.NewDecFromString(userBal.Tradable)
	if err != nil {
		return err
	}
	newTradable, err := math.SafeAddBalance(tradable, amount)
	if err != nil {
		return err
	}
	userBal.Tradable = newTradable.String()
	return k.coreStore.BatchBalanceTable().Save(ctx, userBal)
}

func (k Keeper) retireAndSaveBalance(ctx context.Context, user sdk.AccAddress, batchId uint64, amount math.Dec) error {
	userBal, err := k.coreStore.BatchBalanceTable().Get(ctx, user, batchId)
	if err != nil {
		if ormerrors.IsNotFound(err) {
			userBal = &ecoApi.BatchBalance{
				Address:  user,
				BatchId:  batchId,
				Tradable: "0",
				Retired:  "0",
				Escrowed: "0",
			}
		} else {
			return err
		}
	}
	retired, err := math.NewDecFromString(userBal.Retired)
	if err != nil {
		return err
	}
	newRetired, err := math.SafeAddBalance(retired, amount)
	if err != nil {
		return err
	}
	userBal.Retired = newRetired.String()
	return k.coreStore.BatchBalanceTable().Save(ctx, userBal)
}

func (k Keeper) retireSupply(ctx context.Context, batchId uint64, amount math.Dec) error {
	supply, err := k.coreStore.BatchSupplyTable().Get(ctx, batchId)
	if err != nil {
		return err
	}
	tradable, err := math.NewDecFromString(supply.TradableAmount)
	if err != nil {
		return err
	}

	retired, err := math.NewDecFromString(supply.RetiredAmount)
	if err != nil {
		return err
	}

	tradable, err = math.SafeSubBalance(tradable, amount)
	if err != nil {
		return err
	}

	retired, err = math.SafeAddBalance(retired, amount)
	if err != nil {
		return err
	}

	supply.TradableAmount = tradable.String()
	supply.RetiredAmount = retired.String()

	return k.coreStore.BatchSupplyTable().Update(ctx, supply)
}
