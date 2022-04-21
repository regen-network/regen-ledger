package core

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

// ClassInfo queries for information on a credit class.
func (k Keeper) ClassInfo(ctx context.Context, request *core.QueryClassInfoRequest) (*core.QueryClassInfoResponse, error) {
	class, err := k.stateStore.ClassTable().GetById(ctx, request.ClassId)
	if err != nil {
		return nil, err
	}

	admin := sdk.AccAddress(class.Admin)

	info := core.ClassInfo{
		Id:               class.Id,
		Admin:            admin.String(),
		Metadata:         class.Metadata,
		CreditTypeAbbrev: class.CreditTypeAbbrev,
	}

	return &core.QueryClassInfoResponse{Class: &info}, nil
}
