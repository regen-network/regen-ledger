package keeper

import (
	"context"
	"fmt"

	sdkMath "cosmossdk.io/math"

	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/basket/v1"
	"github.com/regen-network/regen-ledger/types/math"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	basekeeper "github.com/regen-network/regen-ledger/x/ecocredit/base/keeper"
	basetypes "github.com/regen-network/regen-ledger/x/ecocredit/base/types/v1"
	basketsub "github.com/regen-network/regen-ledger/x/ecocredit/basket"
	types "github.com/regen-network/regen-ledger/x/ecocredit/basket/types/v1"
)

func (k Keeper) Take(ctx context.Context, msg *types.MsgTake) (*types.MsgTakeResponse, error) {
	basket, err := k.stateStore.BasketTable().GetByBasketDenom(ctx, msg.BasketDenom)
	if err != nil {
		if ormerrors.IsNotFound(err) {
			return nil, sdkerrors.ErrNotFound.Wrapf("basket %s not found", msg.BasketDenom)
		}
		return nil, err
	}

	creditType, err := k.baseStore.CreditTypeTable().Get(ctx, basket.CreditTypeAbbrev)
	if err != nil {
		return nil, err
	}

	retire := msg.RetireOnTake
	if !basket.DisableAutoRetire && !retire {
		return nil, basketsub.ErrCantDisableRetire
	}

	amountBasketTokens, ok := sdkMath.NewIntFromString(msg.Amount)
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

	err = k.bankKeeper.SendCoinsFromAccountToModule(sdkCtx, acct, basketsub.BasketSubModuleName, basketCoins)
	if err != nil {
		return nil, err
	}

	err = k.bankKeeper.BurnCoins(sdkCtx, basketsub.BasketSubModuleName, basketCoins)
	if err != nil {
		return nil, err
	}

	amountBasketCreditsDec, err := math.NewDecFromString(msg.Amount)
	if err != nil {
		return nil, err
	}

	multiplier := math.NewDecFinite(1, int32(creditType.Precision))
	amountCreditsNeeded, err := amountBasketCreditsDec.QuoExact(multiplier)
	if err != nil {
		return nil, err
	}

	var credits []*types.BasketCredit
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

		// retirement_location is deprecated but still supported
		var retirementJurisdiction string
		if len(msg.RetirementJurisdiction) != 0 {
			retirementJurisdiction = msg.RetirementJurisdiction
		} else {
			retirementJurisdiction = msg.RetirementLocation
		}

		cmp := balance.Cmp(amountCreditsNeeded)
		if cmp > 0 {
			credits = append(credits, &types.BasketCredit{
				BatchDenom: basketBalance.BatchDenom,
				Amount:     amountCreditsNeeded.String(),
			})

			err = k.addCreditBalance(
				ctx,
				acct,
				basketBalance.BatchDenom,
				amountCreditsNeeded,
				retire,
				retirementJurisdiction,
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
			credits = append(credits, &types.BasketCredit{
				BatchDenom: basketBalance.BatchDenom,
				Amount:     balance.String(),
			})

			err = k.addCreditBalance(
				ctx,
				acct,
				basketBalance.BatchDenom,
				balance,
				retire,
				retirementJurisdiction,
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

	err = sdk.UnwrapSDKContext(ctx).EventManager().EmitTypedEvent(&types.EventTake{
		Owner:       msg.Owner,
		BasketDenom: msg.BasketDenom,
		Credits:     credits,    // deprecated
		Amount:      msg.Amount, // deprecated
	})
	return &types.MsgTakeResponse{
		Credits: credits,
	}, err
}

func (k Keeper) addCreditBalance(ctx context.Context, owner sdk.AccAddress, batchDenom string, amount math.Dec, retire bool, jurisdiction string) error {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	batch, err := k.baseStore.BatchTable().GetByDenom(ctx, batchDenom)
	if err != nil {
		return err
	}
	if !retire {
		if err = basekeeper.AddAndSaveBalance(ctx, k.baseStore.BatchBalanceTable(), owner, batch.Key, amount); err != nil {
			return err
		}
		return sdkCtx.EventManager().EmitTypedEvent(&basetypes.EventTransfer{
			Sender:         k.moduleAddress.String(), // basket submodule
			Recipient:      owner.String(),
			BatchDenom:     batchDenom,
			TradableAmount: amount.String(),
		})
	}

	if err = basekeeper.RetireAndSaveBalance(ctx, k.baseStore.BatchBalanceTable(), owner, batch.Key, amount); err != nil {
		return err
	}
	if err = basekeeper.RetireSupply(ctx, k.baseStore.BatchSupplyTable(), batch.Key, amount); err != nil {
		return err
	}
	err = sdkCtx.EventManager().EmitTypedEvent(&basetypes.EventTransfer{
		Sender:        k.moduleAddress.String(), // basket submodule
		Recipient:     owner.String(),
		BatchDenom:    batchDenom,
		RetiredAmount: amount.String(),
	})
	if err != nil {
		return err
	}
	return sdkCtx.EventManager().EmitTypedEvent(&basetypes.EventRetire{
		Owner:        owner.String(),
		BatchDenom:   batchDenom,
		Amount:       amount.String(),
		Jurisdiction: jurisdiction,
	})
}
