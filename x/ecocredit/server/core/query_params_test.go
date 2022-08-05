package core

import (
	"testing"

	"github.com/golang/mock/gomock"
	"gotest.tools/v3/assert"

	"github.com/cosmos/cosmos-sdk/types"

	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

func TestQuery_Params(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	gmAny := gomock.Any()
	s.paramsKeeper.EXPECT().GetParamSet(gmAny, gmAny).SetArg(1, core.Params{
		CreditClassFee:       types.NewCoins(types.NewInt64Coin("foo", 30)),
		AllowedClassCreators: []string{s.addr.String()},
		AllowlistEnabled:     false,
	})

	res, err := s.k.Params(s.ctx, &core.QueryParamsRequest{})
	assert.NilError(t, err)
	assert.Equal(t, false, res.Params.AllowlistEnabled)
	assert.Equal(t, s.addr.String(), res.Params.AllowedClassCreators[0])
}
