package data

import (
	"testing"

	"gotest.tools/v3/assert"

	"github.com/stretchr/testify/require"
)

func TestContentHash_Graph_ToIRI(t *testing.T) {
	hash := []byte("abcdefghijklmnopqrstuvwxyz123456")

	tests := []struct {
		name string
		chg  ContentHash_Graph
		want string
	}{
		{
			"valid graph",
			ContentHash_Graph{
				Hash:                      hash,
				DigestAlgorithm:           DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256,
				CanonicalizationAlgorithm: GraphCanonicalizationAlgorithm_GRAPH_CANONICALIZATION_ALGORITHM_URDNA2015,
				MerkleTree:                GraphMerkleTree_GRAPH_MERKLE_TREE_NONE_UNSPECIFIED,
			},
			"regen:13toVgf5aZqSVSeJQv562xkkeoe3rr3bJWa29PHVKVf77VAkVMcDvVd.rdf",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			iri, err := tt.chg.ToIRI()
			require.NoError(t, err)
			assert.Equal(t, iri, tt.want)
		})
	}
}

func TestContentHash_Raw_ToIRI(t *testing.T) {
	hash := []byte("abcdefghijklmnopqrstuvwxyz123456")

	tests := []struct {
		name string
		chr  ContentHash_Raw
		want string
	}{
		{
			"valid media bin",
			ContentHash_Raw{
				Hash:            hash,
				DigestAlgorithm: DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256,
				MediaType:       RawMediaType_RAW_MEDIA_TYPE_UNSPECIFIED,
			},
			"regen:113gdjFKcVCt13Za6vN7TtbgMM6LMSjRnu89BMCxeuHdkJ1hWUmy.bin",
		},
		{
			"valid media txt",
			ContentHash_Raw{
				Hash:            hash,
				DigestAlgorithm: DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256,
				MediaType:       RawMediaType_RAW_MEDIA_TYPE_TEXT_PLAIN,
			},
			"regen:113gdjFKcVCt13Za6vN7TtbgMM6LMSjRnu89BMCxeuHdkJ1hWUmy.txt",
		},
		{
			"valid media csv",
			ContentHash_Raw{
				Hash:            hash,
				DigestAlgorithm: DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256,
				MediaType:       RawMediaType_RAW_MEDIA_TYPE_CSV,
			},
			"regen:113gdjFKcVCt13Za6vN7TtbgMM6LMSjRnu89BMCxeuHdkJ1hWUmy.csv",
		},
		{
			"valid media json",
			ContentHash_Raw{
				Hash:            hash,
				DigestAlgorithm: DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256,
				MediaType:       RawMediaType_RAW_MEDIA_TYPE_JSON,
			},
			"regen:113gdjFKcVCt13Za6vN7TtbgMM6LMSjRnu89BMCxeuHdkJ1hWUmy.json",
		},
		{
			"valid media xml",
			ContentHash_Raw{
				Hash:            hash,
				DigestAlgorithm: DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256,
				MediaType:       RawMediaType_RAW_MEDIA_TYPE_XML,
			},
			"regen:113gdjFKcVCt13Za6vN7TtbgMM6LMSjRnu89BMCxeuHdkJ1hWUmy.xml",
		},
		{
			"valid media pdf",
			ContentHash_Raw{
				Hash:            hash,
				DigestAlgorithm: DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256,
				MediaType:       RawMediaType_RAW_MEDIA_TYPE_PDF,
			},
			"regen:113gdjFKcVCt13Za6vN7TtbgMM6LMSjRnu89BMCxeuHdkJ1hWUmy.pdf",
		},
		{
			"valid media tiff",
			ContentHash_Raw{
				Hash:            hash,
				DigestAlgorithm: DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256,
				MediaType:       RawMediaType_RAW_MEDIA_TYPE_TIFF,
			},
			"regen:113gdjFKcVCt13Za6vN7TtbgMM6LMSjRnu89BMCxeuHdkJ1hWUmy.tiff",
		},
		{
			"valid media jpg",
			ContentHash_Raw{
				Hash:            hash,
				DigestAlgorithm: DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256,
				MediaType:       RawMediaType_RAW_MEDIA_TYPE_JPG,
			},
			"regen:113gdjFKcVCt13Za6vN7TtbgMM6LMSjRnu89BMCxeuHdkJ1hWUmy.jpg",
		},
		{
			"valid media png",
			ContentHash_Raw{
				Hash:            hash,
				DigestAlgorithm: DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256,
				MediaType:       RawMediaType_RAW_MEDIA_TYPE_PNG,
			},
			"regen:113gdjFKcVCt13Za6vN7TtbgMM6LMSjRnu89BMCxeuHdkJ1hWUmy.png",
		},
		{
			"valid media svg",
			ContentHash_Raw{
				Hash:            hash,
				DigestAlgorithm: DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256,
				MediaType:       RawMediaType_RAW_MEDIA_TYPE_SVG,
			},
			"regen:113gdjFKcVCt13Za6vN7TtbgMM6LMSjRnu89BMCxeuHdkJ1hWUmy.svg",
		},
		{
			"valid media webp",
			ContentHash_Raw{
				Hash:            hash,
				DigestAlgorithm: DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256,
				MediaType:       RawMediaType_RAW_MEDIA_TYPE_WEBP,
			},
			"regen:113gdjFKcVCt13Za6vN7TtbgMM6LMSjRnu89BMCxeuHdkJ1hWUmy.webp",
		},
		{
			"valid media avif",
			ContentHash_Raw{
				Hash:            hash,
				DigestAlgorithm: DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256,
				MediaType:       RawMediaType_RAW_MEDIA_TYPE_AVIF,
			},
			"regen:113gdjFKcVCt13Za6vN7TtbgMM6LMSjRnu89BMCxeuHdkJ1hWUmy.avif",
		},
		{
			"valid media gif",
			ContentHash_Raw{
				Hash:            hash,
				DigestAlgorithm: DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256,
				MediaType:       RawMediaType_RAW_MEDIA_TYPE_GIF,
			},
			"regen:113gdjFKcVCt13Za6vN7TtbgMM6LMSjRnu89BMCxeuHdkJ1hWUmy.gif",
		},
		{
			"valid media apng",
			ContentHash_Raw{
				Hash:            hash,
				DigestAlgorithm: DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256,
				MediaType:       RawMediaType_RAW_MEDIA_TYPE_APNG,
			},
			"regen:113gdjFKcVCt13Za6vN7TtbgMM6LMSjRnu89BMCxeuHdkJ1hWUmy.apng",
		},
		{
			"valid media mpeg",
			ContentHash_Raw{
				Hash:            hash,
				DigestAlgorithm: DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256,
				MediaType:       RawMediaType_RAW_MEDIA_TYPE_MPEG,
			},
			"regen:113gdjFKcVCt13Za6vN7TtbgMM6LMSjRnu89BMCxeuHdkJ1hWUmy.mpeg",
		},
		{
			"valid media mp4",
			ContentHash_Raw{
				Hash:            hash,
				DigestAlgorithm: DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256,
				MediaType:       RawMediaType_RAW_MEDIA_TYPE_MP4,
			},
			"regen:113gdjFKcVCt13Za6vN7TtbgMM6LMSjRnu89BMCxeuHdkJ1hWUmy.mp4",
		},
		{
			"valid media webm",
			ContentHash_Raw{
				Hash:            hash,
				DigestAlgorithm: DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256,
				MediaType:       RawMediaType_RAW_MEDIA_TYPE_WEBM,
			},
			"regen:113gdjFKcVCt13Za6vN7TtbgMM6LMSjRnu89BMCxeuHdkJ1hWUmy.webm",
		},
		{
			"valid media ogg",
			ContentHash_Raw{
				Hash:            hash,
				DigestAlgorithm: DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256,
				MediaType:       RawMediaType_RAW_MEDIA_TYPE_OGG,
			},
			"regen:113gdjFKcVCt13Za6vN7TtbgMM6LMSjRnu89BMCxeuHdkJ1hWUmy.ogg",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			iri, err := tt.chr.ToIRI()
			require.NoError(t, err)
			assert.Equal(t, iri, tt.want)
		})
	}
}

