package core

import (
	"context"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

// CreditClassFees queries credit class creation fees.
func (k Keeper) CreditClassFees(ctx context.Context, request *core.QueryCreditClassFeesRequest) (*core.QueryCreditClassFeesResponse, error) {
	result, err := k.stateStore.ClassFeesTable().Get(ctx)
	if err != nil {
		return nil, err
	}

	classFee, ok := types.ProtoCoinsToCoins(result.Fees)
	if !ok {
		return nil, sdkerrors.ErrInvalidType.Wrap("credit class fee")
	}

	return &core.QueryCreditClassFeesResponse{
		Fees: classFee,
	}, nil
}
