package core

import (
	"testing"

	"gotest.tools/v3/assert"

	"github.com/cosmos/cosmos-sdk/testutil/testdata"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

func TestQuery_Projects_By_Admin(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	_, _, admin2 := testdata.KeyTestPubAddr()

	// insert class
	classKey, err := s.stateStore.ClassTable().InsertReturningID(s.ctx, &api.Class{
		Id: "C01",
	})
	assert.NilError(t, err)

	// create two projects
	project := &api.Project{
		Id:           "C01-001",
		ClassKey:     classKey,
		Admin:        s.addr,
		Jurisdiction: "US-CA",
		Metadata:     "data",
	}

	err = s.stateStore.ProjectTable().Insert(s.ctx, project)
	assert.NilError(t, err)

	project = &api.Project{
		Id:           "C01-002",
		ClassKey:     classKey,
		Admin:        s.addr,
		Jurisdiction: "US-CA",
		Metadata:     "data",
	}

	err = s.stateStore.ProjectTable().Insert(s.ctx, project)
	assert.NilError(t, err)

	// create project with different admin
	project = &api.Project{
		Id:           "C01-003",
		ClassKey:     classKey,
		Admin:        admin2,
		Jurisdiction: "US-CA",
		Metadata:     "data",
	}

	err = s.stateStore.ProjectTable().Insert(s.ctx, project)
	assert.NilError(t, err)

	// query project by admin admin1 expect 2 projects
	res, err := s.k.ProjectsByAdmin(s.ctx, &core.QueryProjectsByAdminRequest{Admin: s.addr.String()})
	assert.NilError(t, err)
	assert.Equal(t, len(res.Projects), 2)

	// query project by admin admin2 expect 1 project
	res, err = s.k.ProjectsByAdmin(s.ctx, &core.QueryProjectsByAdminRequest{Admin: admin2.String()})
	assert.NilError(t, err)
	assert.Equal(t, len(res.Projects), 1)
}
