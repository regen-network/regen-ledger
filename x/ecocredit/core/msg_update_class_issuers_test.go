package core

import (
	"testing"

	"github.com/regen-network/regen-ledger/types/testutil"
	"github.com/stretchr/testify/require"
)

func TestMsgUpdateClassIssuers(t *testing.T) {
	t.Parallel()

	a1 := testutil.GenAddress()
	a2 := testutil.GenAddress()

	tests := map[string]struct {
		src    MsgUpdateClassIssuers
		expErr bool
	}{
		"valid": {
			src:    MsgUpdateClassIssuers{Admin: a2, ClassId: "C01", AddIssuers: []string{a1}, RemoveIssuers: []string{a2}},
			expErr: false,
		},
		"invalid: no issuers": {
			src:    MsgUpdateClassIssuers{Admin: a2, ClassId: "C01"},
			expErr: true,
		},
		"invalid: no class ID": {
			src:    MsgUpdateClassIssuers{Admin: a2, ClassId: "", AddIssuers: []string{a1}},
			expErr: true,
		},
		"invalid: bad admin address": {
			src:    MsgUpdateClassIssuers{Admin: "//????.!", ClassId: "C01", AddIssuers: []string{a1}},
			expErr: true,
		},
		"invalid: bad class ID": {
			src:    MsgUpdateClassIssuers{Admin: a1, ClassId: "s.1%?#%", AddIssuers: []string{a1}},
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
