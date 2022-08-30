package core

import (
	"testing"

	ecocreditv1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
	"gotest.tools/v3/assert"
)

func TestQuery_AllowlistEnabledTest(t *testing.T) {
	t.Parallel()
	s := setupBase(t)

	res, err := s.k.CreditClassAllowlistEnabled(s.ctx, &core.QueryCreditClassAllowlistEnabledRequest{})
	assert.NilError(t, err)
	assert.Equal(t, res.AllowlistEnabled, false)

	err = s.stateStore.AllowListEnabledTable().Save(s.ctx, &ecocreditv1.AllowListEnabled{
		Enabled: true,
	})
	assert.NilError(t, err)

	res, err = s.k.CreditClassAllowlistEnabled(s.ctx, &core.QueryCreditClassAllowlistEnabledRequest{})
	assert.NilError(t, err)
	assert.Equal(t, res.AllowlistEnabled, true)
}
