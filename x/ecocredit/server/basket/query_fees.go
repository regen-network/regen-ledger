package basket

import (
	"context"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	baskettypes "github.com/regen-network/regen-ledger/x/ecocredit/basket"

	"github.com/regen-network/regen-ledger/types"
)

func (k Keeper) BasketFees(ctx context.Context, req *baskettypes.QueryBasketFeesRequest) (*baskettypes.QueryBasketFeesResponse, error) {

	fees, err := k.stateStore.BasketFeesTable().Get(ctx)
	if err != nil {
		return nil, err
	}

	result, ok := types.ProtoCoinsToCoins(fees.Fees)
	if !ok {
		return nil, sdkerrors.ErrInvalidType.Wrapf("failed to get basket fees")
	}

	return &baskettypes.QueryBasketFeesResponse{
		Fees: result,
	}, nil
}
