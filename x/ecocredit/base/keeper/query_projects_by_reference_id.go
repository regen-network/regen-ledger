package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/orm/model/ormlist"
	sdk "github.com/cosmos/cosmos-sdk/types"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	regenerrors "github.com/regen-network/regen-ledger/errors"
	"github.com/regen-network/regen-ledger/types/ormutil"
	types "github.com/regen-network/regen-ledger/x/ecocredit/base/types/v1"
)

// ProjectsByReferenceId queries projects by reference id.
func (k Keeper) ProjectsByReferenceId(ctx context.Context, req *types.QueryProjectsByReferenceIdRequest) (*types.QueryProjectsByReferenceIdResponse, error) { //nolint:revive,stylecheck
	if req.ReferenceId == "" {
		return nil, regenerrors.ErrInvalidArgument.Wrap("reference-id is empty")
	}

	pg, err := ormutil.GogoPageReqToPulsarPageReq(req.Pagination)
	if err != nil {
		return nil, regenerrors.ErrInvalidArgument.Wrap(err.Error())
	}

	it, err := k.stateStore.ProjectTable().List(ctx, api.ProjectReferenceIdIndexKey{}.WithReferenceId(req.ReferenceId), ormlist.Paginate(pg))
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
			return nil, regenerrors.ErrNotFound.Wrapf("class with key: %d", project.ClassKey)
		}

		info := &types.ProjectInfo{
			Id:           project.Id,
			Admin:        sdk.AccAddress(project.Admin).String(),
			ClassId:      class.Id,
			Jurisdiction: project.Jurisdiction,
			Metadata:     project.Metadata,
			ReferenceId:  project.ReferenceId,
		}

		projects = append(projects, info)
	}

	pr, err := ormutil.PulsarPageResToGogoPageRes(it.PageResponse())
	if err != nil {
		return nil, regenerrors.ErrInternal.Wrap(err.Error())
	}

	return &types.QueryProjectsByReferenceIdResponse{
		Projects:   projects,
		Pagination: pr,
	}, nil
}
