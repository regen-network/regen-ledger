package core

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/regen-network/regen-ledger/types/testutil"
)

func TestMsgMintBatchCredits(t *testing.T) {
	t.Parallel()
	require := require.New(t)
	issuer := testutil.GenAddress()

	tcs := []struct {
		name string
		err  string
		m    MsgMintBatchCredits
	}{
		{"invalid issuer", "issuer", MsgMintBatchCredits{Issuer: "invalid"}},
		{"invalid batch denom", "invalid batch denom", MsgMintBatchCredits{Issuer: issuer, BatchDenom: "XXX"}},
		{"missing origin tx", "origin tx cannot be empty",
			MsgMintBatchCredits{Issuer: issuer, BatchDenom: batchDenom}},
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
