package keeper

import (
	"context"

	sdkbase "cosmossdk.io/api/cosmos/base/v1beta1"

	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	api "github.com/regen-network/regen-ledger/api/v2/regen/ecocredit/v1"
	regentypes "github.com/regen-network/regen-ledger/types/v2"
	types "github.com/regen-network/regen-ledger/x/ecocredit/v3/base/types/v1"
)

func (k Keeper) UpdateClassFee(ctx context.Context, req *types.MsgUpdateClassFee) (*types.MsgUpdateClassFeeResponse, error) {
	if k.authority.String() != req.Authority {
		return nil, govtypes.ErrInvalidSigner.Wrapf("invalid authority: expected %s, got %s", k.authority, req.Authority)
	}

	var classFee *sdkbase.Coin
	if req.Fee != nil && req.Fee.IsPositive() {
		classFee = regentypes.CoinToCosmosAPILegacy(*req.Fee)
	}

	if err := k.stateStore.ClassFeeTable().Save(ctx, &api.ClassFee{
		Fee: classFee,
	}); err != nil {
		return nil, err
	}

	return &types.MsgUpdateClassFeeResponse{}, nil
}
