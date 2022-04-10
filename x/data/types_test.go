package data

import (
	"strconv"
	"testing"

	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
)

type contentHash struct {
	t   gocuke.TestingT
	ch  *ContentHash
	err error
}

func TestTypes(t *testing.T) {
	gocuke.NewRunner(t, &contentHash{}).Path("./features/types_content_hash.feature").Run()
}

func (s *contentHash) Before(t gocuke.TestingT) {
	s.t = t
}

func (s *contentHash) AnEmptyContentHash() {
	s.ch = &ContentHash{}
}

func (s *contentHash) AnEmptyRawContentHash() {
	s.ch = &ContentHash{
		Raw: &ContentHash_Raw{},
	}
}

func (s *contentHash) AnEmptyGraphContentHash() {
	s.ch = &ContentHash{
		Graph: &ContentHash_Graph{},
	}
}

func (s *contentHash) ARawContentHashOf(a string, b string, c string) {
	length, err := strconv.Atoi(a)
	require.NoError(s.t, err)

	digest, err := strconv.Atoi(b)
	require.NoError(s.t, err)

	media, err := strconv.Atoi(c)
	require.NoError(s.t, err)

	s.ch = &ContentHash{
		Raw: &ContentHash_Raw{
			Hash:            make([]byte, length),
			DigestAlgorithm: DigestAlgorithm(digest),
			MediaType:       RawMediaType(media),
		},
	}
}

func (s *contentHash) AGraphContentHashOf(a string, b string, c string, d string) {
	length, err := strconv.Atoi(a)
	require.NoError(s.t, err)

	digest, err := strconv.Atoi(b)
	require.NoError(s.t, err)

	canon, err := strconv.Atoi(c)
	require.NoError(s.t, err)

	merkle, err := strconv.Atoi(d)
	require.NoError(s.t, err)

	s.ch = &ContentHash{
		Graph: &ContentHash_Graph{
			Hash:                      make([]byte, length),
			DigestAlgorithm:           DigestAlgorithm(digest),
			CanonicalizationAlgorithm: GraphCanonicalizationAlgorithm(canon),
			MerkleTree:                GraphMerkleTree(merkle),
		},
	}
}

func (s *contentHash) TheContentHashIsValidated() {
	s.err = s.ch.Validate()
}

func (s *contentHash) AnErrorOf(a string) {
	if a == "" {
		require.NoError(s.t, s.err)
	} else {
		require.EqualError(s.t, s.err, a)
	}
}
