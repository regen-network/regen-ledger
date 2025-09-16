package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	types "github.com/regen-network/regen-ledger/x/ecocredit/v4/base/types/v1"
)

// UpdateProjectAdmin updates the project admin.
func (k Keeper) UpdateProjectAdmin(ctx context.Context, req *types.MsgUpdateProjectAdmin) (*types.MsgUpdateProjectAdminResponse, error) {
	if err := req.ValidateBasic(); err != nil {
		return nil, err
	}

	adminBz, err := k.ac.StringToBytes(req.Admin)
	if err != nil {
		return nil, sdkerrors.ErrInvalidAddress.Wrapf("admin: %s", err)
	}

	adminAddr := sdk.AccAddress(adminBz)
	newadminBz, err := k.ac.StringToBytes(req.NewAdmin)
	if err != nil {
		return nil, sdkerrors.ErrInvalidAddress.Wrapf("new admin: %s", err)
	}

	project, err := k.stateStore.ProjectTable().GetById(ctx, req.ProjectId)
	if err != nil {
		return nil, sdkerrors.ErrInvalidRequest.Wrapf(
			"could not get project with id %s: %s", req.ProjectId, err,
		)
	}
	if !sdk.AccAddress(project.Admin).Equals(adminAddr) {
		return nil, sdkerrors.ErrUnauthorized.Wrapf(
			"%s is not the admin of project %s", req.Admin, req.ProjectId,
		)
	}
	project.Admin = newadminBz
	if err := k.stateStore.ProjectTable().Update(ctx, project); err != nil {
		return nil, err
	}

	if err := sdk.UnwrapSDKContext(ctx).EventManager().EmitTypedEvent(&types.EventUpdateProjectAdmin{
		ProjectId: project.Id,
	}); err != nil {
		return nil, err
	}

	return &types.MsgUpdateProjectAdminResponse{}, nil
}
