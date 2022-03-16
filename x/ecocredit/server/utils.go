package server

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/params/types"
	ecocreditv1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/types/math"
	"github.com/regen-network/regen-ledger/x/ecocredit"
)

type ParamKeeper interface {
	GetParamSet(ctx sdk.Context, ps types.ParamSet)
}



// GetCreditTypeFromBatchDenom extracts the classId from a batch denom string, then retrieves it from the params.
func GetCreditTypeFromBatchDenom(ctx context.Context, store ecocreditv1.StateStore, k ParamKeeper, denom string) (ecocredit.CreditType, error) {
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
