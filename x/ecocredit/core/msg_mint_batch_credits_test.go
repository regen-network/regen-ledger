package core

import (
	"testing"

	"github.com/regen-network/regen-ledger/types/testutil"
	"github.com/stretchr/testify/require"
	"github.com/thanhpk/randstr"
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
		{"missing origin tx", "origin_tx is required",
			MsgMintBatchCredits{Issuer: issuer, BatchDenom: batchDenom}},

		{"good-no-note", "",
			MsgMintBatchCredits{Issuer: issuer, BatchDenom: batchDenom, OriginTx: &batchOriginTx,
				Issuance: batchIssuances}},
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
		{
			"empty id",
			"origin_tx.id must be",
			OriginTx{},
		},
		{
			"id must start with alphanumeric",
			"origin_tx.id must be",
			OriginTx{Id: "-foo"},
		},
		{
			"id contains invalid characters",
			"origin_tx.id must be",
			OriginTx{Id: "foo!"},
		},
		{
			"id too long",
			"origin_tx.id must be",
			OriginTx{Id: randstr.String(129)},
		},
		{
			"empty source",
			"origin_tx.source must be",
			OriginTx{Id: "0x123"},
		},
		{
			"source must start with alphanumeric",
			"origin_tx.source must be",
			OriginTx{
				Id:     "0x123",
				Source: "-foo",
			},
		},
		{
			"source contains invalid characters",
			"origin_tx.source must be",
			OriginTx{
				Id:     "0x123",
				Source: "foo!",
			},
		},
		{
			"source too long",
			"origin_tx.source must be",
			OriginTx{
				Id:     "0x123",
				Source: randstr.String(33)},
		},
		{
			"invalid contract",
			"origin_tx.contract must",
			OriginTx{
				Source:   "polygon",
				Id:       "0x123",
				Contract: "foo",
			},
		},
		{
			"note too long",
			"origin_tx.note must",
			OriginTx{
				Source: "polygon",
				Id:     "0x123",
				Note:   randstr.String(513),
			},
		},
		{
			"valid polygon",
			"",
			OriginTx{
				Source: "polygon",
				Id:     "0x7a70692a348e8688f54ab2bdfe87d925d8cc88932520492a11eaa02dc128243e",
			},
		},
		{
			"valid verra",
			"",
			OriginTx{
				Source: "verra",
				Id:     "0001-000001-000100-VCS-VCU-003-VER-US-0003-01012020-31122020-1",
			},
		},
		{
			"valid with contract",
			"",
			OriginTx{
				Source:   "polygon",
				Id:       "0x7a70692a348e8688f54ab2bdfe87d925d8cc88932520492a11eaa02dc128243e",
				Contract: "0x0e65079a29d7793ab5ca500c2d88e60ee99ba606",
			},
		},
		{
			"valid with note",
			"",
			OriginTx{
				Source: "polygon",
				Id:     "0x7a70692a348e8688f54ab2bdfe87d925d8cc88932520492a11eaa02dc128243e",
				Note:   randstr.String(512),
			},
		},
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
