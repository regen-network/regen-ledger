package core

import (
	ecocreditv1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	v1 "github.com/regen-network/regen-ledger/x/ecocredit/v1"
	"gotest.tools/v3/assert"
	"testing"
)

func TestQuery_CreditTypes(t *testing.T) {
	t.Parallel()
	s := setupBase(t)

	// insert a few credit types
	assert.NilError(t, s.stateStore.CreditTypeStore().Insert(s.ctx, &ecocreditv1.CreditType{
		Abbreviation: "C",
		Name:         "carbon",
		Unit:         "a ton",
		Precision:    6,
	}))
	assert.NilError(t, s.stateStore.CreditTypeStore().Insert(s.ctx, &ecocreditv1.CreditType{
		Abbreviation: "F",
		Name:         "foobar",
		Unit:         "foo per inch",
		Precision:    18,
	}))

	// base query should return all types
	res, err := s.k.CreditTypes(s.ctx, &v1.QueryCreditTypesRequest{})
	assert.NilError(t, err)
	assert.Equal(t, 2, len(res.CreditTypes))
	assert.Equal(t, uint32(6), res.CreditTypes[0].Precision)
	assert.Equal(t, "foobar", res.CreditTypes[1].Name)
}