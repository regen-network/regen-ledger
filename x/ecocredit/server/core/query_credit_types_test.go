package core

import (
	"testing"

	"gotest.tools/v3/assert"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

func TestQuery_CreditTypes(t *testing.T) {
	t.Parallel()
	s := setupBase(t)

	// insert a few credit types
	assert.NilError(t, s.stateStore.CreditTypeStore().Insert(s.ctx, &api.CreditType{
		Abbreviation: "C",
		Name:         "carbon",
		Unit:         "a ton",
		Precision:    6,
	}))
	assert.NilError(t, s.stateStore.CreditTypeStore().Insert(s.ctx, &api.CreditType{
		Abbreviation: "F",
		Name:         "foobar",
		Unit:         "foo per inch",
		Precision:    18,
	}))

	// base query should return all types
	res, err := s.k.CreditTypes(s.ctx, &core.QueryCreditTypesRequest{})
	assert.NilError(t, err)
	assert.Equal(t, 2, len(res.CreditTypes))
	assert.Equal(t, uint32(6), res.CreditTypes[0].Precision)
	assert.Equal(t, "foobar", res.CreditTypes[1].Name)
}
