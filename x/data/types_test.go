package data

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestContentHash_Validate(t *testing.T) {
	tests := []struct {
		name string
		ch   ContentHash
		err  string
	}{
		{
			"invalid empty",
			ContentHash{},
			"content hash must be one of raw type or graph type: invalid request",
		},
		{
			"invalid both types",
			ContentHash{
				Raw: &ContentHash_Raw{
					Hash:            make([]byte, 32),
					DigestAlgorithm: DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256,
					MediaType:       RawMediaType_RAW_MEDIA_TYPE_UNSPECIFIED,
				},
				Graph: &ContentHash_Graph{
					Hash:                      make([]byte, 32),
					DigestAlgorithm:           DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256,
					CanonicalizationAlgorithm: GraphCanonicalizationAlgorithm_GRAPH_CANONICALIZATION_ALGORITHM_URDNA2015,
					MerkleTree:                GraphMerkleTree_GRAPH_MERKLE_TREE_NONE_UNSPECIFIED,
				},
			},
			"content hash must be one of raw type or graph type: invalid request",
		},
		{
			"valid raw type",
			ContentHash{
				Raw: &ContentHash_Raw{
					Hash:            make([]byte, 32),
					DigestAlgorithm: DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256,
				},
			},
			"",
		},
		{
			"valid graph type",
			ContentHash{
				Graph: &ContentHash_Graph{
					Hash:                      make([]byte, 32),
					DigestAlgorithm:           DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256,
					CanonicalizationAlgorithm: GraphCanonicalizationAlgorithm_GRAPH_CANONICALIZATION_ALGORITHM_URDNA2015,
				},
			},
			"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.ch.Validate()
			if err != nil {
				require.EqualError(t, err, tt.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestContentHash_Raw_Validate(t *testing.T) {
	tests := []struct {
		name string
		chr  *ContentHash_Raw
		err  string
	}{
		{
			"invalid digest unknown",
			&ContentHash_Raw{
				DigestAlgorithm: -1,
			},
			"invalid or unknown data.DigestAlgorithm -1: invalid request",
		},
		{
			"invalid digest unspecified",
			&ContentHash_Raw{
				DigestAlgorithm: DigestAlgorithm_DIGEST_ALGORITHM_UNSPECIFIED,
			},
			"invalid or unknown data.DigestAlgorithm DIGEST_ALGORITHM_UNSPECIFIED: invalid request",
		},
		{
			"invalid hash empty",
			&ContentHash_Raw{
				DigestAlgorithm: DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256,
			},
			"expected 32 bytes for DIGEST_ALGORITHM_BLAKE2B_256, got 0: invalid request",
		},
		{
			"invalid hash length",
			&ContentHash_Raw{
				Hash:            make([]byte, 16),
				DigestAlgorithm: DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256,
			},
			"expected 32 bytes for DIGEST_ALGORITHM_BLAKE2B_256, got 16: invalid request",
		},
		{
			"invalid media unknown",
			&ContentHash_Raw{
				Hash:            make([]byte, 32),
				DigestAlgorithm: DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256,
				MediaType:       -1,
			},
			"unknown data.RawMediaType -1: invalid request",
		},
		{
			"valid media unspecified",
			&ContentHash_Raw{
				Hash:            make([]byte, 32),
				DigestAlgorithm: DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256,
			},
			"",
		},
		{
			"valid",
			&ContentHash_Raw{
				Hash:            make([]byte, 32),
				DigestAlgorithm: DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256,
				MediaType:       RawMediaType_RAW_MEDIA_TYPE_PDF,
			},
			"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.chr.Validate()
			if err != nil {
				require.EqualError(t, err, tt.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestContentHash_Graph_Validate(t *testing.T) {
	tests := []struct {
		name string
		chg  *ContentHash_Graph
		err  string
	}{
		{
			"invalid digest unknown",
			&ContentHash_Graph{
				DigestAlgorithm: -1,
			},
			"invalid or unknown data.DigestAlgorithm -1: invalid request",
		},
		{
			"invalid digest unspecified",
			&ContentHash_Graph{
				DigestAlgorithm: DigestAlgorithm_DIGEST_ALGORITHM_UNSPECIFIED,
			},
			"invalid or unknown data.DigestAlgorithm DIGEST_ALGORITHM_UNSPECIFIED: invalid request",
		},
		{
			"invalid hash empty",
			&ContentHash_Graph{
				DigestAlgorithm: DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256,
			},
			"expected 32 bytes for DIGEST_ALGORITHM_BLAKE2B_256, got 0: invalid request",
		},
		{
			"invalid hash length",
			&ContentHash_Graph{
				Hash:            make([]byte, 16),
				DigestAlgorithm: DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256,
			},
			"expected 32 bytes for DIGEST_ALGORITHM_BLAKE2B_256, got 16: invalid request",
		},
		{
			"invalid canonical unknown",
			&ContentHash_Graph{
				Hash:                      make([]byte, 32),
				DigestAlgorithm:           DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256,
				CanonicalizationAlgorithm: -1,
			},
			"unknown data.GraphCanonicalizationAlgorithm -1: invalid request",
		},
		{
			"invalid canonical unspecified",
			&ContentHash_Graph{
				Hash:                      make([]byte, 32),
				DigestAlgorithm:           DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256,
				CanonicalizationAlgorithm: GraphCanonicalizationAlgorithm_GRAPH_CANONICALIZATION_ALGORITHM_UNSPECIFIED,
			},
			"invalid data.GraphCanonicalizationAlgorithm GRAPH_CANONICALIZATION_ALGORITHM_UNSPECIFIED: invalid request",
		},
		{
			"invalid merkle unknown",
			&ContentHash_Graph{
				Hash:                      make([]byte, 32),
				DigestAlgorithm:           DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256,
				CanonicalizationAlgorithm: GraphCanonicalizationAlgorithm_GRAPH_CANONICALIZATION_ALGORITHM_URDNA2015,
				MerkleTree:                -1,
			},
			"unknown data.GraphMerkleTree -1: invalid request",
		},
		{
			"valid",
			&ContentHash_Graph{
				Hash:                      make([]byte, 32),
				DigestAlgorithm:           DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256,
				CanonicalizationAlgorithm: GraphCanonicalizationAlgorithm_GRAPH_CANONICALIZATION_ALGORITHM_URDNA2015,
			},
			"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.chg.Validate()
			if err != nil {
				require.EqualError(t, err, tt.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestDigestAlgorithm_Validate(t *testing.T) {
	tests := []struct {
		name string
		da   DigestAlgorithm
		hash []byte
		err  string
	}{
		{
			"invalid unknown",
			-1,
			make([]byte, 32),
			"invalid or unknown data.DigestAlgorithm -1: invalid request",
		},
		{
			"invalid unspecified",
			DigestAlgorithm_DIGEST_ALGORITHM_UNSPECIFIED,
			make([]byte, 32),
			"invalid or unknown data.DigestAlgorithm DIGEST_ALGORITHM_UNSPECIFIED: invalid request",
		},
		{
			"invalid hash length",
			DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256,
			make([]byte, 16),
			"expected 32 bytes for DIGEST_ALGORITHM_BLAKE2B_256, got 16: invalid request",
		},

		{
			"valid",
			DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256,
			make([]byte, 32),
			"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.da.Validate(tt.hash)
			if err != nil {
				require.EqualError(t, err, tt.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestRawMediaType_Validate(t *testing.T) {
	tests := []struct {
		name    string
		rmt     RawMediaType
		wantErr string
	}{

		{
			"invalid unknown",
			-1,
			"unknown data.RawMediaType -1: invalid request",
		},
		{
			"valid unspecified",
			RawMediaType_RAW_MEDIA_TYPE_UNSPECIFIED,
			"",
		},
		{
			"valid",
			RawMediaType_RAW_MEDIA_TYPE_PDF,
			"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.rmt.Validate()
			if err != nil {
				require.EqualError(t, err, tt.wantErr)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestGraphCanonicalizationAlgorithm_Validate(t *testing.T) {
	tests := []struct {
		name    string
		gca     GraphCanonicalizationAlgorithm
		wantErr string
	}{
		{
			"invalid unknown",
			-1,
			"unknown data.GraphCanonicalizationAlgorithm -1: invalid request",
		},
		{
			"invalid unspecified",
			GraphCanonicalizationAlgorithm_GRAPH_CANONICALIZATION_ALGORITHM_UNSPECIFIED,
			"invalid data.GraphCanonicalizationAlgorithm GRAPH_CANONICALIZATION_ALGORITHM_UNSPECIFIED: invalid request",
		},
		{
			"valid",
			GraphCanonicalizationAlgorithm_GRAPH_CANONICALIZATION_ALGORITHM_URDNA2015,
			"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.gca.Validate()
			if err != nil {
				require.EqualError(t, err, tt.wantErr)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestGraphMerkleTree_Validate(t *testing.T) {
	tests := []struct {
		name    string
		gmt     GraphMerkleTree
		wantErr string
	}{
		{
			"invalid unknown",
			-1,
			"unknown data.GraphMerkleTree -1: invalid request",
		},
		{
			"valid",
			GraphMerkleTree_GRAPH_MERKLE_TREE_NONE_UNSPECIFIED,
			"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.gmt.Validate()
			if err != nil {
				require.EqualError(t, err, tt.wantErr)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
