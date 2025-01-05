package keeper

import (
	"context"

	api "github.com/regen-network/regen-ledger/api/v2/regen/ecocredit/marketplace/v1"
	regenerrors "github.com/regen-network/regen-ledger/types/v2/errors"
	"github.com/regen-network/regen-ledger/types/v2/ormutil"
	types "github.com/regen-network/regen-ledger/x/ecocredit/v3/marketplace/types/v1"
)

func (k Keeper) AllowedDenoms(ctx context.Context, req *types.QueryAllowedDenomsRequest) (*types.QueryAllowedDenomsResponse, error) {
	if req == nil {
		return nil, regenerrors.ErrInvalidArgument.Wrap("empty request")
	}

	pg := ormutil.PageReqToOrmPaginate(req.Pagination)
	it, err := k.stateStore.AllowedDenomTable().List(ctx, &api.AllowedDenomPrimaryKey{}, pg)
	if err != nil {
		return nil, err
	}
	defer it.Close()

	allowedDenoms := make([]*types.AllowedDenom, 0)
	for it.Next() {
		allowedDenom, err := it.Value()
		if err != nil {
			return nil, err
		}
		var ad types.AllowedDenom
		if err = ormutil.PulsarToGogoSlow(allowedDenom, &ad); err != nil {
			return nil, err
		}
		allowedDenoms = append(allowedDenoms, &ad)
	}
	pr := ormutil.PageResToCosmosTypes(it.PageResponse())
	return &types.QueryAllowedDenomsResponse{AllowedDenoms: allowedDenoms, Pagination: pr}, nil
}
