package core

import (
	"testing"

	"gotest.tools/v3/assert"

	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	"github.com/cosmos/cosmos-sdk/types/query"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
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
		Id:                  "P01",
		ClassKey:            classKey,
		ProjectJurisdiction: "US-CA",
		Metadata:            "data",
	}

	// insert two projects under "C01" credit class
	assert.NilError(t, s.stateStore.ProjectTable().Insert(s.ctx, project))
	assert.NilError(t, s.stateStore.ProjectTable().Insert(s.ctx, &api.Project{
		Id:       "P02",
		ClassKey: classKey,
	}))

	// query projects by "C01" credit class
	res, err := s.k.Projects(s.ctx, &core.QueryProjectsRequest{
		ClassId:    "C01",
		Pagination: &query.PageRequest{Limit: 1, CountTotal: true},
	})
	assert.NilError(t, err)
	assert.Equal(t, 1, len(res.Projects))
	assert.Equal(t, project.Id, res.Projects[0].Id)
	assert.Equal(t, "C01", res.Projects[0].ClassId)
	assert.Equal(t, project.ProjectJurisdiction, res.Projects[0].Jurisdiction)
	assert.Equal(t, uint64(2), res.Pagination.Total)

	// query projects by unknown credit class
	_, err = s.k.Projects(s.ctx, &core.QueryProjectsRequest{ClassId: "F01"})
	assert.ErrorContains(t, err, ormerrors.NotFound.Error())
}
