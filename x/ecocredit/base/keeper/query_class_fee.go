package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	regentypes "github.com/regen-network/regen-ledger/types/v2"
	regenerrors "github.com/regen-network/regen-ledger/types/v2/errors"
	types "github.com/regen-network/regen-ledger/x/ecocredit/v3/base/types/v1"
)

// ClassFee queries credit class creation fees.
func (k Keeper) ClassFee(ctx context.Context, _ *types.QueryClassFeeRequest) (*types.QueryClassFeeResponse, error) {
	classFee, err := k.stateStore.ClassFeeTable().Get(ctx)
	if err != nil {
		return nil, regenerrors.ErrInternal.Wrapf("failed to get class fee: %s", err.Error())
	}

	var fee sdk.Coin
	if classFee.Fee != nil {
		fee = regentypes.CoinFromCosmosApiLegacy(classFee.Fee)
	}

	return &types.QueryClassFeeResponse{Fee: &fee}, nil
}
