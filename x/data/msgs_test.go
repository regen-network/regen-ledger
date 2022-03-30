package data

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/cosmos/cosmos-sdk/testutil/testdata"
)

func TestMsgAnchorRequest_ValidateBasic(t *testing.T) {
	_, _, addr := testdata.KeyTestPubAddr()
	type fields struct {
		Sender string
		Hash   *ContentHash
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr string
	}{
		{
			name: "good",
			fields: fields{
				Sender: addr.String(),
				Hash: &ContentHash{Raw: &ContentHash_Raw{
					Hash:            make([]byte, 32),
					DigestAlgorithm: DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256,
					MediaType:       RawMediaType_RAW_MEDIA_TYPE_UNSPECIFIED,
				}},
			},
			wantErr: "",
		},
		{
			name: "nil",
			fields: fields{
				Sender: addr.String(),
				Hash:   nil,
			},
			wantErr: "hash cannot be empty: invalid request",
		},
		{
			name: "bad",
			fields: fields{
				Sender: addr.String(),
				Hash: &ContentHash{Raw: &ContentHash_Raw{
					Hash:            make([]byte, 31),
					DigestAlgorithm: DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256,
					MediaType:       RawMediaType_RAW_MEDIA_TYPE_UNSPECIFIED,
				}},
			},
			wantErr: "expected 32 bytes for DIGEST_ALGORITHM_BLAKE2B_256, got 31: invalid request",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MsgAnchor{
				Sender: tt.fields.Sender,
				Hash:   tt.fields.Hash,
			}
			err := m.ValidateBasic()
			if len(tt.wantErr) != 0 {
				require.EqualError(t, err, tt.wantErr)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestMsgAttestRequest_ValidateBasic(t *testing.T) {
	_, _, addr := testdata.KeyTestPubAddr()
	type fields struct {
		Attestors []string
		Hash      *ContentHash_Graph
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr string
	}{
		{
			"good",
			fields{
				Attestors: []string{addr.String()},
				Hash: &ContentHash_Graph{
					Hash:                      make([]byte, 32),
					DigestAlgorithm:           DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256,
					CanonicalizationAlgorithm: GraphCanonicalizationAlgorithm_GRAPH_CANONICALIZATION_ALGORITHM_URDNA2015,
					MerkleTree:                GraphMerkleTree_GRAPH_MERKLE_TREE_NONE_UNSPECIFIED,
				},
			},
			"",
		},
		{
			"bad",
			fields{
				Attestors: nil,
				Hash: &ContentHash_Graph{
					Hash:                      make([]byte, 32),
					DigestAlgorithm:           DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256,
					CanonicalizationAlgorithm: GraphCanonicalizationAlgorithm_GRAPH_CANONICALIZATION_ALGORITHM_UNSPECIFIED,
					MerkleTree:                GraphMerkleTree_GRAPH_MERKLE_TREE_NONE_UNSPECIFIED,
				},
			},
			"invalid data.GraphCanonicalizationAlgorithm GRAPH_CANONICALIZATION_ALGORITHM_UNSPECIFIED: invalid request",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MsgAttest{
				Attestors: tt.fields.Attestors,
				Hash:      tt.fields.Hash,
			}
			err := m.ValidateBasic()
			if len(tt.wantErr) != 0 {
				require.EqualError(t, err, tt.wantErr)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestMsgDefineResolver_ValidateBasic(t *testing.T) {
	_, _, addr := testdata.KeyTestPubAddr()
	type fields struct {
		Manager     string
		ResolverUrl string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr string
	}{
		{
			"valid message",
			fields{
				Manager:     addr.String(),
				ResolverUrl: "https://foo.bar",
			},
			"",
		},
		{
			"invalid manager",
			fields{
				Manager:     "foo",
				ResolverUrl: "https://foo.bar",
			},
			"decoding bech32 failed: invalid bech32 string length 3: invalid address",
		},
		{
			"invalid resolver url",
			fields{
				Manager:     addr.String(),
				ResolverUrl: "foobar",
			},
			"invalid resolver url: invalid request",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MsgDefineResolver{
				Manager:     tt.fields.Manager,
				ResolverUrl: tt.fields.ResolverUrl,
			}
			err := m.ValidateBasic()
			if len(tt.wantErr) != 0 {
				require.EqualError(t, err, tt.wantErr)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestMsgRegisterResolver_ValidateBasic(t *testing.T) {
	_, _, addr := testdata.KeyTestPubAddr()
	validData := []*ContentHash{
		{
			Raw: &ContentHash_Raw{
				Hash:            make([]byte, 32),
				DigestAlgorithm: DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256,
				MediaType:       RawMediaType_RAW_MEDIA_TYPE_UNSPECIFIED,
			},
		},
	}

	type fields struct {
		Manager string
		Data    []*ContentHash
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr string
	}{
		{
			"valid message",
			fields{
				Manager: addr.String(),
				Data:    validData,
			},
			"",
		},
		{
			"invalid manager",
			fields{
				Manager: "foo",
				Data:    validData,
			},
			"decoding bech32 failed: invalid bech32 string length 3: invalid address",
		},
		{
			"data cannot be empty",
			fields{
				Manager: addr.String(),
				Data:    []*ContentHash{},
			},
			"data cannot be empty: invalid request",
		},
		{
			"invalid content hash",
			fields{
				Manager: addr.String(),
				Data: []*ContentHash{
					{
						Raw: &ContentHash_Raw{
							Hash:            make([]byte, 31),
							DigestAlgorithm: DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256,
							MediaType:       RawMediaType_RAW_MEDIA_TYPE_UNSPECIFIED,
						},
					},
				},
			},
			"expected 32 bytes for DIGEST_ALGORITHM_BLAKE2B_256, got 31: invalid request",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MsgRegisterResolver{
				Manager: tt.fields.Manager,
				Data:    tt.fields.Data,
			}
			err := m.ValidateBasic()
			if len(tt.wantErr) != 0 {
				require.EqualError(t, err, tt.wantErr)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
