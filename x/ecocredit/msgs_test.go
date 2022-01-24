package ecocredit

import (
	"math/rand"
	"testing"
	"time"

	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/stretchr/testify/require"
)

var (
	s = rand.NewSource(1)
	r = rand.New(s)
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
				Admin:          addr1.String(),
				Issuers:        []string{addr1.String(), addr2.String()},
				CreditTypeName: "carbon",
				Metadata:       []byte("hello"),
			},
			expErr: false,
		},
		"valid msg without metadata": {
			src: MsgCreateClass{
				Admin:          addr1.String(),
				CreditTypeName: "carbon",
				Issuers:        []string{addr1.String(), addr2.String()},
			},
			expErr: false,
		},
		"invalid without admin": {
			src:    MsgCreateClass{},
			expErr: true,
		},
		"invalid without issuers": {
			src: MsgCreateClass{
				Admin:          addr1.String(),
				CreditTypeName: "carbon",
			},
			expErr: true,
		},
		"invalid with wrong issuers": {
			src: MsgCreateClass{
				Admin:          addr1.String(),
				CreditTypeName: "carbon",
				Issuers:        []string{"xyz", "xyz1"},
			},
			expErr: true,
		},
		"invalid with wrong admin": {
			src: MsgCreateClass{
				Admin:          "wrongAdmin",
				CreditTypeName: "carbon",
				Issuers:        []string{addr1.String(), addr2.String()},
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
		"invalid metadata maxlength is exceeded": {
			src: MsgCreateClass{
				Admin:          addr1.String(),
				CreditTypeName: "carbon",
				Issuers:        []string{addr1.String(), addr2.String()},
				Metadata:       []byte(simtypes.RandStringOfLength(r, 288)),
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

func TestMsgCreateProject(t *testing.T) {
	_, _, addr1 := testdata.KeyTestPubAddr()

	testCases := []struct {
		name   string
		src    MsgCreateProject
		expErr bool
	}{
		{
			"valid msg with project id",
			MsgCreateProject{
				Issuer:          addr1.String(),
				ClassId:         "A00",
				Metadata:        []byte("hello"),
				ProjectLocation: "AB-CDE FG1 345",
				ProjectId:       "A0",
			},
			false,
		},
		{
			"valid msg without project id",
			MsgCreateProject{
				Issuer:          addr1.String(),
				ClassId:         "A00",
				Metadata:        []byte("hello"),
				ProjectLocation: "AB-CDE FG1 345",
			},
			false,
		},
		{
			"invalid issuer",
			MsgCreateProject{
				Issuer:          "invalid address",
				ClassId:         "A00",
				Metadata:        []byte("hello"),
				ProjectLocation: "AB-CDE FG1 345",
				ProjectId:       "A0",
			},
			true,
		},
		{
			"invalid project id",
			MsgCreateProject{
				Issuer:          addr1.String(),
				ClassId:         "A00",
				Metadata:        []byte("hello"),
				ProjectLocation: "AB-CDE FG1 345",
				ProjectId:       "A",
			},
			true,
		},
		{
			"invalid class id",
			MsgCreateProject{
				Issuer:          addr1.String(),
				ClassId:         "ABCD",
				Metadata:        []byte("hello"),
				ProjectLocation: "AB-CDE FG1 345",
				ProjectId:       "AB",
			},
			true,
		},
		{
			"invalid project location",
			MsgCreateProject{
				Issuer:          addr1.String(),
				ClassId:         "A01",
				Metadata:        []byte("hello"),
				ProjectLocation: "abcd",
				ProjectId:       "AB",
			},
			true,
		},
		{
			"invalid: metadata is too large",
			MsgCreateProject{
				Issuer:          addr1.String(),
				ClassId:         "A01",
				Metadata:        []byte(simtypes.RandStringOfLength(r, 288)),
				ProjectLocation: "AB-CDE FG1 345",
				ProjectId:       "AB",
			},
			true,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
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
				Metadata: []byte("hello"),
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
				Metadata:  []byte("hello"),
			},
			expErr: true,
		},
		"invalid msg without class id": {
			src: MsgCreateBatch{
				Issuer:    addr1.String(),
				StartDate: &startDate,
				EndDate:   &endDate,
				Metadata:  []byte("hello"),
			},
			expErr: true,
		},
		"invalid msg without start date": {
			src: MsgCreateBatch{
				Issuer:    addr1.String(),
				ProjectId: "C01",
				EndDate:   &endDate,
				Metadata:  []byte("hello"),
			},
			expErr: true,
		},
		"invalid msg without enddate": {
			src: MsgCreateBatch{
				Issuer:    addr1.String(),
				ProjectId: "C01",
				StartDate: &startDate,
				Metadata:  []byte("hello"),
			},
			expErr: true,
		},
		"invalid msg with enddate < startdate": {
			src: MsgCreateBatch{
				Issuer:    addr1.String(),
				ProjectId: "C01",
				StartDate: &endDate,
				EndDate:   &startDate,
				Metadata:  []byte("hello"),
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
				Metadata:  []byte(simtypes.RandStringOfLength(r, 288)),
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
						BatchDenom:         "A00-00000000-00000000-000",
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
						BatchDenom:     "A00-00000000-00000000-000",
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
						BatchDenom: "A00-00000000-00000000-000",
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
						BatchDenom: "A00-00000000-00000000-000",
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
						BatchDenom: "A00-00000000-00000000-000",
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
						BatchDenom: "A00-00000000-00000000-000",
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
						BatchDenom: "A00-00000000-00000000-000",
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
						BatchDenom: "A00-00000000-00000000-000",
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
						BatchDenom: "A00-00000000-00000000-000",
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
						BatchDenom: "A00-00000000-00000000-000",
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
						BatchDenom: "A00-00000000-00000000-000",
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
						BatchDenom: "A00-00000000-00000000-000",
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
						BatchDenom: "A00-00000000-00000000-000",
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
						BatchDenom: "A00-00000000-00000000-000",
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

func TestMsgUpdateClassAdmin(t *testing.T) {
	_, _, admin := testdata.KeyTestPubAddr()
	_, _, newAdmin := testdata.KeyTestPubAddr()

	tests := map[string]struct {
		src    MsgUpdateClassAdmin
		expErr bool
	}{
		"valid": {
			src:    MsgUpdateClassAdmin{Admin: admin.String(), NewAdmin: newAdmin.String(), ClassId: "C01"},
			expErr: false,
		},
		"invalid: same address": {
			src:    MsgUpdateClassAdmin{Admin: admin.String(), NewAdmin: admin.String(), ClassId: "C01"},
			expErr: true,
		},
		"invalid: bad ClassID": {
			src:    MsgUpdateClassAdmin{Admin: admin.String(), NewAdmin: newAdmin.String(), ClassId: "asl;dfjkdjk???fgs;dfljgk"},
			expErr: true,
		},
		"invalid: bad admin addr": {
			src:    MsgUpdateClassAdmin{Admin: "?!@%)(87", NewAdmin: newAdmin.String(), ClassId: "C02"},
			expErr: true,
		},
		"invalid: bad NewAdmin addr": {
			src:    MsgUpdateClassAdmin{Admin: admin.String(), NewAdmin: "?!?@%?@$#6", ClassId: "C02"},
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

func TestMsgUpdateClassIssuers(t *testing.T) {
	_, _, a1 := testdata.KeyTestPubAddr()
	_, _, a2 := testdata.KeyTestPubAddr()

	tests := map[string]struct {
		src    MsgUpdateClassIssuers
		expErr bool
	}{
		"valid": {
			src:    MsgUpdateClassIssuers{Admin: a2.String(), ClassId: "C01", Issuers: []string{a1.String()}},
			expErr: false,
		},
		"invalid: no issuers": {
			src:    MsgUpdateClassIssuers{Admin: a2.String(), ClassId: "C01", Issuers: []string{}},
			expErr: true,
		},
		"invalid: no class ID": {
			src:    MsgUpdateClassIssuers{Admin: a2.String(), ClassId: "", Issuers: []string{a1.String()}},
			expErr: true,
		},
		"invalid: bad admin address": {
			src:    MsgUpdateClassIssuers{Admin: "//????.!", ClassId: "C01", Issuers: []string{a1.String()}},
			expErr: true,
		},
		"invalid: bad class ID": {
			src:    MsgUpdateClassIssuers{Admin: a1.String(), ClassId: "s.1%?#%", Issuers: []string{a1.String()}},
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

func TestMsgUpdateClassMetadata(t *testing.T) {
	_, _, a1 := testdata.KeyTestPubAddr()

	tests := map[string]struct {
		src    MsgUpdateClassMetadata
		expErr bool
	}{
		"valid": {
			src:    MsgUpdateClassMetadata{Admin: a1.String(), ClassId: "C01", Metadata: []byte("hello world")},
			expErr: false,
		},
		"invalid: bad admin address": {
			src:    MsgUpdateClassMetadata{Admin: "???a!#)(%", ClassId: "C01", Metadata: []byte("hello world")},
			expErr: true,
		},
		"invalid: bad class ID": {
			src:    MsgUpdateClassMetadata{Admin: a1.String(), ClassId: "6012949", Metadata: []byte("hello world")},
			expErr: true,
		},
		"invalid: no class ID": {
			src:    MsgUpdateClassMetadata{Admin: a1.String()},
			expErr: true,
		},
		"invalid: metadata too large": {
			src:    MsgUpdateClassMetadata{Admin: a1.String(), ClassId: "C01", Metadata: []byte(simtypes.RandStringOfLength(r, 288))},
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

func TestMsgSell(t *testing.T) {
	_, _, a1 := testdata.KeyTestPubAddr()

	validExpiration := time.Date(2030, 01, 01, 0, 0, 0, 0, time.UTC)

	tests := map[string]struct {
		src    MsgSell
		expErr bool
	}{
		"valid": {
			src: MsgSell{
				Owner: a1.String(),
				Orders: []*MsgSell_Order{
					{
						BatchDenom: "A00-00000000-00000000-000",
						Quantity:   "1.5",
						AskPrice: &sdk.Coin{
							Denom:  "uregen",
							Amount: sdk.NewInt(20),
						},
						DisableAutoRetire: true,
						Expiration:        &validExpiration,
					},
				},
			},
			expErr: false,
		},
		"invalid: bad owner address": {
			src: MsgSell{
				Owner: "foobar",
				Orders: []*MsgSell_Order{
					{
						BatchDenom: "A00-00000000-00000000-000",
						Quantity:   "1.5",
						AskPrice: &sdk.Coin{
							Denom:  "uregen",
							Amount: sdk.NewInt(20),
						},
						DisableAutoRetire: true,
					},
				},
			},
			expErr: true,
		},
		"invalid: bad batch denom": {
			src: MsgSell{
				Owner: a1.String(),
				Orders: []*MsgSell_Order{
					{
						BatchDenom: "foobar",
						Quantity:   "1.5",
						AskPrice: &sdk.Coin{
							Denom:  "uregen",
							Amount: sdk.NewInt(20),
						},
						DisableAutoRetire: true,
					},
				},
			},
			expErr: true,
		},
		"invalid: bad quantity": {
			src: MsgSell{
				Owner: a1.String(),
				Orders: []*MsgSell_Order{
					{
						BatchDenom: "A00-00000000-00000000-000",
						Quantity:   "-1.5",
						AskPrice: &sdk.Coin{
							Denom:  "uregen",
							Amount: sdk.NewInt(20),
						},
						DisableAutoRetire: true,
					},
				},
			},
			expErr: true,
		},
		"invalid: bad ask price": {
			src: MsgSell{
				Owner: a1.String(),
				Orders: []*MsgSell_Order{
					{
						BatchDenom: "A00-00000000-00000000-000",
						Quantity:   "1.5",
						AskPrice: &sdk.Coin{
							Denom:  "uregen",
							Amount: sdk.NewInt(-20),
						},
						DisableAutoRetire: true,
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

func TestMsgUpdateSellOrders(t *testing.T) {
	_, _, a1 := testdata.KeyTestPubAddr()

	validExpiration := time.Date(2030, 01, 01, 0, 0, 0, 0, time.UTC)

	tests := map[string]struct {
		src    MsgUpdateSellOrders
		expErr bool
	}{
		"valid": {
			src: MsgUpdateSellOrders{
				Owner: a1.String(),
				Updates: []*MsgUpdateSellOrders_Update{
					{
						NewQuantity: "1.5",
						NewAskPrice: &sdk.Coin{
							Denom:  "uregen",
							Amount: sdk.NewInt(20),
						},
						DisableAutoRetire: true,
						NewExpiration:     &validExpiration,
					},
				},
			},
			expErr: false,
		},
		"invalid: bad owner address": {
			src: MsgUpdateSellOrders{
				Owner: "foobar",
				Updates: []*MsgUpdateSellOrders_Update{
					{
						NewQuantity: "1.5",
						NewAskPrice: &sdk.Coin{
							Denom:  "uregen",
							Amount: sdk.NewInt(20),
						},
						DisableAutoRetire: true,
					},
				},
			},
			expErr: true,
		},
		"invalid: bad quantity": {
			src: MsgUpdateSellOrders{
				Owner: a1.String(),
				Updates: []*MsgUpdateSellOrders_Update{
					{
						NewQuantity: "-1.5",
						NewAskPrice: &sdk.Coin{
							Denom:  "uregen",
							Amount: sdk.NewInt(20),
						},
						DisableAutoRetire: true,
					},
				},
			},
			expErr: true,
		},
		"invalid: bad ask price": {
			src: MsgUpdateSellOrders{
				Owner: a1.String(),
				Updates: []*MsgUpdateSellOrders_Update{
					{
						NewQuantity: "1.5",
						NewAskPrice: &sdk.Coin{
							Denom:  "uregen",
							Amount: sdk.NewInt(-20),
						},
						DisableAutoRetire: true,
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

func TestMsgBuy(t *testing.T) {
	_, _, a1 := testdata.KeyTestPubAddr()

	validExpiration := time.Date(2030, 01, 01, 0, 0, 0, 0, time.UTC)

	tests := map[string]struct {
		src    MsgBuy
		expErr bool
	}{
		"valid": {
			src: MsgBuy{
				Buyer: a1.String(),
				Orders: []*MsgBuy_Order{
					{
						Quantity: "1.5",
						BidPrice: &sdk.Coin{
							Denom:  "uregen",
							Amount: sdk.NewInt(20),
						},
						DisableAutoRetire:  true,
						DisablePartialFill: true,
						Expiration:         &validExpiration,
					},
				},
			},
			expErr: false,
		},
		"invalid: bad owner address": {
			src: MsgBuy{
				Buyer: "foobar",
				Orders: []*MsgBuy_Order{
					{
						Quantity: "1.5",
						BidPrice: &sdk.Coin{
							Denom:  "uregen",
							Amount: sdk.NewInt(20),
						},
						DisableAutoRetire:  true,
						DisablePartialFill: true,
					},
				},
			},
			expErr: true,
		},
		"invalid: bad quantity": {
			src: MsgBuy{
				Buyer: a1.String(),
				Orders: []*MsgBuy_Order{
					{
						Quantity: "-1.5",
						BidPrice: &sdk.Coin{
							Denom:  "uregen",
							Amount: sdk.NewInt(20),
						},
						DisableAutoRetire:  true,
						DisablePartialFill: true,
					},
				},
			},
			expErr: true,
		},
		"invalid: bad bid price": {
			src: MsgBuy{
				Buyer: a1.String(),
				Orders: []*MsgBuy_Order{
					{
						Quantity: "1.5",
						BidPrice: &sdk.Coin{
							Denom:  "uregen",
							Amount: sdk.NewInt(-20),
						},
						DisableAutoRetire:  true,
						DisablePartialFill: true,
					},
				},
			},
			expErr: true,
		},
		"invalid: bad retirement location": {
			src: MsgBuy{
				Buyer: a1.String(),
				Orders: []*MsgBuy_Order{
					{
						Quantity: "1.5",
						BidPrice: &sdk.Coin{
							Denom:  "uregen",
							Amount: sdk.NewInt(20),
						},
						DisableAutoRetire:  true,
						DisablePartialFill: true,
						RetirementLocation: "foo",
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

func TestMsgAllowAskDenom(t *testing.T) {
	_, _, a1 := testdata.KeyTestPubAddr()

	tests := map[string]struct {
		src    MsgAllowAskDenom
		expErr bool
	}{
		"valid": {
			src: MsgAllowAskDenom{
				RootAddress:  a1.String(),
				Denom:        "uregen",
				DisplayDenom: "regen",
				Exponent:     6,
			},
			expErr: false,
		},
		"invalid address": {
			src: MsgAllowAskDenom{
				RootAddress:  "foobar",
				Denom:        "uregen",
				DisplayDenom: "regen",
				Exponent:     6,
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
