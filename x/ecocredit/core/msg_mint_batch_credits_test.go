package core

import (
	fmt "fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/thanhpk/randstr"
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
		{"invalid denom", MsgMintBatchCredits{Issuer: issuer, BatchDenom: "XXX"}, "invalid denom"},
		{"invalid note",
			MsgMintBatchCredits{Issuer: issuer, BatchDenom: batchDenom, Note: randstr.String(514)}, "note must"},
		{"missing origin tx",
			MsgMintBatchCredits{Issuer: issuer, BatchDenom: batchDenom}, "origin_tx is required"},

		{"good-no-note", MsgMintBatchCredits{Issuer: issuer, BatchDenom: batchDenom, OriginTx: &batchOrigTx, Issuance: batchIssuances}, ""},
		{"good-note", MsgMintBatchCredits{Issuer: issuer, BatchDenom: batchDenom, OriginTx: &batchOrigTx, Note: randstr.String(300),
			Issuance: batchIssuances}, ""},
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
