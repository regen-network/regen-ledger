package core

import (
	"context"
	"fmt"

	"github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

// Bridge cancel credits, removing them from the supply and balance of the holder
func (k Keeper) Bridge(ctx context.Context, req *core.MsgBridge) (*core.MsgBridgeResponse, error) {
	_, err := k.Cancel(ctx, &core.MsgCancel{
		Owner:   req.Owner,
		Credits: req.Credits,
		Reason:  fmt.Sprintf("bridge-%s", req.Target),
	})
	if err != nil {
		return nil, err
	}

	sdkCtx := types.UnwrapSDKContext(ctx)
	for _, credit := range req.Credits {
		if err = sdkCtx.EventManager().EmitTypedEvent(&core.EventBridge{
			Target:    req.Target,
			Recipient: req.Recipient,
			Contract:  req.Contract,
			Amount:    credit.Amount,
		}); err != nil {
			return nil, err
		}
	}

	return &core.MsgBridgeResponse{}, nil
}
