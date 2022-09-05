package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/orm/model/ormlist"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/types/ormutil"
	types "github.com/regen-network/regen-ledger/x/ecocredit/base/types/v1"
)

// ProjectsByClass queries all projects from a given credit class.
func (k Keeper) ProjectsByClass(ctx context.Context, request *types.QueryProjectsByClassRequest) (*types.QueryProjectsByClassResponse, error) {
	pg, err := ormutil.GogoPageReqToPulsarPageReq(request.Pagination)
	if err != nil {
		return nil, err
	}

	cInfo, err := k.stateStore.ClassTable().GetById(ctx, request.ClassId)
	if err != nil {
		return nil, sdkerrors.ErrInvalidRequest.Wrapf("could not get class with id %s: %s", request.ClassId, err.Error())
	}

	it, err := k.stateStore.ProjectTable().List(ctx, api.ProjectClassKeyIdIndexKey{}.WithClassKey(cInfo.Key), ormlist.Paginate(pg))
	if err != nil {
		return nil, err
	}
	defer it.Close()

	projects := make([]*types.ProjectInfo, 0)
	for it.Next() {
		project, err := it.Value()
		if err != nil {
			return nil, err
		}

		admin := sdk.AccAddress(project.Admin)

		class, err := k.stateStore.ClassTable().Get(ctx, project.ClassKey)
		if err != nil {
			return nil, err
		}

		info := types.ProjectInfo{
			Id:           project.Id,
			Admin:        admin.String(),
			ClassId:      class.Id,
			Jurisdiction: project.Jurisdiction,
			Metadata:     project.Metadata,
			ReferenceId:  project.ReferenceId,
		}

		projects = append(projects, &info)
	}

	pr, err := ormutil.PulsarPageResToGogoPageRes(it.PageResponse())
	if err != nil {
		return nil, err
	}

	return &types.QueryProjectsByClassResponse{
		Projects:   projects,
		Pagination: pr,
	}, nil
}
