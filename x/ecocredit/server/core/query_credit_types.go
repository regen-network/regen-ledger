package core

import (
	"context"

	"github.com/regen-network/regen-ledger/x/ecocredit/core"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// CreditTypes queries the list of allowed types that credit classes can have.
func (k Keeper) CreditTypes(ctx context.Context, _ *core.QueryCreditTypesRequest) (*core.QueryCreditTypesResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	var creditTypes []*core.CreditType
	k.paramsKeeper.Get(sdkCtx, core.KeyCreditTypes, &creditTypes)
	return &core.QueryCreditTypesResponse{CreditTypes: creditTypes}, nil
}
