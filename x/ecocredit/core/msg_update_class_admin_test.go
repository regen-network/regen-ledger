package core

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/cosmos/cosmos-sdk/testutil/testdata"
)

func TestMsgUpdateClassAdmin(t *testing.T) {
	t.Parallel()

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
