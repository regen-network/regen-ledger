package testsuite

import (
	"crypto"

	"github.com/regen-network/regen-ledger/x/data"
)

func (s *IntegrationTestSuite) createIRIAndGraphHash(content []byte) (string, *data.ContentHash) {
	require := s.Require()

	hash := crypto.BLAKE2b_256.New()

	_, err := hash.Write(content)
	require.NoError(err)

	digest := hash.Sum(nil)

	ch := data.ContentHash{
		Graph: &data.ContentHash_Graph{
			Hash:                      digest,
			DigestAlgorithm:           data.DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256,
			CanonicalizationAlgorithm: data.GraphCanonicalizationAlgorithm_GRAPH_CANONICALIZATION_ALGORITHM_URDNA2015,
		},
	}

	iri, err := ch.GetGraph().ToIRI()
	require.NoError(err)

	return iri, &ch
}