func TestMediaType_ToExtension(t *testing.T) {
	// ensure every valid media type has an extension
	for mt := range RawMediaType_name {
		_, err := RawMediaType(mt).ToExtension()
		require.NoError(t, err)
	}

	_, err := RawMediaType(-1).ToExtension()
	require.Error(t, err)
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
			wantErr: "failed to parse IRI regen:esV713VcRqk5TWxDgKQjGSpN4aXL4a9XTzbWRduCMQDqq2zo3TtX.rdf: invalid version: invalid IRI",
		},
		{
			name:    "invalid media extension",
			iri:     "regen:114DDL1RtVwKpfqgaPfAG153ckiKfuPEgTT7tEGs1Hic5sC9dCta.abc",
			wantErr: "failed to resolve media type for extension abc, expected bin: invalid media extension",
		},

		{
			name: "valid media bin",
			iri:  "regen:113gdjFKcVCt13Za6vN7TtbgMM6LMSjRnu89BMCxeuHdkJ1hWUmy.bin",
			wantHash: &ContentHash{Raw: &ContentHash_Raw{
				Hash:            hash,
				DigestAlgorithm: DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256,
				MediaType:       RawMediaType_RAW_MEDIA_TYPE_UNSPECIFIED,
			}},
		},
		{
			name: "valid media txt",
			iri:  "regen:113gdjFKcVCt13Za6vN7TtbgMM6LMSjRnu89BMCxeuHdkJ1hWUmy.txt",
			wantHash: &ContentHash{Raw: &ContentHash_Raw{
				Hash:            hash,
				DigestAlgorithm: DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256,
				MediaType:       RawMediaType_RAW_MEDIA_TYPE_TEXT_PLAIN,
			}},
		},
		{
			name: "valid media csv",
			iri:  "regen:113gdjFKcVCt13Za6vN7TtbgMM6LMSjRnu89BMCxeuHdkJ1hWUmy.csv",
			wantHash: &ContentHash{Raw: &ContentHash_Raw{
				Hash:            hash,
				DigestAlgorithm: DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256,
				MediaType:       RawMediaType_RAW_MEDIA_TYPE_CSV,
			}},
		},
		{
			name: "valid media json",
			iri:  "regen:113gdjFKcVCt13Za6vN7TtbgMM6LMSjRnu89BMCxeuHdkJ1hWUmy.json",
			wantHash: &ContentHash{Raw: &ContentHash_Raw{
				Hash:            hash,
				DigestAlgorithm: DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256,
				MediaType:       RawMediaType_RAW_MEDIA_TYPE_JSON,
			}},
		},
		{
			name: "valid media xml",
			iri:  "regen:113gdjFKcVCt13Za6vN7TtbgMM6LMSjRnu89BMCxeuHdkJ1hWUmy.xml",
			wantHash: &ContentHash{Raw: &ContentHash_Raw{
				Hash:            hash,
				DigestAlgorithm: DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256,
				MediaType:       RawMediaType_RAW_MEDIA_TYPE_XML,
			}},
		},
		{
			name: "valid raw media pdf",
			iri:  "regen:113gdjFKcVCt13Za6vN7TtbgMM6LMSjRnu89BMCxeuHdkJ1hWUmy.pdf",
			wantHash: &ContentHash{Raw: &ContentHash_Raw{
				Hash:            hash,
				DigestAlgorithm: DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256,
				MediaType:       RawMediaType_RAW_MEDIA_TYPE_PDF,
			}},
		},
		{
			name: "valid media tiff",
			iri:  "regen:113gdjFKcVCt13Za6vN7TtbgMM6LMSjRnu89BMCxeuHdkJ1hWUmy.tiff",
			wantHash: &ContentHash{Raw: &ContentHash_Raw{
				Hash:            hash,
				DigestAlgorithm: DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256,
				MediaType:       RawMediaType_RAW_MEDIA_TYPE_TIFF,
			}},
		},
		{
			name: "valid media jpg",
			iri:  "regen:113gdjFKcVCt13Za6vN7TtbgMM6LMSjRnu89BMCxeuHdkJ1hWUmy.jpg",
			wantHash: &ContentHash{Raw: &ContentHash_Raw{
				Hash:            hash,
				DigestAlgorithm: DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256,
				MediaType:       RawMediaType_RAW_MEDIA_TYPE_JPG,
			}},
		},
		{
			name: "valid media png",
			iri:  "regen:113gdjFKcVCt13Za6vN7TtbgMM6LMSjRnu89BMCxeuHdkJ1hWUmy.png",
			wantHash: &ContentHash{Raw: &ContentHash_Raw{
				Hash:            hash,
				DigestAlgorithm: DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256,
				MediaType:       RawMediaType_RAW_MEDIA_TYPE_PNG,
			}},
		},
		{
			name: "valid media svg",
			iri:  "regen:113gdjFKcVCt13Za6vN7TtbgMM6LMSjRnu89BMCxeuHdkJ1hWUmy.svg",
			wantHash: &ContentHash{Raw: &ContentHash_Raw{
				Hash:            hash,
				DigestAlgorithm: DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256,
				MediaType:       RawMediaType_RAW_MEDIA_TYPE_SVG,
			}},
		},
		{
			name: "valid media webp",
			iri:  "regen:113gdjFKcVCt13Za6vN7TtbgMM6LMSjRnu89BMCxeuHdkJ1hWUmy.webp",
			wantHash: &ContentHash{Raw: &ContentHash_Raw{
				Hash:            hash,
				DigestAlgorithm: DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256,
				MediaType:       RawMediaType_RAW_MEDIA_TYPE_WEBP,
			}},
		},
		{
			name: "valid media avif",
			iri:  "regen:113gdjFKcVCt13Za6vN7TtbgMM6LMSjRnu89BMCxeuHdkJ1hWUmy.avif",
			wantHash: &ContentHash{Raw: &ContentHash_Raw{
				Hash:            hash,
				DigestAlgorithm: DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256,
				MediaType:       RawMediaType_RAW_MEDIA_TYPE_AVIF,
			}},
		},
		{
			name: "valid media gif",
			iri:  "regen:113gdjFKcVCt13Za6vN7TtbgMM6LMSjRnu89BMCxeuHdkJ1hWUmy.gif",
			wantHash: &ContentHash{Raw: &ContentHash_Raw{
				Hash:            hash,
				DigestAlgorithm: DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256,
				MediaType:       RawMediaType_RAW_MEDIA_TYPE_GIF,
			}},
		},
		{
			name: "valid media apng",
			iri:  "regen:113gdjFKcVCt13Za6vN7TtbgMM6LMSjRnu89BMCxeuHdkJ1hWUmy.apng",
			wantHash: &ContentHash{Raw: &ContentHash_Raw{
				Hash:            hash,
				DigestAlgorithm: DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256,
				MediaType:       RawMediaType_RAW_MEDIA_TYPE_APNG,
			}},
		},
		{
			name: "valid media mpeg",
			iri:  "regen:113gdjFKcVCt13Za6vN7TtbgMM6LMSjRnu89BMCxeuHdkJ1hWUmy.mpeg",
			wantHash: &ContentHash{Raw: &ContentHash_Raw{
				Hash:            hash,
				DigestAlgorithm: DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256,
				MediaType:       RawMediaType_RAW_MEDIA_TYPE_MPEG,
			}},
		},
		{
			name: "valid media mp4",
			iri:  "regen:113gdjFKcVCt13Za6vN7TtbgMM6LMSjRnu89BMCxeuHdkJ1hWUmy.mp4",
			wantHash: &ContentHash{Raw: &ContentHash_Raw{
				Hash:            hash,
				DigestAlgorithm: DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256,
				MediaType:       RawMediaType_RAW_MEDIA_TYPE_MP4,
			}},
		},
		{
			name: "valid media webm",
			iri:  "regen:113gdjFKcVCt13Za6vN7TtbgMM6LMSjRnu89BMCxeuHdkJ1hWUmy.webm",
			wantHash: &ContentHash{Raw: &ContentHash_Raw{
				Hash:            hash,
				DigestAlgorithm: DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256,
				MediaType:       RawMediaType_RAW_MEDIA_TYPE_WEBM,
			}},
		},
		{
			name: "valid media ogg",
			iri:  "regen:113gdjFKcVCt13Za6vN7TtbgMM6LMSjRnu89BMCxeuHdkJ1hWUmy.ogg",
			wantHash: &ContentHash{Raw: &ContentHash_Raw{
				Hash:            hash,
				DigestAlgorithm: DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256,
				MediaType:       RawMediaType_RAW_MEDIA_TYPE_OGG,
			}},
		},
		{
			name: "valid graph",
			iri:  "regen:13toVgf5aZqSVSeJQv562xkkeoe3rr3bJWa29PHVKVf77VAkVMcDvVd.rdf",
			wantHash: &ContentHash{Graph: &ContentHash_Graph{
				Hash:                      hash,
				DigestAlgorithm:           DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256,
				CanonicalizationAlgorithm: GraphCanonicalizationAlgorithm_GRAPH_CANONICALIZATION_ALGORITHM_URDNA2015,
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
