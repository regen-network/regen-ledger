package keeper

import (
	"context"

	types "github.com/regen-network/regen-ledger/x/ecocredit/base/types/v1"
)

// CreditType queries credit type information by abbreviation.
func (k Keeper) CreditType(ctx context.Context, request *types.QueryCreditTypeRequest) (*types.QueryCreditTypeResponse, error) {
	creditType, err := k.stateStore.CreditTypeTable().Get(ctx, request.Abbreviation)
	if err != nil {
		return nil, err
	}

	return &types.QueryCreditTypeResponse{
		CreditType: &types.CreditType{
			Abbreviation: creditType.Abbreviation,
			Name:         creditType.Name,
			Unit:         creditType.Unit,
			Precision:    creditType.Precision,
		},
	}, nil
}
