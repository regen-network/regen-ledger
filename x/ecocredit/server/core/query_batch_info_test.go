package core

import (
	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	ecocreditv1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	v1 "github.com/regen-network/regen-ledger/x/ecocredit/v1"
	"gotest.tools/v3/assert"
	"testing"
)

func TestQuery_BatchInfo(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	batchDenom := "C01-20200101-20220101-001"
	assert.NilError(t, s.stateStore.BatchInfoStore().Insert(s.ctx, &ecocreditv1.BatchInfo{
		ProjectId:  1,
		BatchDenom: batchDenom,
		Metadata:   nil,
		StartDate:  nil,
		EndDate:    nil,
	}))

	// invalid query
	_, err := s.k.BatchInfo(s.ctx, &v1.QueryBatchInfoRequest{BatchDenom: "A00-00000000-00000000-000"})
	assert.ErrorContains(t, err, ormerrors.NotFound.Error())

	// good query
	res, err := s.k.BatchInfo(s.ctx, &v1.QueryBatchInfoRequest{BatchDenom: batchDenom})
	assert.NilError(t, err)
	assert.Equal(t, uint64(1), res.Info.ProjectId)
}
