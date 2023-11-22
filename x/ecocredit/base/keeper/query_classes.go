package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/orm/model/ormlist"
	sdk "github.com/cosmos/cosmos-sdk/types"

	api "github.com/regen-network/regen-ledger/api/v2/regen/ecocredit/v1"
	regenerrors "github.com/regen-network/regen-ledger/types/v2/errors"
	"github.com/regen-network/regen-ledger/types/v2/ormutil"
	types "github.com/regen-network/regen-ledger/x/ecocredit/v3/base/types/v1"
)

// Classes queries for all credit classes with pagination.
func (k Keeper) Classes(ctx context.Context, request *types.QueryClassesRequest) (*types.QueryClassesResponse, error) {
	pg, err := ormutil.GogoPageReqToPulsarPageReq(request.Pagination)
	if err != nil {
		return nil, regenerrors.ErrInvalidArgument.Wrap(err.Error())
	}

	it, err := k.stateStore.ClassTable().List(ctx, &api.ClassPrimaryKey{}, ormlist.Paginate(pg))
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

	pr, err := ormutil.PulsarPageResToGogoPageRes(it.PageResponse())
	if err != nil {
		return nil, regenerrors.ErrInternal.Wrap(err.Error())
	}

	return &types.QueryClassesResponse{
		Classes:    classes,
		Pagination: pr,
	}, err
}
