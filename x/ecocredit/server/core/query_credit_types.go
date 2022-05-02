package core

import (
	"context"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/types/ormutil"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

// CreditTypes queries the list of allowed types that credit classes can have.
func (k Keeper) CreditTypes(ctx context.Context, _ *core.QueryCreditTypesRequest) (*core.QueryCreditTypesResponse, error) {
	it, err := k.stateStore.CreditTypeTable().List(ctx, &api.CreditTypePrimaryKey{})
	if err != nil {
		return nil, err
	}
	defer it.Close()

	creditTypes := make([]*core.CreditType, 0)
	for it.Next() {
		ct, err := it.Value()
		if err != nil {
			return nil, err
		}
		var creditType core.CreditType
		if err := ormutil.PulsarToGogoSlow(ct, &creditType); err != nil {
			return nil, err
		}
		creditTypes = append(creditTypes, &creditType)
	}
	return &core.QueryCreditTypesResponse{CreditTypes: creditTypes}, nil
}
