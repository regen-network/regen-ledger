package marketplace

import (
	"context"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	v1 "github.com/regen-network/regen-ledger/x/ecocredit/marketplace"
	"github.com/regen-network/regen-ledger/x/ecocredit/server/basket"
)

func (k Keeper) SellOrder(ctx context.Context, req *v1.QuerySellOrderRequest) (*v1.QuerySellOrderResponse, error) {
	order, err := k.stateStore.SellOrderTable().Get(ctx, req.SellOrderId)
	if err != nil {
		return nil, sdkerrors.ErrInvalidRequest.Wrapf("could not get sell order with id %d: %s", req.SellOrderId, err.Error())
	}
	var so v1.SellOrder
	if err = basket.PulsarToGogoSlow(order, &so); err != nil {
		return nil, err
	}
	return &v1.QuerySellOrderResponse{SellOrder: &so}, nil
}
