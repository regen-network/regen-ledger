package core

import (
	"context"
	"fmt"

	"github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

// Bridge cancel credits, removing them from the supply and balance of the holder
func (k Keeper) Bridge(ctx context.Context, req *core.MsgBridge) (*core.MsgBridgeResponse, error) {

	creditsToCancel := make([]*core.MsgCancel_CancelCredits, len(req.Credits))
	for i, credit := range req.Credits {
		creditsToCancel[i] = &core.MsgCancel_CancelCredits{
			BatchDenom: credit.BatchDenom,
			Amount:     credit.Amount,
		}
	}

	_, err := k.Cancel(ctx, &core.MsgCancel{
		Holder:  req.Holder,
		Credits: creditsToCancel,
		Reason:  fmt.Sprintf("bridge-%s", req.Target),
	})
	if err != nil {
		return nil, err
	}

	sdkCtx := types.UnwrapSDKContext(ctx)
	if err = sdkCtx.EventManager().EmitTypedEvent(&core.EventBridge{
		Target:    req.Target,
		Recipient: req.Recipient,
		Contract:  req.Contract,
	}); err != nil {
		return nil, err
	}

	return &core.MsgBridgeResponse{}, nil
}
