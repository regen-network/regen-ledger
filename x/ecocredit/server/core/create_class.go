package core

import (
	"context"
	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	ecocreditv1beta1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1beta1"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/regen-network/regen-ledger/x/ecocredit/v1beta1"
)

// TODO: Revisit this once we have proper gas fee framework.
// Tracking issues https://github.com/cosmos/cosmos-sdk/issues/9054, https://github.com/cosmos/cosmos-sdk/discussions/9072
const gasCostPerIteration = uint64(10)

func (k Keeper) CreateClass(goCtx context.Context, req *v1beta1.MsgCreateClass) (*v1beta1.MsgCreateClassResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(goCtx)
	// Charge the admin a fee to create the credit class
	adminAddress, err := sdk.AccAddressFromBech32(req.Admin)
	if err != nil {
		return nil, err
	}

	// TODO: remove params https://github.com/regen-network/regen-ledger/issues/729
	var params ecocredit.Params
	k.params.GetParamSet(sdkCtx, &params)
	if params.AllowlistEnabled && !k.isCreatorAllowListed(sdkCtx, params.AllowedClassCreators, adminAddress) {
		return nil, sdkerrors.ErrUnauthorized.Wrapf("%s is not allowed to create credit classes", adminAddress.String())
	}

	err = k.chargeCreditClassFee(sdkCtx, adminAddress)
	if err != nil {
		return nil, err
	}

	creditType, err := k.getCreditType(sdkCtx, req.CreditTypeName)
	if err != nil {
		return nil, err
	}

	var seq uint64 = 1
	classSeq, err := k.stateStore.ClassSequenceStore().Get(goCtx, creditType.Abbreviation)
	if err != nil {
		if !ormerrors.IsNotFound(err) {
			return nil, err
		}
	} else {
		seq = classSeq.NextClassId
	}
	if err = k.stateStore.ClassSequenceStore().Save(goCtx, &ecocreditv1beta1.ClassSequence{
		CreditType:  creditType.Abbreviation,
		NextClassId: seq + 1,
	}); err != nil {
		return nil, err
	}

	classID := ecocredit.FormatClassID(creditType.Abbreviation, seq)

	rowId, err := k.stateStore.ClassInfoStore().InsertReturningID(goCtx, &ecocreditv1beta1.ClassInfo{
		Name:       classID,
		Admin:      adminAddress,
		Metadata:   req.Metadata,
		CreditType: creditType.Abbreviation,
	})
	if err != nil {
		return nil, err
	}

	for _, issuer := range req.Issuers {
		issuer, _ := sdk.AccAddressFromBech32(issuer)
		if err = k.stateStore.ClassIssuerStore().Insert(goCtx, &ecocreditv1beta1.ClassIssuer{
			ClassId: rowId,
			Issuer:  issuer,
		}); err != nil {
			return nil, err
		}
	}

	err = sdkCtx.EventManager().EmitTypedEvent(&v1beta1.EventCreateClass{
		ClassId: classID,
		Admin:   req.Admin,
	})
	if err != nil {
		return nil, err
	}

	return &v1beta1.MsgCreateClassResponse{ClassId: classID}, nil
}

func (k Keeper) isCreatorAllowListed(ctx sdk.Context, allowlist []string, designer sdk.AccAddress) bool {
	for _, addr := range allowlist {
		ctx.GasMeter().ConsumeGas(gasCostPerIteration, "credit class creators allowlist")
		allowListedAddr, _ := sdk.AccAddressFromBech32(addr)
		if designer.Equals(allowListedAddr) {
			return true
		}
	}
	return false
}

func (k Keeper) chargeCreditClassFee(ctx sdk.Context, creatorAddr sdk.AccAddress) error {
	creditClassFee := k.getCreditClassFee(ctx)

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

func (k Keeper) getCreditClassFee(ctx sdk.Context) sdk.Coins {
	var params ecocredit.Params
	k.params.GetParamSet(ctx, &params)
	return params.CreditClassFee
}

func (k Keeper) getCreditType(ctx sdk.Context, creditTypeName string) (ecocreditv1beta1.CreditType, error) {
	creditTypes := k.getAllCreditTypes(ctx)
	creditTypeName = ecocredit.NormalizeCreditTypeName(creditTypeName)
	for _, creditType := range creditTypes {
		// credit type name's stored via params have enforcement on normalization, so we can be sure they will already
		// be normalized here.
		if creditType.Name == creditTypeName {
			return ecocreditv1beta1.CreditType{
				Abbreviation: creditType.Abbreviation,
				Name:         creditType.Name,
				Unit:         creditType.Unit,
				Precision:    creditType.Precision,
			}, nil
		}
	}
	return ecocreditv1beta1.CreditType{}, sdkerrors.ErrInvalidType.Wrapf("%s is not a valid credit type", creditTypeName)
}

func (k Keeper) getAllCreditTypes(ctx sdk.Context) []*ecocredit.CreditType {
	var params ecocredit.Params
	k.params.GetParamSet(ctx, &params)
	return params.CreditTypes
}
