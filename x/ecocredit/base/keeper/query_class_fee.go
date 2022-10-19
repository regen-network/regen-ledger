package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	regentypes "github.com/regen-network/regen-ledger/types"
	types "github.com/regen-network/regen-ledger/x/ecocredit/base/types/v1"
)

// ClassFee queries credit class creation fees.
func (k Keeper) ClassFee(ctx context.Context, request *types.QueryClassFeeRequest) (*types.QueryClassFeeResponse, error) {
	classFee, err := k.stateStore.ClassFeeTable().Get(ctx)
	if err != nil {
		return nil, err
	}

	var fee sdk.Coin
	if classFee.Fee != nil {
		var ok bool
		fee, ok = regentypes.ProtoCoinToCoin(classFee.Fee)
		if !ok {
			return nil, sdkerrors.ErrInvalidType.Wrapf("failed to parse class fee")
		}
	}

	return &types.QueryClassFeeResponse{
		Fee: &fee,
	}, nil
}
