package basket

import (
	"context"
	"fmt"

	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/basket/v1"
	"github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/types/math"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	baskettypes "github.com/regen-network/regen-ledger/x/ecocredit/basket"
	coretypes "github.com/regen-network/regen-ledger/x/ecocredit/core"
	"github.com/regen-network/regen-ledger/x/ecocredit/server/core"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k Keeper) Take(ctx context.Context, msg *baskettypes.MsgTake) (*baskettypes.MsgTakeResponse, error) {
	basket, err := k.stateStore.BasketTable().GetByBasketDenom(ctx, msg.BasketDenom)
	if err != nil {
		if ormerrors.IsNotFound(err) {
			return nil, sdkerrors.ErrNotFound.Wrapf("basket %s not found", msg.BasketDenom)
		}
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

	ownerBalance := k.bankKeeper.GetBalance(sdkCtx, acct, basket.BasketDenom)
	if ownerBalance.IsNil() || ownerBalance.IsLT(basketCoins[0]) {
		return nil, sdkerrors.ErrInsufficientFunds.Wrapf("insufficient balance for basket denom %s", basket.BasketDenom)
	}

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
				msg.RetirementJurisdiction,
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
				msg.RetirementJurisdiction,
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
		Credits:     credits,    // deprecated
		Amount:      msg.Amount, // deprecated
	})
	return &baskettypes.MsgTakeResponse{
		Credits: credits,
	}, err
}

func (k Keeper) addCreditBalance(ctx context.Context, owner sdk.AccAddress, batchDenom string, amount math.Dec, basketDenom string, retire bool, retirementJurisdiction string) error {
	sdkCtx := types.UnwrapSDKContext(ctx)
	batch, err := k.coreStore.BatchTable().GetByDenom(ctx, batchDenom)
	if err != nil {
		return err
	}
	if !retire {
		if err = core.AddAndSaveBalance(ctx, k.coreStore.BatchBalanceTable(), owner, batch.Key, amount); err != nil {
			return err
		}
		return sdkCtx.EventManager().EmitTypedEvent(&coretypes.EventTransfer{
			Sender:         k.moduleAddress.String(), // basket submodule
			Recipient:      owner.String(),
			BatchDenom:     batchDenom,
			TradableAmount: amount.String(),
		})
	} else {
		if err = core.RetireAndSaveBalance(ctx, k.coreStore.BatchBalanceTable(), owner, batch.Key, amount); err != nil {
			return err
		}
		if err = core.RetireSupply(ctx, k.coreStore.BatchSupplyTable(), batch.Key, amount); err != nil {
			return err
		}
		err = sdkCtx.EventManager().EmitTypedEvent(&coretypes.EventTransfer{
			Sender:        k.moduleAddress.String(), // basket submodule
			Recipient:     owner.String(),
			BatchDenom:    batchDenom,
			RetiredAmount: amount.String(),
		})
		if err != nil {
			return err
		}
		return sdkCtx.EventManager().EmitTypedEvent(&coretypes.EventRetire{
			Owner:        owner.String(),
			BatchDenom:   batchDenom,
			Amount:       amount.String(),
			Jurisdiction: retirementJurisdiction,
		})
	}
}
