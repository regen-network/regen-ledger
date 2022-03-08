package core

import (
	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	ecocreditv1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	v1 "github.com/regen-network/regen-ledger/x/ecocredit/v1"
	"gotest.tools/v3/assert"
	"testing"
)

func TestQuery_Supply(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	batchDenom := "C01-20200101-20220101-001"
	tradable := "10.54321"
	retired := "50.3214"
	cancelled := "0.3215"

	// make a batch and some supply
	assert.NilError(t, s.stateStore.BatchInfoStore().Insert(s.ctx, &ecocreditv1.BatchInfo{
		ProjectId:  1,
		BatchDenom: batchDenom,
		Metadata:   nil,
		StartDate:  nil,
		EndDate:    nil,
	}))
	assert.NilError(t, s.stateStore.BatchSupplyStore().Insert(s.ctx, &ecocreditv1.BatchSupply{
		BatchId:         1,
		TradableAmount:  tradable,
		RetiredAmount:   retired,
		CancelledAmount: cancelled,
	}))

	// valid query
	res, err := s.k.Supply(s.ctx, &v1.QuerySupplyRequest{BatchDenom: batchDenom})
	assert.NilError(t, err)
	assert.Equal(t, tradable, res.TradableSupply)
	assert.Equal(t, retired, res.RetiredSupply)
	assert.Equal(t, cancelled, res.CancelledAmount)

	// bad denom query
	_, err = s.k.Supply(s.ctx, &v1.QuerySupplyRequest{BatchDenom: "A00-00000000-00000000-001"})
	assert.ErrorContains(t, err, ormerrors.NotFound.Error())
}
