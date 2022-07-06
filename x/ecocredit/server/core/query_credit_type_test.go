package core

import (
	"testing"

	"gotest.tools/v3/assert"

	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

func TestQuery_CreditType(t *testing.T) {
	t.Parallel()
	s := setupBase(t) // setupBase gives us 1 default credit type, we add another here for testing
	assert.NilError(t, s.stateStore.CreditTypeTable().Insert(s.ctx, &api.CreditType{
		Abbreviation: "BIO",
		Name:         "biodiversity",
		Unit:         "acres",
		Precision:    6,
	}))

	res, err := s.k.CreditType(s.ctx, &core.QueryCreditTypeRequest{
		Abbreviation: "BIO",
	})
	assert.NilError(t, err)
	assert.Equal(t, res.CreditType.Precision, uint32(6))
	assert.Equal(t, res.CreditType.Abbreviation, "BIO")
	assert.Equal(t, res.CreditType.Name, "biodiversity")

	// query credit type by unknown abbreviation
	_, err = s.k.CreditType(s.ctx, &core.QueryCreditTypeRequest{
		Abbreviation: "D",
	})
	assert.ErrorIs(t, err, ormerrors.NotFound)
}
