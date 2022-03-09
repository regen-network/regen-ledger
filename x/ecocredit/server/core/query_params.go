package core

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	v1 "github.com/regen-network/regen-ledger/x/ecocredit/v1"
)

// Params queries the ecocredit module parameters.
// TODO: remove params https://github.com/regen-network/regen-ledger/issues/729
// Currently this is an ugly hack that grabs v1alpha types and converts them into v1beta types.
// will be gone with #729.
func (k Keeper) Params(ctx context.Context, _ *v1.QueryParamsRequest) (*v1.QueryParamsResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	var params ecocredit.Params
	k.params.GetParamSet(sdkCtx, &params)
	v1beta1types := make([]*v1.CreditType, len(params.CreditTypes))
	for i, typ := range params.CreditTypes {
		v1beta1types[i] = &v1.CreditType{
			Abbreviation: typ.Abbreviation,
			Name:         typ.Name,
			Unit:         typ.Unit,
			Precision:    typ.Precision,
		}
	}
	v1beta1Params := v1.Params{
		CreditClassFee:       params.CreditClassFee,
		AllowedClassCreators: params.AllowedClassCreators,
		AllowlistEnabled:     params.AllowlistEnabled,
		CreditTypes:          v1beta1types,
	}
	return &v1.QueryParamsResponse{Params: &v1beta1Params}, nil
}
