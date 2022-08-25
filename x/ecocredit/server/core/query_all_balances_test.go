package core

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"gotest.tools/v3/assert"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

func TestQueryAllBalances(t *testing.T) {
	t.Parallel()
	s := setupBase(t)

	bKey1, err := s.stateStore.BatchTable().InsertReturningID(s.ctx, &api.Batch{Denom: "C01-001-20200101-20220101-001"})
	assert.NilError(t, err)
	bKey2, err := s.stateStore.BatchTable().InsertReturningID(s.ctx, &api.Batch{Denom: "C02-001-20200101-20220101-001"})
	assert.NilError(t, err)

	balance1 := &api.BatchBalance{Address: s.addr, BatchKey: bKey1, TradableAmount: "15", RetiredAmount: "15", EscrowedAmount: "15"}
	balance2 := &api.BatchBalance{Address: s.addr2, BatchKey: bKey2, TradableAmount: "19", RetiredAmount: "20", EscrowedAmount: "33"}

	assert.NilError(t, s.stateStore.BatchBalanceTable().Insert(s.ctx, balance1))
	assert.NilError(t, s.stateStore.BatchBalanceTable().Insert(s.ctx, balance2))

	res, err := s.k.AllBalances(s.ctx, &core.QueryAllBalancesRequest{Pagination: &query.PageRequest{Limit: 2, CountTotal: true}})
	assert.NilError(t, err)
	assert.Equal(t, len(res.Balances), 2)
	assert.Equal(t, res.Pagination.Total, uint64(2))

	// response balances don't always come back in the same order, so check here before asserting equal
	if res.Balances[0].Address == sdk.AccAddress(balance1.Address).String() {
		assertBalanceEqual(s.ctx, t, s.k, res.Balances[0], balance1)
		assertBalanceEqual(s.ctx, t, s.k, res.Balances[1], balance2)
	} else {
		assertBalanceEqual(s.ctx, t, s.k, res.Balances[0], balance2)
		assertBalanceEqual(s.ctx, t, s.k, res.Balances[1], balance1)
	}
}
