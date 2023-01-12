package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	regenerrors "github.com/regen-network/regen-ledger/types/v2/errors"
	types "github.com/regen-network/regen-ledger/x/ecocredit/v3/base/types/v1"
)

// Project queries project info from the given project name.
func (k Keeper) Project(ctx context.Context, request *types.QueryProjectRequest) (*types.QueryProjectResponse, error) {
	project, err := k.stateStore.ProjectTable().GetById(ctx, request.ProjectId)
	if err != nil {
		return nil, regenerrors.ErrNotFound.Wrapf("could not get project with id %s: %s", request.ProjectId, err.Error())
	}

	admin := sdk.AccAddress(project.Admin)

	class, err := k.stateStore.ClassTable().Get(ctx, project.ClassKey)
	if err != nil {
		return nil, regenerrors.ErrNotFound.Wrapf("could not get class with key %d: %s", project.ClassKey, err.Error())
	}

	info := types.ProjectInfo{
		Id:           project.Id,
		Admin:        admin.String(),
		ClassId:      class.Id,
		Jurisdiction: project.Jurisdiction,
		Metadata:     project.Metadata,
		ReferenceId:  project.ReferenceId,
	}

	return &types.QueryProjectResponse{Project: &info}, nil
}
