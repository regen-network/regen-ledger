package keeper

import (
	"context"

	sdkbase "cosmossdk.io/api/cosmos/base/v1beta1"
	"cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	api "github.com/regen-network/regen-ledger/api/v2/regen/ecocredit/basket/v1"
	regentypes "github.com/regen-network/regen-ledger/types/v2"
	types "github.com/regen-network/regen-ledger/x/ecocredit/v4/basket/types/v1"
)

// UpdateBasketFee is an RPC to handle basket.UpdateBasketFee
func (k Keeper) UpdateBasketFee(ctx context.Context, req *types.MsgUpdateBasketFee) (*types.MsgUpdateBasketFeeResponse, error) {
	if err := req.ValidateBasic(); err != nil {
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

	var basketFee *sdkbase.Coin
	if req.Fee != nil && req.Fee.IsPositive() {
		basketFee = regentypes.CoinToCosmosAPILegacy(*req.Fee)
	}

	if err := k.stateStore.BasketFeeTable().Save(ctx, &api.BasketFee{
		Fee: basketFee,
	}); err != nil {
		return nil, err
	}

	return &types.MsgUpdateBasketFeeResponse{}, nil
}
