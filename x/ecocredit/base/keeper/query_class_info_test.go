package keeper

import (
	"testing"

	"gotest.tools/v3/assert"

	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	types "github.com/regen-network/regen-ledger/x/ecocredit/base/types/v1"
)

func TestQuery_ClassInfo(t *testing.T) {
	t.Parallel()
	s := setupBase(t)

	class := &api.Class{
		Id:               "C01",
		Admin:            s.addr,
		Metadata:         "data",
		CreditTypeAbbrev: "C",
	}

	err := s.stateStore.ClassTable().Insert(s.ctx, class)
	assert.NilError(t, err)

	// query class by the "C01" class id
	res, err := s.k.Class(s.ctx, &types.QueryClassRequest{ClassId: class.Id})
	assert.NilError(t, err)
	assert.Equal(t, class.Id, res.Class.Id)
	assert.Equal(t, s.addr.String(), res.Class.Admin)
	assert.Equal(t, class.Metadata, res.Class.Metadata)
	assert.Equal(t, class.CreditTypeAbbrev, res.Class.CreditTypeAbbrev)

	// query class by an unknown class id
	_, err = s.k.Class(s.ctx, &types.QueryClassRequest{ClassId: "C02"})
	assert.ErrorContains(t, err, ormerrors.NotFound.Error())
}
