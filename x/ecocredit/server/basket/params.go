package basket

import (
	"context"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/basket/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/regen-network/regen-ledger/x/ecocredit/basket"
)

// UpdateBasketFee is a governance only function that allows for the removal and addition of fees users can pay to create a basket
func (k Keeper) UpdateBasketFee(ctx context.Context, req *basket.MsgUpdateBasketFee) (*basket.MsgUpdateBasketFeeResponse, error) {
	if err := ecocredit.AssertGovernance(req.RootAddress, k.accountKeeper); err != nil {
		return nil, err
	}

	store := k.stateStore.BasketFeeTable()

	for _, denom := range req.RemoveFees {
		if err := store.Delete(ctx, &api.BasketFee{
			Denom: denom,
		}); err != nil {
			return nil, err
		}
	}

	for _, fee := range req.AddFees {
		if err := store.Insert(ctx, &api.BasketFee{
			Denom:  fee.Denom,
			Amount: fee.Amount.String(),
		}); err != nil {
			return nil, err
		}
	}

	return &basket.MsgUpdateBasketFeeResponse{}, nil
}
