package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	types "github.com/regen-network/regen-ledger/x/ecocredit/v3/basket/types/v1"
)

// UpdateCurator is an RPC to handle basket.UpdateCurator
func (k Keeper) UpdateCurator(ctx context.Context, req *types.MsgUpdateCurator) (*types.MsgUpdateCuratorResponse, error) {
	basket, err := k.stateStore.BasketTable().GetByBasketDenom(ctx, req.Denom)
	if err != nil {
		if ormerrors.IsNotFound(err) {
			return nil, sdkerrors.ErrNotFound.Wrapf("basket with denom %s", req.Denom)
		}
		return nil, err
	}

	curator, err := sdk.AccAddressFromBech32(req.Curator)
	if err != nil {
		return nil, err
	}

	if !curator.Equals(sdk.AccAddress(basket.Curator)) {
		return nil, sdkerrors.ErrUnauthorized.Wrapf("expected curator %s got %s", sdk.AccAddress(basket.Curator).String(), req.Curator)
	}

	newCurator, err := sdk.AccAddressFromBech32(req.NewCurator)
	if err != nil {
		return nil, err
	}

	basket.Curator = newCurator
	if err := k.stateStore.BasketTable().Update(ctx, basket); err != nil {
		return nil, sdkerrors.ErrInvalidRequest.Wrapf("unable to update basket %s", req.Denom)
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)
	if err = sdkCtx.EventManager().EmitTypedEvent(&types.EventUpdateCurator{
		Denom: basket.BasketDenom,
	}); err != nil {
		return nil, err
	}

	return &types.MsgUpdateCuratorResponse{}, nil
}
