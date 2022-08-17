package basket

import (
	"context"

	v1beta1 "github.com/cosmos/cosmos-sdk/api/cosmos/base/v1beta1"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	basketv1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/basket/v1"
	baskettypes "github.com/regen-network/regen-ledger/x/ecocredit/basket"
)

// Create is an RPC to handle basket.UpdateBasketFee
func (k Keeper) UpdateBasketFee(ctx context.Context, req *baskettypes.MsgUpdateBasketFee) (*baskettypes.MsgUpdateBasketFeeResponse, error) {
	panic("OOPS")
	if k.authority.String() != req.Authority {
		return nil, govtypes.ErrInvalidSigner.Wrapf("invalid authority: expected %s, got %s", k.authority, req.Authority)
	}

	basketFee := make([]*v1beta1.Coin, req.BasketFee.Len())
	for i, coin := range req.BasketFee {
		basketFee[i] = &v1beta1.Coin{
			Denom:  coin.Denom,
			Amount: coin.Amount.String(),
		}
	}

	if err := k.stateStore.BasketFeeTable().Save(ctx, &basketv1.BasketFee{
		Fee: basketFee,
	}); err != nil {
		return nil, err
	}

	return &baskettypes.MsgUpdateBasketFeeResponse{}, nil
}
