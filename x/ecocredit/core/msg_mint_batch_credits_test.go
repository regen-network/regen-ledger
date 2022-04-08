package core

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	"github.com/stretchr/testify/require"
)

func TestMsgMintBatchCredits(t *testing.T) {
	t.Parallel()
	require := require.New(t)
	issuer := genAddress()

	tcs := []struct {
		name string
		m    MsgMintBatchCredits
		err  string
	}{
		{"invalid issuer", MsgMintBatchCredits{Issuer: "invalid"}, "issuer"},
		{"invalid denom", MsgMintBatchCredits{Issuer: issuer, BatchDenom: "XXX"}, ""},

		{"good", MsgMintBatchCredits{Issuer: issuer}, ""},
	}
	for _, tc := range tcs {
		err := tc.m.ValidateBasic()
		if tc.err == "" {
			require.NoError(err, tc.name)
		} else {
			require.ErrorContains(err, tc.err, tc.name)
		}
	}
}
