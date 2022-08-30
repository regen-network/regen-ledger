package keeper

import (
	"context"

	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	regentypes "github.com/regen-network/regen-ledger/types"
	types "github.com/regen-network/regen-ledger/x/ecocredit/base/types/v1"
)

func (k Keeper) UpdateClassFees(ctx context.Context, req *types.MsgUpdateClassFees) (*types.MsgUpdateClassFeesResponse, error) {
	if k.authority.String() != req.Authority {
		return nil, govtypes.ErrInvalidSigner.Wrapf("invalid authority: expected %s, got %s", k.authority, req.Authority)
	}

	classFee := regentypes.CoinsToProtoCoins(req.Fees)
	if err := k.stateStore.ClassFeesTable().Save(ctx, &api.ClassFees{
		Fees: classFee,
	}); err != nil {
		return nil, err
	}

	return &types.MsgUpdateClassFeesResponse{}, nil
}
