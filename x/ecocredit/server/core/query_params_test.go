package core

import (
	"testing"

	"gotest.tools/v3/assert"

	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

func TestQuery_Params(t *testing.T) {
	t.Parallel()
	s := setupBase(t)

	// TODO: add params tests
	_, err := s.k.Params(s.ctx, &core.QueryParamsRequest{})
	assert.NilError(t, err)
}
