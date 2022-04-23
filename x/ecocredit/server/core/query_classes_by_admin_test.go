package core

import (
	"testing"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
	"gotest.tools/v3/assert"

	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
)

func TestQueryClassesByAdmin(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	_, _, addr := testdata.KeyTestPubAddr()
	_, _, noClasses := testdata.KeyTestPubAddr()
	assert.NilError(t, s.stateStore.ClassTable().Insert(s.ctx, &api.Class{Id: "C01", Admin: s.addr, CreditTypeAbbrev: "C"}))
	assert.NilError(t, s.stateStore.ClassTable().Insert(s.ctx, &api.Class{Id: "C02", Admin: s.addr, CreditTypeAbbrev: "C"}))
	assert.NilError(t, s.stateStore.ClassTable().Insert(s.ctx, &api.Class{Id: "C03", Admin: addr, CreditTypeAbbrev: "C"}))

	// valid query
	res, err := s.k.ClassesByAdmin(s.ctx, &core.QueryClassesByAdminRequest{Admin: s.addr.String(), Pagination: &query.PageRequest{Limit: 1, CountTotal: true}})
	assert.NilError(t, err)
	assert.Equal(t, len(res.Classes), 1)
	assert.Equal(t, "C01", res.Classes[0].Id)
	assert.Equal(t, uint64(2), res.Pagination.Total)

	// should be empty
	res, err = s.k.ClassesByAdmin(s.ctx, &core.QueryClassesByAdminRequest{Admin: noClasses.String(), Pagination: &query.PageRequest{Limit: 10, CountTotal: true}})
	assert.NilError(t, err)
	assert.Equal(t, 0, len(res.Classes))

	// invalid address
	_, err = s.k.ClassesByAdmin(s.ctx, &core.QueryClassesByAdminRequest{Admin: "invalid_address"})
	assert.ErrorContains(t, err, sdkerrors.ErrInvalidAddress.Error())
}
