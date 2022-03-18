package marketplace

import (
	"context"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/regen-network/regen-ledger/types/ormutil"
	"github.com/regen-network/regen-ledger/x/ecocredit/marketplace"
)

func (k Keeper) SellOrder(ctx context.Context, req *marketplace.QuerySellOrderRequest) (*marketplace.QuerySellOrderResponse, error) {
	order, err := k.stateStore.SellOrderTable().Get(ctx, req.SellOrderId)
	if err != nil {
		return nil, sdkerrors.ErrInvalidRequest.Wrapf("could not get sell order with id %d: %s", req.SellOrderId, err.Error())
	}
	var so marketplace.SellOrder
	if err = ormutil.PulsarToGogoSlow(order, &so); err != nil {
		return nil, err
	}
	return &marketplace.QuerySellOrderResponse{SellOrder: &so}, nil
}
