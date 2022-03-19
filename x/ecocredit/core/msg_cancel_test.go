package core

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/cosmos/cosmos-sdk/testutil/testdata"
)

func TestMsgCancel(t *testing.T) {
	t.Parallel()

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
