package core

import (
	"context"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

// ClassInfo queries for information on a credit class.
func (k Keeper) ClassInfo(ctx context.Context, request *core.QueryClassInfoRequest) (*core.QueryClassInfoResponse, error) {
	classInfo, err := k.stateStore.ClassInfoStore().GetByName(ctx, request.ClassId)
	if err != nil {
		return nil, err
	}

	var ci core.ClassInfo
	if err = PulsarToGogoSlow(classInfo, &ci); err != nil {
		return nil, err
	}
	return &core.QueryClassInfoResponse{Info: &ci}, nil
}
