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
		"invalid: bad denom": {
			src: MsgBuy{
				Buyer: a1.String(),
				Orders: []*MsgBuy_Order{
					{
						Quantity: "1.5",
						BidPrice: &sdk.Coin{
							Denom:  "$$$$$",
							Amount: sdk.NewInt(20),
						},
						DisableAutoRetire:  true,
						DisablePartialFill: true,
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

func TestMsgTakeFromBasket_ValidateBasic(t *testing.T) {
	_, _, addr := testdata.KeyTestPubAddr()
	validLocation := "AB-CDE FG1 345"
	validDenom := "FooBar-Basket"

	type fields struct {
		Owner              string
		BasketDenom        string
		Amount             string
		RetirementLocation string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "valid",
			fields: fields{
				Owner:              addr.String(),
				BasketDenom:        validDenom,
				Amount:             "10.23510",
				RetirementLocation: validLocation,
			},
		},
		{
			name: "bad owner",
			fields: fields{
				Owner:              "foo",
				BasketDenom:        validDenom,
				Amount:             "10.23510",
				RetirementLocation: validLocation,
			},
			wantErr: true,
		},
		{
			name: "bad denom",
			fields: fields{
				Owner:              addr.String(),
				BasketDenom:        "foo.bar",
				Amount:             "10.23510",
				RetirementLocation: validLocation,
			},
			wantErr: true,
		},
		{
			name: "bad amount",
			fields: fields{
				Owner:              addr.String(),
				BasketDenom:        validDenom,
				Amount:             "-4.20",
				RetirementLocation: validLocation,
			},
			wantErr: true,
		},
		{
			name: "bad location",
			fields: fields{
				Owner:              addr.String(),
				BasketDenom:        validDenom,
				Amount:             "10.23510",
				RetirementLocation: "oops",
			},
			wantErr: true,
		},
		{
			name: "valid - no location is fine",
			fields: fields{
				Owner:              addr.String(),
				BasketDenom:        validDenom,
				Amount:             "10.23510",
				RetirementLocation: "",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MsgTakeFromBasket{
				Owner:              tt.fields.Owner,
				BasketDenom:        tt.fields.BasketDenom,
				Amount:             tt.fields.Amount,
				RetirementLocation: tt.fields.RetirementLocation,
			}
			if err := m.ValidateBasic(); (err != nil) != tt.wantErr {
				t.Errorf("ValidateBasic() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMsgPickFromBasket_ValidateBasic(t *testing.T) {
	_, _, addr := testdata.KeyTestPubAddr()
	validLocation := "AB-CDE FG1 345"
	validDenom := "A00-00000000-00000000-000"

	type fields struct {
		Owner              string
		BasketDenom        string
		Credits            []*BasketCredit
		RetirementLocation string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "valid",
			fields: fields{
				Owner:              addr.String(),
				BasketDenom:        validDenom,
				Credits:            []*BasketCredit{{BatchDenom: validDenom, TradableAmount: "4.20"}},
				RetirementLocation: validLocation,
			},
		},
		{
			name: "bad addr",
			fields: fields{
				Owner:              "foo",
				BasketDenom:        validDenom,
				Credits:            []*BasketCredit{{BatchDenom: validDenom, TradableAmount: "4.20"}},
				RetirementLocation: validLocation,
			},
			wantErr: true,
		},
		{
			name: "bad denom",
			fields: fields{
				Owner:              addr.String(),
				BasketDenom:        "foo.bar//",
				Credits:            []*BasketCredit{{BatchDenom: validDenom, TradableAmount: "4.20"}},
				RetirementLocation: validLocation,
			},
			wantErr: true,
		},
		{
			name: "no credits",
			fields: fields{
				Owner:              addr.String(),
				BasketDenom:        validDenom,
				Credits:            nil,
				RetirementLocation: validLocation,
			},
			wantErr: true,
		},
		{
			name: "bad basket credit denom",
			fields: fields{
				Owner:              addr.String(),
				BasketDenom:        validDenom,
				Credits:            []*BasketCredit{{BatchDenom: "alwkef", TradableAmount: "1.20"}},
				RetirementLocation: validLocation,
			},
			wantErr: true,
		},
		{
			name: "bad basket credit tradable amount",
			fields: fields{
				Owner:              addr.String(),
				BasketDenom:        validDenom,
				Credits:            []*BasketCredit{{BatchDenom: validDenom, TradableAmount: "-1.20"}},
				RetirementLocation: validLocation,
			},
			wantErr: true,
		},
		{
			name: "bad location",
			fields: fields{
				Owner:              addr.String(),
				BasketDenom:        validDenom,
				Credits:            []*BasketCredit{{BatchDenom: validDenom, TradableAmount: "4.20"}},
				RetirementLocation: "foo",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MsgPickFromBasket{
				Owner:              tt.fields.Owner,
				BasketDenom:        tt.fields.BasketDenom,
				Credits:            tt.fields.Credits,
				RetirementLocation: tt.fields.RetirementLocation,
			}
			if err := m.ValidateBasic(); (err != nil) != tt.wantErr {
				t.Errorf("ValidateBasic() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMsgAddToBasket_ValidateBasic(t *testing.T) {
	_, _, addr := testdata.KeyTestPubAddr()
	validDenom := "A00-00000000-00000000-000"

	type fields struct {
		Owner       string
		BasketDenom string
		Credits     []*BasketCredit
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "valid",
			fields: fields{
				Owner:       addr.String(),
				BasketDenom: validDenom,
				Credits:     []*BasketCredit{{validDenom, "12.402"}},
			},
		},
		{
			name: "bad addr",
			fields: fields{
				Owner:       "foo",
				BasketDenom: validDenom,
				Credits:     []*BasketCredit{{validDenom, "12.402"}},
			},
			wantErr: true,
		},
		{
			name: "bad denom",
			fields: fields{
				Owner:       addr.String(),
				BasketDenom: "foo.bar//--",
				Credits:     []*BasketCredit{{validDenom, "12.402"}},
			},
			wantErr: true,
		},
		{
			name: "no credits",
			fields: fields{
				Owner:       addr.String(),
				BasketDenom: validDenom,
				Credits:     nil,
			},
			wantErr: true,
		},
		{
			name: "bad credit denom",
			fields: fields{
				Owner:       addr.String(),
				BasketDenom: validDenom,
				Credits:     []*BasketCredit{{"foo", "12.402"}},
			},
			wantErr: true,
		},
		{
			name: "bad credit tradable amount",
			fields: fields{
				Owner:       addr.String(),
				BasketDenom: validDenom,
				Credits:     []*BasketCredit{{validDenom, "-1.30"}},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MsgAddToBasket{
				Owner:       tt.fields.Owner,
				BasketDenom: tt.fields.BasketDenom,
				Credits:     tt.fields.Credits,
			}
			if err := m.ValidateBasic(); (err != nil) != tt.wantErr {
				t.Errorf("ValidateBasic() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMsgCreateBasket_ValidateBasic(t *testing.T) {
	_, _, addr := testdata.KeyTestPubAddr()
	validLocation := "AB-CDE FG1 345"
	_ = validLocation
	validDenom := "A00-00000000-00000000-000"
	_ = validDenom

	type fields struct {
		Curator           string
		Name              string
		DisplayName       string
		Exponent          uint32
		BasketCriteria    *Filter
		DisableAutoRetire bool
		AllowPicking      bool
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "valid",
			fields: fields{
				Curator:           addr.String(),
				Name:              "MyVeryCoolBasket",
				DisplayName:       "cool BASKET, inc.",
				Exponent:          10,
				BasketCriteria:    &Filter{Sum: &Filter_Owner{Owner: addr.String()}},
				DisableAutoRetire: false,
				AllowPicking:      false,
			},
		},
		{
			name: "valid nested",
			fields: fields{
				Curator:           addr.String(),
				Name:              "MyVeryCoolBasket",
				DisplayName:       "cool BASKET, inc.",
				Exponent:          10,
				BasketCriteria:    &Filter{Sum: &Filter_And_{And: &Filter_And{Filters: []*Filter{{Sum: &Filter_Issuer{Issuer: addr.String()}}}}}},
				DisableAutoRetire: false,
				AllowPicking:      false,
			},
		},
		{
			name: "bad addr",
			fields: fields{
				Curator:           "oops",
				Name:              "my very cool basket",
				DisplayName:       "cool BASKET, inc.",
				Exponent:          10,
				BasketCriteria:    &Filter{Sum: &Filter_Owner{Owner: addr.String()}},
				DisableAutoRetire: false,
				AllowPicking:      false,
			},
			wantErr: true,
		},
		{
			name: "bad filter (not an address)",
			fields: fields{
				Curator:           addr.String(),
				Name:              "my very cool basket",
				DisplayName:       "cool BASKET, inc.",
				Exponent:          10,
				BasketCriteria:    &Filter{Sum: &Filter_Owner{Owner: "foo"}},
				DisableAutoRetire: false,
				AllowPicking:      false,
			},
			wantErr: true,
		},
		{
			name: "bad filter nested",
			fields: fields{
				Curator:           addr.String(),
				Name:              "my very cool basket",
				DisplayName:       "cool BASKET, inc.",
				Exponent:          10,
				BasketCriteria:    &Filter{Sum: &Filter_And_{And: &Filter_And{Filters: []*Filter{{Sum: &Filter_Issuer{Issuer: addr.String()}}, {Sum: &Filter_Owner{Owner: "foo"}}}}}},
				DisableAutoRetire: false,
				AllowPicking:      false,
			},
			wantErr: true,
		},
		{
			name: "bad filter location",
			fields: fields{
				Curator:           addr.String(),
				Name:              "my very cool basket",
				DisplayName:       "cool BASKET, inc.",
				Exponent:          10,
				BasketCriteria:    &Filter{Sum: &Filter_ProjectLocation{ProjectLocation: "not a location"}},
				DisableAutoRetire: false,
				AllowPicking:      false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MsgCreateBasket{
				Curator:           tt.fields.Curator,
				Name:              tt.fields.Name,
				DisplayName:       tt.fields.DisplayName,
				Exponent:          tt.fields.Exponent,
				BasketCriteria:    tt.fields.BasketCriteria,
				DisableAutoRetire: tt.fields.DisableAutoRetire,
				AllowPicking:      tt.fields.AllowPicking,
			}
			if err := m.ValidateBasic(); (err != nil) != tt.wantErr {
				t.Errorf("ValidateBasic() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
