package core

import (
	"testing"

	"gotest.tools/v3/assert"

	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	"github.com/cosmos/cosmos-sdk/testutil/testdata"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

func TestQuery_Balance(t *testing.T) {
	t.Parallel()
	s := setupBase(t)

	batchDenom := "C01-001-20200101-20210101-001"

	// insert batch
	bKey, err := s.stateStore.BatchTable().InsertReturningID(s.ctx, &api.Batch{
		Denom: batchDenom,
	})
	assert.NilError(t, err)

	balance := &api.BatchBalance{
		BatchKey: bKey,
		Address:  s.addr,
		Tradable: "10.54321",
		Retired:  "50.3214",
	}

	// insert balance for s.addr
	assert.NilError(t, s.stateStore.BatchBalanceTable().Insert(s.ctx, balance))

	// query balance for s.addr
	res, err := s.k.Balance(s.ctx, &core.QueryBalanceRequest{
		Account:    s.addr.String(),
		BatchDenom: batchDenom,
	})
	assert.NilError(t, err)
	assert.Equal(t, s.addr.String(), res.Balance.Address)
	assert.Equal(t, batchDenom, res.Balance.BatchDenom)
	assert.Equal(t, balance.Tradable, res.Balance.Tradable)
	assert.Equal(t, balance.Retired, res.Balance.Retired)

	_, _, noBalance := testdata.KeyTestPubAddr()

	// query balance for address with no balance
	res, err = s.k.Balance(s.ctx, &core.QueryBalanceRequest{
		Account:    noBalance.String(),
		BatchDenom: batchDenom,
	})
	assert.NilError(t, err)
	assert.Equal(t, noBalance.String(), res.Balance.Address)
	assert.Equal(t, batchDenom, res.Balance.BatchDenom)
	assert.Equal(t, "0", res.Balance.Tradable)
	assert.Equal(t, "0", res.Balance.Retired)

	// query balance with unknown batch denom
	_, err = s.k.Balance(s.ctx, &core.QueryBalanceRequest{
		Account:    s.addr.String(),
		BatchDenom: "A00-00000000-00000000-001",
	})
	assert.ErrorContains(t, err, ormerrors.NotFound.Error())
}
