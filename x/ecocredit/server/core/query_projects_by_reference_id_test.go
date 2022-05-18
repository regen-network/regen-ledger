package core

import (
	"testing"

	"gotest.tools/v3/assert"

	"github.com/cosmos/cosmos-sdk/types/query"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

func TestQuery_ProjectsByReferenceId(t *testing.T) {
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
		Metadata:     "data",
		ReferenceId:  "R1",
	}

	// insert two projects with "R1" reference id and one without reference id
	assert.NilError(t, s.stateStore.ProjectTable().Insert(s.ctx, project))
	assert.NilError(t, s.stateStore.ProjectTable().Insert(s.ctx, &api.Project{
		Id:          "C01-002",
		ClassKey:    classKey,
		ReferenceId: "R1",
	}))
	assert.NilError(t, s.stateStore.ProjectTable().Insert(s.ctx, &api.Project{
		Id:       "C01-003",
		ClassKey: classKey,
	}))

	// query projects by "R1" reference id
	res, err := s.k.ProjectsByReferenceId(s.ctx, &core.QueryProjectsByReferenceIdRequest{
		ReferenceId: "R1",
		Pagination:  &query.PageRequest{Limit: 1, CountTotal: true},
	})
	assert.NilError(t, err)
	assert.Equal(t, 1, len(res.Projects))
	assert.Equal(t, project.Id, res.Projects[0].Id)
	assert.Equal(t, "C01", res.Projects[0].ClassId)
	assert.Equal(t, "R1", res.Projects[0].ReferenceId)
	assert.Equal(t, project.Jurisdiction, res.Projects[0].Jurisdiction)
	assert.Equal(t, uint64(2), res.Pagination.Total)

	// query projects by unknown reference id
	res, err = s.k.ProjectsByReferenceId(s.ctx, &core.QueryProjectsByReferenceIdRequest{ReferenceId: "RR2"})
	assert.Equal(t, len(res.Projects), 0)
	assert.NilError(t, err)
}
