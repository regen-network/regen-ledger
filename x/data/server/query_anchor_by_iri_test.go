package server

import (
	"bytes"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/timestamppb"

	api "github.com/regen-network/regen-ledger/api/v2/regen/data/v1"
	"github.com/regen-network/regen-ledger/x/data/v2"
)

func TestQuery_AnchorByIRI(t *testing.T) {
	t.Parallel()
	s := setupBase(t)

	id := []byte{0}
	ch := &data.ContentHash{Graph: &data.ContentHash_Graph{
		Hash:                      bytes.Repeat([]byte{0}, 32),
		DigestAlgorithm:           data.DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256,
		CanonicalizationAlgorithm: data.GraphCanonicalizationAlgorithm_GRAPH_CANONICALIZATION_ALGORITHM_URDNA2015,
		MerkleTree:                data.GraphMerkleTree_GRAPH_MERKLE_TREE_NONE_UNSPECIFIED,
	}}
	iri, err := ch.ToIRI()
	require.NoError(t, err)

	// insert data id
	err = s.server.stateStore.DataIDTable().Insert(s.ctx, &api.DataID{
		Id:  id,
		Iri: iri,
	})
	require.NoError(t, err)

	timestamp := timestamppb.New(time.Now().UTC())

	// insert data anchor
	err = s.server.stateStore.DataAnchorTable().Insert(s.ctx, &api.DataAnchor{
		Id:        id,
		Timestamp: timestamp,
	})
	require.NoError(t, err)

	// query data anchor with valid iri
	res, err := s.server.AnchorByIRI(s.ctx, &data.QueryAnchorByIRIRequest{
		Iri: iri,
	})
	require.NoError(t, err)
	require.Equal(t, iri, res.Anchor.Iri)
	require.Equal(t, ch, res.Anchor.ContentHash)
	require.Equal(t, timestamp.Seconds, res.Anchor.Timestamp.Seconds)
	require.Equal(t, timestamp.Nanos, res.Anchor.Timestamp.Nanos)

	// query data anchor with empty iri
	_, err = s.server.AnchorByIRI(s.ctx, &data.QueryAnchorByIRIRequest{})
	require.EqualError(t, err, "IRI cannot be empty: invalid argument")

	// query data anchor with invalid iri
	_, err = s.server.AnchorByIRI(s.ctx, &data.QueryAnchorByIRIRequest{
		Iri: "foo",
	})
	require.EqualError(t, err, "failed to parse IRI: failed to parse IRI foo: regen: prefix required: invalid IRI: invalid argument")

	// query data anchor with iri that has not been anchored
	_, err = s.server.AnchorByIRI(s.ctx, &data.QueryAnchorByIRIRequest{
		Iri: "regen:13toVfvdftodu8c1Jc4TXxCnq7XLRAe4p9MgDKF2VeFKMx9eZXMgGnB.rdf",
	})
	require.EqualError(t, err, "data record with IRI: regen:13toVfvdftodu8c1Jc4TXxCnq7XLRAe4p9MgDKF2VeFKMx9eZXMgGnB.rdf: not found")
}
