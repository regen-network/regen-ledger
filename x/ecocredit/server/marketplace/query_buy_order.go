package marketplace

import (
	"context"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/regen-network/regen-ledger/types/ormutil"
	v1 "github.com/regen-network/regen-ledger/x/ecocredit/marketplace"
)

// BuyOrder queries a single buy order.
func (k Keeper) BuyOrder(ctx context.Context, request *v1.QueryBuyOrderRequest) (*v1.QueryBuyOrderResponse, error) {
	order, err := k.stateStore.BuyOrderTable().Get(ctx, request.BuyOrderId)
	if err != nil {
		return nil, sdkerrors.ErrInvalidRequest.Wrapf("could not get buy order %d: %s", request.BuyOrderId, err.Error())
	}
	var gogoOrder v1.BuyOrder
	if err = ormutil.PulsarToGogoSlow(order, &gogoOrder); err != nil {
		return nil, err
	}
	return &v1.QueryBuyOrderResponse{BuyOrder: &gogoOrder}, nil
}
