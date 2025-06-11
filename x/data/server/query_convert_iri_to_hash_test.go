package server

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/regen-network/regen-ledger/x/data/v3"
)

func TestQuery_ConvertIRIToHash(t *testing.T) {
	t.Parallel()
	s := setupBase(t)

	ch1 := &data.ContentHash{Graph: &data.ContentHash_Graph{
		Hash:                      bytes.Repeat([]byte{0}, 32),
		DigestAlgorithm:           uint32(data.DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256),
		CanonicalizationAlgorithm: uint32(data.GraphCanonicalizationAlgorithm_GRAPH_CANONICALIZATION_ALGORITHM_RDFC_1_0),
	}}
	iri1, err := ch1.ToIRI()
	require.NoError(t, err)

	// convert hash to iri
	res, err := s.server.ConvertIRIToHash(s.ctx, &data.ConvertIRIToHashRequest{
		Iri: iri1,
	})
	require.NoError(t, err)
	require.Equal(t, ch1, res.ContentHash)

	// query attestations with empty content hash
	_, err = s.server.ConvertIRIToHash(s.ctx, &data.ConvertIRIToHashRequest{})
	require.EqualError(t, err, "IRI cannot be empty: invalid argument")

	// query attestations with invalid content hash
	_, err = s.server.ConvertIRIToHash(s.ctx, &data.ConvertIRIToHashRequest{
		Iri: "foo",
	})
	require.EqualError(t, err, "failed to parse IRI foo: regen: prefix required: invalid IRI: invalid argument")
}
