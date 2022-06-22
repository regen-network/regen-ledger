package core

import (
	"context"
	"fmt"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
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

	for _, credit := range req.Credits {
		it, err := k.stateStore.BatchOriginTxTable().List(
			ctx,
			api.BatchOriginTxBatchDenomIndexKey{}.WithBatchDenom(credit.BatchDenom),
		)
		// all contract addresses should be the same for a given batch, therefore,
		// we only want the first origin tx that matches the batch denom
		var originTx *api.BatchOriginTx
		if it.Next() {
			var err error
			originTx, err = it.Value()
			if err != nil {
				return nil, err
			}
		}
		it.Close()

		// if no matching origin tx was found then we error because we only support
		// bridging credits from credit batches that were created as a result of a
		// bridge operation (i.e. only bridging previously bridged credits)
		if originTx == nil {
			return nil, sdkerrors.ErrInvalidRequest.Wrap(
				"only credits previously bridged from another chain are supported at this time",
			)
		}

		sdkCtx := types.UnwrapSDKContext(ctx)
		if err = sdkCtx.EventManager().EmitTypedEvent(&core.EventBridge{
			Target:    req.Target,
			Recipient: req.Recipient,
			Contract:  originTx.Contract,
			Amount:    credit.Amount,
		}); err != nil {
			return nil, err
		}
	}

	return &core.MsgBridgeResponse{}, nil
}
