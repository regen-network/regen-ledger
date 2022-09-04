package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	regentypes "github.com/regen-network/regen-ledger/types"
	types "github.com/regen-network/regen-ledger/x/ecocredit/basket/types/v1"
)

func (k Keeper) BasketFee(ctx context.Context, req *types.QueryBasketFeeRequest) (*types.QueryBasketFeeResponse, error) {
	basketFee, err := k.stateStore.BasketFeeTable().Get(ctx)
	if err != nil {
		return nil, err
	}

	var fee sdk.Coin

	if basketFee.Fee != nil {
		var ok bool
		fee, ok = regentypes.ProtoCoinToCoin(basketFee.Fee)
		if !ok {
			return nil, sdkerrors.ErrInvalidType.Wrapf("failed to parse basket fee")
		}
	}

	return &types.QueryBasketFeeResponse{
		Fee: &fee,
	}, nil
}
