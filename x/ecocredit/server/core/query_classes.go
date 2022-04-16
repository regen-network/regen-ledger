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

	it, err := k.stateStore.ClassInfoTable().List(ctx, &api.ClassInfoPrimaryKey{}, ormlist.Paginate(pg))
	if err != nil {
		return nil, err
	}

	classes := make([]*core.ClassInfoEntry, 0)
	for it.Next() {
		info, err := it.Value()
		if err != nil {
			return nil, err
		}

		admin := sdk.AccAddress(info.Admin)

		entry := core.ClassInfoEntry{
			Id:               info.Name,
			Admin:            admin.String(),
			Metadata:         info.Metadata,
			CreditTypeAbbrev: info.CreditType,
		}

		classes = append(classes, &entry)
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
