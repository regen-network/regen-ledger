package core

import (
	"github.com/cosmos/cosmos-sdk/types"
	"github.com/golang/mock/gomock"
	ecocreditv1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	v1 "github.com/regen-network/regen-ledger/x/ecocredit/v1"
	"gotest.tools/v3/assert"
	"testing"
)

func TestQuery_Params(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	any := gomock.Any()

	assert.NilError(t, s.stateStore.CreditTypeStore().Insert(s.ctx, &ecocreditv1.CreditType{
		Abbreviation: "C",
		Name:         "carbon",
		Unit:         "a ton",
		Precision:    6,
	}))

	s.paramsKeeper.EXPECT().GetParamSet(any, any).SetArg(1, ecocredit.Params{
		CreditClassFee:       types.NewCoins(types.NewInt64Coin("foo", 30)),
		AllowedClassCreators: []string{s.addr.String()},
		AllowlistEnabled:     false,
		CreditTypes:          []*ecocredit.CreditType{{
			Abbreviation: "C",
			Name: "carbon",
			Unit: "a ton",
			Precision: 6,
		}},
	})

	res, err := s.k.Params(s.ctx, &v1.QueryParamsRequest{})
	assert.NilError(t, err)
	assert.Equal(t,false, res.Params.AllowlistEnabled)
	assert.Equal(t, s.addr.String(), res.Params.AllowedClassCreators[0])
}