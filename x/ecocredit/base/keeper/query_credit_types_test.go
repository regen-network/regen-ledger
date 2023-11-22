package keeper

import (
	"testing"

	"gotest.tools/v3/assert"

	api "github.com/regen-network/regen-ledger/api/v2/regen/ecocredit/v1"
	types "github.com/regen-network/regen-ledger/x/ecocredit/v3/base/types/v1"
)

func TestQuery_CreditTypes(t *testing.T) {
	t.Parallel()
	s := setupBase(t) // setupBase gives us 1 default credit type, we add another here for testing
	assert.NilError(t, s.stateStore.CreditTypeTable().Insert(s.ctx, &api.CreditType{
		Abbreviation: "C",
		Name:         "carbon",
		Unit:         "metric ton CO2 equivalent",
		Precision:    6,
	}))
	assert.NilError(t, s.stateStore.CreditTypeTable().Insert(s.ctx, &api.CreditType{
		Abbreviation: "BIO",
		Name:         "biodiversity",
		Unit:         "acres",
		Precision:    6,
	}))
	// base query should return all types
	res, err := s.k.CreditTypes(s.ctx, &types.QueryCreditTypesRequest{})
	assert.NilError(t, err)
	assert.Equal(t, 2, len(res.CreditTypes))
	assert.Equal(t, uint32(6), res.CreditTypes[1].Precision)
	assert.Equal(t, "carbon", res.CreditTypes[1].Name)
	assert.Equal(t, "C", res.CreditTypes[1].Abbreviation)
	assert.Equal(t, "BIO", res.CreditTypes[0].Abbreviation)
}
