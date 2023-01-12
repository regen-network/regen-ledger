package keeper

import (
	"context"

	sdkv1beta1 "github.com/cosmos/cosmos-sdk/api/cosmos/base/v1beta1"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	api "github.com/regen-network/regen-ledger/api/v2/regen/ecocredit/basket/v1"
	regentypes "github.com/regen-network/regen-ledger/types/v2"
	types "github.com/regen-network/regen-ledger/x/ecocredit/v3/basket/types/v1"
)

// UpdateBasketFee is an RPC to handle basket.UpdateBasketFee
func (k Keeper) UpdateBasketFee(ctx context.Context, req *types.MsgUpdateBasketFee) (*types.MsgUpdateBasketFeeResponse, error) {
	if k.authority.String() != req.Authority {
		return nil, govtypes.ErrInvalidSigner.Wrapf("invalid authority: expected %s, got %s", k.authority, req.Authority)
	}

	var basketFee *sdkv1beta1.Coin

	if req.Fee != nil && req.Fee.IsPositive() {
		basketFee = regentypes.CoinToProtoCoin(*req.Fee)
	}

	if err := k.stateStore.BasketFeeTable().Save(ctx, &api.BasketFee{
		Fee: basketFee,
	}); err != nil {
		return nil, err
	}

	return &types.MsgUpdateBasketFeeResponse{}, nil
}
