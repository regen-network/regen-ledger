package keeper

import (
	"testing"

	"gotest.tools/v3/assert"

	"github.com/stretchr/testify/require"

	"github.com/cosmos/cosmos-sdk/types/query"

	api "github.com/regen-network/regen-ledger/api/v2/regen/ecocredit/marketplace/v1"
	"github.com/regen-network/regen-ledger/types/v2/ormutil"
	types "github.com/regen-network/regen-ledger/x/ecocredit/v3/marketplace/types/v1"
)

func TestQueryAllowedDenoms(t *testing.T) {
	t.Parallel()
	s := setupBase(t, 0)

	// nil request
	_, err := s.k.AllowedDenoms(s.ctx, nil)
	require.Error(t, err)
	require.ErrorContains(t, err, "invalid argument")

	// no allowed denoms present
	_, err = s.k.AllowedDenoms(s.ctx, &types.QueryAllowedDenomsRequest{})
	assert.NilError(t, err)

	allowedDenom := api.AllowedDenom{
		BankDenom:    "uregen",
		DisplayDenom: "regen",
		Exponent:     18,
	}
	err = s.marketStore.AllowedDenomTable().Insert(s.ctx, &allowedDenom)
	assert.NilError(t, err)

	allowedDenomOsmo := api.AllowedDenom{
		BankDenom:    "uosmo",
		DisplayDenom: "osmo",
		Exponent:     18,
	}
	err = s.marketStore.AllowedDenomTable().Insert(s.ctx, &allowedDenomOsmo)
	assert.NilError(t, err)

	var gogoAllowedDenom types.AllowedDenom
	assert.NilError(t, ormutil.PulsarToGogoSlow(&allowedDenomOsmo, &gogoAllowedDenom))

	res, err := s.k.AllowedDenoms(s.ctx, &types.QueryAllowedDenomsRequest{Pagination: &query.PageRequest{CountTotal: true, Limit: 1}})
	assert.NilError(t, err)
	assert.Check(t, res.Pagination != nil)
	assert.Equal(t, res.Pagination.Total, uint64(2))
	assert.Equal(t, 1, len(res.AllowedDenoms))
	assert.DeepEqual(t, res.AllowedDenoms[0], &gogoAllowedDenom)
}
