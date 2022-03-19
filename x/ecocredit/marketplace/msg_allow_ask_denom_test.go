package marketplace

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/cosmos/cosmos-sdk/testutil/testdata"
)

func TestMsgAllowAskDenom(t *testing.T) {
	t.Parallel()

	_, _, a1 := testdata.KeyTestPubAddr()

	tests := map[string]struct {
		src    MsgAllowAskDenom
		expErr bool
	}{
		"valid": {
			src: MsgAllowAskDenom{
				RootAddress:  a1.String(),
				Denom:        "uregen",
				DisplayDenom: "regen",
				Exponent:     6,
			},
			expErr: false,
		},
		"invalid address": {
			src: MsgAllowAskDenom{
				RootAddress:  "foobar",
				Denom:        "uregen",
				DisplayDenom: "regen",
				Exponent:     6,
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
