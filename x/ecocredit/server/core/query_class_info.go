package core

import (
	"context"
	"github.com/regen-network/regen-ledger/types"
	v1 "github.com/regen-network/regen-ledger/x/ecocredit/v1"
)

// ClassInfo queries for information on a credit class.
func (k Keeper) ClassInfo(ctx context.Context, request *v1.QueryClassInfoRequest) (*v1.QueryClassInfoResponse, error) {
	classInfo, err := k.stateStore.ClassInfoStore().GetByName(ctx, request.ClassId)
	if err != nil {
		return nil, err
	}

	bz, err := types.DecodeMetadata(classInfo.Metadata)
	if err != nil {
		return nil, err
	}
	classInfo.Metadata = string(bz)

	var ci v1.ClassInfo
	if err = PulsarToGogoSlow(classInfo, &ci); err != nil {
		return nil, err
	}

	return &v1.QueryClassInfoResponse{Info: &ci}, nil
}
