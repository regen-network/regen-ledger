package keeper

import (
	"context"

	sdkv1beta1 "github.com/cosmos/cosmos-sdk/api/cosmos/base/v1beta1"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	api "github.com/regen-network/regen-ledger/api/v2/regen/ecocredit/v1"
	regentypes "github.com/regen-network/regen-ledger/types/v2"
	types "github.com/regen-network/regen-ledger/x/ecocredit/v3/base/types/v1"
)

func (k Keeper) UpdateProjectFee(ctx context.Context, req *types.MsgUpdateProjectFee) (*types.MsgUpdateProjectFeeResponse, error) {
	if k.authority.String() != req.Authority {
		return nil, govtypes.ErrInvalidSigner.Wrapf("invalid authority: expected %s, got %s", k.authority, req.Authority)
	}

	var classFee *sdkv1beta1.Coin
	if req.Fee != nil && req.Fee.IsPositive() {
		classFee = regentypes.CoinToProtoCoin(*req.Fee)
	}

	if err := k.stateStore.ProjectFeeTable().Save(ctx, &api.ProjectFee{
		Fee: classFee,
	}); err != nil {
		return nil, err
	}

	return &types.MsgUpdateProjectFeeResponse{}, nil
}
