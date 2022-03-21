package marketplace

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func TestMsgBuy(t *testing.T) {
	t.Parallel()

	_, _, a1 := testdata.KeyTestPubAddr()

	validExpiration := time.Date(2030, 01, 01, 0, 0, 0, 0, time.UTC)

	tests := map[string]struct {
		src    MsgBuy
		expErr bool
	}{
		"valid": {
			src: MsgBuy{
				Buyer: a1.String(),
				Orders: []*MsgBuy_Order{
					{
						Quantity: "1.5",
						BidPrice: &sdk.Coin{
							Denom:  "uregen",
							Amount: sdk.NewInt(20),
						},
						DisableAutoRetire:  false,
						DisablePartialFill: true,
						Expiration:         &validExpiration,
						RetirementLocation: "US-WA",
					},
				},
			},
			expErr: false,
		},
		"invalid: bad owner address": {
			src: MsgBuy{
				Buyer: "foobar",
				Orders: []*MsgBuy_Order{
					{
						Quantity: "1.5",
						BidPrice: &sdk.Coin{
							Denom:  "uregen",
							Amount: sdk.NewInt(20),
						},
						DisableAutoRetire:  true,
						DisablePartialFill: true,
					},
				},
			},
			expErr: true,
		},
		"invalid: bad quantity": {
			src: MsgBuy{
				Buyer: a1.String(),
				Orders: []*MsgBuy_Order{
					{
						Quantity: "-1.5",
						BidPrice: &sdk.Coin{
							Denom:  "uregen",
							Amount: sdk.NewInt(20),
						},
						DisableAutoRetire:  true,
						DisablePartialFill: true,
					},
				},
			},
			expErr: true,
		},
		"invalid: bad bid price": {
			src: MsgBuy{
				Buyer: a1.String(),
				Orders: []*MsgBuy_Order{
					{
						Quantity: "1.5",
						BidPrice: &sdk.Coin{
							Denom:  "uregen",
							Amount: sdk.NewInt(-20),
						},
						DisableAutoRetire:  true,
						DisablePartialFill: true,
					},
				},
			},
			expErr: true,
		},
		"invalid: bad retirement location": {
			src: MsgBuy{
				Buyer: a1.String(),
				Orders: []*MsgBuy_Order{
					{
						Quantity: "1.5",
						BidPrice: &sdk.Coin{
							Denom:  "uregen",
							Amount: sdk.NewInt(20),
						},
						DisableAutoRetire:  false,
						DisablePartialFill: true,
						RetirementLocation: "foo",
					},
				},
			},
			expErr: true,
		},
		"invalid: retirement location required when DisableAutoRetire is false": {
			src: MsgBuy{
				Buyer: a1.String(),
				Orders: []*MsgBuy_Order{
					{
						Quantity: "1.5",
						BidPrice: &sdk.Coin{
							Denom:  "uregen",
							Amount: sdk.NewInt(20),
						},
						DisableAutoRetire:  false,
						DisablePartialFill: true,
						RetirementLocation: "",
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
