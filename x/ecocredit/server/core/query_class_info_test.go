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

	classId := "C01"
	metadata := "data"
	creditType := "C"

	err := s.stateStore.ClassInfoTable().Insert(s.ctx, &api.ClassInfo{
		Id:               classId,
		Admin:            s.addr,
		Metadata:         metadata,
		CreditTypeAbbrev: creditType,
	})
	assert.NilError(t, err)

	// query a valid class
	res, err := s.k.ClassInfo(s.ctx, &core.QueryClassInfoRequest{ClassId: classId})
	assert.NilError(t, err)
	assert.Equal(t, classId, res.Class.Id)
	assert.Equal(t, s.addr.String(), res.Class.Admin)
	assert.Equal(t, metadata, res.Class.Metadata)
	assert.Equal(t, creditType, res.Class.CreditTypeAbbrev)

	// query an invalid class
	_, err = s.k.ClassInfo(s.ctx, &core.QueryClassInfoRequest{ClassId: "C02"})
	assert.ErrorContains(t, err, ormerrors.NotFound.Error())

	// query a valid class
	res, err = s.k.ClassInfo(s.ctx, &core.QueryClassInfoRequest{ClassId: "C01"})
	assert.NilError(t, err)
	assert.Equal(t, "C01", res.Class.Id)
	assert.Equal(t, s.addr.String(), res.Class.Admin)
}
