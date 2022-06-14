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
	contract := "0x06012c8cf97bead5deae237070f9587f8e7a266d"

	tests := map[string]struct {
		src    MsgBridge
		expErr bool
	}{
		"valid msg": {
			src: MsgBridge{
				Owner: addr1,
				Credits: []*Credits{
					{
						BatchDenom: batchDenom,
						Amount:     "10",
					},
				},
				Target:    "polygon",
				Contract:  contract,
				Recipient: recipient,
			},
			expErr: false,
		},
		"invalid msg without owner": {
			src: MsgBridge{
				Credits: []*Credits{
					{
						BatchDenom: batchDenom,
						Amount:     "10",
					},
				},
				Target:    "polygon",
				Recipient: recipient,
				Contract:  contract,
			},
			expErr: true,
		},
		"invalid msg with wrong owner address": {
			src: MsgBridge{
				Owner: "wrong owner",
				Credits: []*Credits{
					{
						BatchDenom: batchDenom,
						Amount:     "10",
					},
				},
			},
			expErr: true,
		},
		"invalid msg without credits": {
			src: MsgBridge{
				Owner: addr1,
			},
			expErr: true,
		},
		"invalid msg without Credits.BatchDenom": {
			src: MsgBridge{
				Owner: addr1,
				Credits: []*Credits{
					{
						Amount: "10",
					},
				},
			},
			expErr: true,
		},
		"invalid msg without Credits.Amount": {
			src: MsgBridge{
				Owner: addr1,
				Credits: []*Credits{
					{
						BatchDenom: batchDenom,
					},
				},
			},
			expErr: true,
		},
		"invalid msg with wrong Credits.Amount": {
			src: MsgBridge{
				Owner: addr1,
				Credits: []*Credits{
					{
						BatchDenom: batchDenom,
						Amount:     "abc",
					},
				},
			},
			expErr: true,
		},
		"invalid msg without bridge target": {
			src: MsgBridge{
				Owner: addr1,
				Credits: []*Credits{
					{
						BatchDenom: batchDenom,
						Amount:     "10",
					},
				},
				Contract:  contract,
				Recipient: recipient,
			},
			expErr: true,
		},
		"invalid msg without bridge contract": {
			src: MsgBridge{
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
			expErr: true,
		},
		"invalid msg without bridge recipient address": {
			src: MsgBridge{
				Owner: addr1,
				Credits: []*Credits{
					{
						BatchDenom: batchDenom,
						Amount:     "10",
					},
				},
				Target:   "polygon",
				Contract: contract,
			},
			expErr: true,
		},
		"invalid bridge recipient address": {
			src: MsgBridge{
				Owner: addr1,
				Credits: []*Credits{
					{
						BatchDenom: batchDenom,
						Amount:     "10",
					},
				},
				Target:    "polygon",
				Recipient: addr1,
				Contract:  contract,
			},
			expErr: true,
		},
		"invalid bridge target": {
			src: MsgBridge{
				Owner: addr1,
				Credits: []*Credits{
					{
						BatchDenom: batchDenom,
						Amount:     "10",
					},
				},
				Target:    "polygon1",
				Recipient: recipient,
				Contract:  contract,
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
