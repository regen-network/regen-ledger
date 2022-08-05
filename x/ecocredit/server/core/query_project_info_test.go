package core

import (
	"testing"

	"gotest.tools/v3/assert"

	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

func TestQuery_Project(t *testing.T) {
	t.Parallel()
	s := setupBase(t)

	// insert class
	classKey, err := s.stateStore.ClassTable().InsertReturningID(s.ctx, &api.Class{
		Id: "C01",
	})
	assert.NilError(t, err)

	project := &api.Project{
		Id:           "C01-001",
		ClassKey:     classKey,
		Jurisdiction: "US-CA",
		Metadata:     "data",
		ReferenceId:  "R01",
	}

	// insert project
	err = s.stateStore.ProjectTable().Insert(s.ctx, project)
	assert.NilError(t, err)

	// query project by "C01-001" project id
	res, err := s.k.Project(s.ctx, &core.QueryProjectRequest{ProjectId: "C01-001"})
	assert.NilError(t, err)
	assert.Equal(t, project.Id, res.Project.Id)
	assert.Equal(t, "C01", res.Project.ClassId)
	assert.Equal(t, project.Jurisdiction, res.Project.Jurisdiction)
	assert.Equal(t, project.Metadata, res.Project.Metadata)
	assert.Equal(t, project.ReferenceId, res.Project.ReferenceId)

	// query project by unknown project id
	_, err = s.k.Project(s.ctx, &core.QueryProjectRequest{ProjectId: "F01"})
	assert.ErrorContains(t, err, ormerrors.NotFound.Error())
}
