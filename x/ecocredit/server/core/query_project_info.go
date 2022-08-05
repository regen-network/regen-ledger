package core

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

// Project queries project info from the given project name.
func (k Keeper) Project(ctx context.Context, request *core.QueryProjectRequest) (*core.QueryProjectResponse, error) {
	project, err := k.stateStore.ProjectTable().GetById(ctx, request.ProjectId)
	if err != nil {
		return nil, sdkerrors.ErrInvalidRequest.Wrapf("could not get project with id %s: %s", request.ProjectId, err.Error())
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
		Jurisdiction: project.Jurisdiction,
		Metadata:     project.Metadata,
		ReferenceId:  project.ReferenceId,
	}

	return &core.QueryProjectResponse{Project: &info}, nil
}
