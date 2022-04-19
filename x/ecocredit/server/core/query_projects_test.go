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

	// create a class and 2 projects from it
	err := s.stateStore.ClassInfoTable().Insert(s.ctx, &api.ClassInfo{
		Id:               "C01",
		Admin:            s.addr,
		Metadata:         "",
		CreditTypeAbbrev: "C",
	})
	assert.NilError(t, err)
	err = s.stateStore.ProjectInfoTable().Insert(s.ctx, &api.ProjectInfo{
		Id:                  "P01",
		ClassKey:            1,
		ProjectJurisdiction: "US-CA",
		Metadata:            "",
	})
	assert.NilError(t, err)
	err = s.stateStore.ProjectInfoTable().Insert(s.ctx, &api.ProjectInfo{
		Id:                  "P02",
		ClassKey:            1,
		ProjectJurisdiction: "US-CA",
		Metadata:            "",
	})
	assert.NilError(t, err)

	// base query
	res, err := s.k.Projects(s.ctx, &core.QueryProjectsRequest{ClassId: "C01"})
	assert.NilError(t, err)
	assert.Equal(t, 2, len(res.Projects))
	assert.Equal(t, "US-CA", res.Projects[0].ProjectJurisdiction)

	// invalid query
	_, err = s.k.Projects(s.ctx, &core.QueryProjectsRequest{ClassId: "F01"})
	assert.ErrorContains(t, err, ormerrors.NotFound.Error())

	// paginated query
	res, err = s.k.Projects(s.ctx, &core.QueryProjectsRequest{
		ClassId:    "C01",
		Pagination: &query.PageRequest{Limit: 1, CountTotal: true},
	})
	assert.NilError(t, err)
	assert.Equal(t, 1, len(res.Projects))
	assert.Equal(t, uint64(2), res.Pagination.Total)
}
