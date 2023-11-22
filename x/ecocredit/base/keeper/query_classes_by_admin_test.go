package keeper

import (
	"testing"

	"gotest.tools/v3/assert"

	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	"github.com/cosmos/cosmos-sdk/types/query"

	api "github.com/regen-network/regen-ledger/api/v2/regen/ecocredit/v1"
	regenerrors "github.com/regen-network/regen-ledger/types/v2/errors"
	types "github.com/regen-network/regen-ledger/x/ecocredit/v3/base/types/v1"
)

func TestQueryClassesByAdmin(t *testing.T) {
	t.Parallel()
	s := setupBase(t)

	class := &api.Class{
		Id:               "C01",
		Admin:            s.addr,
		Metadata:         "data",
		CreditTypeAbbrev: "C",
	}

	// insert two classes with s.addr as the admin
	assert.NilError(t, s.stateStore.ClassTable().Insert(s.ctx, class))
	assert.NilError(t, s.stateStore.ClassTable().Insert(s.ctx, &api.Class{
		Id:    "C02",
		Admin: s.addr,
	}))

	// query classes by the admin s.addr
	res, err := s.k.ClassesByAdmin(s.ctx, &types.QueryClassesByAdminRequest{
		Admin:      s.addr.String(),
		Pagination: &query.PageRequest{Limit: 1, CountTotal: true},
	})
	assert.NilError(t, err)
	assert.Equal(t, 1, len(res.Classes))
	assert.Equal(t, class.Id, res.Classes[0].Id)
	assert.Equal(t, s.addr.String(), res.Classes[0].Admin)
	assert.Equal(t, class.Metadata, res.Classes[0].Metadata)
	assert.Equal(t, class.CreditTypeAbbrev, res.Classes[0].CreditTypeAbbrev)
	assert.Equal(t, uint64(2), res.Pagination.Total)

	_, _, notAdmin := testdata.KeyTestPubAddr()

	// query classes by an unknown admin address
	res, err = s.k.ClassesByAdmin(s.ctx, &types.QueryClassesByAdminRequest{
		Admin: notAdmin.String(),
	})
	assert.NilError(t, err)
	assert.Equal(t, 0, len(res.Classes))

	// query classes by an invalid admin address
	_, err = s.k.ClassesByAdmin(s.ctx, &types.QueryClassesByAdminRequest{Admin: "foobar"})
	assert.ErrorContains(t, err, regenerrors.ErrInvalidArgument.Error())
}
