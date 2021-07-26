package ecocredit

import (
	"testing"
	"time"

	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	"github.com/stretchr/testify/require"
)

func TestMsgCreateClass(t *testing.T) {
	_, _, addr1 := testdata.KeyTestPubAddr()
	_, _, addr2 := testdata.KeyTestPubAddr()
	tests := map[string]struct {
		src    MsgCreateClass
		expErr bool
	}{
		"valid msg": {
			src: MsgCreateClass{
				Designer: addr1.String(),
				Issuers:  []string{addr1.String(), addr2.String()},
				Metadata: []byte("hello"),
			},
			expErr: false,
		},
		"valid msg without metadata": {
			src: MsgCreateClass{
				Designer: addr1.String(),
				Issuers:  []string{addr1.String(), addr2.String()},
			},
			expErr: false,
		},
		"invalid without designer": {
			src:    MsgCreateClass{},
			expErr: true,
		},
		"invalid without issuers": {
			src: MsgCreateClass{
				Designer: addr1.String(),
			},
			expErr: true,
		},
		"invalid with wrong issuers": {
			src: MsgCreateClass{
				Designer: addr1.String(),
				Issuers:  []string{"xyz", "xyz1"},
			},
			expErr: true,
		},
	}

	for msg, test := range tests {
		t.Run(msg, func(t *testing.T) {
			err := test.src.ValidateBasic()
			if test.expErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestMsgCreateBatch(t *testing.T) {
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
				ClassId:   "ID",
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
				ProjectLocation: "AB-CDE FG1 345",
				Metadata:        []byte("hello"),
			},
			expErr: false,
		},
		"valid msg with minimal fields": {
			src: MsgCreateBatch{
				Issuer:          addr1.String(),
				ClassId:         "ID",
				StartDate:       &startDate,
				EndDate:         &endDate,
				ProjectLocation: "AB-CDE FG1 345",
			},
			expErr: false,
		},
		"valid msg without Issuance.TradableAmount (assumes zero by default)": {
			src: MsgCreateBatch{
				Issuer:    addr1.String(),
				ClassId:   "ID",
				StartDate: &startDate,
				EndDate:   &endDate,
				Issuance: []*MsgCreateBatch_BatchIssuance{
					{
						Recipient:          addr2.String(),
						RetiredAmount:      "50",
						RetirementLocation: "ST-UVW XY Z12",
					},
				},
				ProjectLocation: "AB-CDE FG1 345",
			},
			expErr: false,
		},
		"invalid msg with wrong Issuance.TradableAmount": {
			src: MsgCreateBatch{
				Issuer:    addr1.String(),
				ClassId:   "ID",
				StartDate: &startDate,
				EndDate:   &endDate,
				Issuance: []*MsgCreateBatch_BatchIssuance{
					{
						Recipient:      addr2.String(),
						TradableAmount: "abc",
					},
				},
				ProjectLocation: "AB-CDE FG1 345",
			},
			expErr: true,
		},
		"valid msg without Issuance.RetiredAmount (assumes zero by default)": {
			src: MsgCreateBatch{
				Issuer:    addr1.String(),
				ClassId:   "ID",
				StartDate: &startDate,
				EndDate:   &endDate,
				Issuance: []*MsgCreateBatch_BatchIssuance{
					{
						Recipient: addr2.String(),
					},
				},
				ProjectLocation: "AB-CDE FG1 345",
			},
			expErr: false,
		},
		"invalid msg with Issuance.RetirementLocation": {
			src: MsgCreateBatch{
				Issuer:    addr1.String(),
				ClassId:   "ID",
				StartDate: &startDate,
				EndDate:   &endDate,
				Issuance: []*MsgCreateBatch_BatchIssuance{
					{
						Recipient:          addr2.String(),
						RetiredAmount:      "50",
						RetirementLocation: "wrong location",
					},
				},
				ProjectLocation: "AB-CDE FG1 345",
			},
			expErr: true,
		},
		"invalid msg with wrong Issuance.RetiredAmount": {
			src: MsgCreateBatch{
				Issuer:    addr1.String(),
				ClassId:   "ID",
				StartDate: &startDate,
				EndDate:   &endDate,
				Issuance: []*MsgCreateBatch_BatchIssuance{
					{
						Recipient:     addr2.String(),
						RetiredAmount: "abc",
					},
				},
				ProjectLocation: "AB-CDE FG1 345",
			},
			expErr: true,
		},
		"invalid msg with wrong ProjectLocation": {
			src: MsgCreateBatch{
				Issuer:          addr1.String(),
				ClassId:         "ID",
				StartDate:       &startDate,
				EndDate:         &endDate,
				ProjectLocation: "wrong location",
			},
			expErr: true,
		},
		"invalid msg without issuer": {
			src: MsgCreateBatch{
				ClassId:         "ID",
				StartDate:       &startDate,
				EndDate:         &endDate,
				ProjectLocation: "AB-CDE FG1 345",
				Metadata:        []byte("hello"),
			},
			expErr: true,
		},
		"invalid msg without class id": {
			src: MsgCreateBatch{
				Issuer:          addr1.String(),
				StartDate:       &startDate,
				EndDate:         &endDate,
				ProjectLocation: "AB-CDE FG1 345",
				Metadata:        []byte("hello"),
			},
			expErr: true,
		},
		"invalid msg without start date": {
			src: MsgCreateBatch{
				Issuer:          addr1.String(),
				ClassId:         "ID",
				EndDate:         &endDate,
				ProjectLocation: "AB-CDE FG1 345",
				Metadata:        []byte("hello"),
			},
			expErr: true,
		},
		"invalid msg without enddate": {
			src: MsgCreateBatch{
				Issuer:          addr1.String(),
				ClassId:         "ID",
				StartDate:       &startDate,
				ProjectLocation: "AB-CDE FG1 345",
				Metadata:        []byte("hello"),
			},
			expErr: true,
		},
		"invalid msg without project location": {
			src: MsgCreateBatch{
				Issuer:    addr1.String(),
				ClassId:   "ID",
				StartDate: &startDate,
				EndDate:   &endDate,
				Metadata:  []byte("hello"),
			},
			expErr: true,
		},
		"invalid msg with enddate < startdate": {
			src: MsgCreateBatch{
				Issuer:    addr1.String(),
				ClassId:   "ID",
				StartDate: &endDate,
				EndDate:   &startDate,
				Metadata:  []byte("hello"),
			},
			expErr: true,
		},
	}

	for msg, test := range tests {
		t.Run(msg, func(t *testing.T) {
			err := test.src.ValidateBasic()
			if test.expErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
