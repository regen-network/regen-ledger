package core

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	ecocreditv1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

func (k Keeper) RemoveClassCreator(ctx context.Context, req *core.MsgRemoveClassCreator) (*core.MsgRemoveClassCreatorResponse, error) {
	if k.authority.String() != req.Authority {
		return nil, govtypes.ErrInvalidSigner.Wrapf("invalid authority: expected %s, got %s", k.authority, req.Authority)
	}

	for _, creator := range req.Creators {
		creatorAddr, err := sdk.AccAddressFromBech32(creator)
		if err != nil {
			return nil, err
		}

		if err := k.stateStore.AllowedClassCreatorTable().Delete(ctx, &ecocreditv1.AllowedClassCreator{
			Address: creatorAddr,
		}); err != nil {
			return nil, sdkerrors.ErrInvalidRequest.Wrapf("unable to remove %s from class creator list", creator)
		}
	}

	return &core.MsgRemoveClassCreatorResponse{}, nil
}
