package core

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/regen-network/regen-ledger/types/testutil"
)

func TestMsgSend(t *testing.T) {
	t.Parallel()

	addr1 := testutil.GenAddress()
	addr2 := testutil.GenAddress()

	tests := map[string]struct {
		src    MsgSendBulk
		expErr bool
	}{
		"valid msg": {
			src: MsgSendBulk{
				Sender:    addr1,
				Recipient: addr2,
				Credits: []*MsgSendBulk_SendCredits{
					{
						BatchDenom:             batchDenom,
						TradableAmount:         "10",
						RetiredAmount:          "10",
						RetirementJurisdiction: "ST-UVW XY Z12",
					},
				},
			},
			expErr: false,
		},
		"invalid msg with Credits.RetiredAmount negative value": {
			src: MsgSendBulk{
				Sender:    addr1,
				Recipient: addr2,
				Credits: []*MsgSendBulk_SendCredits{
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
			src: MsgSendBulk{
				Sender:    addr1,
				Recipient: addr2,
			},
			expErr: true,
		},
		"invalid msg without sender": {
			src: MsgSendBulk{
				Recipient: addr2,
				Credits: []*MsgSendBulk_SendCredits{
					{
						BatchDenom:             "some_denom",
						TradableAmount:         "10",
						RetiredAmount:          "10",
						RetirementJurisdiction: "ST-UVW XY Z12",
					},
				},
			},
			expErr: true,
		},
		"invalid msg without recipient": {
			src: MsgSendBulk{
				Sender: addr1,
				Credits: []*MsgSendBulk_SendCredits{
					{
						BatchDenom:             "some_denom",
						TradableAmount:         "10",
						RetiredAmount:          "10",
						RetirementJurisdiction: "ST-UVW XY Z12",
					},
				},
			},
			expErr: true,
		},
		"invalid msg without Credits.BatchDenom": {
			src: MsgSendBulk{
				Sender:    addr1,
				Recipient: addr2,
				Credits: []*MsgSendBulk_SendCredits{
					{
						TradableAmount:         "10",
						RetiredAmount:          "10",
						RetirementJurisdiction: "ST-UVW XY Z12",
					},
				},
			},
			expErr: true,
		},
		"invalid msg without Credits.TradableAmount set": {
			src: MsgSendBulk{
				Sender:    addr1,
				Recipient: addr2,
				Credits: []*MsgSendBulk_SendCredits{
					{
						BatchDenom:             "some_denom",
						RetiredAmount:          "10",
						RetirementJurisdiction: "ST-UVW XY Z12",
					},
				},
			},
			expErr: true,
		},
		"invalid msg without Credits.RetiredAmount set": {
			src: MsgSendBulk{
				Sender:    addr1,
				Recipient: addr2,
				Credits: []*MsgSendBulk_SendCredits{
					{
						BatchDenom:             "some_denom",
						TradableAmount:         "10",
						RetirementJurisdiction: "ST-UVW XY Z12",
					},
				},
			},
			expErr: true,
		},
		"invalid msg without Credits.RetirementJurisdiction": {
			src: MsgSendBulk{
				Sender:    addr1,
				Recipient: addr2,
				Credits: []*MsgSendBulk_SendCredits{
					{
						BatchDenom:     "some_denom",
						TradableAmount: "10",
						RetiredAmount:  "10",
					},
				},
			},
			expErr: true,
		},
		"valid msg without Credits.RetirementJurisdiction(When RetiredAmount is zero)": {
			src: MsgSendBulk{
				Sender:    addr1,
				Recipient: addr2,
				Credits: []*MsgSendBulk_SendCredits{
					{
						BatchDenom:     batchDenom,
						TradableAmount: "10",
						RetiredAmount:  "0",
					},
				},
			},
			expErr: false,
		},
		"invalid msg with wrong sender": {
			src: MsgSendBulk{
				Sender:    "wrongSender",
				Recipient: addr2,
				Credits: []*MsgSendBulk_SendCredits{
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
			src: MsgSendBulk{
				Sender:    addr1,
				Recipient: "wrongRecipient",
				Credits: []*MsgSendBulk_SendCredits{
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
