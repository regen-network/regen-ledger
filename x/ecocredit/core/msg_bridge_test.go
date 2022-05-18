package core

import (
	"testing"

	"github.com/regen-network/regen-ledger/types/testutil"
	"github.com/stretchr/testify/require"
)

func TestMsgBridge(t *testing.T) {
	t.Parallel()

	addr1 := testutil.GenAddress()
	addr2 := testutil.GenAddress()
	contract := "0x06012c8cf97bead5deae237070f9587f8e7a266d"

	tests := map[string]struct {
		src    MsgBridge
		expErr bool
	}{
		"valid msg": {
			src: MsgBridge{
				MsgCancel: &MsgCancel{
					Holder: addr1,
					Credits: []*MsgCancel_CancelCredits{
						{
							BatchDenom: batchDenom,
							Amount:     "10",
						},
					},
					Reason: "reason",
				},
				BridgeTarget:    addr1,
				BridgeContract:  contract,
				BridgeRecipient: addr2,
			},
			expErr: false,
		},
		"invalid msg without holder": {
			src: MsgBridge{
				MsgCancel: &MsgCancel{
					Credits: []*MsgCancel_CancelCredits{
						{
							BatchDenom: batchDenom,
							Amount:     "10",
						},
					},
					Reason: "reason",
				},
				BridgeTarget:    addr1,
				BridgeRecipient: contract,
				BridgeContract:  addr2,
			},
			expErr: true,
		},
		"invalid msg with wrong holder address": {
			src: MsgBridge{
				MsgCancel: &MsgCancel{
					Holder: "wrongHolder",
					Credits: []*MsgCancel_CancelCredits{
						{
							BatchDenom: batchDenom,
							Amount:     "10",
						},
					},
					Reason: "reason",
				},
			},
			expErr: true,
		},
		"invalid msg without credits": {
			src: MsgBridge{
				MsgCancel: &MsgCancel{
					Holder: addr1,
				},
			},
			expErr: true,
		},
		"invalid msg without Credits.BatchDenom": {
			src: MsgBridge{
				MsgCancel: &MsgCancel{
					Holder: addr1,
					Credits: []*MsgCancel_CancelCredits{
						{
							Amount: "10",
						},
					},
				},
			},
			expErr: true,
		},
		"invalid msg without Credits.Amount": {
			src: MsgBridge{
				MsgCancel: &MsgCancel{
					Holder: addr1,
					Credits: []*MsgCancel_CancelCredits{
						{
							BatchDenom: batchDenom,
						},
					},
				},
			},
			expErr: true,
		},
		"invalid msg with wrong Credits.Amount": {
			src: MsgBridge{
				MsgCancel: &MsgCancel{
					Holder: addr1,
					Credits: []*MsgCancel_CancelCredits{
						{
							BatchDenom: batchDenom,
							Amount:     "abc",
						},
					},
				},
			},
			expErr: true,
		},
		"invalid msg without bridge target": {
			src: MsgBridge{
				MsgCancel: &MsgCancel{
					Holder: addr1,
					Credits: []*MsgCancel_CancelCredits{
						{
							BatchDenom: batchDenom,
							Amount:     "10",
						},
					},
					Reason: "reason",
				},
				BridgeContract:  contract,
				BridgeRecipient: addr2,
			},
			expErr: true,
		},
		"invalid msg without bridge contract": {
			src: MsgBridge{
				MsgCancel: &MsgCancel{
					Holder: addr1,
					Credits: []*MsgCancel_CancelCredits{
						{
							BatchDenom: batchDenom,
							Amount:     "10",
						},
					},
					Reason: "reason",
				},
				BridgeTarget:    addr1,
				BridgeRecipient: addr2,
			},
			expErr: true,
		},
		"invalid msg without bridge recipient address": {
			src: MsgBridge{
				MsgCancel: &MsgCancel{
					Holder: addr1,
					Credits: []*MsgCancel_CancelCredits{
						{
							BatchDenom: batchDenom,
							Amount:     "10",
						},
					},
					Reason: "reason",
				},
				BridgeTarget:   addr1,
				BridgeContract: contract,
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
