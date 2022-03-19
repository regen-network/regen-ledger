package core

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/cosmos/cosmos-sdk/testutil/testdata"
)

func TestMsgUpdateClassIssuers(t *testing.T) {
	t.Parallel()

	_, _, a1 := testdata.KeyTestPubAddr()
	_, _, a2 := testdata.KeyTestPubAddr()

	tests := map[string]struct {
		src    MsgUpdateClassIssuers
		expErr bool
	}{
		"valid": {
			src:    MsgUpdateClassIssuers{Admin: a2.String(), ClassId: "C01", AddIssuers: []string{a1.String()}, RemoveIssuers: []string{a2.String()}},
			expErr: false,
		},
		"invalid: no issuers": {
			src:    MsgUpdateClassIssuers{Admin: a2.String(), ClassId: "C01"},
			expErr: true,
		},
		"invalid: no class ID": {
			src:    MsgUpdateClassIssuers{Admin: a2.String(), ClassId: "", AddIssuers: []string{a1.String()}},
			expErr: true,
		},
		"invalid: bad admin address": {
			src:    MsgUpdateClassIssuers{Admin: "//????.!", ClassId: "C01", AddIssuers: []string{a1.String()}},
			expErr: true,
		},
		"invalid: bad class ID": {
			src:    MsgUpdateClassIssuers{Admin: a1.String(), ClassId: "s.1%?#%", AddIssuers: []string{a1.String()}},
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
