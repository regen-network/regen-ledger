package core

import (
	"context"
	ecocreditv1beta1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1beta1"
	"github.com/regen-network/regen-ledger/x/ecocredit/v1beta1"
)

// CreditTypes queries the list of allowed types that credit classes can have.
func (k Keeper) CreditTypes(ctx context.Context, _ *v1beta1.QueryCreditTypesRequest) (*v1beta1.QueryCreditTypesResponse, error) {
	creditTypes := make([]*v1beta1.CreditType, 0)
	it, err := k.stateStore.CreditTypeStore().List(ctx, ecocreditv1beta1.CreditTypePrimaryKey{})
	if err != nil {
		return nil, err
	}
	for it.Next() {
		ct, err := it.Value()
		if err != nil {
			return nil, err
		}
		var cType v1beta1.CreditType
		if err = PulsarToGogoSlow(ct, &cType); err != nil {
			return nil, err
		}
		creditTypes = append(creditTypes, &cType)
	}
	return &v1beta1.QueryCreditTypesResponse{CreditTypes: creditTypes}, nil
}
