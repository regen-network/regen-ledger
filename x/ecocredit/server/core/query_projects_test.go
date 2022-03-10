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
	err := s.stateStore.ClassInfoStore().Insert(s.ctx, &api.ClassInfo{
		Name:       "C01",
		Admin:      s.addr,
		Metadata:   nil,
		CreditType: "C",
	})
	assert.NilError(t, err)
	err = s.stateStore.ProjectInfoStore().Insert(s.ctx, &api.ProjectInfo{
		Name:            "P01",
		ClassId:         1,
		ProjectLocation: "US-CA",
		Metadata:        nil,
	})
	assert.NilError(t, err)
	err = s.stateStore.ProjectInfoStore().Insert(s.ctx, &api.ProjectInfo{
		Name:            "P02",
		ClassId:         1,
		ProjectLocation: "US-CA",
		Metadata:        nil,
	})
	assert.NilError(t, err)

	// base query
	res, err := s.k.Projects(s.ctx, &core.QueryProjectsRequest{ClassId: "C01"})
	assert.NilError(t, err)
	assert.Equal(t, 2, len(res.Projects))
	assert.Equal(t, "US-CA", res.Projects[0].ProjectLocation)

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
