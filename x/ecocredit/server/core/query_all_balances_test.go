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

	// using these addresses here to have consistent ordering
	addr1 := sdk.AccAddress("foo")
	addr2 := sdk.AccAddress("bar")

	bKey1, err := s.stateStore.BatchTable().InsertReturningID(s.ctx, &api.Batch{Denom: "C01-001-20200101-20220101-001"})
	assert.NilError(t, err)
	bKey2, err := s.stateStore.BatchTable().InsertReturningID(s.ctx, &api.Batch{Denom: "C02-001-20200101-20220101-001"})
	assert.NilError(t, err)

	balance1 := &api.BatchBalance{Address: addr1, BatchKey: bKey1, TradableAmount: "15", RetiredAmount: "15", EscrowedAmount: "15"}
	balance2 := &api.BatchBalance{Address: addr1, BatchKey: bKey2, TradableAmount: "1", RetiredAmount: "2", EscrowedAmount: "3"}
	balance3 := &api.BatchBalance{Address: addr2, BatchKey: bKey2, TradableAmount: "19", RetiredAmount: "20", EscrowedAmount: "33"}

	balances := []*api.BatchBalance{balance3, balance1, balance2} // the order that ORM sorts the entries

	assert.NilError(t, s.stateStore.BatchBalanceTable().Insert(s.ctx, balance1))
	assert.NilError(t, s.stateStore.BatchBalanceTable().Insert(s.ctx, balance2))
	assert.NilError(t, s.stateStore.BatchBalanceTable().Insert(s.ctx, balance3))

	res, err := s.k.AllBalances(s.ctx, &core.QueryAllBalancesRequest{Pagination: &query.PageRequest{Limit: 10, CountTotal: true}})
	assert.NilError(t, err)
	assert.Equal(t, len(res.Balances), 3)
	assert.Equal(t, res.Pagination.Total, uint64(3))

	for i, bal := range res.Balances {
		assertBalanceEqual(s.ctx, t, s.k, bal, balances[i])
	}
}