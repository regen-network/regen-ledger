package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	api "github.com/regen-network/regen-ledger/api/v2/regen/ecocredit/v1"
	types "github.com/regen-network/regen-ledger/x/ecocredit/v3/base/types/v1"
)

func (k Keeper) AddClassCreator(ctx context.Context, req *types.MsgAddClassCreator) (*types.MsgAddClassCreatorResponse, error) {
	if k.authority.String() != req.Authority {
		return nil, govtypes.ErrInvalidSigner.Wrapf("invalid authority: expected %s, got %s", k.authority, req.Authority)
	}

	creatorAddr, err := sdk.AccAddressFromBech32(req.Creator)
	if err != nil {
		return nil, err
	}

	found, err := k.stateStore.AllowedClassCreatorTable().Has(ctx, creatorAddr)
	if err != nil {
		return nil, err
	}

	if found {
		return nil, sdkerrors.ErrInvalidRequest.Wrapf("class creator %s already exists", req.Creator)
	}

	if err := k.stateStore.AllowedClassCreatorTable().Insert(ctx, &api.AllowedClassCreator{
		Address: creatorAddr,
	}); err != nil {
		return nil, err
	}

	return &types.MsgAddClassCreatorResponse{}, nil
}
