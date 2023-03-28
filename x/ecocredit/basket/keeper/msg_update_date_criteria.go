package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	types "github.com/regen-network/regen-ledger/x/ecocredit/v3/basket/types/v1"
)

// UpdateDateCriteria is an RPC to handle basket.MsgUpdateDateCriteria
func (k Keeper) UpdateDateCriteria(ctx context.Context, msg *types.MsgUpdateDateCriteria) (*types.MsgUpdateDateCriteriaResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	if k.authority.String() != msg.Authority {
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
