package core

import (
	"context"

	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// CreditTypes queries the list of allowed types that credit classes can have.
func (k Keeper) CreditTypes(ctx context.Context, _ *core.QueryCreditTypesRequest) (*core.QueryCreditTypesResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	var params ecocredit.Params
	k.paramsKeeper.GetParamSet(sdkCtx, &params)

	cTypes := make([]*core.CreditType, len(params.CreditTypes))
	for i, ct := range params.CreditTypes {
		cTypes[i] = &core.CreditType{
			Abbreviation: ct.Abbreviation,
			Name:         ct.Name,
			Unit:         ct.Unit,
			Precision:    ct.Precision,
		}
	}
	return &core.QueryCreditTypesResponse{CreditTypes: cTypes}, nil
}
