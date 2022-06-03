package marketplace

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func TestMsgUpdateSellOrders(t *testing.T) {
	t.Parallel()

	_, _, a1 := testdata.KeyTestPubAddr()

	validExpiration := time.Date(2030, 01, 01, 0, 0, 0, 0, time.UTC)

	tests := map[string]struct {
		src    MsgUpdateSellOrders
		expErr bool
	}{
		"valid": {
			src: MsgUpdateSellOrders{
				Seller: a1.String(),
				Updates: []*MsgUpdateSellOrders_Update{
					{
						NewQuantity: "1.5",
						NewAskPrice: &sdk.Coin{
							Denom:  "uregen",
							Amount: sdk.NewInt(20),
						},
						DisableAutoRetire: true,
						NewExpiration:     &validExpiration,
					},
				},
			},
			expErr: false,
		},
		"invalid: bad seller address": {
			src: MsgUpdateSellOrders{
				Seller: "foobar",
				Updates: []*MsgUpdateSellOrders_Update{
					{
						NewQuantity: "1.5",
						NewAskPrice: &sdk.Coin{
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
			src: MsgUpdateSellOrders{
				Seller: a1.String(),
				Updates: []*MsgUpdateSellOrders_Update{
					{
						NewQuantity: "-1.5",
						NewAskPrice: &sdk.Coin{
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
			src: MsgUpdateSellOrders{
				Seller: a1.String(),
				Updates: []*MsgUpdateSellOrders_Update{
					{
						NewQuantity: "1.5",
						NewAskPrice: &sdk.Coin{
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
