package keeper

import (
	"context"

	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	api "github.com/regen-network/regen-ledger/api/v2/regen/ecocredit/v1"
	types "github.com/regen-network/regen-ledger/x/ecocredit/v4/base/types/v1"
)

func (k Keeper) AddClassCreator(ctx context.Context, req *types.MsgAddClassCreator) (*types.MsgAddClassCreatorResponse, error) {

	creatorBz, err := k.ac.StringToBytes(req.Creator)
	if err != nil {
		return nil, err
	}
	authorityBz, err := k.ac.StringToBytes(req.Authority)
	if err != nil {
		return nil, errors.Wrapf(err, "invalid authority address")
	}

	authority := sdk.AccAddress(authorityBz)
	if !authority.Equals(k.authority) {
		return nil, govtypes.ErrInvalidSigner.Wrapf("invalid authority: expected %s, got %s", k.authority, req.Authority)
	}

	found, err := k.stateStore.AllowedClassCreatorTable().Has(ctx, creatorBz)
	if err != nil {
		return nil, err
	}

	if found {
		return nil, sdkerrors.ErrInvalidRequest.Wrapf("class creator %s already exists", req.Creator)
	}

	if err := k.stateStore.AllowedClassCreatorTable().Insert(ctx, &api.AllowedClassCreator{
		Address: creatorBz,
	}); err != nil {
		return nil, err
	}

	return &types.MsgAddClassCreatorResponse{}, nil
}
