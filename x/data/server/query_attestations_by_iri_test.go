package server

import (
	"bytes"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/cosmos/cosmos-sdk/types/query"

	api "github.com/regen-network/regen-ledger/api/v2/regen/data/v1"
	"github.com/regen-network/regen-ledger/x/data/v2"
)

func TestQuery_AttestationsByIRI(t *testing.T) {
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

	// insert data ids (one with attestations and one without)
	err = s.server.stateStore.DataIDTable().Insert(s.ctx, &api.DataID{Id: id1, Iri: iri1})
	require.NoError(t, err)
	err = s.server.stateStore.DataIDTable().Insert(s.ctx, &api.DataID{Id: id2, Iri: iri2})
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
		Id:        id1,
		Attestor:  s.addrs[1],
		Timestamp: timestamp,
	})
	require.NoError(t, err)

	// query attestations with valid iri
	res, err := s.server.AttestationsByIRI(s.ctx, &data.QueryAttestationsByIRIRequest{
		Iri:        iri1,
		Pagination: &query.PageRequest{Limit: 1, CountTotal: true},
	})
	require.NoError(t, err)

	// check pagination
	require.Len(t, res.Attestations, 1)
	require.Equal(t, uint64(2), res.Pagination.Total)

	// check attestation properties
	require.Equal(t, iri1, res.Attestations[0].Iri)
	require.Equal(t, timestamp.Seconds, res.Attestations[0].Timestamp.Seconds)
	require.Equal(t, timestamp.Nanos, res.Attestations[0].Timestamp.Nanos)

	// order of attestations may vary
	if s.addrs[0].String() == res.Attestations[0].Attestor {
		require.Equal(t, s.addrs[0].String(), res.Attestations[0].Attestor)
	} else {
		require.Equal(t, s.addrs[1].String(), res.Attestations[0].Attestor)
	}

	// query attestations with iri that has no attestations
	res, err = s.server.AttestationsByIRI(s.ctx, &data.QueryAttestationsByIRIRequest{
		Iri: iri2,
	})
	require.NoError(t, err)
	require.Empty(t, res.Attestations)

	// query attestations with empty iri
	_, err = s.server.AttestationsByIRI(s.ctx, &data.QueryAttestationsByIRIRequest{})
	require.EqualError(t, err, "IRI cannot be empty: invalid argument")

	// query attestations with invalid iri
	_, err = s.server.AttestationsByIRI(s.ctx, &data.QueryAttestationsByIRIRequest{
		Iri: "foo",
	})
	require.EqualError(t, err, "failed to parse IRI: failed to parse IRI foo: regen: prefix required: invalid IRI: invalid argument")

	// query attestations with iri that has not been anchored
	_, err = s.server.AttestationsByIRI(s.ctx, &data.QueryAttestationsByIRIRequest{
		Iri: "regen:13toVfw5KEeQwbmV733E3j9HwhVCQTxB7ojFPjGdmr7HX3kuSASGXxV.rdf",
	})
	require.EqualError(t, err, "data record with IRI: regen:13toVfw5KEeQwbmV733E3j9HwhVCQTxB7ojFPjGdmr7HX3kuSASGXxV.rdf: not found")
}
