package server

import (
	"bytes"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/cosmos/cosmos-sdk/types/query"

	api "github.com/regen-network/regen-ledger/api/regen/data/v1"
	"github.com/regen-network/regen-ledger/x/data"
)

func TestQuery_AttestationsByAttestor(t *testing.T) {
	t.Parallel()
	s := setupBase(t)

	id1 := []byte{0}
	id2 := []byte{1}
	ch := &data.ContentHash{Graph: &data.ContentHash_Graph{
		Hash:                      bytes.Repeat([]byte{0}, 32),
		DigestAlgorithm:           data.DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256,
		CanonicalizationAlgorithm: data.GraphCanonicalizationAlgorithm_GRAPH_CANONICALIZATION_ALGORITHM_URDNA2015,
	}}
	iri, err := ch.ToIRI()
	require.NoError(t, err)

	// insert data ids
	err = s.server.stateStore.DataIDTable().Insert(s.ctx, &api.DataID{Id: id1, Iri: iri})
	require.NoError(t, err)
	err = s.server.stateStore.DataIDTable().Insert(s.ctx, &api.DataID{Id: id2})
	require.NoError(t, err)

	timestamp := timestamppb.New(time.Now().UTC())

	// insert attestations
	err = s.server.stateStore.DataAttestorTable().Insert(s.ctx, &api.DataAttestor{
		Id:        id1,
		Attestor:  s.addrs[0],
		Timestamp: timestamp,
	})
	require.NoError(t, err)
	err = s.server.stateStore.DataAttestorTable().Insert(s.ctx, &api.DataAttestor{
		Id:        id2,
		Attestor:  s.addrs[0],
		Timestamp: timestamp,
	})
	require.NoError(t, err)

	// query attestations with valid attestor
	res, err := s.server.AttestationsByAttestor(s.ctx, &data.QueryAttestationsByAttestorRequest{
		Attestor:   s.addrs[0].String(),
		Pagination: &query.PageRequest{Limit: 1, CountTotal: true},
	})
	require.NoError(t, err)

	// check pagination
	require.Len(t, res.Attestations, 1)
	require.Equal(t, uint64(2), res.Pagination.Total)

	// check attestation properties
	require.Equal(t, iri, res.Attestations[0].Iri)
	require.Equal(t, s.addrs[0].String(), res.Attestations[0].Attestor)
	require.Equal(t, timestamp.Seconds, res.Attestations[0].Timestamp.Seconds)
	require.Equal(t, timestamp.Nanos, res.Attestations[0].Timestamp.Nanos)

	// query attestations with address that has no attestations
	res, err = s.server.AttestationsByAttestor(s.ctx, &data.QueryAttestationsByAttestorRequest{
		Attestor: s.addrs[1].String(),
	})
	require.NoError(t, err)
	require.Empty(t, res.Attestations)

	// query attestations with empty attestor
	_, err = s.server.AttestationsByAttestor(s.ctx, &data.QueryAttestationsByAttestorRequest{})
	require.EqualError(t, err, "attestor cannot be empty: invalid argument")

	// query attestations with invalid attestor address
	_, err = s.server.AttestationsByAttestor(s.ctx, &data.QueryAttestationsByAttestorRequest{
		Attestor: "foo",
	})
	require.EqualError(t, err, "attestor: decoding bech32 failed: invalid bech32 string length 3: invalid argument")
}
