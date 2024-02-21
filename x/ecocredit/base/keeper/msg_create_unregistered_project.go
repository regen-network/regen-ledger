package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	types "github.com/regen-network/regen-ledger/x/ecocredit/v3/base/types/v1"
)

func (k Keeper) CreateUnregisteredProject(ctx context.Context, msg *types.MsgCreateUnregisteredProject) (*types.MsgCreateUnregisteredProjectResponse, error) {
	admin, err := sdk.AccAddressFromBech32(msg.Admin)
	if err != nil {
		return nil, err
	}

	err = k.deductCreateProjectFee(ctx, admin, msg.Fee)
	if err != nil {
		return nil, err
	}

	project, projectID, err := k.createNewProject(ctx)
	if err != nil {
		return nil, err
	}

	project.Admin = admin
	project.Jurisdiction = msg.Jurisdiction
	project.Metadata = msg.Metadata
	if msg.ReferenceId != "" {
		panic("reject reference ID")
	}

	if err = k.stateStore.ProjectTable().Save(ctx, project); err != nil {
		return nil, err
	}

	return &types.MsgCreateUnregisteredProjectResponse{ProjectId: projectID}, nil
}
