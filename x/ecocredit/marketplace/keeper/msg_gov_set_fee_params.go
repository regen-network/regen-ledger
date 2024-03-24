package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	marketplacev1 "github.com/regen-network/regen-ledger/api/v2/regen/ecocredit/marketplace/v1"
	types "github.com/regen-network/regen-ledger/x/ecocredit/v3/marketplace/types/v1"
)

func (k Keeper) GovSetFeeParams(ctx context.Context, msg *types.MsgGovSetFeeParams) (*types.MsgGovSetFeeParamsResponse, error) {
	authority, err := sdk.AccAddressFromBech32(msg.Authority)
	if err != nil {
		return nil, err
	}

	if !authority.Equals(k.authority) {
		return nil, sdkerrors.ErrUnauthorized
	}

	// convert from gogo to protoreflect
	var feeParams marketplacev1.FeeParams
	err = gogoToProtoReflect(msg.Fees, &feeParams)
	if err != nil {
		return nil, err
	}

	err = k.stateStore.FeeParamsTable().Save(ctx, &feeParams)
	if err != nil {
		return nil, err
	}

	return &types.MsgGovSetFeeParamsResponse{}, nil
}
