package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	api "github.com/regen-network/regen-ledger/api/v2/regen/ecocredit/v1"
	regenerrors "github.com/regen-network/regen-ledger/types/v2/errors"
	"github.com/regen-network/regen-ledger/types/v2/ormutil"
	types "github.com/regen-network/regen-ledger/x/ecocredit/v4/base/types/v1"
)

// ClassesByAdmin queries for all classes with a specific admin address.
func (k Keeper) ClassesByAdmin(ctx context.Context, req *types.QueryClassesByAdminRequest) (*types.QueryClassesByAdminResponse, error) {
	admin, err := sdk.AccAddressFromBech32(req.Admin)
	if err != nil {
		return nil, regenerrors.ErrInvalidArgument.Wrapf("invalid admin: %s", err.Error())
	}

	pg := ormutil.PageReqToOrmPaginate(req.Pagination)
	it, err := k.stateStore.ClassTable().List(ctx, api.ClassAdminIndexKey{}.WithAdmin(admin), pg)
	if err != nil {
		return nil, err
	}
	defer it.Close()

	adminString := admin.String()
	classes := make([]*types.ClassInfo, 0)
	for it.Next() {
		class, err := it.Value()
		if err != nil {
			return nil, err
		}
		info := types.ClassInfo{
			Id:               class.Id,
			Admin:            adminString,
			Metadata:         class.Metadata,
			CreditTypeAbbrev: class.CreditTypeAbbrev,
		}
		classes = append(classes, &info)
	}

	pr := ormutil.PageResToCosmosTypes(it.PageResponse())
	return &types.QueryClassesByAdminResponse{Classes: classes, Pagination: pr}, nil
}
