package keeper

import (
	"testing"

	"gotest.tools/v3/assert"

	"github.com/cosmos/cosmos-sdk/types/query"

	api "github.com/regen-network/regen-ledger/api/v2/regen/ecocredit/v1"
	types "github.com/regen-network/regen-ledger/x/ecocredit/v3/base/types/v1"
)

func TestQuery_Projects(t *testing.T) {
	t.Parallel()
	s := setupBase(t)

	// insert credit class
	classKey, err := s.stateStore.ClassTable().InsertReturningID(s.ctx, &api.Class{
		Id: "C01",
	})
	assert.NilError(t, err)

	project := &api.Project{
		Id:           "C01-001",
		ClassKey:     classKey,
		Jurisdiction: "US-CA",
		Metadata:     "metadata",
	}

	// insert two projects
	assert.NilError(t, s.stateStore.ProjectTable().Insert(s.ctx, project))
	assert.NilError(t, s.stateStore.ProjectTable().Insert(s.ctx, &api.Project{
		Id:       "C01-002",
		ClassKey: classKey,
	}))

	// query projects with pagination
	res, err := s.k.Projects(s.ctx, &types.QueryProjectsRequest{
		Pagination: &query.PageRequest{Limit: 1, CountTotal: true},
	})
	assert.NilError(t, err)

	// check pagination
	assert.Equal(t, 1, len(res.Projects))
	assert.Equal(t, uint64(2), res.Pagination.Total)

	// check project properties
	assert.Equal(t, project.Id, res.Projects[0].Id)
	assert.Equal(t, "C01", res.Projects[0].ClassId)
	assert.Equal(t, project.Jurisdiction, res.Projects[0].Jurisdiction)
	assert.Equal(t, project.Metadata, res.Projects[0].Metadata)
}
