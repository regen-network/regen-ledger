package core

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

// SealBatch sets the Open field to false in a batch IFF the requester address matches the batch issuer address.
// This method is a no-op for batches which already have Open set to false.
func (k Keeper) SealBatch(ctx context.Context, req *core.MsgSealBatch) (*core.MsgSealBatchResponse, error) {
	issuer, err := sdk.AccAddressFromBech32(req.Issuer)
	if err != nil {
		return nil, err
	}

	batch, err := k.stateStore.BatchTable().GetByDenom(ctx, req.BatchDenom)
	if err != nil {
		return nil, sdkerrors.ErrInvalidRequest.Wrapf("could not get batch with denom %s: %s", req.BatchDenom, err.Error())
	}

	if !sdk.AccAddress(batch.Issuer).Equals(issuer) {
		return nil, sdkerrors.ErrUnauthorized.Wrap("only the batch issuer can seal the batch")
	}

	if !batch.Open {
		return &core.MsgSealBatchResponse{}, nil
	}

	batch.Open = false
	if err := k.stateStore.BatchTable().Update(ctx, batch); err != nil {
		return nil, err
	}

	if err := sdk.UnwrapSDKContext(ctx).EventManager().EmitTypedEvent(&core.EventSealBatch{BatchDenom: batch.Denom}); err != nil {
		return nil, err
	}

	return &core.MsgSealBatchResponse{}, nil
}
