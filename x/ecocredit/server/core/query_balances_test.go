package core

import (
	"context"
	"testing"

	"gotest.tools/v3/assert"

	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

func TestQuery_Balances(t *testing.T) {
	t.Parallel()
	s := setupBase(t)

	bKey1, err := s.stateStore.BatchTable().InsertReturningID(s.ctx, &api.Batch{Denom: "C01-20200101-20220101-001"})
	assert.NilError(t, err)
	bKey2, err := s.stateStore.BatchTable().InsertReturningID(s.ctx, &api.Batch{Denom: "C02-20200101-20220101-001"})
	assert.NilError(t, err)

	balance1 := &api.BatchBalance{Address: s.addr, BatchKey: bKey1, TradableAmount: "15", RetiredAmount: "15", EscrowedAmount: "15"}
	balance2 := &api.BatchBalance{Address: s.addr, BatchKey: bKey2, TradableAmount: "19", RetiredAmount: "20", EscrowedAmount: "33"}

	assert.NilError(t, s.stateStore.BatchBalanceTable().Insert(s.ctx, balance1))
	assert.NilError(t, s.stateStore.BatchBalanceTable().Insert(s.ctx, balance2))

	// query balances for s.addr
	res, err := s.k.Balances(s.ctx, &core.QueryBalancesRequest{
		Account:    s.addr.String(),
		Pagination: &query.PageRequest{Limit: 1, CountTotal: true},
	})
	assert.NilError(t, err)
	assert.Equal(t, 1, len(res.Balances))
	assertBalanceEqual(t, s.ctx, s.k, res.Balances[0], balance1)
	assert.Equal(t, uint64(2), res.Pagination.Total)

	_, _, noBalAddr := testdata.KeyTestPubAddr()

	// query balances for address with no balance
	res, err = s.k.Balances(s.ctx, &core.QueryBalancesRequest{
		Account: noBalAddr.String(),
	})
	assert.NilError(t, err)
	assert.Equal(t, 0, len(res.Balances))
}

func assertBalanceEqual(t *testing.T, ctx context.Context, k Keeper, received *core.BatchBalanceInfo, balance *api.BatchBalance) {
	addr := sdk.AccAddress(balance.Address)

	batch, err := k.stateStore.BatchTable().Get(ctx, balance.BatchKey)
	assert.NilError(t, err)

	info := core.BatchBalanceInfo{
		Address:        addr.String(),
		BatchDenom:     batch.Denom,
		TradableAmount: balance.TradableAmount,
		RetiredAmount:  balance.RetiredAmount,
		EscrowedAmount: balance.EscrowedAmount,
	}

	assert.DeepEqual(t, info, *received)
}
