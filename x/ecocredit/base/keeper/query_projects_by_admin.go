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

func (k Keeper) ProjectsByAdmin(ctx context.Context, req *types.QueryProjectsByAdminRequest) (*types.QueryProjectsByAdminResponse, error) {
	admin, err := sdk.AccAddressFromBech32(req.Admin)
	if err != nil {
		return nil, regenerrors.ErrInvalidArgument.Wrap(err.Error())
	}

	pg, err := ormutil.GogoPageReqToPulsarPageReq(req.Pagination)
	if err != nil {
		return nil, regenerrors.ErrInvalidArgument.Wrap(err.Error())
	}

	it, err := k.stateStore.ProjectTable().List(ctx, api.ProjectAdminIndexKey{}.WithAdmin(admin), ormlist.Paginate(pg))
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

		class, err := k.stateStore.ClassTable().Get(ctx, project.ClassKey)
		if err != nil {
			return nil, regenerrors.ErrNotFound.Wrapf("unable to get class with key: %d: %s", project.ClassKey, err.Error())
		}

		projects = append(projects, &types.ProjectInfo{
			Id:           project.Id,
			Admin:        req.Admin,
			ClassId:      class.Id,
			Jurisdiction: project.Jurisdiction,
			Metadata:     project.Metadata,
			ReferenceId:  project.ReferenceId,
		})
	}

	pr, err := ormutil.PulsarPageResToGogoPageRes(it.PageResponse())
	if err != nil {
		return nil, regenerrors.ErrInternal.Wrap(err.Error())
	}

	return &types.QueryProjectsByAdminResponse{Projects: projects, Pagination: pr}, nil
}
