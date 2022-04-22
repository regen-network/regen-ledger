package core

import (
	"testing"

	"gotest.tools/v3/assert"

	"github.com/cosmos/cosmos-sdk/types/query"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

func TestQuery_Classes(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	err := s.stateStore.ClassTable().Insert(s.ctx, &api.Class{
		Id:               "C01",
		Admin:            s.addr,
		Metadata:         "",
		CreditTypeAbbrev: "C",
	})
	assert.NilError(t, err)
	err = s.stateStore.ClassTable().Insert(s.ctx, &api.Class{
		Id:               "C02",
		Admin:            s.addr,
		Metadata:         "",
		CreditTypeAbbrev: "C",
	})
	assert.NilError(t, err)

	// requesting all
	res, err := s.k.Classes(s.ctx, &core.QueryClassesRequest{})
	assert.NilError(t, err)
	assert.Equal(t, 2, len(res.Classes))
	assert.Equal(t, "C01", res.Classes[0].Id)
	assert.Equal(t, "C02", res.Classes[1].Id)

	// request with pagination
	res, err = s.k.Classes(s.ctx, &core.QueryClassesRequest{Pagination: &query.PageRequest{
		Limit:      1,
		CountTotal: true,
		Reverse:    true,
	}})
	assert.NilError(t, err)
	assert.Equal(t, 1, len(res.Classes))
	assert.Equal(t, uint64(2), res.Pagination.Total)
	assert.Equal(t, "C02", res.Classes[0].Id)
}
