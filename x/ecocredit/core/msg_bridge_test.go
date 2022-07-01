package core

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/regen-network/regen-ledger/types/testutil"
)

func TestMsgBridge(t *testing.T) {
	t.Parallel()

	addr1 := testutil.GenAddress()
	recipient := "0x323b5d4c32345ced77393b3530b1eed0f346429d"

	tests := map[string]struct {
		msg       MsgBridge
		expErr    bool
		expErrMsg string
	}{
		"valid msg": {
			msg: MsgBridge{
				Owner: addr1,
				Credits: []*Credits{
					{
						BatchDenom: batchDenom,
						Amount:     "10",
					},
				},
				Target:    "polygon",
				Recipient: recipient,
			},
			expErr: false,
		},
		"invalid msg without owner": {
			msg: MsgBridge{
				Credits: []*Credits{
					{
						BatchDenom: batchDenom,
						Amount:     "10",
					},
				},
				Target:    "polygon",
				Recipient: recipient,
			},
			expErr:    true,
			expErrMsg: "empty address string is not allowed",
		},
		"invalid msg with invalid owner address": {
			msg: MsgBridge{
				Owner: "wrong owner",
				Credits: []*Credits{
					{
						BatchDenom: batchDenom,
						Amount:     "10",
					},
				},
			},
			expErr:    true,
			expErrMsg: "decoding bech32 failed",
		},
		"invalid msg without credits": {
			msg: MsgBridge{
				Owner: addr1,
			},
			expErr:    true,
			expErrMsg: "credits should not be empty",
		},
		"invalid msg without Credits.BatchDenom": {
			msg: MsgBridge{
				Owner: addr1,
				Credits: []*Credits{
					{
						Amount: "10",
					},
				},
			},
			expErr:    true,
			expErrMsg: "invalid batch denom",
		},
		"invalid msg without Credits.Amount": {
			msg: MsgBridge{
				Owner: addr1,
				Credits: []*Credits{
					{
						BatchDenom: batchDenom,
					},
				},
			},
			expErr:    true,
			expErrMsg: "invalid decimal string",
		},
		"invalid msg with invalid Credits.Amount": {
			msg: MsgBridge{
				Owner: addr1,
				Credits: []*Credits{
					{
						BatchDenom: batchDenom,
						Amount:     "abc",
					},
				},
			},
			expErr:    true,
			expErrMsg: "invalid decimal string",
		},
		"invalid msg without bridge target": {
			msg: MsgBridge{
				Owner: addr1,
				Credits: []*Credits{
					{
						BatchDenom: batchDenom,
						Amount:     "10",
					},
				},
				Recipient: recipient,
			},
			expErr:    true,
			expErrMsg: "expected polygon",
		},
		"invalid msg with invalid bridge target": {
			msg: MsgBridge{
				Owner: addr1,
				Credits: []*Credits{
					{
						BatchDenom: batchDenom,
						Amount:     "10",
					},
				},
				Target:    "polygon1",
				Recipient: recipient,
			},
			expErr:    true,
			expErrMsg: "expected polygon",
		},
		"invalid msg without bridge recipient address": {
			msg: MsgBridge{
				Owner: addr1,
				Credits: []*Credits{
					{
						BatchDenom: batchDenom,
						Amount:     "10",
					},
				},
				Target: "polygon",
			},
			expErr:    true,
			expErrMsg: "not a valid ethereum address",
		},
		"invalid msg with invalid bridge recipient address": {
			msg: MsgBridge{
				Owner: addr1,
				Credits: []*Credits{
					{
						BatchDenom: batchDenom,
						Amount:     "10",
					},
				},
				Target:    "polygon",
				Recipient: addr1,
			},
			expErr:    true,
			expErrMsg: "not a valid ethereum address",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			err := test.msg.ValidateBasic()
			if test.expErr {
				require.ErrorContains(t, err, test.expErrMsg)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
