package core

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"gotest.tools/v3/assert"

	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	"github.com/cosmos/cosmos-sdk/types/query"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

func TestQuery_Balances(t *testing.T) {
	t.Parallel()
	s := setupBase(t)

	b1, err := s.stateStore.BatchInfoTable().InsertReturningID(s.ctx, &api.BatchInfo{BatchDenom: "C01-20200101-20220101-001"})
	assert.NilError(t, err)
	b2, err := s.stateStore.BatchInfoTable().InsertReturningID(s.ctx, &api.BatchInfo{BatchDenom: "C02-20200101-20220101-001"})
	assert.NilError(t, err)
	b3, err := s.stateStore.BatchInfoTable().InsertReturningID(s.ctx, &api.BatchInfo{BatchDenom: "C03-20200101-20220101-001"})
	assert.NilError(t, err)

	balance1 := &api.BatchBalance{Address: s.addr, BatchId: b1, Tradable: "15", Retired: "15", Escrowed: "15"}
	balance2 := &api.BatchBalance{Address: s.addr, BatchId: b2, Tradable: "19", Retired: "20", Escrowed: "33"}

	assert.NilError(t, s.stateStore.BatchBalanceTable().Insert(s.ctx, balance1))
	assert.NilError(t, s.stateStore.BatchBalanceTable().Insert(s.ctx, balance2))
	assert.NilError(t, s.stateStore.BatchBalanceTable().Insert(s.ctx, &api.BatchBalance{
		Address:  s.addr,
		BatchId:  b3,
		Tradable: "4",
		Retired:  "5",
		Escrowed: "6",
	}))

	res, err := s.k.Balances(s.ctx, &core.QueryBalancesRequest{
		Account:    s.addr.String(),
		Pagination: &query.PageRequest{CountTotal: true, Limit: 2},
	})
	assert.NilError(t, err)
	assert.Equal(t, 2, len(res.Balances))
	assertBalanceEqual(t, s.ctx, s.k, res.Balances[0], balance1)
	assertBalanceEqual(t, s.ctx, s.k, res.Balances[1], balance2)
	assert.Equal(t, uint64(3), res.Pagination.Total)

	_, _, noBalAddr := testdata.KeyTestPubAddr()
	res, err = s.k.Balances(s.ctx, &core.QueryBalancesRequest{Account: noBalAddr.String()})
	assert.NilError(t, err)
	assert.Equal(t, 0, len(res.Balances))
}

func assertBalanceEqual(t *testing.T, ctx context.Context, k Keeper, received *core.BatchBalanceEntry, balance *api.BatchBalance) {
	addr := sdk.AccAddress(balance.Address)

	batch, err := k.stateStore.BatchInfoTable().Get(ctx, balance.BatchId)
	assert.NilError(t, err)

	entry := core.BatchBalanceEntry{
		Address:    addr.String(),
		BatchDenom: batch.BatchDenom,
		Tradable:   balance.Tradable,
		Retired:    balance.Retired,
		Escrowed:   balance.Escrowed,
	}

	assert.DeepEqual(t, entry, *received)
}
