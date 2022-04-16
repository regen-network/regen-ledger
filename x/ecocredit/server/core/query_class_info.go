package core

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

// ClassInfo queries for information on a credit class.
func (k Keeper) ClassInfo(ctx context.Context, request *core.QueryClassInfoRequest) (*core.QueryClassInfoResponse, error) {
	class, err := k.stateStore.ClassInfoTable().GetByName(ctx, request.ClassId)
	if err != nil {
		return nil, err
	}

	admin := sdk.AccAddress(class.Admin)

	entry := core.ClassInfoEntry{
		Id:               class.Name,
		Admin:            admin.String(),
		Metadata:         class.Metadata,
		CreditTypeAbbrev: class.CreditType,
	}

	return &core.QueryClassInfoResponse{Class: &entry}, nil
}
