package core

import (
	"context"

	"github.com/cosmos/cosmos-sdk/orm/model/ormlist"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/types/ormutil"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

// Classes queries for all credit classes with pagination.
func (k Keeper) Classes(ctx context.Context, request *core.QueryClassesRequest) (*core.QueryClassesResponse, error) {
	pg, err := ormutil.GogoPageReqToPulsarPageReq(request.Pagination)
	if err != nil {
		return nil, err
	}
	it, err := k.stateStore.ClassTable().List(ctx, &api.ClassPrimaryKey{}, ormlist.Paginate(pg))
	if err != nil {
		return nil, err
	}

	infos := make([]*core.Class, 0)
	for it.Next() {
		info, err := it.Value()
		if err != nil {
			return nil, err
		}

		var ci core.Class
		if err = ormutil.PulsarToGogoSlow(info, &ci); err != nil {
			return nil, err
		}
		infos = append(infos, &ci)
	}
	pr, err := ormutil.PulsarPageResToGogoPageRes(it.PageResponse())
	if err != nil {
		return nil, err
	}
	return &core.QueryClassesResponse{
		Classes:    infos,
		Pagination: pr,
	}, err
}
