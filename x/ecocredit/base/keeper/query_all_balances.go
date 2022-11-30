package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/orm/model/ormlist"
	sdk "github.com/cosmos/cosmos-sdk/types"

	api "github.com/regen-network/regen-ledger/api/v2/regen/ecocredit/v1"
	regenerrors "github.com/regen-network/regen-ledger/types/v2/errors"
	"github.com/regen-network/regen-ledger/types/v2/ormutil"
	types "github.com/regen-network/regen-ledger/x/ecocredit/v3/base/types/v1"
)

// AllBalances queries all credit balances from state with optional pagination.
// NOTE: If no pagination is given in the request, responses will be limited by the Cosmos SDK's default limit (100).
func (k Keeper) AllBalances(ctx context.Context, req *types.QueryAllBalancesRequest) (*types.QueryAllBalancesResponse, error) {
	pg, err := ormutil.GogoPageReqToPulsarPageReq(req.Pagination)
	if err != nil {
		return nil, regenerrors.ErrInvalidArgument.Wrapf(err.Error())
	}
	it, err := k.stateStore.BatchBalanceTable().List(ctx, api.BatchBalancePrimaryKey{}, ormlist.Paginate(pg))
	if err != nil {
		return nil, err
	}
	defer it.Close()

	var res types.QueryAllBalancesResponse
	for it.Next() {
		balance, err := it.Value()
		if err != nil {
			return nil, err
		}

		batch, err := k.stateStore.BatchTable().Get(ctx, balance.BatchKey)
		if err != nil {
			return nil, regenerrors.ErrNotFound.Wrapf("batch with key: %d", balance.BatchKey)
		}

		res.Balances = append(res.Balances, &types.BatchBalanceInfo{
			Address:        sdk.AccAddress(balance.Address).String(),
			BatchDenom:     batch.Denom,
			TradableAmount: balance.TradableAmount,
			RetiredAmount:  balance.RetiredAmount,
			EscrowedAmount: balance.EscrowedAmount,
		})
	}
	res.Pagination, err = ormutil.PulsarPageResToGogoPageRes(it.PageResponse())
	if err != nil {
		return nil, regenerrors.ErrInternal.Wrap(err.Error())
	}
	return &res, nil
}
