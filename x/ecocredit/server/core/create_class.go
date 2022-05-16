package core

import (
	"context"

	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

// CreateClass creates a new class of ecocredit
//
// The admin is charged a fee for creating the class. This is controlled by
// the global parameter CreditClassFee, which can be updated through the
// governance process.
func (k Keeper) CreateClass(goCtx context.Context, req *core.MsgCreateClass) (*core.MsgCreateClassResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(goCtx)
	adminAddress, err := sdk.AccAddressFromBech32(req.Admin)
	if err != nil {
		return nil, err
	}

	if err := k.assertCanCreateClass(sdkCtx, adminAddress); err != nil {
		return nil, err
	}

	// TODO: remove params https://github.com/regen-network/regen-ledger/issues/729
	var fee sdk.Coins
	k.paramsKeeper.Get(sdkCtx, core.KeyCreditClassFee, &fee)

	feeAmt := fee.AmountOf(req.Fee.Denom)
	if feeAmt.IsZero() {
		return nil, sdkerrors.ErrInvalidRequest.Wrapf("%s is not allowed to be used in credit class fees", req.Fee.Denom)
	}
	if req.Fee.Amount.LT(feeAmt) {
		return nil, sdkerrors.ErrInsufficientFee.Wrapf("expected %v%s for fee, got %v", feeAmt, req.Fee.Denom, req.Fee)
	}

	// Charge the admin a fee to create the credit class
	err = k.chargeCreditClassFee(sdkCtx, adminAddress, sdk.Coins{sdk.Coin{Denom: req.Fee.Denom, Amount: feeAmt}})
	if err != nil {
		return nil, err
	}

	creditType, err := k.stateStore.CreditTypeTable().Get(goCtx, req.CreditTypeAbbrev)
	if err != nil {
		return nil, sdkerrors.ErrInvalidRequest.Wrapf("could not get credit type with abbreviation %s: %s", req.CreditTypeAbbrev, err.Error())
	}

	// default the sequence to 1 for the `not found` case.
	// will get overwritten by the actual sequence if it exists.
	var seq uint64 = 1
	classSeq, err := k.stateStore.ClassSequenceTable().Get(goCtx, creditType.Abbreviation)
	if err != nil {
		if !ormerrors.IsNotFound(err) {
			return nil, err
		}
	} else {
		seq = classSeq.NextSequence
	}
	if err = k.stateStore.ClassSequenceTable().Save(goCtx, &api.ClassSequence{
		CreditTypeAbbrev: creditType.Abbreviation,
		NextSequence:     seq + 1,
	}); err != nil {
		return nil, err
	}

	classID := core.FormatClassId(creditType.Abbreviation, seq)

	key, err := k.stateStore.ClassTable().InsertReturningID(goCtx, &api.Class{
		Id:               classID,
		Admin:            adminAddress,
		Metadata:         req.Metadata,
		CreditTypeAbbrev: creditType.Abbreviation,
	})
	if err != nil {
		return nil, err
	}

	for _, issuer := range req.Issuers {
		issuer, err := sdk.AccAddressFromBech32(issuer)
		if err != nil {
			return nil, err
		}
		if err = k.stateStore.ClassIssuerTable().Insert(goCtx, &api.ClassIssuer{
			ClassKey: key,
			Issuer:   issuer,
		}); err != nil {
			return nil, err
		}

		sdkCtx.GasMeter().ConsumeGas(ecocredit.GasCostPerIteration, "ecocredit/core/MsgCreateClass issuer iteration")
	}

	err = sdkCtx.EventManager().EmitTypedEvent(&core.EventCreateClass{
		ClassId: classID,
	})
	if err != nil {
		return nil, err
	}

	return &core.MsgCreateClassResponse{ClassId: classID}, nil
}

func (k Keeper) assertCanCreateClass(sdkCtx sdk.Context, adminAddress sdk.AccAddress) error {
	var allowListEnabled bool
	k.paramsKeeper.Get(sdkCtx, core.KeyAllowlistEnabled, &allowListEnabled)
	if allowListEnabled {
		var allowList []string
		k.paramsKeeper.Get(sdkCtx, core.KeyAllowedClassCreators, &allowList)
		if !k.isCreatorAllowListed(sdkCtx, allowList, adminAddress) {
			return sdkerrors.ErrUnauthorized.Wrapf("%s is not allowed to create credit classes", adminAddress.String())
		}
	}
	return nil
}

func (k Keeper) isCreatorAllowListed(ctx sdk.Context, allowlist []string, designer sdk.AccAddress) bool {
	for _, addr := range allowlist {
		allowListedAddr, _ := sdk.AccAddressFromBech32(addr)
		if designer.Equals(allowListedAddr) {
			return true
		}
		ctx.GasMeter().ConsumeGas(ecocredit.GasCostPerIteration, "ecocredit/core/MsgCreateClass address iteration")
	}
	return false
}

func (k Keeper) chargeCreditClassFee(ctx sdk.Context, creatorAddr sdk.AccAddress, creditClassFee sdk.Coins) error {
	// Move the fee to the ecocredit module's account
	err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, creatorAddr, ecocredit.ModuleName, creditClassFee)
	if err != nil {
		return err
	}

	// Burn the coins
	// TODO: Update this implementation based on the discussion at
	// https://github.com/regen-network/regen-ledger/issues/351
	err = k.bankKeeper.BurnCoins(ctx, ecocredit.ModuleName, creditClassFee)
	if err != nil {
		return err
	}

	return nil
}
