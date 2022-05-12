package core

import (
	"context"

	"github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

// Bridge cancel credits, removing them from the supply and balance of the holder
func (k Keeper) Bridge(ctx context.Context, req *core.MsgBridge) (*core.MsgBridgeResponse, error) {
	_, err := k.Cancel(ctx, req.MsgCancel)
	if err != nil {
		return nil, err
	}

	sdkCtx := types.UnwrapSDKContext(ctx)
	if err = sdkCtx.EventManager().EmitTypedEvent(&core.EventBridge{
		BridgeTarget:    req.BridgeTarget,
		BridgeRecipient: req.BridgeRecipient,
		BridgeContract:  req.BridgeContract,
	}); err != nil {
		return nil, err
	}

	return &core.MsgBridgeResponse{}, nil
}
