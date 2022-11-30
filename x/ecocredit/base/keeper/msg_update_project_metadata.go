package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	types "github.com/regen-network/regen-ledger/x/ecocredit/v3/base/types/v1"
)

// UpdateProjectMetadata updates the project metadata.
func (k Keeper) UpdateProjectMetadata(ctx context.Context, req *types.MsgUpdateProjectMetadata) (*types.MsgUpdateProjectMetadataResponse, error) {
	admin, err := sdk.AccAddressFromBech32(req.Admin)
	if err != nil {
		return nil, err
	}
	project, err := k.stateStore.ProjectTable().GetById(ctx, req.ProjectId)
	if err != nil {
		return nil, sdkerrors.ErrInvalidRequest.Wrapf(
			"could not get project with id %s: %s", req.ProjectId, err,
		)
	}
	if !sdk.AccAddress(project.Admin).Equals(admin) {
		return nil, sdkerrors.ErrUnauthorized.Wrapf(
			"%s is not the admin of project %s", req.Admin, req.ProjectId,
		)
	}
	project.Metadata = req.NewMetadata
	if err := k.stateStore.ProjectTable().Update(ctx, project); err != nil {
		return nil, err
	}

	if err := sdk.UnwrapSDKContext(ctx).EventManager().EmitTypedEvent(&types.EventUpdateProjectMetadata{
		ProjectId: project.Id,
	}); err != nil {
		return nil, err
	}

	return &types.MsgUpdateProjectMetadataResponse{}, nil
}
