package core

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

// Params queries the ecocredit module parameters.
// TODO: remove params https://github.com/regen-network/regen-ledger/issues/729
// Currently this is an ugly hack that grabs v1alpha types and converts them into v1beta types.
// will be gone with #729.
func (k Keeper) Params(ctx context.Context, _ *core.QueryParamsRequest) (*core.QueryParamsResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	var params ecocredit.Params
	k.paramsKeeper.GetParamSet(sdkCtx, &params)
	v1beta1types := make([]*core.CreditType, len(params.CreditTypes))
	for i, typ := range params.CreditTypes {
		v1beta1types[i] = &core.CreditType{
			Abbreviation: typ.Abbreviation,
			Name:         typ.Name,
			Unit:         typ.Unit,
			Precision:    typ.Precision,
		}
	}
	v1beta1Params := core.Params{
		CreditClassFee:       params.CreditClassFee,
		AllowedClassCreators: params.AllowedClassCreators,
		AllowlistEnabled:     params.AllowlistEnabled,
		CreditTypes:          v1beta1types,
	}
	return &core.QueryParamsResponse{Params: &v1beta1Params}, nil
}
