package keeper

import (
	"testing"

	"gotest.tools/v3/assert"

	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	"github.com/cosmos/cosmos-sdk/types/query"

	api "github.com/regen-network/regen-ledger/api/v2/regen/ecocredit/v1"
	types "github.com/regen-network/regen-ledger/x/ecocredit/v3/base/types/v1"
)

func TestQuery_ProjectsByClass(t *testing.T) {
	t.Parallel()
	s := setupBase(t)

	// insert credit class
	clsId, err := s.stateStore.ClassTable().InsertReturningID(s.ctx, &api.Class{
		Id: "C01",
	})
	assert.NilError(t, err)

	project := &api.Project{
		Id:           "C01-001",
		Jurisdiction: "US-CA",
		Metadata:     "metadata",
	}

	// insert two projects under "C01" credit class
	id, err := s.stateStore.ProjectTable().InsertReturningID(s.ctx, project)
	assert.NilError(t, err)
	err = s.stateStore.ProjectEnrollmentTable().Insert(s.ctx, &api.ProjectEnrollment{
		ProjectKey: id,
		ClassKey:   clsId,
		Status:     api.ProjectEnrollmentStatus_PROJECT_ENROLLMENT_STATUS_ACCEPTED,
	})
	assert.NilError(t, err)
	id, err = s.stateStore.ProjectTable().InsertReturningID(s.ctx, &api.Project{
		Id: "C01-002",
	})
	assert.NilError(t, err)
	err = s.stateStore.ProjectEnrollmentTable().Insert(s.ctx, &api.ProjectEnrollment{
		ProjectKey: id,
		ClassKey:   clsId,
		Status:     api.ProjectEnrollmentStatus_PROJECT_ENROLLMENT_STATUS_ACCEPTED,
	})
	assert.NilError(t, err)

	// query projects by "C01" credit class
	res, err := s.k.ProjectsByClass(s.ctx, &types.QueryProjectsByClassRequest{
		ClassId:    "C01",
		Pagination: &query.PageRequest{Limit: 1, CountTotal: true},
	})
	assert.NilError(t, err)

	// check pagination
	assert.Equal(t, 1, len(res.Projects))
	assert.Equal(t, uint64(2), res.Pagination.Total)

	// check project properties
	assert.Equal(t, project.Id, res.Projects[0].Id)
	assert.Equal(t, project.Jurisdiction, res.Projects[0].Jurisdiction)
	assert.Equal(t, project.Metadata, res.Projects[0].Metadata)

	// query projects by unknown credit class
	_, err = s.k.ProjectsByClass(s.ctx, &types.QueryProjectsByClassRequest{ClassId: "F01"})
	assert.ErrorContains(t, err, ormerrors.NotFound.Error())
}
