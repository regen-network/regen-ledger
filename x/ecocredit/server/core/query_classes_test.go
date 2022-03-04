package core

import (
	"github.com/cosmos/cosmos-sdk/types/query"
	ecocreditv1beta1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1beta1"
	"github.com/regen-network/regen-ledger/x/ecocredit/v1beta1"
	"gotest.tools/v3/assert"
	"testing"
)

func TestQuery_Classes(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	err := s.stateStore.ClassInfoStore().Insert(s.ctx, &ecocreditv1beta1.ClassInfo{
		Name:       "C01",
		Admin:      s.addr,
		Metadata:   nil,
		CreditType: "C",
	})
	assert.NilError(t, err)
	err = s.stateStore.ClassInfoStore().Insert(s.ctx, &ecocreditv1beta1.ClassInfo{
		Name:       "C02",
		Admin:      s.addr,
		Metadata:   nil,
		CreditType: "C",
	})
	assert.NilError(t, err)

	// requesting all
	res, err := s.k.Classes(s.ctx, &v1beta1.QueryClassesRequest{})
	assert.NilError(t, err)
	assert.Equal(t, 2, len(res.Classes))
	assert.Equal(t, "C01", res.Classes[0].Name)
	assert.Equal(t, "C02", res.Classes[1].Name)

	// request with pagination
	res, err = s.k.Classes(s.ctx, &v1beta1.QueryClassesRequest{Pagination: &query.PageRequest{
		Limit:      1,
		CountTotal: true,
		Reverse:    true,
	}})
	assert.NilError(t, err)
	assert.Equal(t, 1, len(res.Classes))
	assert.Equal(t, uint64(2), res.Pagination.Total)
	assert.Equal(t, "C02", res.Classes[0].Name)
}
