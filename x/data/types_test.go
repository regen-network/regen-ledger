package data

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestContentHash_Validate(t *testing.T) {
	tests := []struct {
		name    string
		sum     isContentHash_Sum
		wantErr string
	}{
		{
			"nil",
			nil,
			"invalid data.ContentHash type <nil>: invalid request",
		},
		{
			"good raw",
			&ContentHash_Raw_{Raw: &ContentHash_Raw{
				Hash:            make([]byte, 32),
				DigestAlgorithm: DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256,
				MediaType:       MediaType_MEDIA_TYPE_UNSPECIFIED,
			}},
			"",
		},
		{
			"good graph",
			&ContentHash_Graph_{Graph: &ContentHash_Graph{
				Hash:                      make([]byte, 32),
				DigestAlgorithm:           DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256,
				CanonicalizationAlgorithm: GraphCanonicalizationAlgorithm_GRAPH_CANONICALIZATION_ALGORITHM_URDNA2015,
				MerkleTree:                GraphMerkleTree_GRAPH_MERKLE_TREE_NONE_UNSPECIFIED,
			}},
			"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ch := ContentHash{
				Sum: tt.sum,
			}
			err := ch.Validate()
			if len(tt.wantErr) != 0 {
				require.EqualError(t, err, tt.wantErr)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestContentHash_Raw_Validate(t *testing.T) {
	type fields struct {
		Hash            []byte
		DigestAlgorithm DigestAlgorithm
		MediaType       MediaType
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			"good",
			fields{
				Hash:            make([]byte, 32),
				DigestAlgorithm: DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256,
				MediaType:       MediaType_MEDIA_TYPE_UNSPECIFIED,
			},
			false,
		},
		{
			"bad",
			fields{},
			true,
		},
		{
			"bad mediatype",
			fields{
				MediaType: -1,
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			chr := ContentHash_Raw{
				Hash:            tt.fields.Hash,
				DigestAlgorithm: tt.fields.DigestAlgorithm,
				MediaType:       tt.fields.MediaType,
			}
			if err := chr.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestContentHash_Graph_Validate(t *testing.T) {
	type fields struct {
		Hash                      []byte
		DigestAlgorithm           DigestAlgorithm
		CanonicalizationAlgorithm GraphCanonicalizationAlgorithm
		MerkleTree                GraphMerkleTree
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			"good",
			fields{
				Hash:                      make([]byte, 32),
				DigestAlgorithm:           DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256,
				CanonicalizationAlgorithm: GraphCanonicalizationAlgorithm_GRAPH_CANONICALIZATION_ALGORITHM_URDNA2015,
				MerkleTree:                GraphMerkleTree_GRAPH_MERKLE_TREE_NONE_UNSPECIFIED,
			},
			false,
		},
		{
			"bad hash",
			fields{
				Hash:                      make([]byte, 31),
				DigestAlgorithm:           DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256,
				CanonicalizationAlgorithm: GraphCanonicalizationAlgorithm_GRAPH_CANONICALIZATION_ALGORITHM_URDNA2015,
				MerkleTree:                GraphMerkleTree_GRAPH_MERKLE_TREE_NONE_UNSPECIFIED,
			},
			true,
		},
		{
			"bad digest",
			fields{
				Hash:                      make([]byte, 32),
				DigestAlgorithm:           DigestAlgorithm_DIGEST_ALGORITHM_UNSPECIFIED,
				CanonicalizationAlgorithm: GraphCanonicalizationAlgorithm_GRAPH_CANONICALIZATION_ALGORITHM_URDNA2015,
				MerkleTree:                GraphMerkleTree_GRAPH_MERKLE_TREE_NONE_UNSPECIFIED,
			},
			true,
		},
		{
			"bad alg",
			fields{
				Hash:                      make([]byte, 32),
				DigestAlgorithm:           DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256,
				CanonicalizationAlgorithm: GraphCanonicalizationAlgorithm_GRAPH_CANONICALIZATION_ALGORITHM_UNSPECIFIED,
				MerkleTree:                GraphMerkleTree_GRAPH_MERKLE_TREE_NONE_UNSPECIFIED,
			},
			true,
		},
		{
			"bad merkle tree",
			fields{
				Hash:                      make([]byte, 32),
				DigestAlgorithm:           DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256,
				CanonicalizationAlgorithm: GraphCanonicalizationAlgorithm_GRAPH_CANONICALIZATION_ALGORITHM_URDNA2015,
				MerkleTree:                -1,
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			chg := ContentHash_Graph{
				Hash:                      tt.fields.Hash,
				DigestAlgorithm:           tt.fields.DigestAlgorithm,
				CanonicalizationAlgorithm: tt.fields.CanonicalizationAlgorithm,
				MerkleTree:                tt.fields.MerkleTree,
			}
			if err := chg.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDigestAlgorithm_Validate(t *testing.T) {
	tests := []struct {
		name      string
		algorithm DigestAlgorithm
		hash      []byte
		wantErr   string
	}{
		{
			"right len",
			DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256,
			make([]byte, 32),
			"",
		},
		{
			"wrong len",
			DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256,
			make([]byte, 31),
			"expected 32 bytes for DIGEST_ALGORITHM_BLAKE2B_256, got 31: invalid request",
		},
		{
			"unspecified",
			DigestAlgorithm_DIGEST_ALGORITHM_UNSPECIFIED,
			make([]byte, 32),
			"invalid or unknown data.DigestAlgorithm DIGEST_ALGORITHM_UNSPECIFIED: invalid request",
		},
		{
			"bad algorithm",
			-1,
			make([]byte, 32),
			"invalid or unknown data.DigestAlgorithm -1: invalid request",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.algorithm.Validate(tt.hash)
			if len(tt.wantErr) != 0 {
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
		x       GraphCanonicalizationAlgorithm
		wantErr string
	}{
		{
			"urdna2015",
			GraphCanonicalizationAlgorithm_GRAPH_CANONICALIZATION_ALGORITHM_URDNA2015,
			"",
		},
		{
			"unspecified",
			GraphCanonicalizationAlgorithm_GRAPH_CANONICALIZATION_ALGORITHM_UNSPECIFIED,
			"invalid data.GraphCanonicalizationAlgorithm GRAPH_CANONICALIZATION_ALGORITHM_UNSPECIFIED: invalid request",
		},
		{
			"bad",
			-1,
			"unknown data.GraphCanonicalizationAlgorithm -1: invalid request",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.x.Validate()
			if len(tt.wantErr) != 0 {
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
		x       GraphMerkleTree
		wantErr string
	}{
		{
			"unspecified",
			GraphMerkleTree_GRAPH_MERKLE_TREE_NONE_UNSPECIFIED,
			"",
		},
		{
			"bad",
			-1,
			"unknown data.GraphMerkleTree -1: invalid request",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.x.Validate()
			if len(tt.wantErr) != 0 {
				require.EqualError(t, err, tt.wantErr)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestMediaType_Validate(t *testing.T) {
	tests := []struct {
		name    string
		x       MediaType
		wantErr string
	}{
		{
			"good",
			MediaType_MEDIA_TYPE_PDF,
			"",
		},
		{
			"unspecified",
			MediaType_MEDIA_TYPE_UNSPECIFIED,
			"",
		},
		{
			"bad",
			-1,
			"unknown data.MediaType -1: invalid request",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.x.Validate()
			if len(tt.wantErr) != 0 {
				require.EqualError(t, err, tt.wantErr)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
