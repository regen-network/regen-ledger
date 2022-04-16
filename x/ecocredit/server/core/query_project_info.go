package core

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

// ProjectInfo queries project info from the given project name.
func (k Keeper) ProjectInfo(ctx context.Context, request *core.QueryProjectInfoRequest) (*core.QueryProjectInfoResponse, error) {
	project, err := k.stateStore.ProjectInfoTable().GetByName(ctx, request.ProjectId)
	if err != nil {
		return nil, err
	}

	admin := sdk.AccAddress(project.Admin)

	class, err := k.stateStore.ClassInfoTable().Get(ctx, project.ClassId)
	if err != nil {
		return nil, err
	}

	entry := core.ProjectInfoEntry{
		Id:              project.Name,
		Admin:           admin.String(),
		ClassId:         class.Name,
		ProjectLocation: project.ProjectLocation,
		Metadata:        project.Metadata,
	}

	return &core.QueryProjectInfoResponse{Project: &entry}, nil
}
