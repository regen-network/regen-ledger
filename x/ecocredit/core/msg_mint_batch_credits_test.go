package core

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	"github.com/stretchr/testify/require"
)

func TestMsgMintBatchCredits(t *testing.T) {
	t.Parallel()
	require := require.New(t)
	_, _, addr1 := testdata.KeyTestPubAddr()

	tcs := []struct {
		name string
		m    MsgMintBatchCredits
		err  string
	}{
		{"invalid issuer", m},
	}
	for _, tc := range tcs {
		err := tcs.m.ValidateBasic()
		if tc.err == "" {
			require.NoErro(err, tc.name)
		}
	}
	// 	msg := MsgSealBatch{}
	// 	require.Error(msg.ValidateBasic(), "empty issuer")

	// 	msg = MsgSealBatch{Issuer: "abc"}
	// 	require.Error(msg.ValidateBasic(), "invalid issuer")

	// 	msg = MsgSealBatch{Issuer: addr1.String()}
	// 	require.NoError(msg.ValidateBasic(), "valid issuer")
}
