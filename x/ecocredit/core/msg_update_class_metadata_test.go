package core

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func TestMsgUpdateClassMetadata(t *testing.T) {
	t.Parallel()

	_, _, a1 := testdata.KeyTestPubAddr()

	tests := map[string]struct {
		src    MsgUpdateClassMetadata
		expErr bool
	}{
		"valid": {
			src:    MsgUpdateClassMetadata{Admin: a1.String(), ClassId: "C01", Metadata: "hello world"},
			expErr: false,
		},
		"invalid: bad admin address": {
			src:    MsgUpdateClassMetadata{Admin: "???a!#)(%", ClassId: "C01", Metadata: "hello world"},
			expErr: true,
		},
		"invalid: bad class ID": {
			src:    MsgUpdateClassMetadata{Admin: a1.String(), ClassId: "6012949", Metadata: "hello world"},
			expErr: true,
		},
		"invalid: no class ID": {
			src:    MsgUpdateClassMetadata{Admin: a1.String()},
			expErr: true,
		},
		"invalid: metadata too large": {
			src:    MsgUpdateClassMetadata{Admin: a1.String(), ClassId: "C01", Metadata: simtypes.RandStringOfLength(r, 288)},
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
