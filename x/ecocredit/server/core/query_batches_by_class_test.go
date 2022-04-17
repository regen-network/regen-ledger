package core

import (
	"strings"
	"testing"

	"gotest.tools/v3/assert"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

func TestQuery_BatchesByClass(t *testing.T) {
	t.Parallel()
	s := setupBase(t)

	// insert class
	assert.NilError(t, s.stateStore.ClassInfoTable().Insert(s.ctx, &api.ClassInfo{
		Name: "C01",
	}))

	// insert project
	projectKey, err := s.stateStore.ProjectInfoTable().InsertReturningID(s.ctx, &api.ProjectInfo{
		Name: "P01",
	})
	assert.NilError(t, err)

	// insert three batches under the project
	assert.NilError(t, s.stateStore.BatchInfoTable().Insert(s.ctx, &api.BatchInfo{
		ProjectId:  projectKey,
		BatchDenom: "C01-20200101-20200102-001",
		Metadata:   "",
		StartDate:  nil,
		EndDate:    nil,
	}))
	assert.NilError(t, s.stateStore.BatchInfoTable().Insert(s.ctx, &api.BatchInfo{
		ProjectId:  projectKey,
		BatchDenom: "C01-20190203-20200102-002",
		Metadata:   "",
		StartDate:  nil,
		EndDate:    nil,
	}))
	assert.NilError(t, s.stateStore.BatchInfoTable().Insert(s.ctx, &api.BatchInfo{
		ProjectId:  projectKey,
		BatchDenom: "C01-20500404-20900102-003",
		Metadata:   "",
		StartDate:  nil,
		EndDate:    nil,
	}))

	// classes that SHOULD NOT show up from a query for "C01"
	assert.NilError(t, s.stateStore.BatchInfoTable().Insert(s.ctx, &api.BatchInfo{
		ProjectId:  projectKey,
		BatchDenom: "C011-20500404-20900102-003",
		Metadata:   "",
		StartDate:  nil,
		EndDate:    nil,
	}))
	assert.NilError(t, s.stateStore.BatchInfoTable().Insert(s.ctx, &api.BatchInfo{
		ProjectId:  projectKey,
		BatchDenom: "BIO1-20500404-20900102-003",
		Metadata:   "",
		StartDate:  nil,
		EndDate:    nil,
	}))

	res, err := s.k.BatchesByClass(s.ctx, &core.QueryBatchesByClassRequest{
		ClassId:    "C01",
		Pagination: nil,
	})
	assert.NilError(t, err)
	assert.Equal(t, 3, len(res.Batches))
	for _, batch := range res.Batches {
		assert.Check(t, strings.Contains(batch.BatchDenom, "C01"))
	}
}
