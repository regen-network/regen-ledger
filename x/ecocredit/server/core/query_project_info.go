package core

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

// ProjectInfo queries project info from the given project name.
func (k Keeper) ProjectInfo(ctx context.Context, request *core.QueryProjectInfoRequest) (*core.QueryProjectInfoResponse, error) {
	project, err := k.stateStore.ProjectTable().GetById(ctx, request.ProjectId)
	if err != nil {
		return nil, err
	}

	admin := sdk.AccAddress(project.Admin)

	class, err := k.stateStore.ClassTable().Get(ctx, project.ClassKey)
	if err != nil {
		return nil, err
	}

	info := core.ProjectInfo{
		Id:           project.Id,
		Admin:        admin.String(),
		ClassId:      class.Id,
		Jurisdiction: project.ProjectJurisdiction,
		Metadata:     project.Metadata,
	}

	return &core.QueryProjectInfoResponse{Project: &info}, nil
}
