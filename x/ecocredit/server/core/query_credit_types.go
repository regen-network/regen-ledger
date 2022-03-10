package core

import (
	"context"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

// CreditTypes queries the list of allowed types that credit classes can have.
func (k Keeper) CreditTypes(ctx context.Context, _ *core.QueryCreditTypesRequest) (*core.QueryCreditTypesResponse, error) {
	creditTypes := make([]*core.CreditType, 0)
	it, err := k.stateStore.CreditTypeStore().List(ctx, api.CreditTypePrimaryKey{})
	if err != nil {
		return nil, err
	}
	defer it.Close()
	for it.Next() {
		ct, err := it.Value()
		if err != nil {
			return nil, err
		}
		var cType core.CreditType
		if err = PulsarToGogoSlow(ct, &cType); err != nil {
			return nil, err
		}
		creditTypes = append(creditTypes, &cType)
	}
	return &core.QueryCreditTypesResponse{CreditTypes: creditTypes}, nil
}
