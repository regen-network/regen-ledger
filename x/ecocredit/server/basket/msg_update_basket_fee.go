package basket

import (
	"context"

	v1beta1 "github.com/cosmos/cosmos-sdk/api/cosmos/base/v1beta1"

	basketv1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/basket/v1"
	baskettypes "github.com/regen-network/regen-ledger/x/ecocredit/basket"
)

func (k Keeper) UpdateBasketFee(ctx context.Context, req *baskettypes.MsgUpdateBasketFee) (*baskettypes.MsgUpdateBasketFeeResponse, error) {

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
