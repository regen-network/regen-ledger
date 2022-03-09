package core

import (
	"context"
	"github.com/regen-network/regen-ledger/types"
	v1 "github.com/regen-network/regen-ledger/x/ecocredit/v1"
)

// ProjectInfo queries project info from the given project name.
func (k Keeper) ProjectInfo(ctx context.Context, request *v1.QueryProjectInfoRequest) (*v1.QueryProjectInfoResponse, error) {
	info, err := k.stateStore.ProjectInfoStore().GetByName(ctx, request.ProjectId)
	if err != nil {
		return nil, err
	}

	bz, err := types.DecodeMetadata(info.Metadata)
	if err != nil {
		return nil, err
	}
	info.Metadata = string(bz)

	var pi v1.ProjectInfo
	if err = PulsarToGogoSlow(info, &pi); err != nil {
		return nil, err
	}
	return &v1.QueryProjectInfoResponse{Info: &pi}, nil
}
