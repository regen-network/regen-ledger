package core

import (
	"testing"

	"gotest.tools/v3/assert"

	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

func TestQuery_ClassInfo(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	err := s.stateStore.ClassInfoTable().Insert(s.ctx, &api.ClassInfo{
		Id:               "C01",
		Admin:            s.addr,
		Metadata:         "",
		CreditTypeAbbrev: "C",
	})
	assert.NilError(t, err)

	// query an invalid class
	_, err = s.k.ClassInfo(s.ctx, &core.QueryClassInfoRequest{ClassId: "C02"})
	assert.ErrorContains(t, err, ormerrors.NotFound.Error())

	// query a valid class
	res, err := s.k.ClassInfo(s.ctx, &core.QueryClassInfoRequest{ClassId: "C01"})
	assert.NilError(t, err)
	assert.Equal(t, "C01", res.Info.Id)
	assert.DeepEqual(t, s.addr.Bytes(), res.Info.Admin)
}
