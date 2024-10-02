package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	api "github.com/regen-network/regen-ledger/api/v2/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/types/v2/ormutil"
	types "github.com/regen-network/regen-ledger/x/ecocredit/v3/base/types/v1"
)

// Classes queries for all credit classes with pagination.
func (k Keeper) Classes(ctx context.Context, req *types.QueryClassesRequest) (*types.QueryClassesResponse, error) {
	pg := ormutil.PageReqToOrmPaginate(req.Pagination)
	it, err := k.stateStore.ClassTable().List(ctx, &api.ClassPrimaryKey{}, pg)
	if err != nil {
		return nil, err
	}
	defer it.Close()

	classes := make([]*types.ClassInfo, 0)
	for it.Next() {
		class, err := it.Value()
		if err != nil {
			return nil, err
		}

		admin := sdk.AccAddress(class.Admin)
		info := types.ClassInfo{
			Id:               class.Id,
			Admin:            admin.String(),
			Metadata:         class.Metadata,
			CreditTypeAbbrev: class.CreditTypeAbbrev,
		}
		classes = append(classes, &info)
	}

	return &types.QueryClassesResponse{
		Classes:    classes,
		Pagination: ormutil.PageResToCosmosTypes(it.PageResponse()),
	}, err
}
