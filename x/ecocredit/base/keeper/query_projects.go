package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	api "github.com/regen-network/regen-ledger/api/v2/regen/ecocredit/v1"
	regenerrors "github.com/regen-network/regen-ledger/types/v2/errors"
	"github.com/regen-network/regen-ledger/types/v2/ormutil"
	types "github.com/regen-network/regen-ledger/x/ecocredit/v3/base/types/v1"
)

// Projects queries all projects.
func (k Keeper) Projects(ctx context.Context, req *types.QueryProjectsRequest) (*types.QueryProjectsResponse, error) {
	pg := ormutil.PageReqToOrmPaginate(req.Pagination)
	it, err := k.stateStore.ProjectTable().List(ctx, api.ProjectIdIndexKey{}, pg)
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
			return nil, regenerrors.ErrNotFound.Wrapf("class with key: %d", project.ClassKey)
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

	return &types.QueryProjectsResponse{
		Projects:   projects,
		Pagination: ormutil.PageResToCosmosTypes(it.PageResponse()),
	}, nil
}
