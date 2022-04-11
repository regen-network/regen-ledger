package marketplace

import (
	"testing"

	"gotest.tools/v3/assert"

	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/regen-network/regen-ledger/types/math"
)

func TestMsgBuyDirect_ValidateBasic(t *testing.T) {
	type fields struct {
		Buyer              string
		SellOrderId        uint64
		Quantity           string
		PricePerCredit     *sdk.Coin
		DisableAutoRetire  bool
		RetirementLocation string
	}
	validCoin := sdk.NewInt64Coin("ufoo", 31)
	_, _, addr := testdata.KeyTestPubAddr()
	tests := []struct {
		name   string
		fields fields
		errMsg string
	}{
		{
			name: "valid",
			fields: fields{
				Buyer:             addr.String(),
				SellOrderId:       1,
				Quantity:          "45.32",
				PricePerCredit:    &validCoin,
				DisableAutoRetire: true,
			},
		},
		{
			name: "valid retirement location",
			fields: fields{
				Buyer:              addr.String(),
				SellOrderId:        1,
				Quantity:           "45",
				PricePerCredit:     &validCoin,
				DisableAutoRetire:  false,
				RetirementLocation: "US-NY"},
		},
		{
			name:   "invalid addr",
			fields: fields{Buyer: "foobar"},
			errMsg: sdkerrors.ErrInvalidAddress.Error(),
		},
		{
			name: "invalid order id",
			fields: fields{
				Buyer:       addr.String(),
				SellOrderId: 0},
			errMsg: "0 is not a valid sell order id",
		},
		{
			name: "invalid quantity",
			fields: fields{
				Buyer:       addr.String(),
				SellOrderId: 1,
				Quantity:    "45.3xyz"},
			errMsg: math.ErrInvalidDecString.Error(),
		},
		{
			name: "no price per credit",
			fields: fields{
				Buyer:             addr.String(),
				SellOrderId:       1,
				Quantity:          "45",
				DisableAutoRetire: true,
				PricePerCredit:    nil},
			errMsg: "must specify price per credit",
		},
		{
			name: "invalid coin",
			fields: fields{
				Buyer:             addr.String(),
				SellOrderId:       1,
				Quantity:          "45",
				DisableAutoRetire: true,
				PricePerCredit:    &sdk.Coin{Denom: "foo3=21.", Amount: sdk.NewInt(3)}},
			errMsg: "invalid denom",
		},
		{
			name: "no retirement location when AutoRetiring",
			fields: fields{
				Buyer:             addr.String(),
				SellOrderId:       1,
				Quantity:          "45",
				PricePerCredit:    &validCoin,
				DisableAutoRetire: false},
			errMsg: "when DisableAutoRetire is false, a valid retirement location must be provided",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := MsgBuyDirect{
				Buyer: tt.fields.Buyer,
				Orders: []*MsgBuyDirect_Order{
					{
						SellOrderId:        tt.fields.SellOrderId,
						Quantity:           tt.fields.Quantity,
						BidPrice:           tt.fields.PricePerCredit,
						DisableAutoRetire:  tt.fields.DisableAutoRetire,
						RetirementLocation: tt.fields.RetirementLocation,
					},
				},
			}
			err := m.ValidateBasic()
			if len(tt.errMsg) == 0 {
				assert.NilError(t, err)
			} else {
				assert.Check(t, err != nil)
				assert.ErrorContains(t, err, tt.errMsg)
			}
		})
	}
}
