package server

import (
	"context"

	ecocreditAPI "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/types/math"
	"github.com/regen-network/regen-ledger/x/ecocredit"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

// GetCreditType searches for a credit type that matches the given abbreviation within a credit type slice.
func GetCreditType(ctAbbrev string, creditTypes []*ecocredit.CreditType) (ecocredit.CreditType, error) {
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

// GetCreditTypeFromBatchDenom extracts the classId from a batch denom string, then retrieves it from the params.
func GetCreditTypeFromBatchDenom(ctx context.Context, store ecocreditAPI.StateStore, k ecocredit.ParamKeeper, denom string) (ecocredit.CreditType, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	classId := ecocredit.GetClassIdFromBatchDenom(denom)
	classInfo, err := store.ClassInfoTable().GetByName(ctx, classId)
	if err != nil {
		return ecocredit.CreditType{}, err
	}
	p := &ecocredit.Params{}
	k.GetParamSet(sdkCtx, p)
	return GetCreditType(classInfo.CreditType, p.CreditTypes)
}

// GetNonNegativeFixedDecs takes an arbitrary amount of decimal strings, and returns their corresponding fixed decimals
// in a slice.
func GetNonNegativeFixedDecs(precision uint32, decimals ...string) ([]math.Dec, error) {
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

// AssertGovernance asserts the address is equal to the governance module address
func AssertGovernance(addr sdk.AccAddress, k ecocredit.AccountKeeper) error {
	if !k.GetModuleAddress(govtypes.ModuleName).Equals(addr) {
		return sdkerrors.ErrUnauthorized.Wrapf("params can only be updated via governance")
	}
	return nil
}
