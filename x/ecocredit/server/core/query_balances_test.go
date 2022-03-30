package core

import (
	"testing"

	"gotest.tools/v3/assert"

	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	"github.com/cosmos/cosmos-sdk/types/query"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/types/ormutil"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

func TestQuery_Balances(t *testing.T) {
	t.Parallel()
	s := setupBase(t)

	balance1 := &api.BatchBalance{Address: s.addr, BatchId: 1, Tradable: "15", Retired: "15", Escrowed: "15"}
	balance2 := &api.BatchBalance{Address: s.addr, BatchId: 2, Tradable: "19", Retired: "20", Escrowed: "33"}
	assert.NilError(t, s.stateStore.BatchBalanceTable().Insert(s.ctx, balance1))
	assert.NilError(t, s.stateStore.BatchBalanceTable().Insert(s.ctx, balance2))
	assert.NilError(t, s.stateStore.BatchBalanceTable().Insert(s.ctx, &api.BatchBalance{
		Address:  s.addr,
		BatchId:  3,
		Tradable: "4",
		Retired:  "5",
		Escrowed: "6",
	}))

	res, err := s.k.Balances(s.ctx, &core.QueryBalancesRequest{Account: s.addr.String(), Pagination: &query.PageRequest{CountTotal: true, Limit: 2}})
	assert.NilError(t, err)
	assert.Equal(t, 2, len(res.Balances))
	assertBalanceEqual(t, res.Balances[0], balance1)
	assertBalanceEqual(t, res.Balances[1], balance2)
	assert.Equal(t, uint64(3), res.Pagination.Total)

	_, _, noBalAddr := testdata.KeyTestPubAddr()
	res, err = s.k.Balances(s.ctx, &core.QueryBalancesRequest{Account: noBalAddr.String()})
	assert.NilError(t, err)
	assert.Equal(t, 0, len(res.Balances))
}

func assertBalanceEqual(t *testing.T, balance *core.BatchBalance, batchBalance *api.BatchBalance) {
	var bal core.BatchBalance
	assert.NilError(t, ormutil.PulsarToGogoSlow(batchBalance, &bal))
	assert.DeepEqual(t, bal, *balance)
}
