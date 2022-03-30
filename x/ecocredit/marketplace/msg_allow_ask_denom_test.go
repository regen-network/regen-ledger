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
				RootAddress: a1.String(),
				AddDenoms: []*MsgAllowAskDenom_DenomInfo{
					{Denom: "uregen", DisplayDenom: "regen", Exponent: 6},
				},
				RemoveDenoms: []string{"uregen"},
			},
			expErr: false,
		},
		"invalid address": {
			src: MsgAllowAskDenom{
				RootAddress: "foobar",
			},
			expErr: true,
		},
		"none specified": {
			src: MsgAllowAskDenom{
				RootAddress: a1.String(),
			},
			expErr: true,
		},
		"invalid denom in add_denoms": {
			src: MsgAllowAskDenom{
				RootAddress: a1.String(),
				AddDenoms: []*MsgAllowAskDenom_DenomInfo{
					{Denom: "r/e-ge!n", DisplayDenom: "foo", Exponent: 3},
				},
			},
			expErr: true,
		},
		"invalid denom in remove_denoms": {
			src: MsgAllowAskDenom{
				RootAddress:  a1.String(),
				RemoveDenoms: []string{"r/e!g)(n"},
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
