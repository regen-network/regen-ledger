package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/regen-network/regen-ledger/api/v2/orm/types/ormerrors"
	types "github.com/regen-network/regen-ledger/x/ecocredit/v4/basket/types/v1"
)

// UpdateDateCriteria is an RPC to handle basket.MsgUpdateDateCriteria
func (k Keeper) UpdateDateCriteria(ctx context.Context, msg *types.MsgUpdateDateCriteria) (*types.MsgUpdateDateCriteriaResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	authorityBz, err := k.ac.StringToBytes(msg.Authority)
	if err != nil {
		return nil, sdkerrors.ErrInvalidAddress.Wrapf("invalid authority address: %s", err)
	}
	authority := sdk.AccAddress(authorityBz)

	if !authority.Equals(k.authority) {
		return nil, govtypes.ErrInvalidSigner.Wrapf("invalid authority: expected %s, got %s", k.authority, msg.Authority)
	}

	basket, err := k.stateStore.BasketTable().GetByBasketDenom(ctx, msg.Denom)
	if err != nil {
		return nil, ormerrors.NotFound.Wrapf("basket with denom %s does not exist", msg.Denom)
	}

	basket.DateCriteria = msg.NewDateCriteria.ToAPI()

	err = k.stateStore.BasketTable().Update(ctx, basket)
	if err != nil {
		return nil, err
	}

	err = sdkCtx.EventManager().EmitTypedEvent(&types.EventUpdateDateCriteria{
		Denom: msg.Denom,
	})

	return &types.MsgUpdateDateCriteriaResponse{}, err
}
