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
	classKey, err := s.stateStore.ClassInfoTable().InsertReturningID(s.ctx, &api.ClassInfo{
		Id: "C01",
	})
	assert.NilError(t, err)

	// insert project
	err = s.stateStore.ProjectInfoTable().Insert(s.ctx, &api.ProjectInfo{
		Id:            "P01",
		ClassKey:         classKey,
		ProjectJurisdiction: "US-CA",
		Metadata:        "data",
	})
	assert.NilError(t, err)

	// valid query
	res, err := s.k.ProjectInfo(s.ctx, &core.QueryProjectInfoRequest{ProjectId: "P01"})
	assert.NilError(t, err)
	assert.Equal(t, "P01", res.Project.Id)
	assert.Equal(t, "C01", res.Project.ClassId)
	assert.Equal(t, "US-CA", res.Project.ProjectLocation)
	assert.Equal(t, "data", res.Project.Metadata)

	// invalid query
	_, err = s.k.ProjectInfo(s.ctx, &core.QueryProjectInfoRequest{ProjectId: "F01"})
	assert.ErrorContains(t, err, ormerrors.NotFound.Error())
}
