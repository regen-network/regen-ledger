package keeper

import (
	"testing"

	"gotest.tools/v3/assert"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	types "github.com/regen-network/regen-ledger/x/ecocredit/base/types/v1"
)

func TestQuery_AllowlistEnabledTest(t *testing.T) {
	t.Parallel()
	s := setupBase(t)

	res, err := s.k.CreditClassAllowlistEnabled(s.ctx, &types.QueryCreditClassAllowlistEnabledRequest{})
	assert.NilError(t, err)
	assert.Equal(t, res.AllowlistEnabled, false)

	err = s.stateStore.AllowListEnabledTable().Save(s.ctx, &api.AllowListEnabled{
		Enabled: true,
	})
	assert.NilError(t, err)

	res, err = s.k.CreditClassAllowlistEnabled(s.ctx, &types.QueryCreditClassAllowlistEnabledRequest{})
	assert.NilError(t, err)
	assert.Equal(t, res.AllowlistEnabled, true)
}
