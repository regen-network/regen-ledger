package core

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/regen-network/regen-ledger/types/math"
	"github.com/regen-network/regen-ledger/x/ecocredit"
)

// assertClassIssuer makes sure that the issuer is part of issuers of given classID.
// Returns ErrUnauthorized otherwise.
func (k Keeper) assertClassIssuer(goCtx context.Context, classID uint64, issuer string) error {
	addr, err := sdk.AccAddressFromBech32(issuer)
	if err != nil {
		return err
	}
	found, err := k.stateStore.ClassIssuerStore().Has(goCtx, classID, addr)
	if err != nil {
		return err
	}
	if !found {
		return sdkerrors.ErrUnauthorized.Wrapf("%s is not an issuer for the class", issuer)
	}
	return nil
}

// getCreditType searches for a credit type that matches the given abbreviation within a credit type slice.
func (k Keeper) getCreditType(ctAbbrev string, creditTypes []*ecocredit.CreditType) (ecocredit.CreditType, error) {
	//creditTypeName = ecocredit.NormalizeCreditTypeName(creditTypeName)
	for _, creditType := range creditTypes {
		// credit type name's stored via params have enforcement on normalization, so we can be sure they will already
		// be normalized here.
		if creditType.Abbreviation == ctAbbrev {
			return *creditType, nil
		}
	}
	return ecocredit.CreditType{}, sdkerrors.ErrInvalidType.Wrapf("%s is not a valid credit type", ctAbbrev)
}

// getCreditTypeFromBatchDenom extracts the classId from a batch denom string, then retrieves it from the params.
func (k Keeper) getCreditTypeFromBatchDenom(ctx context.Context, denom string) (ecocredit.CreditType, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	classId := ecocredit.GetClassIdFromBatchDenom(denom)
	classInfo, err := k.stateStore.ClassInfoStore().GetByName(ctx, classId)
	if err != nil {
		return ecocredit.CreditType{}, err
	}
	p := &ecocredit.Params{}
	k.params.GetParamSet(sdkCtx, p)
	return k.getCreditType(classInfo.CreditType, p.CreditTypes)
}

// getNonNegativeFixedDecs takes an arbitrary amount of decimal strings, and returns their corresponding fixed decimals
// in a slice.
func getNonNegativeFixedDecs(precision uint32, decimals ...string) ([]math.Dec, error) {
	decs := make([]math.Dec, len(decimals))
	for i, decimal := range decimals {
		dec, err := math.NewNonNegativeFixedDecFromString(decimal, precision)
		if err != nil {
			return nil, err
		}
		decs[i] = dec
	}
	return decs, nil
}
