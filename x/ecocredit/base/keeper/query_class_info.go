package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	regenerrors "github.com/regen-network/regen-ledger/types/v2/errors"
	types "github.com/regen-network/regen-ledger/x/ecocredit/v3/base/types/v1"
)

// Class queries for information on a credit class.
func (k Keeper) Class(ctx context.Context, request *types.QueryClassRequest) (*types.QueryClassResponse, error) {
	class, err := k.stateStore.ClassTable().GetById(ctx, request.ClassId)
	if err != nil {
		return nil, regenerrors.ErrNotFound.Wrapf("could not get class with id %s: %s", request.ClassId, err.Error())
	}

	admin := sdk.AccAddress(class.Admin)

	info := types.ClassInfo{
		Id:               class.Id,
		Admin:            admin.String(),
		Metadata:         class.Metadata,
		CreditTypeAbbrev: class.CreditTypeAbbrev,
	}

	return &types.QueryClassResponse{Class: &info}, nil
}
