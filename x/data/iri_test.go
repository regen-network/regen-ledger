package data

import (
	"testing"

	"gotest.tools/v3/assert"

	"github.com/stretchr/testify/require"
)

func TestContentHash_Graph_ToIRI(t *testing.T) {
	hash := []byte("abcdefghijklmnopqrstuvwxyz123456")

	tests := []struct {
		name    string
		chg     ContentHash_Graph
		wantErr bool
		want    string
	}{
		{
			"hash too long",
			ContentHash_Graph{
				Hash:                      []byte("abcdefghijklmnopqrstuvwxyz1234567abcdefghijklmnopqrstuvwxyz1234567"),
				DigestAlgorithm:           2,
				CanonicalizationAlgorithm: 2,
				MerkleTree:                1,
			},
			true,
			"",
		},
		{
			"missing digest algorithm",
			ContentHash_Graph{
				Hash:                      hash,
				DigestAlgorithm:           0,
				CanonicalizationAlgorithm: 1,
				MerkleTree:                0,
			},
			true,
			"",
		},
		{
			"missing canonicalization algorithm",
			ContentHash_Graph{
				Hash:                      hash,
				DigestAlgorithm:           1,
				CanonicalizationAlgorithm: 0,
				MerkleTree:                0,
			},
			true,
			"",
		},
		{
			"no merkle tree",
			ContentHash_Graph{
				Hash:                      hash,
				DigestAlgorithm:           1,
				CanonicalizationAlgorithm: 1,
				MerkleTree:                0,
			},
			false,
			"regen:13toVgf5aZqSVSeJQv562xkkeoe3rr3bJWa29PHVKVf77VAkVMcDvVd.rdf",
		},
		{
			"have merkle tree",
			ContentHash_Graph{
				Hash:                      hash,
				DigestAlgorithm:           2,
				CanonicalizationAlgorithm: 2,
				MerkleTree:                1,
			},
			false,
			"regen:13uTWjemEQUTrKexY94vQytMLSQnSbJrVuEjU7rQJggC5q6W41qPueS.rdf",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			iri, err := tt.chg.ToIRI()
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, iri, tt.want)
			}
		})
	}
}

func TestContentHash_Raw_ToIRI(t *testing.T) {
	hash := []byte("abcdefghijklmnopqrstuvwxyz123456")

	tests := []struct {
		name    string
		chr     ContentHash_Raw
		wantErr bool
		want    string
	}{
		{
			"valid raw data",
			ContentHash_Raw{
				Hash:            hash,
				DigestAlgorithm: 2,
				FileExtension:   "jpg",
			},
			false,
			"regen:115dNuUmeLdZEBP9opwjTqvH8GCx56SGfgmDHbMeYtwe8neW4iXb.jpg",
		},
		{
			"hash too short",
			ContentHash_Raw{
				Hash:            []byte("abc"),
				DigestAlgorithm: 2,
				FileExtension:   "jpg",
			},
			true,
			"",
		},
		{
			"hash too long",
			ContentHash_Raw{
				Hash:            []byte("abcdefghijklmnopqrstuvwxyz1234567abcdefghijklmnopqrstuvwxyz1234567"),
				DigestAlgorithm: 2,
				FileExtension:   "jpg",
			},
			true,
			"",
		},
		{
			"missing digest algorithm",
			ContentHash_Raw{
				Hash:            hash,
				DigestAlgorithm: 0,
				FileExtension:   "jpg",
			},
			true,
			"",
		},
		{
			"ext too short",
			ContentHash_Raw{
				Hash:            hash,
				DigestAlgorithm: 2,
				FileExtension:   "j",
			},
			true,
			"",
		},
		{
			"ext too long",
			ContentHash_Raw{
				Hash:            hash,
				DigestAlgorithm: 2,
				FileExtension:   "abcdefg",
			},
			true,
			"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			iri, err := tt.chr.ToIRI()
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, iri, tt.want)
			}
		})
	}
}

