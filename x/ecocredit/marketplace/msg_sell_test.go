package marketplace

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func TestMsgSell(t *testing.T) {
	t.Parallel()

	_, _, a1 := testdata.KeyTestPubAddr()

	validExpiration := time.Date(2030, 01, 01, 0, 0, 0, 0, time.UTC)

	tests := map[string]struct {
		src    MsgSell
		expErr bool
	}{
		"valid": {
			src: MsgSell{
				Owner: a1.String(),
				Orders: []*MsgSell_Order{
					{
						BatchDenom: "A00-00000000-00000000-000",
						Quantity:   "1.5",
						AskPrice: &sdk.Coin{
							Denom:  "uregen",
							Amount: sdk.NewInt(20),
						},
						DisableAutoRetire: true,
						Expiration:        &validExpiration,
					},
				},
			},
			expErr: false,
		},
		"invalid: bad owner address": {
			src: MsgSell{
				Owner: "foobar",
				Orders: []*MsgSell_Order{
					{
						BatchDenom: "A00-00000000-00000000-000",
						Quantity:   "1.5",
						AskPrice: &sdk.Coin{
							Denom:  "uregen",
							Amount: sdk.NewInt(20),
						},
						DisableAutoRetire: true,
					},
				},
			},
			expErr: true,
		},
		"invalid: bad batch denom": {
			src: MsgSell{
				Owner: a1.String(),
				Orders: []*MsgSell_Order{
					{
						BatchDenom: "foobar",
						Quantity:   "1.5",
						AskPrice: &sdk.Coin{
							Denom:  "uregen",
							Amount: sdk.NewInt(20),
						},
						DisableAutoRetire: true,
					},
				},
			},
			expErr: true,
		},
		"invalid: bad quantity": {
			src: MsgSell{
				Owner: a1.String(),
				Orders: []*MsgSell_Order{
					{
						BatchDenom: "A00-00000000-00000000-000",
						Quantity:   "-1.5",
						AskPrice: &sdk.Coin{
							Denom:  "uregen",
							Amount: sdk.NewInt(20),
						},
						DisableAutoRetire: true,
					},
				},
			},
			expErr: true,
		},
		"invalid: bad ask price": {
			src: MsgSell{
				Owner: a1.String(),
				Orders: []*MsgSell_Order{
					{
						BatchDenom: "A00-00000000-00000000-000",
						Quantity:   "1.5",
						AskPrice: &sdk.Coin{
							Denom:  "uregen",
							Amount: sdk.NewInt(-20),
						},
						DisableAutoRetire: true,
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
