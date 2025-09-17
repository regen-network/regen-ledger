package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	marketplacev1 "github.com/regen-network/regen-ledger/api/v2/regen/ecocredit/marketplace/v1"
	types "github.com/regen-network/regen-ledger/x/ecocredit/v4/marketplace/types/v1"
)

func (k Keeper) GovSetFeeParams(ctx context.Context, msg *types.MsgGovSetFeeParams) (*types.MsgGovSetFeeParamsResponse, error) {
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	authorityBz, err := k.ac.StringToBytes(msg.Authority)
	if err != nil {
		return nil, fmt.Errorf("invalid authority address: %w", err)
	}

	authority := sdk.AccAddress(authorityBz)
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
