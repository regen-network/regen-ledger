package keeper

import (
	"context"

	basev1beta1 "github.com/cosmos/cosmos-sdk/api/cosmos/base/v1beta1"
	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	regentypes "github.com/regen-network/regen-ledger/types/v2"
	"github.com/regen-network/regen-ledger/x/ecocredit/v3"
)

func (k Keeper) deductFee(ctx context.Context, payer sdk.AccAddress, fee *sdk.Coin, minFee *basev1beta1.Coin) error {
	if minFee == nil {
		return nil
	}

	requiredFee, ok := regentypes.ProtoCoinToCoin(minFee)
	if !ok {
		return sdkerrors.ErrInvalidType.Wrap("fee")
	}

	if requiredFee.IsZero() {
		return nil
	}

	// check if fee is empty
	if fee == nil {
		return sdkerrors.ErrInsufficientFee.Wrapf(
			"fee cannot be empty: must be %s", requiredFee,
		)
	}

	// check if fee is the correct denom
	if fee.Denom != requiredFee.Denom {
		return sdkerrors.ErrInsufficientFee.Wrapf(
			"fee must be %s, got %s", requiredFee, fee,
		)
	}

	// check if fee is greater than or equal to required fee
	if !fee.IsGTE(requiredFee) {
		return sdkerrors.ErrInsufficientFee.Wrapf(
			"fee must be %s, got %s", requiredFee, fee,
		)
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)
	// check payer balance against required fee
	payerBalance := k.bankKeeper.GetBalance(sdkCtx, payer, requiredFee.Denom)
	if payerBalance.IsNil() || payerBalance.IsLT(requiredFee) {
		return sdkerrors.ErrInsufficientFunds.Wrapf(
			"insufficient balance %s for bank denom %s", payerBalance.Amount, requiredFee.Denom,
		)
	}

	feeCoins := sdk.Coins{requiredFee}

	err := k.bankKeeper.SendCoinsFromAccountToModule(sdkCtx, payer, ecocredit.ModuleName, feeCoins)
	if err != nil {
		return err
	}

	if requiredFee.Denom == "uregen" {
		err = k.bankKeeper.BurnCoins(sdkCtx, ecocredit.ModuleName, feeCoins)
		if err != nil {
			return err
		}
	}

	return nil
}

func (k Keeper) deductCreateProjectFee(ctx context.Context, address sdk.AccAddress, fee *sdk.Coin) error {
	projectFee, err := k.stateStore.ProjectFeeTable().Get(ctx)
	if ormerrors.IsNotFound(err) {
		// no fee set, so no fee to deduct
		return nil
	}
	if err != nil {
		return err
	}

	return k.deductFee(ctx, address, fee, projectFee.Fee)
}
