package core

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func TestMsgCreateBatch(t *testing.T) {
	t.Parallel()

	_, _, addr1 := testdata.KeyTestPubAddr()
	_, _, addr2 := testdata.KeyTestPubAddr()

	startDate := time.Unix(10000, 10000).UTC()
	endDate := time.Unix(10000, 10050).UTC()

	tests := map[string]struct {
		src    MsgCreateBatch
		expErr bool
	}{
		"valid msg": {
			src: MsgCreateBatch{
				Issuer:    addr1.String(),
				ProjectId: "C01",
				StartDate: &startDate,
				EndDate:   &endDate,
				Issuance: []*MsgCreateBatch_BatchIssuance{
					{
						Recipient:          addr2.String(),
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
				Issuer:    addr1.String(),
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
				Issuer:    addr1.String(),
				ProjectId: "C01",
				StartDate: &startDate,
				EndDate:   &endDate,
				Issuance: []*MsgCreateBatch_BatchIssuance{
					{
						Recipient:          addr2.String(),
						RetiredAmount:      "50",
						RetirementLocation: "ST-UVW XY Z12",
					},
				},
			},
			expErr: false,
		},
		"invalid msg with wrong Issuance.TradableAmount": {
			src: MsgCreateBatch{
				Issuer:    addr1.String(),
				ProjectId: "C01",
				StartDate: &startDate,
				EndDate:   &endDate,
				Issuance: []*MsgCreateBatch_BatchIssuance{
					{
						Recipient:      addr2.String(),
						TradableAmount: "abc",
					},
				},
			},
			expErr: true,
		},
		"valid msg without Issuance.RetiredAmount (assumes zero by default)": {
			src: MsgCreateBatch{
				Issuer:    addr1.String(),
				ProjectId: "C01",
				StartDate: &startDate,
				EndDate:   &endDate,
				Issuance: []*MsgCreateBatch_BatchIssuance{
					{
						Recipient: addr2.String(),
					},
				},
			},
			expErr: false,
		},
		"invalid msg with wrong Issuance.RetiredAmount": {
			src: MsgCreateBatch{
				Issuer:    addr1.String(),
				ProjectId: "C01",
				StartDate: &startDate,
				EndDate:   &endDate,
				Issuance: []*MsgCreateBatch_BatchIssuance{
					{
						Recipient:     addr2.String(),
						RetiredAmount: "abc",
					},
				},
			},
			expErr: true,
		},
		"invalid msg with wrong Issuance.RetirementLocation": {
			src: MsgCreateBatch{
				Issuer:    addr1.String(),
				ProjectId: "C01",
				StartDate: &startDate,
				EndDate:   &endDate,
				Issuance: []*MsgCreateBatch_BatchIssuance{
					{
						Recipient:          addr2.String(),
						RetiredAmount:      "50",
						RetirementLocation: "wrong location",
					},
				},
			},
			expErr: true,
		},
		"invalid msg without Issuance.RetirementLocation": {
			src: MsgCreateBatch{
				Issuer:    addr1.String(),
				ProjectId: "C01",
				StartDate: &startDate,
				EndDate:   &endDate,
				Issuance: []*MsgCreateBatch_BatchIssuance{
					{
						Recipient:     addr2.String(),
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
				Issuer:    addr1.String(),
				StartDate: &startDate,
				EndDate:   &endDate,
				Metadata:  "hello",
			},
			expErr: true,
		},
		"invalid msg without start date": {
			src: MsgCreateBatch{
				Issuer:    addr1.String(),
				ProjectId: "C01",
				EndDate:   &endDate,
				Metadata:  "hello",
			},
			expErr: true,
		},
		"invalid msg without enddate": {
			src: MsgCreateBatch{
				Issuer:    addr1.String(),
				ProjectId: "C01",
				StartDate: &startDate,
				Metadata:  "hello",
			},
			expErr: true,
		},
		"invalid msg with enddate < startdate": {
			src: MsgCreateBatch{
				Issuer:    addr1.String(),
				ProjectId: "C01",
				StartDate: &endDate,
				EndDate:   &startDate,
				Metadata:  "hello",
			},
			expErr: true,
		},
		"invalid with wrong recipient": {
			src: MsgCreateBatch{
				Issuer:    addr1.String(),
				ProjectId: "C01",
				StartDate: &startDate,
				EndDate:   &endDate,
				Issuance: []*MsgCreateBatch_BatchIssuance{
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
				Issuer:    addr1.String(),
				ProjectId: "C01",
				StartDate: &startDate,
				EndDate:   &endDate,
				Issuance: []*MsgCreateBatch_BatchIssuance{
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
				Issuer:    addr1.String(),
				ProjectId: "C01",
				StartDate: &startDate,
				EndDate:   &endDate,
				Metadata:  simtypes.RandStringOfLength(r, 288),
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
