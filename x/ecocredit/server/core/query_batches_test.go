package core

import (
	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	"github.com/cosmos/cosmos-sdk/types/query"
	ecocreditv1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	v1 "github.com/regen-network/regen-ledger/x/ecocredit/v1"
	"gotest.tools/v3/assert"
	"testing"
)

func TestQuery_Batches(t *testing.T) {
	t.Parallel()
	s := setupBase(t)

	// make a project and two batches
	assert.NilError(t, s.stateStore.ProjectInfoStore().Insert(s.ctx, &ecocreditv1.ProjectInfo{
		Name:            "P01",
		ClassId:         1,
		ProjectLocation: "US-CA",
		Metadata:        nil,
	}))
	assert.NilError(t, s.stateStore.BatchInfoStore().Insert(s.ctx, &ecocreditv1.BatchInfo{
		ProjectId:  1,
		BatchDenom: "C01-20200101-20220101-001",
		Metadata:   nil,
		StartDate:  nil,
		EndDate:    nil,
	}))
	assert.NilError(t, s.stateStore.BatchInfoStore().Insert(s.ctx, &ecocreditv1.BatchInfo{
		ProjectId:  1,
		BatchDenom: "C01-20200101-20220101-002",
		Metadata:   nil,
		StartDate:  nil,
		EndDate:    nil,
	}))

	// invalid query
	_, err := s.k.Batches(s.ctx, &v1.QueryBatchesRequest{ProjectId: "F01"})
	assert.ErrorContains(t, err, ormerrors.NotFound.Error())

	//  base query
	res, err := s.k.Batches(s.ctx, &v1.QueryBatchesRequest{ProjectId: "P01"})
	assert.NilError(t, err)
	assert.Equal(t, 2, len(res.Batches))
	assert.Equal(t, "C01-20200101-20220101-001", res.Batches[0].BatchDenom)

	// paginated query
	res, err = s.k.Batches(s.ctx, &v1.QueryBatchesRequest{
		ProjectId: "P01",
		Pagination: &query.PageRequest{Limit: 1, CountTotal: true},
	})
	assert.NilError(t, err)
	assert.Equal(t, 1, len(res.Batches))
	assert.Equal(t, uint64(2), res.Pagination.Total)
}
