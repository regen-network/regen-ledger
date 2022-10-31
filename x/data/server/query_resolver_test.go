package server

import (
	"testing"

	"github.com/stretchr/testify/require"

	api "github.com/regen-network/regen-ledger/api/regen/data/v1"
	"github.com/regen-network/regen-ledger/x/data"
)

func TestQuery_Resolver(t *testing.T) {
	t.Parallel()
	s := setupBase(t)

	// insert resolvers
	id, err := s.server.stateStore.ResolverTable().InsertReturningID(s.ctx, &api.Resolver{
		Url:     testURL,
		Manager: s.addrs[0],
	})
	require.NoError(t, err)

	// query resolver with valid id
	res, err := s.server.Resolver(s.ctx, &data.QueryResolverRequest{Id: id})
	require.NoError(t, err)
	require.Equal(t, id, res.Resolver.Id)
	require.Equal(t, s.addrs[0].String(), res.Resolver.Manager)
	require.Equal(t, testURL, res.Resolver.Url)

	// query resolvers with id that does not exist
	_, err = s.server.Resolver(s.ctx, &data.QueryResolverRequest{
		Id: 404,
	})
	require.EqualError(t, err, "resolver with ID: 404: not found")

	// query resolvers with empty id
	_, err = s.server.Resolver(s.ctx, &data.QueryResolverRequest{})
	require.EqualError(t, err, "ID cannot be empty: invalid argument")
}
