package server

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/cosmos/cosmos-sdk/types/query"

	api "github.com/regen-network/regen-ledger/api/v2/regen/data/v1"
	"github.com/regen-network/regen-ledger/x/data/v2"
)

func TestQuery_ResolversByURL(t *testing.T) {
	t.Parallel()
	s := setupBase(t)

	// insert resolvers
	rid1, err := s.server.stateStore.ResolverTable().InsertReturningID(s.ctx, &api.Resolver{
		Url:     testURL,
		Manager: s.addrs[0],
	})
	require.NoError(t, err)
	err = s.server.stateStore.ResolverTable().Insert(s.ctx, &api.Resolver{
		Url:     testURL,
		Manager: s.addrs[1],
	})
	require.NoError(t, err)

	// query resolvers with valid url
	res, err := s.server.ResolversByURL(s.ctx, &data.QueryResolversByURLRequest{
		Url:        testURL,
		Pagination: &query.PageRequest{Limit: 1, CountTotal: true},
	})
	require.NoError(t, err)

	// check pagination
	require.Len(t, res.Resolvers, 1)
	require.Equal(t, uint64(2), res.Pagination.Total)

	// check resolver properties
	require.Equal(t, rid1, res.Resolvers[0].Id)
	require.Equal(t, s.addrs[0].String(), res.Resolvers[0].Manager)
	require.Equal(t, testURL, res.Resolvers[0].Url)

	// query resolvers with url that has no resolvers
	res, err = s.server.ResolversByURL(s.ctx, &data.QueryResolversByURLRequest{
		Url: "https://bar.baz",
	})
	require.NoError(t, err)
	require.Empty(t, res.Resolvers)

	// query resolvers with empty url
	_, err = s.server.ResolversByURL(s.ctx, &data.QueryResolversByURLRequest{})
	require.EqualError(t, err, "URL cannot be empty: invalid argument")
}
