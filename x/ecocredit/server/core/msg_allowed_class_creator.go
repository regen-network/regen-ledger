package core

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

func (k Keeper) AllowedClassCreator(ctx context.Context, req *core.MsgAllowedClassCreator) (*core.MsgAllowedClassCreatorResponse, error) {
	if k.authority.String() != req.Authority {
		return nil, govtypes.ErrInvalidSigner.Wrapf("invalid authority: expected %s, got %s", k.authority, req.Authority)
	}

	for _, creator := range req.Creators {
		creatorAddr, err := sdk.AccAddressFromBech32(creator)
		if err != nil {
			return nil, err
		}

		found, err := k.stateStore.AllowedClassCreatorTable().Has(ctx, creatorAddr)
		if err != nil {
			return nil, err
		}

		if found {
			return nil, sdkerrors.ErrInvalidRequest.Wrapf("class creator %s already exists", creator)
		}

	}

	return &core.MsgAllowedClassCreatorResponse{}, nil
}
