package core

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

// Class queries for information on a credit class.
func (k Keeper) Class(ctx context.Context, request *core.QueryClassRequest) (*core.QueryClassResponse, error) {
	class, err := k.stateStore.ClassTable().GetById(ctx, request.ClassId)
	if err != nil {
		return nil, sdkerrors.ErrInvalidRequest.Wrapf("could not get class with id %s: %s", request.ClassId, err.Error())
	}

	admin := sdk.AccAddress(class.Admin)

	info := core.ClassInfo{
		Id:               class.Id,
		Admin:            admin.String(),
		Metadata:         class.Metadata,
		CreditTypeAbbrev: class.CreditTypeAbbrev,
	}

	return &core.QueryClassResponse{Class: &info}, nil
}
