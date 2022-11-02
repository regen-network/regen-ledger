package keeper

import (
	"context"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	regenerrors "github.com/regen-network/regen-ledger/errors"
	"github.com/regen-network/regen-ledger/types/ormutil"
	types "github.com/regen-network/regen-ledger/x/ecocredit/base/types/v1"
)

// CreditTypes queries the list of allowed types that credit classes can have.
func (k Keeper) CreditTypes(ctx context.Context, _ *types.QueryCreditTypesRequest) (*types.QueryCreditTypesResponse, error) {
	it, err := k.stateStore.CreditTypeTable().List(ctx, &api.CreditTypePrimaryKey{})
	if err != nil {
		return nil, err
	}
	defer it.Close()

	creditTypes := make([]*types.CreditType, 0)
	for it.Next() {
		ct, err := it.Value()
		if err != nil {
			return nil, err
		}
		var creditType types.CreditType
		if err := ormutil.PulsarToGogoSlow(ct, &creditType); err != nil {
			return nil, regenerrors.ErrInternal.Wrap(err.Error())
		}
		creditTypes = append(creditTypes, &creditType)
	}
	return &types.QueryCreditTypesResponse{CreditTypes: creditTypes}, nil
}
