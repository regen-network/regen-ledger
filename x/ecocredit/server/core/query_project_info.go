package core

import (
	"context"
	v1 "github.com/regen-network/regen-ledger/x/ecocredit/v1"
)

// ProjectInfo queries project info from the given project name.
func (k Keeper) ProjectInfo(ctx context.Context, request *v1.QueryProjectInfoRequest) (*v1.QueryProjectInfoResponse, error) {
	pInfo, err := k.stateStore.ProjectInfoStore().GetByName(ctx, request.ProjectId)
	if err != nil {
		return nil, err
	}
	var pi v1.ProjectInfo
	if err = PulsarToGogoSlow(pInfo, &pi); err != nil {
		return nil, err
	}
	return &v1.QueryProjectInfoResponse{Info: &pi}, nil
}
