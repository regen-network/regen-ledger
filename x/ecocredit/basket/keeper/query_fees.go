package keeper

import (
	"context"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	regentypes "github.com/regen-network/regen-ledger/types"
	types "github.com/regen-network/regen-ledger/x/ecocredit/basket/types/v1"
)

func (k Keeper) BasketFees(ctx context.Context, req *types.QueryBasketFeesRequest) (*types.QueryBasketFeesResponse, error) {

	fees, err := k.stateStore.BasketFeesTable().Get(ctx)
	if err != nil {
		return nil, err
	}

	result, ok := regentypes.ProtoCoinsToCoins(fees.Fees)
	if !ok {
		return nil, sdkerrors.ErrInvalidType.Wrapf("failed to get basket fees")
	}

	return &types.QueryBasketFeesResponse{
		Fees: result,
	}, nil
}
