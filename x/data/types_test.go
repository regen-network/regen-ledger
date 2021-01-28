package data

import "testing"

func TestContentHash_Validate(t *testing.T) {
	tests := []struct {
		name    string
		sum     isContentHash_Sum
		wantErr bool
	}{
		{
			"nil",
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ch := ContentHash{
				Sum: tt.sum,
			}
			if err := ch.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
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
		// TODO: Add test cases.
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
		wantErr   bool
	}{
		{
			"right len",
			DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256,
			make([]byte, 32),
			false,
		},
		{
			"wrong len",
			DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256,
			make([]byte, 31),
			true,
		},
		{
			"unspecified",
			DigestAlgorithm_DIGEST_ALGORITHM_UNSPECIFIED,
			make([]byte, 32),
			true,
		},
		{
			"bad algorithm",
			-1,
			make([]byte, 32),
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.algorithm.Validate(tt.hash); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGraphCanonicalizationAlgorithm_Validate(t *testing.T) {
	tests := []struct {
		name    string
		x       GraphCanonicalizationAlgorithm
		wantErr bool
	}{
		{
			"urdna2015",
			GraphCanonicalizationAlgorithm_GRAPH_CANONICALIZATION_ALGORITHM_URDNA2015,
			false,
		},
		{
			"unspecified",
			GraphCanonicalizationAlgorithm_GRAPH_CANONICALIZATION_ALGORITHM_UNSPECIFIED,
			true,
		},
		{
			"bad",
			-1,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.x.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGraphMerkleTree_Validate(t *testing.T) {
	tests := []struct {
		name    string
		x       GraphMerkleTree
		wantErr bool
	}{
		{
			"unspecified",
			GraphMerkleTree_GRAPH_MERKLE_TREE_NONE_UNSPECIFIED,
			false,
		},
		{
			"bad",
			-1,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.x.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMediaType_Validate(t *testing.T) {
	tests := []struct {
		name    string
		x       MediaType
		wantErr bool
	}{
		{
			"good",
			MediaType_MEDIA_TYPE_PDF,
			false,
		},
		{
			"unspecified",
			MediaType_MEDIA_TYPE_UNSPECIFIED,
			false,
		},
		{
			"bad",
			-1,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.x.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
