package core

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

// UpdateProjectAdmin updates the project admin.
func (k Keeper) UpdateProjectAdmin(ctx context.Context, req *core.MsgUpdateProjectAdmin) (*core.MsgUpdateProjectAdminResponse, error) {
	admin, err := sdk.AccAddressFromBech32(req.Admin)
	if err != nil {
		return nil, err
	}

	newAdmin, err := sdk.AccAddressFromBech32(req.NewAdmin)
	if err != nil {
		return nil, err
	}

	project, err := k.stateStore.ProjectTable().GetById(ctx, req.ProjectId)
	if err != nil {
		return nil, sdkerrors.ErrInvalidRequest.Wrapf("could not get project with id %s: %s", req.ProjectId, err.Error())
	}
	if !sdk.AccAddress(project.Admin).Equals(admin) {
		return nil, sdkerrors.ErrUnauthorized.Wrapf("%s is not the admin of project %s", req.Admin, req.ProjectId)
	}
	project.Admin = newAdmin
	if err := k.stateStore.ProjectTable().Update(ctx, project); err != nil {
		return nil, err
	}

	if err := sdk.UnwrapSDKContext(ctx).EventManager().EmitTypedEvent(&core.EventProjectAdminUpdated{
		ProjectId: project.Id,
		OldAdmin:  req.Admin,
		NewAdmin:  req.NewAdmin,
	}); err != nil {
		return nil, err
	}

	return &core.MsgUpdateProjectAdminResponse{}, nil
}
