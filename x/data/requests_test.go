package data

import (
	"crypto"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/cosmos/cosmos-sdk/testutil/testdata"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func TestMsgAnchorDataRequest_GetSigners(t *testing.T) {
	_, _, addr := testdata.KeyTestPubAddr()

	msg := &MsgAnchorDataRequest{Sender: addr.String()}
	require.Equal(t, []sdk.AccAddress{addr}, msg.GetSigners())

	msg = &MsgAnchorDataRequest{Sender: ""}
	require.Panics(t, func() {
		msg.GetSigners()
	})
}

func TestMsgAnchorDataRequest_ValidateBasic(t *testing.T) {
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
				Sender: "",
				Hash: &ContentHash{Sum: &ContentHash_Raw_{Raw: &ContentHash_Raw{
					Hash:            make([]byte, 32),
					DigestAlgorithm: DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256,
					MediaType:       MediaType_MEDIA_TYPE_UNSPECIFIED,
				},
				}},
			},
			wantErr: "",
		},
		{
			name: "bad",
			fields: fields{
				Sender: "",
				Hash: &ContentHash{Sum: &ContentHash_Raw_{Raw: &ContentHash_Raw{
					Hash:            make([]byte, 31),
					DigestAlgorithm: DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256,
					MediaType:       MediaType_MEDIA_TYPE_UNSPECIFIED,
				},
				}},
			},
			wantErr: "expected 32 bytes for DIGEST_ALGORITHM_BLAKE2B_256, got 31: unknown request",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MsgAnchorDataRequest{
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

func TestMsgSignDataRequest_GetSigners(t *testing.T) {
	_, _, addr := testdata.KeyTestPubAddr()
	_, _, addr2 := testdata.KeyTestPubAddr()

	msg := &MsgSignDataRequest{Signers: []string{addr.String(), addr2.String()}}
	require.Equal(t, []sdk.AccAddress{addr, addr2}, msg.GetSigners())

	msg = &MsgSignDataRequest{Signers: nil}
	require.Empty(t, msg.GetSigners())

	msg = &MsgSignDataRequest{Signers: []string{"abcd"}}
	require.Panics(t, func() {
		msg.GetSigners()
	})
}

func TestMsgSignDataRequest_ValidateBasic(t *testing.T) {
	type fields struct {
		Signers []string
		Hash    *ContentHash_Graph
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr string
	}{
		{
			"good",
			fields{
				Signers: nil,
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
				Signers: nil,
				Hash: &ContentHash_Graph{
					Hash:                      make([]byte, 32),
					DigestAlgorithm:           DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256,
					CanonicalizationAlgorithm: GraphCanonicalizationAlgorithm_GRAPH_CANONICALIZATION_ALGORITHM_UNSPECIFIED,
					MerkleTree:                GraphMerkleTree_GRAPH_MERKLE_TREE_NONE_UNSPECIFIED,
				},
			},
			"invalid data.GraphCanonicalizationAlgorithm GRAPH_CANONICALIZATION_ALGORITHM_UNSPECIFIED: unknown request",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MsgSignDataRequest{
				Signers: tt.fields.Signers,
				Hash:    tt.fields.Hash,
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

func TestMsgStoreRawDataRequest_GetSigners(t *testing.T) {
	_, _, addr := testdata.KeyTestPubAddr()

	msg := &MsgStoreRawDataRequest{Sender: addr.String()}
	require.Equal(t, []sdk.AccAddress{addr}, msg.GetSigners())

	msg = &MsgStoreRawDataRequest{Sender: ""}
	require.Panics(t, func() {
		msg.GetSigners()
	})
}

func TestMsgStoreRawDataRequest_ValidateBasic(t *testing.T) {
	data := []byte("sdf,gh8934tfgno2t09sdghk13y89w87ybdufgbh208phsnbdouguy209367wnb0")
	hash := crypto.BLAKE2b_256.New()
	_, err := hash.Write(data)
	require.NoError(t, err)
	digest := hash.Sum(nil)

	type fields struct {
		Sender  string
		Hash    *ContentHash_Raw
		Content []byte
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr string
	}{
		{
			"good",
			fields{
				Sender: "",
				Hash: &ContentHash_Raw{
					Hash:            digest,
					DigestAlgorithm: DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256,
					MediaType:       MediaType_MEDIA_TYPE_UNSPECIFIED,
				},
				Content: data,
			},
			"",
		},
		{
			"bad",
			fields{
				Sender: "",
				Hash: &ContentHash_Raw{
					Hash:            make([]byte, 32),
					DigestAlgorithm: DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256,
					MediaType:       MediaType_MEDIA_TYPE_UNSPECIFIED,
				},
				Content: data,
			},
			"hash verification failed",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MsgStoreRawDataRequest{
				Sender:  tt.fields.Sender,
				Hash:    tt.fields.Hash,
				Content: tt.fields.Content,
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
