package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	types "github.com/regen-network/regen-ledger/x/ecocredit/base/types/v1"
)

// UpdateBatchMetadata updates the metadata for the batch.
func (k Keeper) UpdateBatchMetadata(ctx context.Context, req *types.MsgUpdateBatchMetadata) (*types.MsgUpdateBatchMetadataResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	reqAddr, err := sdk.AccAddressFromBech32(req.Issuer)
	if err != nil {
		return nil, err
	}

	batchInfo, err := k.stateStore.BatchTable().GetByDenom(ctx, req.BatchDenom)
	if err != nil {
		return nil, sdkerrors.ErrInvalidRequest.Wrapf(
			"could not get credit batch with denom %s: %s", req.BatchDenom, err,
		)
	}

	if !batchInfo.Open {
		return nil, sdkerrors.ErrUnauthorized.Wrapf(
			"credit batch %s is not open", req.BatchDenom,
		)
	}

	issuer := sdk.AccAddress(batchInfo.Issuer)
	if !reqAddr.Equals(issuer) {
		return nil, sdkerrors.ErrUnauthorized.Wrapf(
			"%s is not the issuer of credit batch %s", req.Issuer, req.BatchDenom,
		)
	}

	batchInfo.Metadata = req.NewMetadata
	if err = k.stateStore.BatchTable().Update(ctx, batchInfo); err != nil {
		return nil, err
	}

	if err = sdkCtx.EventManager().EmitTypedEvent(&types.EventUpdateBatchMetadata{
		BatchDenom: req.BatchDenom,
	}); err != nil {
		return nil, err
	}

	return &types.MsgUpdateBatchMetadataResponse{}, err
}