func TestParseIRI(t *testing.T) {
	hash := []byte("abcdefghijklmnopqrstuvwxyz123456")

	tests := []struct {
		name     string
		iri      string
		wantHash *ContentHash
		wantErr  string
	}{
		{
			name:    "invalid prefix",
			iri:     "cosmos:13toVgf5aZqSVSeJQv562xkkeoe3rr3bJWa29PHVKVf77VAkVMcDvVd.rdf",
			wantErr: "failed to parse IRI cosmos:13toVgf5aZqSVSeJQv562xkkeoe3rr3bJWa29PHVKVf77VAkVMcDvVd.rdf: regen: prefix required: invalid IRI",
		},
		{
			name:    "invalid extension",
			iri:     "regen:13toVgf5aZqSVSeJQv562xkkeoe3rr3bJWa29PHVKVf77VAkVMcDvVd",
			wantErr: "failed to parse IRI regen:13toVgf5aZqSVSeJQv562xkkeoe3rr3bJWa29PHVKVf77VAkVMcDvVd: extension required: invalid IRI",
		},
		{
			name:    "invalid checksum",
			iri:     "regen:23toVgf5aZqSVSeJQv562xkkeoe3rr3bJWa29PHVKVf77VAkVMcDvVd.rdf",
			wantErr: "failed to parse IRI regen:23toVgf5aZqSVSeJQv562xkkeoe3rr3bJWa29PHVKVf77VAkVMcDvVd.rdf: checksum error: invalid IRI",
		},
		{
			name:    "invalid version",
			iri:     "regen:esV713VcRqk5TWxDgKQjGSpN4aXL4a9XTzbWRduCMQDqq2zo3TtX.rdf",
			wantErr: "failed to parse IRI regen:esV713VcRqk5TWxDgKQjGSpN4aXL4a9XTzbWRduCMQDqq2zo3TtX.rdf: invalid version 1: invalid IRI",
		},
		{
			name: "valid media bin",
			iri:  "regen:113gdjFKcVCt13Za6vN7TtbgMM6LMSjRnu89BMCxeuHdkJ1hWUmy.bin",
			wantHash: &ContentHash{Raw: &ContentHash_Raw{
				Hash:            hash,
				DigestAlgorithm: uint32(DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256),
				FileExtension:   "bin",
			}},
		},
		{
			name: "valid media txt",
			iri:  "regen:113gdjFKcVCt13Za6vN7TtbgMM6LMSjRnu89BMCxeuHdkJ1hWUmy.txt",
			wantHash: &ContentHash{Raw: &ContentHash_Raw{
				Hash:            hash,
				DigestAlgorithm: uint32(DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256),
				FileExtension:   "txt",
			}},
		},
		{
			name: "valid media csv",
			iri:  "regen:113gdjFKcVCt13Za6vN7TtbgMM6LMSjRnu89BMCxeuHdkJ1hWUmy.csv",
			wantHash: &ContentHash{Raw: &ContentHash_Raw{
				Hash:            hash,
				DigestAlgorithm: uint32(DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256),
				FileExtension:   "csv",
			}},
		},
		{
			name: "valid graph",
			iri:  "regen:13toVgf5aZqSVSeJQv562xkkeoe3rr3bJWa29PHVKVf77VAkVMcDvVd.rdf",
			wantHash: &ContentHash{Graph: &ContentHash_Graph{
				Hash:                      hash,
				DigestAlgorithm:           uint32(DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256),
				CanonicalizationAlgorithm: uint32(GraphCanonicalizationAlgorithm_GRAPH_CANONICALIZATION_ALGORITHM_RDFC_1_0),
			}},
		},
		{
			name: "valid graph with merkle tree",
			iri:  "regen:13uTWjemEQUTrKexY94vQytMLSQnSbJrVuEjU7rQJggC5q6W41qPueS.rdf",
			wantHash: &ContentHash{Graph: &ContentHash_Graph{
				Hash:                      hash,
				DigestAlgorithm:           2,
				CanonicalizationAlgorithm: 2,
				MerkleTree:                1,
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			contentHash, err := ParseIRI(tt.iri)
			if err != nil {
				require.EqualError(t, err, tt.wantErr)
			} else {
				assert.DeepEqual(t, contentHash, tt.wantHash)
			}
		})
	}
}
