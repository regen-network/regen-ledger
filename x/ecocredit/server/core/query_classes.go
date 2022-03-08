package core

import (
	"context"
	"github.com/cosmos/cosmos-sdk/orm/model/ormlist"
	ecocreditv1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	v1 "github.com/regen-network/regen-ledger/x/ecocredit/v1"
)

// Classes queries for all credit classes with pagination.
func (k Keeper) Classes(ctx context.Context, request *v1.QueryClassesRequest) (*v1.QueryClassesResponse, error) {
	pg, err := GogoPageReqToPulsarPageReq(request.Pagination)
	if err != nil {
		return nil, err
	}
	it, err := k.stateStore.ClassInfoStore().List(ctx, &ecocreditv1.ClassInfoPrimaryKey{}, ormlist.Paginate(pg))
	if err != nil {
		return nil, err
	}

	infos := make([]*v1.ClassInfo, 0)
	for it.Next() {
		info, err := it.Value()
		if err != nil {
			return nil, err
		}
		var ci v1.ClassInfo
		if err = PulsarToGogoSlow(info, &ci); err != nil {
			return nil, err
		}
		infos = append(infos, &ci)
	}
	pr, err := PulsarPageResToGogoPageRes(it.PageResponse())
	if err != nil {
		return nil, err
	}
	return &v1.QueryClassesResponse{
		Classes:    infos,
		Pagination: pr,
	}, err
}
