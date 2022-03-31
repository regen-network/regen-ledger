package core

import (
	"testing"

	"github.com/golang/mock/gomock"
	"gotest.tools/v3/assert"

	"github.com/cosmos/cosmos-sdk/types"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

func TestQuery_Params(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	any := gomock.Any()

	assert.NilError(t, s.stateStore.CreditTypeTable().Insert(s.ctx, &api.CreditType{
		Abbreviation: "C",
		Name:         "carbon",
		Unit:         "a ton",
		Precision:    6,
	}))

	s.paramsKeeper.EXPECT().GetParamSet(any, any).SetArg(1, core.Params{
		CreditClassFee:       types.NewCoins(types.NewInt64Coin("foo", 30)),
		AllowedClassCreators: []string{s.addr.String()},
		AllowlistEnabled:     false,
		CreditTypes: []*core.CreditType{{
			Abbreviation: "C",
			Name:         "carbon",
			Unit:         "a ton",
			Precision:    6,
		}},
	})

	res, err := s.k.Params(s.ctx, &core.QueryParamsRequest{})
	assert.NilError(t, err)
	assert.Equal(t, false, res.Params.AllowlistEnabled)
	assert.Equal(t, s.addr.String(), res.Params.AllowedClassCreators[0])
}
