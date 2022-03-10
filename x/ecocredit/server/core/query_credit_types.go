package core

import (
	"context"
	ecocreditv1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	v1 "github.com/regen-network/regen-ledger/x/ecocredit/v1"
)

// CreditTypes queries the list of allowed types that credit classes can have.
func (k Keeper) CreditTypes(ctx context.Context, _ *v1.QueryCreditTypesRequest) (*v1.QueryCreditTypesResponse, error) {
	creditTypes := make([]*v1.CreditType, 0)
	it, err := k.stateStore.CreditTypeStore().List(ctx, ecocreditv1.CreditTypePrimaryKey{})
	if err != nil {
		return nil, err
	}
	defer it.Close()
	for it.Next() {
		ct, err := it.Value()
		if err != nil {
			return nil, err
		}
		var cType v1.CreditType
		if err = PulsarToGogoSlow(ct, &cType); err != nil {
			return nil, err
		}
		creditTypes = append(creditTypes, &cType)
	}
	return &v1.QueryCreditTypesResponse{CreditTypes: creditTypes}, nil
}
