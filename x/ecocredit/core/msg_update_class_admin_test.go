package core

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/regen-network/regen-ledger/types/testutil"
)

func TestMsgUpdateClassAdmin(t *testing.T) {
	t.Parallel()

	admin := testutil.GenAddress()
	newAdmin := testutil.GenAddress()

	tests := map[string]struct {
		src    MsgUpdateClassAdmin
		expErr bool
	}{
		"valid": {
			src:    MsgUpdateClassAdmin{Admin: admin, NewAdmin: newAdmin, ClassId: "C01"},
			expErr: false,
		},
		"invalid: same address": {
			src:    MsgUpdateClassAdmin{Admin: admin, NewAdmin: admin, ClassId: "C01"},
			expErr: true,
		},
		"invalid: bad ClassID": {
			src:    MsgUpdateClassAdmin{Admin: admin, NewAdmin: newAdmin, ClassId: "asl;dfjkdjk???fgs;dfljgk"},
			expErr: true,
		},
		"invalid: bad admin addr": {
			src:    MsgUpdateClassAdmin{Admin: "?!@%)(87", NewAdmin: newAdmin, ClassId: "C02"},
			expErr: true,
		},
		"invalid: bad NewAdmin addr": {
			src:    MsgUpdateClassAdmin{Admin: admin, NewAdmin: "?!?@%?@$#6", ClassId: "C02"},
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
