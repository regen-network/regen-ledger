package core

import (
	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	"github.com/cosmos/cosmos-sdk/types/query"
	ecocreditv1beta1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1beta1"
	"github.com/regen-network/regen-ledger/x/ecocredit/v1beta1"
	"gotest.tools/v3/assert"
	"testing"
)

func TestQuery_Projects(t *testing.T) {
	t.Parallel()
	s := setupBase(t)

	// create a class and 2 projects from it
	err := s.stateStore.ClassInfoStore().Insert(s.ctx, &ecocreditv1beta1.ClassInfo{
		Name:       "C01",
		Admin:      s.addr,
		Metadata:   nil,
		CreditType: "C",
	})
	assert.NilError(t, err)
	err= s.stateStore.ProjectInfoStore().Insert(s.ctx, &ecocreditv1beta1.ProjectInfo{
		Name:            "P01",
		ClassId:         1,
		ProjectLocation: "US-CA",
		Metadata:        nil,
	})
	assert.NilError(t, err)
	err= s.stateStore.ProjectInfoStore().Insert(s.ctx, &ecocreditv1beta1.ProjectInfo{
		Name:            "P02",
		ClassId:         1,
		ProjectLocation: "US-CA",
		Metadata:        nil,
	})
	assert.NilError(t, err)

	// base query
	res, err := s.k.Projects(s.ctx, &v1beta1.QueryProjectsRequest{ClassId: "C01"})
	assert.NilError(t, err)
	assert.Equal(t, 2, len(res.Projects))
	assert.Equal(t, "US-CA", res.Projects[0].ProjectLocation)

	// invalid query
	_, err = s.k.Projects(s.ctx, &v1beta1.QueryProjectsRequest{ClassId: "F01"})
	assert.ErrorContains(t, err, ormerrors.NotFound.Error())

	// paginated query
	res, err = s.k.Projects(s.ctx, &v1beta1.QueryProjectsRequest{
		ClassId:    "C01",
		Pagination: &query.PageRequest{Limit: 1, CountTotal: true},
	})
	assert.NilError(t, err)
	assert.Equal(t, 1, len(res.Projects))
	assert.Equal(t, uint64(2), res.Pagination.Total)
}