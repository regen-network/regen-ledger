package core

import (
	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	ecocreditv1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	v1 "github.com/regen-network/regen-ledger/x/ecocredit/v1"
	"gotest.tools/v3/assert"
	"testing"
)

func TestQuery_Balance(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	noBalanceAddr := genAddrs(1)[0]
	batchDenom := "C01-20200101-20220101-001"
	tradable := "10.54321"
	retired := "50.3214"

	// make a batch and give s.addr some balance
	assert.NilError(t, s.stateStore.BatchInfoStore().Insert(s.ctx, &ecocreditv1.BatchInfo{
		ProjectId:  1,
		BatchDenom: batchDenom,
		Metadata:   nil,
		StartDate:  nil,
		EndDate:    nil,
	}))
	assert.NilError(t, s.stateStore.BatchBalanceStore().Insert(s.ctx, &ecocreditv1.BatchBalance{
		Address:  s.addr,
		BatchId:  1,
		Tradable: tradable,
		Retired:  retired,
	}))

	// valid query
	res, err := s.k.Balance(s.ctx, &v1.QueryBalanceRequest{
		Account:    s.addr.String(),
		BatchDenom: batchDenom,
	})
	assert.NilError(t, err)
	assert.Equal(t, tradable, res.TradableAmount)
	assert.Equal(t, retired, res.RetiredAmount)

	// random addr should just give 0
	res, err = s.k.Balance(s.ctx, &v1.QueryBalanceRequest{
		Account:    noBalanceAddr.String(),
		BatchDenom: batchDenom,
	})
	assert.NilError(t, err)
	assert.Equal(t, "0", res.TradableAmount)
	assert.Equal(t, "0", res.RetiredAmount)

	// query with invalid batch should return not found
	_, err = s.k.Balance(s.ctx, &v1.QueryBalanceRequest{
		Account:    s.addr.String(),
		BatchDenom: "A00-00000000-00000000-001",
	})
	assert.ErrorContains(t, err, ormerrors.NotFound.Error())
}
