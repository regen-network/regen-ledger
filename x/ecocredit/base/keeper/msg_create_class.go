package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	regentypes "github.com/regen-network/regen-ledger/types/v2"
	"github.com/regen-network/regen-ledger/x/ecocredit/v3"
	types "github.com/regen-network/regen-ledger/x/ecocredit/v3/base/types/v1"
)

// CreateClass creates a new class of ecocredit
//
// The admin is charged a fee for creating the class. This is controlled by
// the global parameter CreditClassFee, which can be updated through the
// governance process.
func (k Keeper) CreateClass(goCtx context.Context, req *types.MsgCreateClass) (*types.MsgCreateClassResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(goCtx)
	adminAddress, err := sdk.AccAddressFromBech32(req.Admin)
	if err != nil {
		return nil, err
	}

	if err := k.assertCanCreateClass(goCtx, adminAddress); err != nil {
		return nil, err
	}

	classFee, err := k.stateStore.ClassFeeTable().Get(goCtx)
	if err != nil {
		return nil, err
	}

	// only check and charge fee if required fee is set
	if classFee.Fee != nil {
		requiredFee, ok := regentypes.ProtoCoinToCoin(classFee.Fee)
		if !ok {
			return nil, sdkerrors.ErrInvalidType.Wrap("class fee")
		}

		// check if fee is empty
		if req.Fee == nil {
			return nil, sdkerrors.ErrInsufficientFee.Wrapf(
				"fee cannot be empty: must be %s", requiredFee,
			)
		}

		// check if fee is the correct denom
		if req.Fee.Denom != requiredFee.Denom {
			return nil, sdkerrors.ErrInsufficientFee.Wrapf(
				"fee must be %s, got %s", requiredFee, req.Fee,
			)
		}

		// check if fee is greater than or equal to required fee
		if !req.Fee.IsGTE(requiredFee) {
			return nil, sdkerrors.ErrInsufficientFee.Wrapf(
				"fee must be %s, got %s", requiredFee, req.Fee,
			)
		}

		// check admin balance against required fee
		adminBalance := k.bankKeeper.GetBalance(sdkCtx, adminAddress, requiredFee.Denom)
		if adminBalance.IsNil() || adminBalance.IsLT(requiredFee) {
			return nil, sdkerrors.ErrInsufficientFunds.Wrapf(
				"insufficient balance %s for bank denom %s", adminBalance.Amount, requiredFee.Denom,
			)
		}

		// convert required fee to multiple coins for processing
		requiredFees := sdk.Coins{requiredFee}

		// send coins from account to module and then burn the coins
		err = k.chargeCreditClassFee(sdkCtx, adminAddress, requiredFees)
		if err != nil {
			return nil, err
		}
	}

	evt := &types.EventCreateClass{
		Admin:      req.Admin,
		CreditType: req.CreditTypeAbbrev,
		Metadata:   req.Metadata,
	}

	err = k.eventReducer.Emit(goCtx, evt)
	if err != nil {
		return nil, err
	}

	for _, issuer := range req.Issuers {
		err = k.eventReducer.Emit(goCtx, &types.EventUpdateClassIssuers{
			ClassId: evt.ClassId,
			Issuer:  issuer,
		})
		if err != nil {
			return nil, err
		}

		sdkCtx.GasMeter().ConsumeGas(ecocredit.GasCostPerIteration, "ecocredit/MsgCreateClass issuer iteration")
	}

	return &types.MsgCreateClassResponse{ClassId: evt.ClassId}, nil
}

func (k Keeper) assertCanCreateClass(ctx context.Context, adminAddress sdk.AccAddress) error {
	allowListEnabled, err := k.stateStore.ClassCreatorAllowlistTable().Get(ctx)
	if err != nil {
		return err
	}

	if allowListEnabled.Enabled {
		_, err := k.stateStore.AllowedClassCreatorTable().Get(ctx, adminAddress)
		if err != nil {
			if ormerrors.NotFound.Is(err) {
				return sdkerrors.ErrUnauthorized.Wrapf("%s is not allowed to create credit classes", adminAddress.String())
			}
			return err
		}
	}
	return nil
}

func (k Keeper) chargeCreditClassFee(ctx sdk.Context, creatorAddr sdk.AccAddress, creditClassFee sdk.Coins) error {
	// Move the fee to the ecocredit module's account
	err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, creatorAddr, ecocredit.ModuleName, creditClassFee)
	if err != nil {
		return err
	}

	err = k.bankKeeper.BurnCoins(ctx, ecocredit.ModuleName, creditClassFee)
	if err != nil {
		return err
	}

	return nil
}
