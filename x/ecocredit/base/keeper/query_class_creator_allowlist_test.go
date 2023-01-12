package keeper

import (
	"testing"

	"gotest.tools/v3/assert"

	api "github.com/regen-network/regen-ledger/api/v2/regen/ecocredit/v1"
	types "github.com/regen-network/regen-ledger/x/ecocredit/v3/base/types/v1"
)

func TestQuery_ClassCreatorAllowlist(t *testing.T) {
	t.Parallel()
	s := setupBase(t)

	res, err := s.k.ClassCreatorAllowlist(s.ctx, &types.QueryClassCreatorAllowlistRequest{})
	assert.NilError(t, err)
	assert.Equal(t, res.Enabled, false)

	err = s.stateStore.ClassCreatorAllowlistTable().Save(s.ctx, &api.ClassCreatorAllowlist{
		Enabled: true,
	})
	assert.NilError(t, err)

	res, err = s.k.ClassCreatorAllowlist(s.ctx, &types.QueryClassCreatorAllowlistRequest{})
	assert.NilError(t, err)
	assert.Equal(t, res.Enabled, true)
}
