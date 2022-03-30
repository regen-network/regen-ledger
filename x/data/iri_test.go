package data

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestContentHash_ToIRI(t *testing.T) {
	tests := []struct {
		name    string
		fields  ContentHash
		want    string
		wantErr bool
	}{
		{
			"bad graph",
			ContentHash{Graph: &ContentHash_Graph{}},
			"",
			true,
		},
		{
			"bad raw",
			ContentHash{Raw: &ContentHash_Raw{}},
			"",
			true,
		},
		{
			"nil",
			ContentHash{},
			"",
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ch := ContentHash{}
			got, err := ch.ToIRI()
			if (err != nil) != tt.wantErr {
				t.Errorf("ToIRI() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ToIRI() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContentHash_Graph_ToIRI(t *testing.T) {
	type fields struct {
		Hash                      []byte
		DigestAlgorithm           DigestAlgorithm
		CanonicalizationAlgorithm GraphCanonicalizationAlgorithm
		MerkleTree                GraphMerkleTree
	}

	hash1 := []byte("abcdefghijklmnopqrstuvwxyz123456")

	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{
			"graph",
			fields{
				Hash:                      hash1,
				DigestAlgorithm:           DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256,
				CanonicalizationAlgorithm: GraphCanonicalizationAlgorithm_GRAPH_CANONICALIZATION_ALGORITHM_URDNA2015,
				MerkleTree:                GraphMerkleTree_GRAPH_MERKLE_TREE_NONE_UNSPECIFIED,
			},
			"regen:13toVgf5aZqSVSeJQv562xkkeoe3rr3bJWa29PHVKVf77VAkVMcDvVd.rdf",
			false,
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
			got, err := chg.ToIRI()
			if (err != nil) != tt.wantErr {
				t.Errorf("ToIRI() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ToIRI() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContentHash_Raw_ToIRI(t *testing.T) {
	type fields struct {
		Hash            []byte
		DigestAlgorithm DigestAlgorithm
		MediaType       RawMediaType
	}

	hash1 := []byte("abcdefghijklmnopqrstuvwxyz123456")

	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{
			"pdf",
			fields{
				Hash:            hash1,
				DigestAlgorithm: DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256,
				MediaType:       RawMediaType_RAW_MEDIA_TYPE_PDF,
			},
			"regen:113gdjFKcVCt13Za6vN7TtbgMM6LMSjRnu89BMCxeuHdkJ1hWUmy.pdf",
			false,
		},
		{
			"bad media type",
			fields{
				Hash:            hash1,
				DigestAlgorithm: DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256,
				MediaType:       -1,
			},
			"",
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
			got, err := chr.ToIRI()
			if (err != nil) != tt.wantErr {
				t.Errorf("ToIRI() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ToIRI() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMediaType_ToExtension(t *testing.T) {
	// ensure every good media type has an extension
	for mt := range RawMediaType_name {
		_, err := RawMediaType(mt).ToExtension()
		require.NoError(t, err)
	}

	_, err := RawMediaType(-1).ToExtension()
	require.Error(t, err)
}

func TestParseIRI(t *testing.T) {
	tests := []struct {
		name     string
		iri      string
		wantHash *ContentHash
		wantErr  bool
	}{
		{
			name: "raw",
			iri:  "regen:113gdjFKcVCt13Za6vN7TtbgMM6LMSjRnu89BMCxeuHdkJ1hWUmy.pdf",
			wantHash: &ContentHash{Raw: &ContentHash_Raw{
				Hash:            []byte("abcdefghijklmnopqrstuvwxyz123456"),
				DigestAlgorithm: DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256,
				MediaType:       RawMediaType_RAW_MEDIA_TYPE_PDF,
			}},
		},
		{
			name: "graph",
			iri:  "regen:13toVgf5aZqSVSeJQv562xkkeoe3rr3bJWa29PHVKVf77VAkVMcDvVd.rdf",
			wantHash: &ContentHash{Graph: &ContentHash_Graph{
				Hash:                      []byte("abcdefghijklmnopqrstuvwxyz123456"),
				DigestAlgorithm:           DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256,
				CanonicalizationAlgorithm: GraphCanonicalizationAlgorithm_GRAPH_CANONICALIZATION_ALGORITHM_URDNA2015,
			}},
		},
		{
			name:    "no ext",
			iri:     "regen:13toVgf5aZqSVSeJQv562xkkeoe3rr3bJWa29PHVKVf77VAkVMcDvVd",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			contentHash, err := ParseIRI(tt.iri)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseIRI() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(contentHash, tt.wantHash) {
				t.Errorf("ParseIRI() got = %v, want %v", contentHash, tt.wantHash)
			}
		})
	}
}
