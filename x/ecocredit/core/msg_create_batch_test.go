package core

import (
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/regen-network/regen-ledger/types/testutil"
)

func TestMsgCreateBatch(t *testing.T) {
	t.Parallel()
	issuer := testutil.GenAddress()
	addr2 := testutil.GenAddress()

	startDate := time.Unix(10000, 10000).UTC()
	endDate := time.Unix(10000, 10050).UTC()

	tests := map[string]struct {
		src    MsgCreateBatch
		expErr bool
	}{
		"valid msg": {
			src: MsgCreateBatch{
				Issuer:    issuer,
				ProjectId: "C01",
				StartDate: &startDate,
				EndDate:   &endDate,
				Issuance: []*BatchIssuance{
					{
						Recipient:          addr2,
						TradableAmount:     "1000",
						RetiredAmount:      "50",
						RetirementLocation: "ST-UVW XY Z12",
					},
				},
				Metadata: "hello",
			},
			expErr: false,
		},
		"valid msg with minimal fields": {
			src: MsgCreateBatch{
				Issuer:    issuer,
				ProjectId: "C01",
				StartDate: &startDate,
				EndDate:   &endDate,
			},
			expErr: false,
		},
		"invalid with  wrong issuer": {
			src: MsgCreateBatch{
				Issuer:    "wrongIssuer",
				ProjectId: "C01",
				StartDate: &startDate,
				EndDate:   &endDate,
			},
			expErr: true,
		},
		"valid msg without Issuance.TradableAmount (assumes zero by default)": {
			src: MsgCreateBatch{
				Issuer:    issuer,
				ProjectId: "C01",
				StartDate: &startDate,
				EndDate:   &endDate,
				Issuance: []*BatchIssuance{
					{
						Recipient:          addr2,
						RetiredAmount:      "50",
						RetirementLocation: "ST-UVW XY Z12",
					},
				},
			},
			expErr: false,
		},
		"invalid msg with wrong Issuance.TradableAmount": {
			src: MsgCreateBatch{
				Issuer:    issuer,
				ProjectId: "C01",
				StartDate: &startDate,
				EndDate:   &endDate,
				Issuance: []*BatchIssuance{
					{
						Recipient:      addr2,
						TradableAmount: "abc",
					},
				},
			},
			expErr: true,
		},
		"valid msg without Issuance.RetiredAmount (assumes zero by default)": {
			src: MsgCreateBatch{
				Issuer:    issuer,
				ProjectId: "C01",
				StartDate: &startDate,
				EndDate:   &endDate,
				Issuance: []*BatchIssuance{
					{
						Recipient: addr2,
					},
				},
			},
			expErr: false,
		},
		"invalid msg with wrong Issuance.RetiredAmount": {
			src: MsgCreateBatch{
				Issuer:    issuer,
				ProjectId: "C01",
				StartDate: &startDate,
				EndDate:   &endDate,
				Issuance: []*BatchIssuance{
					{
						Recipient:     addr2,
						RetiredAmount: "abc",
					},
				},
			},
			expErr: true,
		},
		"invalid msg with wrong Issuance.RetirementLocation": {
			src: MsgCreateBatch{
				Issuer:    issuer,
				ProjectId: "C01",
				StartDate: &startDate,
				EndDate:   &endDate,
				Issuance: []*BatchIssuance{
					{
						Recipient:          addr2,
						RetiredAmount:      "50",
						RetirementLocation: "wrong location",
					},
				},
			},
			expErr: true,
		},
		"invalid msg without Issuance.RetirementLocation": {
			src: MsgCreateBatch{
				Issuer:    issuer,
				ProjectId: "C01",
				StartDate: &startDate,
				EndDate:   &endDate,
				Issuance: []*BatchIssuance{
					{
						Recipient:     addr2,
						RetiredAmount: "50",
					},
				},
			},
			expErr: true,
		},
		"invalid msg without issuer": {
			src: MsgCreateBatch{
				ProjectId: "C01",
				StartDate: &startDate,
				EndDate:   &endDate,
				Metadata:  "hello",
			},
			expErr: true,
		},
		"invalid msg without class id": {
			src: MsgCreateBatch{
				Issuer:    issuer,
				StartDate: &startDate,
				EndDate:   &endDate,
				Metadata:  "hello",
			},
			expErr: true,
		},
		"invalid msg without start date": {
			src: MsgCreateBatch{
				Issuer:    issuer,
				ProjectId: "C01",
				EndDate:   &endDate,
				Metadata:  "hello",
			},
			expErr: true,
		},
		"invalid msg without enddate": {
			src: MsgCreateBatch{
				Issuer:    issuer,
				ProjectId: "C01",
				StartDate: &startDate,
				Metadata:  "hello",
			},
			expErr: true,
		},
		"invalid msg with enddate < startdate": {
			src: MsgCreateBatch{
				Issuer:    issuer,
				ProjectId: "C01",
				StartDate: &endDate,
				EndDate:   &startDate,
				Metadata:  "hello",
			},
			expErr: true,
		},
		"invalid with wrong recipient": {
			src: MsgCreateBatch{
				Issuer:    issuer,
				ProjectId: "C01",
				StartDate: &startDate,
				EndDate:   &endDate,
				Issuance: []*BatchIssuance{
					{
						Recipient:          "wrongRecipient",
						RetiredAmount:      "50",
						RetirementLocation: "ST-UVW XY Z12",
					},
				},
			},
			expErr: true,
		},
		"invalid msg without recipient address": {
			src: MsgCreateBatch{
				Issuer:    issuer,
				ProjectId: "C01",
				StartDate: &startDate,
				EndDate:   &endDate,
				Issuance: []*BatchIssuance{
					{
						RetiredAmount:      "50",
						RetirementLocation: "ST-UVW XY Z12",
					},
				},
			},
			expErr: true,
		},
		"invalid metadata maxlength is exceeded": {
			src: MsgCreateBatch{
				Issuer:    issuer,
				ProjectId: "C01",
				StartDate: &startDate,
				EndDate:   &endDate,
				Metadata:  strings.Repeat("x", 288),
			},
			expErr: true,
		},
	}

	for msg, test := range tests {
		t.Run(msg, func(t *testing.T) {
			t.Parallel()

			err := test.src.ValidateBasic()
			if test.expErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
