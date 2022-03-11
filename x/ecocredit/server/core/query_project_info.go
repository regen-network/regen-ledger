package core

import (
	"context"

	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

// ProjectInfo queries project info from the given project name.
func (k Keeper) ProjectInfo(ctx context.Context, request *core.QueryProjectInfoRequest) (*core.QueryProjectInfoResponse, error) {
	pInfo, err := k.stateStore.ProjectInfoStore().GetByName(ctx, request.ProjectId)
	if err != nil {
		return nil, err
	}
	var pi core.ProjectInfo
	if err = PulsarToGogoSlow(pInfo, &pi); err != nil {
		return nil, err
	}
	return &core.QueryProjectInfoResponse{Info: &pi}, nil
}
