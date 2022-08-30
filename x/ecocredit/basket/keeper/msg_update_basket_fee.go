package keeper

import (
	"context"

	sdkv1beta1 "github.com/cosmos/cosmos-sdk/api/cosmos/base/v1beta1"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/basket/v1"
	types "github.com/regen-network/regen-ledger/x/ecocredit/basket/types/v1"
)

// UpdateBasketFees is an RPC to handle basket.UpdateBasketFees
func (k Keeper) UpdateBasketFees(ctx context.Context, req *types.MsgUpdateBasketFees) (*types.MsgUpdateBasketFeesResponse, error) {
	if k.authority.String() != req.Authority {
		return nil, govtypes.ErrInvalidSigner.Wrapf("invalid authority: expected %s, got %s", k.authority, req.Authority)
	}

	basketFee := make([]*sdkv1beta1.Coin, req.BasketFees.Len())
	for i, coin := range req.BasketFees {
		basketFee[i] = &sdkv1beta1.Coin{
			Denom:  coin.Denom,
			Amount: coin.Amount.String(),
		}
	}

	if err := k.stateStore.BasketFeesTable().Save(ctx, &api.BasketFees{
		Fees: basketFee,
	}); err != nil {
		return nil, err
	}

	return &types.MsgUpdateBasketFeesResponse{}, nil
}
