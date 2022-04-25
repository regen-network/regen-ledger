package core

import (
	"context"

	"github.com/cosmos/cosmos-sdk/orm/model/ormlist"
	sdk "github.com/cosmos/cosmos-sdk/types"

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

	classes := make([]*core.ClassInfo, 0)
	for it.Next() {
		class, err := it.Value()
		if err != nil {
			return nil, err
		}

		admin := sdk.AccAddress(class.Admin)

		info := core.ClassInfo{
			Id:               class.Id,
			Admin:            admin.String(),
			Metadata:         class.Metadata,
			CreditTypeAbbrev: class.CreditTypeAbbrev,
		}

		classes = append(classes, &info)
	}

	pr, err := ormutil.PulsarPageResToGogoPageRes(it.PageResponse())
	if err != nil {
		return nil, err
	}

	return &core.QueryClassesResponse{
		Classes:    classes,
		Pagination: pr,
	}, err
}
