package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/orm/model/ormlist"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/marketplace/v1"
	"github.com/regen-network/regen-ledger/types/ormutil"
	types "github.com/regen-network/regen-ledger/x/ecocredit/marketplace/types/v1"
)

func (k Keeper) AllowedDenoms(ctx context.Context, req *types.QueryAllowedDenomsRequest) (*types.QueryAllowedDenomsResponse, error) {
	pg, err := ormutil.GogoPageReqToPulsarPageReq(req.Pagination)
	if err != nil {
		return nil, err
	}

	it, err := k.stateStore.AllowedDenomTable().List(ctx, &api.AllowedDenomPrimaryKey{}, ormlist.Paginate(pg))
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
	pr, err := ormutil.PulsarPageResToGogoPageRes(it.PageResponse())
	if err != nil {
		return nil, err
	}
	return &types.QueryAllowedDenomsResponse{AllowedDenoms: allowedDenoms, Pagination: pr}, nil
}
