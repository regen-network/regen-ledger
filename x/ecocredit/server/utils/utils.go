package utils

import (
	"context"

	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	"github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"

	ecoApi "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/types/math"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

// GetCreditTypeFromBatchDenom extracts the classId from a batch denom string, then retrieves it from the params.
func GetCreditTypeFromBatchDenom(ctx context.Context, store ecoApi.StateStore, k ecocredit.ParamKeeper, denom string) (core.CreditType, error) {
	sdkCtx := types.UnwrapSDKContext(ctx)
	classId := core.GetClassIdFromBatchDenom(denom)
	classInfo, err := store.ClassTable().GetById(ctx, classId)
	if err != nil {
		return core.CreditType{}, err
	}
	p := &core.Params{}
	k.GetParamSet(sdkCtx, p)
	return GetCreditType(classInfo.CreditTypeAbbrev, p.CreditTypes)
}

// GetCreditType searches for a credit type that matches the given abbreviation within a credit type slice.
func GetCreditType(ctAbbrev string, creditTypes []*core.CreditType) (core.CreditType, error) {
	//creditTypeName = ecocredit.NormalizeCreditTypeName(creditTypeName)
	for _, creditType := range creditTypes {
		// credit type name's stored via params have enforcement on normalization, so we can be sure they will already
		// be normalized here.
		if creditType.Abbreviation == ctAbbrev {
			return *creditType, nil
		}
	}
	return core.CreditType{}, errors.ErrInvalidType.Wrapf("%s is not a valid credit type", ctAbbrev)
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

// GetBalance gets the balance from the account, returning a zero value balance if no balance is found.
// NOTE: the default value is not inserted into the balance table in the `not found` case. Calling Update when a default
// value is returned will cause an error. Save should be used when dealing with balances from this function.
func GetBalance(ctx context.Context, table ecoApi.BatchBalanceTable, addr types.AccAddress, key uint64) (*ecoApi.BatchBalance, error) {
	bal, err := table.Get(ctx, addr, key)
	if err != nil {
		if !ormerrors.IsNotFound(err) {
			return nil, err
		}
		bal = &ecoApi.BatchBalance{
			BatchKey: key,
			Address:  addr,
		}
	}
	return bal, nil
}
