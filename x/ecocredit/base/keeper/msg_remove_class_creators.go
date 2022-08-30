package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	types "github.com/regen-network/regen-ledger/x/ecocredit/base/types/v1"
)

func (k Keeper) RemoveClassCreator(ctx context.Context, req *types.MsgRemoveClassCreator) (*types.MsgRemoveClassCreatorResponse, error) {
	if k.authority.String() != req.Authority {
		return nil, govtypes.ErrInvalidSigner.Wrapf("invalid authority: expected %s, got %s", k.authority, req.Authority)
	}

	creatorAddr, err := sdk.AccAddressFromBech32(req.Creator)
	if err != nil {
		return nil, err
	}

	classCreator, err := k.stateStore.AllowedClassCreatorTable().Get(ctx, creatorAddr)
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
