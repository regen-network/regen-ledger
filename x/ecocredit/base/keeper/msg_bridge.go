package keeper

import (
	"context"
	"fmt"
	"strings"

	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	types "github.com/regen-network/regen-ledger/x/ecocredit/base/types/v1"
)

// Bridge cancel credits, removing them from the supply and balance of the holder
func (k Keeper) Bridge(ctx context.Context, req *types.MsgBridge) (*types.MsgBridgeResponse, error) {

	exists, err := k.stateStore.AllowedBridgeChainTable().Has(ctx, strings.ToLower(req.Target))
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, sdkerrors.ErrUnauthorized.Wrapf("%s is not an authorized bridge target", req.Target)
	}

	_, err = k.Cancel(ctx, &types.MsgCancel{
		Owner:   req.Owner,
		Credits: req.Credits,
		Reason:  fmt.Sprintf("bridge-%s", req.Target),
	})
	if err != nil {
		return nil, err
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)
	for _, credit := range req.Credits {
		batch, err := k.stateStore.BatchTable().GetByDenom(ctx, credit.BatchDenom)
		if err != nil {
			return nil, err // we already know the batch exists from Cancel
		}

		// if no batch contract entry is found then we error because we only support
		// bridging credits from credit batches that were created as a result of a
		// bridge operation (i.e. only previously bridged credits)
		batchContract, err := k.stateStore.BatchContractTable().Get(ctx, batch.Key)
		if err != nil {
			if ormerrors.IsNotFound(err) {
				return nil, sdkerrors.ErrInvalidRequest.Wrap(
					"only credits previously bridged from another chain are supported",
				)
			}
			return nil, err
		}

		if err = sdkCtx.EventManager().EmitTypedEvent(&types.EventBridge{
			Target:    req.Target,
			Recipient: req.Recipient,
			Contract:  batchContract.Contract,
			Amount:    credit.Amount,
			Owner:     req.Owner,
		}); err != nil {
			return nil, err
		}
	}

	return &types.MsgBridgeResponse{}, nil
}
