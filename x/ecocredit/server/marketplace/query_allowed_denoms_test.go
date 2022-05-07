package marketplace

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/types/query"
	"gotest.tools/v3/assert"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/marketplace/v1"
	"github.com/regen-network/regen-ledger/types/ormutil"
	"github.com/regen-network/regen-ledger/x/ecocredit/marketplace"
)

func TestQueryAllowedDenoms(t *testing.T) {
	t.Parallel()
	s := setupBase(t)

	allowedDenom := api.AllowedDenom{
		BankDenom:    "uregen",
		DisplayDenom: "regen",
		Exponent:     18,
	}
	err := s.marketStore.AllowedDenomTable().Insert(s.ctx, &allowedDenom)
	assert.NilError(t, err)

	allowedDenomOsmo := api.AllowedDenom{
		BankDenom:    "uosmo",
		DisplayDenom: "osmo",
		Exponent:     18,
	}
	err = s.marketStore.AllowedDenomTable().Insert(s.ctx, &allowedDenomOsmo)
	assert.NilError(t, err)

	var gogoAllowedDenom marketplace.AllowedDenom
	assert.NilError(t, ormutil.PulsarToGogoSlow(&allowedDenomOsmo, &gogoAllowedDenom))

	res, err := s.k.AllowedDenoms(s.ctx, &marketplace.QueryAllowedDenomsRequest{Pagination: &query.PageRequest{CountTotal: true, Limit: 1}})
	assert.NilError(t, err)
	assert.Check(t, res.Pagination != nil)
	assert.Equal(t, res.Pagination.Total, uint64(2))
	assert.Equal(t, 1, len(res.AllowedDenoms))
	assert.DeepEqual(t, res.AllowedDenoms[0], &gogoAllowedDenom)
}
