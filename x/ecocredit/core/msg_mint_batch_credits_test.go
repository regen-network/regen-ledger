package core

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/thanhpk/randstr"

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
		{"invalid denom", "invalid denom", MsgMintBatchCredits{Issuer: issuer, BatchDenom: "XXX"}},
		{"invalid note", "note must",
			MsgMintBatchCredits{Issuer: issuer, BatchDenom: batchDenom, Note: randstr.String(514)}},
		{"missing origin tx", "origin_tx is required",
			MsgMintBatchCredits{Issuer: issuer, BatchDenom: batchDenom}},

		{"good-no-note", "",
			MsgMintBatchCredits{Issuer: issuer, BatchDenom: batchDenom, OriginTx: &batchOrigTx,
				Issuance: batchIssuances}},
		{"good-note", "",
			MsgMintBatchCredits{Issuer: issuer, BatchDenom: batchDenom, OriginTx: &batchOrigTx,
				Note: randstr.String(300), Issuance: batchIssuances}},
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

func TestValidateOriginTx(t *testing.T) {
	t.Parallel()
	require := require.New(t)
	tcs := []struct {
		name string
		err  string
		o    OriginTx
	}{
		{"empty type", "origin_tx.typ must be",
			OriginTx{}},
		{"wrong type", "origin_tx.typ must be",
			OriginTx{Typ: "*xxx"}},
		{"empty tx", "origin_tx.id must be",
			OriginTx{Typ: "Polygon"}},
		{"wrong tx", "origin_tx.id must be",
			OriginTx{Typ: "Polygon", Id: "---"}},

		{"good1", "", OriginTx{Typ: "Polygon", Id: "0x123"}},
		{"good2", "", OriginTx{Typ: "Ethereum", Id: "0x123"}},
	}

	for _, tc := range tcs {
		err := validateOriginTx(&tc.o, true)
		if tc.err == "" {
			require.NoError(err, tc.name)
		} else {
			require.ErrorContains(err, tc.err, tc.name)
		}
	}
}
