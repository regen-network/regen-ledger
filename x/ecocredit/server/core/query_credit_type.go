package core

import (
	"context"

	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

// CreditType queries credit type information by abbreviation.
func (k Keeper) CreditType(ctx context.Context, request *core.QueryCreditTypeRequest) (*core.QueryCreditTypeResponse, error) {
	creditType, err := k.stateStore.CreditTypeTable().Get(ctx, request.Abbreviation)
	if err != nil {
		return nil, err
	}

	return &core.QueryCreditTypeResponse{
		CreditType: &core.CreditType{
			Abbreviation: creditType.Abbreviation,
			Name:         creditType.Name,
			Unit:         creditType.Unit,
			Precision:    creditType.Precision,
		},
	}, nil
}
