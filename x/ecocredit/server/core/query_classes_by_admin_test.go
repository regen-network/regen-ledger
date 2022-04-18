package core

import (
	"testing"

	"gotest.tools/v3/assert"

	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

func TestQueryClassesByAdmin(t *testing.T) {
	t.Parallel()
	s := setupBase(t)

	_, _, addr := testdata.KeyTestPubAddr()
	_, _, noClasses := testdata.KeyTestPubAddr()

	class1 := &api.ClassInfo{
		Name:       "C01",
		Admin:      s.addr,
		Metadata:   "data",
		CreditType: "C",
	}

	assert.NilError(t, s.stateStore.ClassInfoTable().Insert(s.ctx, class1))
	assert.NilError(t, s.stateStore.ClassInfoTable().Insert(s.ctx, &api.ClassInfo{
		Name:       "C02",
		Admin:      s.addr,
		CreditType: "C",
	}))
	assert.NilError(t, s.stateStore.ClassInfoTable().Insert(s.ctx, &api.ClassInfo{
		Name:       "C03",
		Admin:      addr,
		CreditType: "C",
	}))

	// valid query
	res, err := s.k.ClassesByAdmin(s.ctx, &core.QueryClassesByAdminRequest{
		Admin:      s.addr.String(),
		Pagination: &query.PageRequest{Limit: 1, CountTotal: true},
	})
	assert.NilError(t, err)
	assert.Equal(t, len(res.Classes), 1)
	assert.Equal(t, class1.Name, res.Classes[0].Id)
	assert.Equal(t, s.addr.String(), res.Classes[0].Admin)
	assert.Equal(t, class1.Metadata, res.Classes[0].Metadata)
	assert.Equal(t, class1.CreditType, res.Classes[0].CreditTypeAbbrev)
	assert.Equal(t, uint64(2), res.Pagination.Total)

	// should be empty
	res, err = s.k.ClassesByAdmin(s.ctx, &core.QueryClassesByAdminRequest{
		Admin:      noClasses.String(),
		Pagination: &query.PageRequest{Limit: 10, CountTotal: true},
	})
	assert.NilError(t, err)
	assert.Equal(t, 0, len(res.Classes))

	// invalid address
	_, err = s.k.ClassesByAdmin(s.ctx, &core.QueryClassesByAdminRequest{Admin: "invalid_address"})
	assert.ErrorContains(t, err, sdkerrors.ErrInvalidAddress.Error())
}
