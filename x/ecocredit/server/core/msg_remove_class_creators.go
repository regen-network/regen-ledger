package core

import (
	"context"

	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

func (k Keeper) RemoveClassCreators(ctx context.Context, req *core.MsgRemoveClassCreators) (*core.MsgRemoveClassCreatorsResponse, error) {
	if k.authority.String() != req.Authority {
		return nil, govtypes.ErrInvalidSigner.Wrapf("invalid authority: expected %s, got %s", k.authority, req.Authority)
	}

	for _, creator := range req.Creators {
		creatorAddr, err := sdk.AccAddressFromBech32(creator)
		if err != nil {
			return nil, err
		}

		classCreator, err := k.stateStore.AllowedClassCreatorTable().Get(ctx, creatorAddr)
		if err != nil {
			if ormerrors.NotFound.Is(err) {
				return nil, sdkerrors.ErrNotFound.Wrapf("class creator %s", creator)
			}
			return nil, err
		}

		if err := k.stateStore.AllowedClassCreatorTable().Delete(ctx, classCreator); err != nil {
			return nil, sdkerrors.ErrInvalidRequest.Wrapf("unable to remove %s from class creator list", creator)
		}
	}

	return &core.MsgRemoveClassCreatorsResponse{}, nil
}
