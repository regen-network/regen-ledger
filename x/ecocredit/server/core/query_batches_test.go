package core

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	"gotest.tools/v3/assert"

	"github.com/cosmos/cosmos-sdk/types/query"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

func TestQuery_Batches(t *testing.T) {
	t.Parallel()
	s := setupBase(t)

	// insert project
	pKey, err := s.stateStore.ProjectTable().InsertReturningID(s.ctx, &api.Project{
		Id: "P01",
	})
	assert.NilError(t, err)

	batch := &api.Batch{
		ProjectKey: pKey,
		Denom:      "C01-20200101-20220101-001",
		Metadata:   "",
		StartDate:  nil,
		EndDate:    nil,
	}

	// insert two batches that are valid "P01" credit batches
	assert.NilError(t, s.stateStore.BatchTable().Insert(s.ctx, batch))
	assert.NilError(t, s.stateStore.BatchTable().Insert(s.ctx, &api.Batch{
		ProjectKey: pKey,
		Denom:      "C01-20200101-20220101-002",
	}))

	// query batches by "P01" project
	res, err := s.k.Batches(s.ctx, &core.QueryBatchesRequest{
		ProjectId:  "P01",
		Pagination: &query.PageRequest{Limit: 1, CountTotal: true},
	})
	assert.NilError(t, err)
	assert.Equal(t, 1, len(res.Batches))
	assert.Equal(t, uint64(2), res.Pagination.Total)

	// query by unknown project
	_, err = s.k.Batches(s.ctx, &core.QueryBatchesRequest{ProjectId: "F01"})
	assert.ErrorContains(t, err, ormerrors.NotFound.Error())
}
