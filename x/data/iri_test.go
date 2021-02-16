package data

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/cosmos/cosmos-sdk/types"
)

func TestAccAddressToDID(t *testing.T) {
	addr := types.AccAddress("test12345")

	type args struct {
		address         types.AccAddress
		bech32AccPrefix string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"1",
			args{
				address:         addr,
				bech32AccPrefix: "regen",
			},
			"did:regen:1AhJcBqAfzfi4tQ8ZDv",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AccAddressToDID(tt.args.address, tt.args.bech32AccPrefix); got != tt.want {
				t.Errorf("AccAddressToDID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContentHash_ToIRI(t *testing.T) {
	type fields struct {
		Sum isContentHash_Sum
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{
			"bad graph",
			fields{Sum: &ContentHash_Graph_{Graph: &ContentHash_Graph{}}},
			"",
			true,
		},
		{
			"bad raw",
			fields{Sum: &ContentHash_Raw_{Raw: &ContentHash_Raw{}}},
			"",
			true,
		},
		{
			"nil",
			fields{},
			"",
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ch := ContentHash{
				Sum: tt.fields.Sum,
			}
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
	hash1 := []byte("abcdefghijklmnopqrstuvwxyz123456")

	type fields struct {
		Hash            []byte
		DigestAlgorithm DigestAlgorithm
		MediaType       MediaType
	}
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
				MediaType:       MediaType_MEDIA_TYPE_PDF,
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
	for mt := range MediaType_name {
		_, err := MediaType(mt).ToExtension()
		require.NoError(t, err)
	}

	_, err := MediaType(-1).ToExtension()
	require.Error(t, err)
}
