package core

import (
	"testing"

	"gotest.tools/v3/assert"

	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

func TestQuery_ProjectInfo(t *testing.T) {
	t.Parallel()
	s := setupBase(t)

	// insert class
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

	// insert project
	err = s.stateStore.ProjectTable().Insert(s.ctx, project)
	assert.NilError(t, err)

	// query project by "P01" project id
	res, err := s.k.ProjectInfo(s.ctx, &core.QueryProjectInfoRequest{ProjectId: "P01"})
	assert.NilError(t, err)
	assert.Equal(t, project.Id, res.Project.Id)
	assert.Equal(t, "C01", res.Project.ClassId)
	assert.Equal(t, project.ProjectJurisdiction, res.Project.Jurisdiction)
	assert.Equal(t, project.Metadata, res.Project.Metadata)

	// query project by unknown project id
	_, err = s.k.ProjectInfo(s.ctx, &core.QueryProjectInfoRequest{ProjectId: "F01"})
	assert.ErrorContains(t, err, ormerrors.NotFound.Error())
}
