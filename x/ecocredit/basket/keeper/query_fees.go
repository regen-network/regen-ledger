package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	regentypes "github.com/regen-network/regen-ledger/types/v2"
	regenerrors "github.com/regen-network/regen-ledger/types/v2/errors"
	types "github.com/regen-network/regen-ledger/x/ecocredit/v3/basket/types/v1"
)

func (k Keeper) BasketFee(ctx context.Context, _ *types.QueryBasketFeeRequest) (*types.QueryBasketFeeResponse, error) {
	basketFee, err := k.stateStore.BasketFeeTable().Get(ctx)
	if err != nil {
		return nil, regenerrors.ErrInternal.Wrap(err.Error())
	}

	var fee sdk.Coin

	if basketFee.Fee != nil {
		var ok bool
		fee, ok = regentypes.ProtoCoinToCoin(basketFee.Fee)
		if !ok {
			return nil, regenerrors.ErrInternal.Wrap("failed to parse basket fee")
		}
	}

	return &types.QueryBasketFeeResponse{
		Fee: &fee,
	}, nil
}
