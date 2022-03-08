package core

import (
	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	ecocreditv1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	v1 "github.com/regen-network/regen-ledger/x/ecocredit/v1"
	"gotest.tools/v3/assert"
	"testing"
)

func TestQuery_ClassInfo(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	err := s.stateStore.ClassInfoStore().Insert(s.ctx, &ecocreditv1.ClassInfo{
		Name:       "C01",
		Admin:      s.addr,
		Metadata:   nil,
		CreditType: "C",
	})
	assert.NilError(t, err)

	// query an invalid class
	_, err = s.k.ClassInfo(s.ctx, &v1.QueryClassInfoRequest{ClassId: "C02"})
	assert.ErrorContains(t, err, ormerrors.NotFound.Error())

	// query a valid class
	res, err := s.k.ClassInfo(s.ctx, &v1.QueryClassInfoRequest{ClassId: "C01"})
	assert.NilError(t, err)
	assert.Equal(t, "C01", res.Info.Name)
	assert.DeepEqual(t, s.addr.Bytes(), res.Info.Admin)
}