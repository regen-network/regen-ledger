package core

import (
	"context"

	v1beta1 "github.com/cosmos/cosmos-sdk/api/cosmos/base/v1beta1"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	ecocreditv1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

func (k Keeper) UpdateClassFees(ctx context.Context, req *core.MsgUpdateClassFees) (*core.MsgUpdateClassFeesResponse, error) {
	if k.authority.String() != req.Authority {
		return nil, govtypes.ErrInvalidSigner.Wrapf("invalid authority: expected %s, got %s", k.authority, req.Authority)
	}

	classFee := make([]*v1beta1.Coin, req.Fees.Len())
	for i, coin := range req.Fees {
		classFee[i] = &v1beta1.Coin{
			Denom:  coin.Denom,
			Amount: coin.Amount.String(),
		}
	}

	if err := k.stateStore.ClassFeesTable().Save(ctx, &ecocreditv1.ClassFees{
		Fees: classFee,
	}); err != nil {
		return nil, err
	}

	return &core.MsgUpdateClassFeesResponse{}, nil
}