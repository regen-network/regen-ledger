package core

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

// Params queries the ecocredit module parameters.
// TODO: remove params https://github.com/regen-network/regen-ledger/issues/729
// Currently this is an ugly hack that grabs v1alpha types and converts them into v1beta types.
// will be gone with #729.
func (k Keeper) Params(ctx context.Context, _ *core.QueryParamsRequest) (*core.QueryParamsResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	var params core.Params
	k.paramsKeeper.GetParamSet(sdkCtx, &params)
	return &core.QueryParamsResponse{Params: &params}, nil
}
