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
				Admin:      addr1.String(),
				Issuers:    []string{addr1.String(), addr2.String()},
				CreditType: "carbon",
				Metadata:   []byte("hello"),
			},
			expErr: false,
		},
		"valid msg without metadata": {
			src: MsgCreateClass{
				Admin:      addr1.String(),
				CreditType: "carbon",
				Issuers:    []string{addr1.String(), addr2.String()},
			},
			expErr: false,
		},
		"invalid without admin": {
			src:    MsgCreateClass{},
			expErr: true,
		},
		"invalid without issuers": {
			src: MsgCreateClass{
				Admin:      addr1.String(),
				CreditType: "carbon",
			},
			expErr: true,
		},
		"invalid with wrong issuers": {
			src: MsgCreateClass{
				Admin:      addr1.String(),
				CreditType: "carbon",
				Issuers:    []string{"xyz", "xyz1"},
			},
			expErr: true,
		},
		"invalid with wrong admin": {
			src: MsgCreateClass{
				Admin:      "wrongAdmin",
				CreditType: "carbon",
				Issuers:    []string{addr1.String(), addr2.String()},
			},
			expErr: true,
		},
		"invalid with no credit type": {
			src: MsgCreateClass{
				Admin:   addr1.String(),
				Issuers: []string{addr1.String(), addr2.String()},
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
				ClassId:   "C01",
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
				ClassId:         "C01",
				StartDate:       &startDate,
				EndDate:         &endDate,
				ProjectLocation: "AB-CDE FG1 345",
			},
			expErr: false,
		},
		"invalid with  wrong issuer": {
			src: MsgCreateBatch{
				Issuer:          "wrongIssuer",
				ClassId:         "ID",
				StartDate:       &startDate,
				EndDate:         &endDate,
				ProjectLocation: "AB-CDE FG1 345",
			},
			expErr: true,
		},
		"valid msg without Issuance.TradableAmount (assumes zero by default)": {
			src: MsgCreateBatch{
				Issuer:    addr1.String(),
				ClassId:   "C01",
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
				ClassId:   "C01",
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
		"invalid msg with wrong Issuance.RetirementLocation": {
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
		"invalid msg without Issuance.RetirementLocation": {
			src: MsgCreateBatch{
				Issuer:    addr1.String(),
				ClassId:   "ID",
				StartDate: &startDate,
				EndDate:   &endDate,
				Issuance: []*MsgCreateBatch_BatchIssuance{
					{
						Recipient:     addr2.String(),
						RetiredAmount: "50",
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
		"invalid with wrong recipient": {
			src: MsgCreateBatch{
				Issuer:    addr1.String(),
				ClassId:   "ID",
				StartDate: &startDate,
				EndDate:   &endDate,
				Issuance: []*MsgCreateBatch_BatchIssuance{
					{
						Recipient:          "wrongRecipient",
						RetiredAmount:      "50",
						RetirementLocation: "ST-UVW XY Z12",
					},
				},
				ProjectLocation: "AB-CDE FG1 345",
			},
			expErr: true,
		},
		"invalid msg without recipient address": {
			src: MsgCreateBatch{
				Issuer:    addr1.String(),
				ClassId:   "ID",
				StartDate: &startDate,
				EndDate:   &endDate,
				Issuance: []*MsgCreateBatch_BatchIssuance{
					{
						RetiredAmount:      "50",
						RetirementLocation: "ST-UVW XY Z12",
					},
				},
				ProjectLocation: "AB-CDE FG1 345",
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

func TestMsgSend(t *testing.T) {
	_, _, addr1 := testdata.KeyTestPubAddr()
	_, _, addr2 := testdata.KeyTestPubAddr()

	tests := map[string]struct {
		src    MsgSend
		expErr bool
	}{
		"valid msg": {
			src: MsgSend{
				Sender:    addr1.String(),
				Recipient: addr2.String(),
				Credits: []*MsgSend_SendCredits{
					{
						BatchDenom:         "some_denom",
						TradableAmount:     "10",
						RetiredAmount:      "10",
						RetirementLocation: "ST-UVW XY Z12",
					},
				},
			},
			expErr: false,
		},
		"invalid msg with Credits.RetiredAmount negative value": {
			src: MsgSend{
				Sender:    addr1.String(),
				Recipient: addr2.String(),
				Credits: []*MsgSend_SendCredits{
					{
						BatchDenom:     "some_denom",
						TradableAmount: "10",
						RetiredAmount:  "-10",
					},
				},
			},
			expErr: true,
		},
		"invalid msg without credits": {
			src: MsgSend{
				Sender:    addr1.String(),
				Recipient: addr2.String(),
			},
			expErr: true,
		},
		"invalid msg without sender": {
			src: MsgSend{
				Recipient: addr2.String(),
				Credits: []*MsgSend_SendCredits{
					{
						BatchDenom:         "some_denom",
						TradableAmount:     "10",
						RetiredAmount:      "10",
						RetirementLocation: "ST-UVW XY Z12",
					},
				},
			},
			expErr: true,
		},
		"invalid msg without recipient": {
			src: MsgSend{
				Sender: addr1.String(),
				Credits: []*MsgSend_SendCredits{
					{
						BatchDenom:         "some_denom",
						TradableAmount:     "10",
						RetiredAmount:      "10",
						RetirementLocation: "ST-UVW XY Z12",
					},
				},
			},
			expErr: true,
		},
		"invalid msg without Credits.BatchDenom": {
			src: MsgSend{
				Sender:    addr1.String(),
				Recipient: addr2.String(),
				Credits: []*MsgSend_SendCredits{
					{
						TradableAmount:     "10",
						RetiredAmount:      "10",
						RetirementLocation: "ST-UVW XY Z12",
					},
				},
			},
			expErr: true,
		},
		"invalid msg without Credits.TradableAmount set": {
			src: MsgSend{
				Sender:    addr1.String(),
				Recipient: addr2.String(),
				Credits: []*MsgSend_SendCredits{
					{
						BatchDenom:         "some_denom",
						RetiredAmount:      "10",
						RetirementLocation: "ST-UVW XY Z12",
					},
				},
			},
			expErr: true,
		},
		"invalid msg without Credits.RetiredAmount set": {
			src: MsgSend{
				Sender:    addr1.String(),
				Recipient: addr2.String(),
				Credits: []*MsgSend_SendCredits{
					{
						BatchDenom:         "some_denom",
						TradableAmount:     "10",
						RetirementLocation: "ST-UVW XY Z12",
					},
				},
			},
			expErr: true,
		},
		"invalid msg without Credits.RetirementLocation": {
			src: MsgSend{
				Sender:    addr1.String(),
				Recipient: addr2.String(),
				Credits: []*MsgSend_SendCredits{
					{
						BatchDenom:     "some_denom",
						TradableAmount: "10",
						RetiredAmount:  "10",
					},
				},
			},
			expErr: true,
		},
		"valid msg without Credits.RetirementLocation(When RetiredAmount is zero)": {
			src: MsgSend{
				Sender:    addr1.String(),
				Recipient: addr2.String(),
				Credits: []*MsgSend_SendCredits{
					{
						BatchDenom:     "some_denom",
						TradableAmount: "10",
						RetiredAmount:  "0",
					},
				},
			},
			expErr: false,
		},
		"invalid msg with wrong sender": {
			src: MsgSend{
				Sender:    "wrongSender",
				Recipient: addr2.String(),
				Credits: []*MsgSend_SendCredits{
					{
						BatchDenom:     "some_denom",
						TradableAmount: "10",
						RetiredAmount:  "10",
					},
				},
			},
			expErr: true,
		},
		"invalid msg with wrong recipient": {
			src: MsgSend{
				Sender:    addr1.String(),
				Recipient: "wrongRecipient",
				Credits: []*MsgSend_SendCredits{
					{
						BatchDenom:     "some_denom",
						TradableAmount: "10",
						RetiredAmount:  "10",
					},
				},
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

func TestMsgRetire(t *testing.T) {
	_, _, addr1 := testdata.KeyTestPubAddr()

	tests := map[string]struct {
		src    MsgRetire
		expErr bool
	}{
		"valid msg": {
			src: MsgRetire{
				Holder: addr1.String(),
				Credits: []*MsgRetire_RetireCredits{
					{
						BatchDenom: "some_denom",
						Amount:     "10",
					},
				},
				Location: "AB-CDE FG1 345",
			},
			expErr: false,
		},
		"invalid msg without holder": {
			src: MsgRetire{
				Credits: []*MsgRetire_RetireCredits{
					{
						BatchDenom: "some_denom",
						Amount:     "10",
					},
				},
				Location: "AB-CDE FG1 345",
			},
			expErr: true,
		},
		"invalid msg with wrong holder address": {
			src: MsgRetire{
				Holder: "wrongHolder",
				Credits: []*MsgRetire_RetireCredits{
					{
						BatchDenom: "some_denom",
						Amount:     "10",
					},
				},
				Location: "AB-CDE FG1 345",
			},
			expErr: true,
		},
		"invalid msg without credits": {
			src: MsgRetire{
				Holder:   addr1.String(),
				Location: "AB-CDE FG1 345",
			},
			expErr: true,
		},
		"invalid msg without Credits.BatchDenom": {
			src: MsgRetire{
				Holder: addr1.String(),
				Credits: []*MsgRetire_RetireCredits{
					{
						Amount: "10",
					},
				},
				Location: "AB-CDE FG1 345",
			},
			expErr: true,
		},
		"invalid msg without Credits.Amount": {
			src: MsgRetire{
				Holder: addr1.String(),
				Credits: []*MsgRetire_RetireCredits{
					{
						BatchDenom: "some_denom",
					},
				},
				Location: "AB-CDE FG1 345",
			},
			expErr: true,
		},
		"invalid msg with wrong Credits.Amount": {
			src: MsgRetire{
				Holder: addr1.String(),
				Credits: []*MsgRetire_RetireCredits{
					{
						BatchDenom: "some_denom",
						Amount:     "abc",
					},
				},
				Location: "AB-CDE FG1 345",
			},
			expErr: true,
		},
		"invalid msg without location": {
			src: MsgRetire{
				Holder: addr1.String(),
				Credits: []*MsgRetire_RetireCredits{
					{
						BatchDenom: "some_denom",
						Amount:     "10",
					},
				},
			},
			expErr: true,
		},
		"invalid msg with wrong location": {
			src: MsgRetire{
				Holder: addr1.String(),
				Credits: []*MsgRetire_RetireCredits{
					{
						BatchDenom: "some_denom",
						Amount:     "10",
					},
				},
				Location: "wrongLocation",
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

func TestMsgCancel(t *testing.T) {
	_, _, addr1 := testdata.KeyTestPubAddr()

	tests := map[string]struct {
		src    MsgCancel
		expErr bool
	}{
		"valid msg": {
			src: MsgCancel{
				Holder: addr1.String(),
				Credits: []*MsgCancel_CancelCredits{
					{
						BatchDenom: "some_denom",
						Amount:     "10",
					},
				},
			},
			expErr: false,
		},
		"invalid msg without holder": {
			src: MsgCancel{
				Credits: []*MsgCancel_CancelCredits{
					{
						BatchDenom: "some_denom",
						Amount:     "10",
					},
				},
			},
			expErr: true,
		},
		"invalid msg with wrong holder address": {
			src: MsgCancel{
				Holder: "wrongHolder",
				Credits: []*MsgCancel_CancelCredits{
					{
						BatchDenom: "some_denom",
						Amount:     "10",
					},
				},
			},
			expErr: true,
		},
		"invalid msg without credits": {
			src: MsgCancel{
				Holder: addr1.String(),
			},
			expErr: true,
		},
		"invalid msg without Credits.BatchDenom": {
			src: MsgCancel{
				Holder: addr1.String(),
				Credits: []*MsgCancel_CancelCredits{
					{
						Amount: "10",
					},
				},
			},
			expErr: true,
		},
		"invalid msg without Credits.Amount": {
			src: MsgCancel{
				Holder: addr1.String(),
				Credits: []*MsgCancel_CancelCredits{
					{
						BatchDenom: "some_denom",
					},
				},
			},
			expErr: true,
		},
		"invalid msg with wrong Credits.Amount": {
			src: MsgCancel{
				Holder: addr1.String(),
				Credits: []*MsgCancel_CancelCredits{
					{
						BatchDenom: "some_denom",
						Amount:     "abc",
					},
				},
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
