package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	regenerrors "github.com/regen-network/regen-ledger/errors"
	types "github.com/regen-network/regen-ledger/x/ecocredit/base/types/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/server/utils"
)

// Balance queries the balance (both tradable and retired) of a given credit
// batch for a given account.
func (k Keeper) Balance(ctx context.Context, req *types.QueryBalanceRequest) (*types.QueryBalanceResponse, error) {
	batch, err := k.stateStore.BatchTable().GetByDenom(ctx, req.BatchDenom)
	if err != nil {
		return nil, regenerrors.ErrNotFound.Wrapf("could not get batch with denom %s: %s", req.BatchDenom, err.Error())
	}

	addr, err := sdk.AccAddressFromBech32(req.Address)
	if err != nil {
		return nil, regenerrors.ErrInvalidArgument.Wrap(err.Error())
	}

	balance, err := utils.GetBalance(ctx, k.stateStore.BatchBalanceTable(), addr, batch.Key)
	if err != nil {
		return nil, regenerrors.ErrNotFound.Wrapf("unable to get balance with batch key: %d", batch.Key)
	}

	info := types.BatchBalanceInfo{
		Address:        addr.String(),
		BatchDenom:     batch.Denom,
		TradableAmount: balance.TradableAmount,
		RetiredAmount:  balance.RetiredAmount,
		EscrowedAmount: balance.EscrowedAmount,
	}

	return &types.QueryBalanceResponse{Balance: &info}, nil
}
