package keeper

import (
	"testing"

	"gotest.tools/v3/assert"

	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	"github.com/stretchr/testify/require"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	types "github.com/regen-network/regen-ledger/x/ecocredit/base/types/v1"
)

func TestQuery_Supply(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	batchDenom := "C01-001-20200101-20220101-001"
	tradable := "10.54321"
	retired := "50.3214"
	cancelled := "0.3215"

	// make a batch and some supply
	assert.NilError(t, s.stateStore.BatchTable().Insert(s.ctx, &api.Batch{
		ProjectKey: 1,
		Denom:      batchDenom,
		Metadata:   "",
		StartDate:  nil,
		EndDate:    nil,
	}))

	// supply not found
	_, err := s.k.Supply(s.ctx, &types.QuerySupplyRequest{BatchDenom: batchDenom})
	require.Error(t, err)
	assert.Equal(t, "unable to get batch supply for batch: C01-001-20200101-20220101-001: invalid argument", err.Error())

	assert.NilError(t, s.stateStore.BatchSupplyTable().Insert(s.ctx, &api.BatchSupply{
		BatchKey:        1,
		TradableAmount:  tradable,
		RetiredAmount:   retired,
		CancelledAmount: cancelled,
	}))

	// valid query
	res, err := s.k.Supply(s.ctx, &types.QuerySupplyRequest{BatchDenom: batchDenom})
	assert.NilError(t, err)
	assert.Equal(t, tradable, res.TradableAmount)
	assert.Equal(t, retired, res.RetiredAmount)
	assert.Equal(t, cancelled, res.CancelledAmount)

	// bad denom query
	_, err = s.k.Supply(s.ctx, &types.QuerySupplyRequest{BatchDenom: "A00-00000000-00000000-001"})
	assert.ErrorContains(t, err, ormerrors.NotFound.Error())
}
