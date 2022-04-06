package core

import (
	"testing"

	"github.com/golang/mock/gomock"
	"gotest.tools/v3/assert"

	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

func TestQuery_CreditTypes(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	gmAny := gomock.Any()
	s.paramsKeeper.EXPECT().GetParamSet(gmAny, gmAny).Do(func(any interface{}, p *core.Params) {
		p.CreditTypes = []*core.CreditType{{Name: "foobar", Abbreviation: "C", Unit: "tonne", Precision: 6}}
	}).Times(1)

	// base query should return all types
	res, err := s.k.CreditTypes(s.ctx, &core.QueryCreditTypesRequest{})
	assert.NilError(t, err)
	assert.Equal(t, 1, len(res.CreditTypes))
	assert.Equal(t, uint32(6), res.CreditTypes[0].Precision)
	assert.Equal(t, "foobar", res.CreditTypes[0].Name)
}
