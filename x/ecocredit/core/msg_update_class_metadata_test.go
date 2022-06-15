package core

import (
	"strings"
	"testing"

	"github.com/regen-network/regen-ledger/types/testutil"
	"github.com/stretchr/testify/require"
)

func TestMsgUpdateClassMetadata(t *testing.T) {
	t.Parallel()

	a1 := testutil.GenAddress()
	tests := map[string]struct {
		src    MsgUpdateClassMetadata
		expErr bool
	}{
		"valid": {
			src:    MsgUpdateClassMetadata{Admin: a1, ClassId: "C01", NewMetadata: "hello world"},
			expErr: false,
		},
		"invalid: bad admin address": {
			src:    MsgUpdateClassMetadata{Admin: "???a!#)(%", ClassId: "C01", NewMetadata: "hello world"},
			expErr: true,
		},
		"invalid: bad class ID": {
			src:    MsgUpdateClassMetadata{Admin: a1, ClassId: "6012949", NewMetadata: "hello world"},
			expErr: true,
		},
		"invalid: no class ID": {
			src:    MsgUpdateClassMetadata{Admin: a1},
			expErr: true,
		},
		"invalid: metadata too large": {
			src:    MsgUpdateClassMetadata{Admin: a1, ClassId: "C01", NewMetadata: strings.Repeat("x", 288)},
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
