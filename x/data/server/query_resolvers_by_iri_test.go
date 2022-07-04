package server

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/cosmos/cosmos-sdk/types/query"

	api "github.com/regen-network/regen-ledger/api/regen/data/v1"
	"github.com/regen-network/regen-ledger/x/data"
)

func TestQuery_ResolversByIRI(t *testing.T) {
	t.Parallel()
	s := setupBase(t)

	id1 := []byte{0}
	ch1 := &data.ContentHash{Graph: &data.ContentHash_Graph{
		Hash:                      bytes.Repeat([]byte{0}, 32),
		DigestAlgorithm:           data.DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256,
		CanonicalizationAlgorithm: data.GraphCanonicalizationAlgorithm_GRAPH_CANONICALIZATION_ALGORITHM_URDNA2015,
	}}
	iri1, err := ch1.ToIRI()
	require.NoError(t, err)

	id2 := []byte{1}
	ch2 := &data.ContentHash{Graph: &data.ContentHash_Graph{
		Hash:                      bytes.Repeat([]byte{1}, 32),
		DigestAlgorithm:           data.DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256,
		CanonicalizationAlgorithm: data.GraphCanonicalizationAlgorithm_GRAPH_CANONICALIZATION_ALGORITHM_URDNA2015,
	}}
	iri2, err := ch2.ToIRI()
	require.NoError(t, err)

	url := "https://foo.bar"

	// insert data ids
	err = s.server.stateStore.DataIDTable().Insert(s.ctx, &api.DataID{Id: id1, Iri: iri1})
	require.NoError(t, err)
	err = s.server.stateStore.DataIDTable().Insert(s.ctx, &api.DataID{Id: id2, Iri: iri2})
	require.NoError(t, err)

	// insert resolvers
	rid1, err := s.server.stateStore.ResolverTable().InsertReturningID(s.ctx, &api.Resolver{
		Url:     url,
		Manager: s.addrs[0],
	})
	require.NoError(t, err)
	rid2, err := s.server.stateStore.ResolverTable().InsertReturningID(s.ctx, &api.Resolver{
		Url:     url,
		Manager: s.addrs[1],
	})
	require.NoError(t, err)

	// insert registration records
	err = s.server.stateStore.DataResolverTable().Insert(s.ctx, &api.DataResolver{
		Id:         id1,
		ResolverId: rid1,
	})
	require.NoError(t, err)
	err = s.server.stateStore.DataResolverTable().Insert(s.ctx, &api.DataResolver{
		Id:         id1,
		ResolverId: rid2,
	})
	require.NoError(t, err)

	// query resolvers with valid iri
	res, err := s.server.ResolversByIRI(s.ctx, &data.QueryResolversByIRIRequest{
		Iri:        iri1,
		Pagination: &query.PageRequest{Limit: 1, CountTotal: true},
	})
	require.NoError(t, err)

	// check pagination
	require.Len(t, res.Resolvers, 1)
	require.Equal(t, uint64(2), res.Pagination.Total)

	// check resolver properties
	require.Equal(t, rid1, res.Resolvers[0].Id)
	require.Equal(t, s.addrs[0].String(), res.Resolvers[0].Manager)
	require.Equal(t, url, res.Resolvers[0].Url)

	// query resolvers with iri that has not been registered
	res, err = s.server.ResolversByIRI(s.ctx, &data.QueryResolversByIRIRequest{
		Iri: iri2,
	})
	require.NoError(t, err)
	require.Empty(t, res.Resolvers)

	// query resolvers with empty iri
	_, err = s.server.ResolversByIRI(s.ctx, &data.QueryResolversByIRIRequest{})
	require.EqualError(t, err, "IRI cannot be empty: invalid request")

	// query resolvers with invalid iri
	_, err = s.server.ResolversByIRI(s.ctx, &data.QueryResolversByIRIRequest{
		Iri: "foo",
	})
	require.EqualError(t, err, "failed to parse IRI foo: regen: prefix required: invalid IRI")

	// query resolvers with iri that has not been anchored
	_, err = s.server.ResolversByIRI(s.ctx, &data.QueryResolversByIRIRequest{
		Iri: "regen:13toVfw5KEeQwbmV733E3j9HwhVCQTxB7ojFPjGdmr7HX3kuSASGXxV.rdf",
	})
	require.EqualError(t, err, "data record with IRI: not found")
}
