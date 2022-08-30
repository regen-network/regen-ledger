package keeper

import (
	"context"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	regentypes "github.com/regen-network/regen-ledger/types"
	types "github.com/regen-network/regen-ledger/x/ecocredit/base/types/v1"
)

// CreditClassFees queries credit class creation fees.
func (k Keeper) CreditClassFees(ctx context.Context, request *types.QueryCreditClassFeesRequest) (*types.QueryCreditClassFeesResponse, error) {
	result, err := k.stateStore.ClassFeesTable().Get(ctx)
	if err != nil {
		return nil, err
	}

	classFee, ok := regentypes.ProtoCoinsToCoins(result.Fees)
	if !ok {
		return nil, sdkerrors.ErrInvalidType.Wrap("credit class fee")
	}

	return &types.QueryCreditClassFeesResponse{
		Fees: classFee,
	}, nil
}
