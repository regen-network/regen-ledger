package core

import (
	"context"

	"github.com/regen-network/regen-ledger/x/ecocredit/core"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// CreditTypes queries the list of allowed types that credit classes can have.
func (k Keeper) CreditTypes(ctx context.Context, _ *core.QueryCreditTypesRequest) (*core.QueryCreditTypesResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	var params core.Params
	k.paramsKeeper.GetParamSet(sdkCtx, &params)
	return &core.QueryCreditTypesResponse{CreditTypes: params.CreditTypes}, nil
}
