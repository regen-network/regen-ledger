package server

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/regen-network/regen-ledger/x/data"
)

func TestQuery_ConvertHashToIRI(t *testing.T) {
	t.Parallel()
	s := setupBase(t)

	ch1 := &data.ContentHash{Graph: &data.ContentHash_Graph{
		Hash:                      bytes.Repeat([]byte{0}, 32),
		DigestAlgorithm:           data.DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256,
		CanonicalizationAlgorithm: data.GraphCanonicalizationAlgorithm_GRAPH_CANONICALIZATION_ALGORITHM_URDNA2015,
	}}
	iri1, err := ch1.ToIRI()
	require.NoError(t, err)

	// convert hash to iri
	res, err := s.server.ConvertHashToIRI(s.ctx, &data.ConvertHashToIRIRequest{
		ContentHash: ch1,
	})
	require.NoError(t, err)
	require.Equal(t, iri1, res.Iri)

	// query attestations with empty content hash
	_, err = s.server.ConvertHashToIRI(s.ctx, &data.ConvertHashToIRIRequest{})
	require.EqualError(t, err, "content hash cannot be empty: invalid request")

	// query attestations with invalid content hash
	_, err = s.server.ConvertHashToIRI(s.ctx, &data.ConvertHashToIRIRequest{
		ContentHash: &data.ContentHash{},
	})
	require.EqualError(t, err, "invalid data.ContentHash: invalid type")
}
