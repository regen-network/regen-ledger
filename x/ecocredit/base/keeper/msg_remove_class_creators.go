package keeper

import (
	"context"

	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/regen-network/regen-ledger/orm/types/ormerrors"

	types "github.com/regen-network/regen-ledger/x/ecocredit/v4/base/types/v1"
)

func (k Keeper) RemoveClassCreator(ctx context.Context, req *types.MsgRemoveClassCreator) (*types.MsgRemoveClassCreatorResponse, error) {
	authorityBz, err := k.ac.StringToBytes(req.Authority)
	if err != nil {
		return nil, errors.Wrapf(err, "invalid authority address")
	}
	authority := sdk.AccAddress(authorityBz)

	if !authority.Equals(k.authority) {
		return nil, govtypes.ErrInvalidSigner.Wrapf("invalid authority: expected %s, got %s", k.authority, req.Authority)
	}

	creattorBz, err := k.ac.StringToBytes(req.Creator)
	if err != nil {
		return nil, err
	}

	classCreator, err := k.stateStore.AllowedClassCreatorTable().Get(ctx, creattorBz)
	if err != nil {
		if ormerrors.NotFound.Is(err) {
			return nil, sdkerrors.ErrNotFound.Wrapf("class creator %s", req.Creator)
		}
		return nil, err
	}

	if err := k.stateStore.AllowedClassCreatorTable().Delete(ctx, classCreator); err != nil {
		return nil, sdkerrors.ErrInvalidRequest.Wrapf("unable to remove %s from class creator list: %s", req.Creator, err.Error())
	}

	return &types.MsgRemoveClassCreatorResponse{}, nil
}
