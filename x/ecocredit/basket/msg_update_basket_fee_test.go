package basket

import (
	"testing"

	"gotest.tools/v3/assert"

	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func TestMsgUpdateBasketFeeRequest_ValidateBasic(t *testing.T) {
	type fields struct {
		RootAddress  string
		AddFees      []*sdk.Coin
		RemoveDenoms []string
	}
	validCoin := sdk.NewInt64Coin("foo", 10)
	_, _, validAddr := testdata.KeyTestPubAddr()
	tests := []struct {
		name   string
		fields fields
		errMsg string
	}{
		{name: "valid",
			fields: fields{
				RootAddress:  validAddr.String(),
				AddFees:      []*sdk.Coin{&validCoin},
				RemoveDenoms: []string{"bar"},
			},
		},
		{
			name: "bad addr",
			fields: fields{
				RootAddress:  "dlksadjg.xkcm.v",
				AddFees:      nil,
				RemoveDenoms: nil,
			},
			errMsg: sdkerrors.ErrInvalidAddress.Error(),
		},
		{
			name: "no fields specified",
			fields: fields{
				RootAddress:  validAddr.String(),
				AddFees:      nil,
				RemoveDenoms: nil,
			},
			errMsg: "at least one of add_fees or remove_denoms must be specified",
		},
		{
			name: "no bad AddFees denom",
			fields: fields{
				RootAddress:  validAddr.String(),
				AddFees:      []*sdk.Coin{{Denom: "al;skdjg;l", Amount: sdk.NewInt(10)}},
				RemoveDenoms: nil,
			},
			errMsg: "invalid denom",
		},
		{
			name: "no bad AddFees amount",
			fields: fields{
				RootAddress:  validAddr.String(),
				AddFees:      []*sdk.Coin{{Denom: "foo", Amount: sdk.NewInt(0)}},
				RemoveDenoms: nil,
			},
			errMsg: "fee must be greater than zero",
		},
		{
			name: "no bad RemoveDenoms denom",
			fields: fields{
				RootAddress:  validAddr.String(),
				AddFees:      nil,
				RemoveDenoms: []string{"x,,3290ikx.!"},
			},
			errMsg: "invalid denom",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MsgUpdateBasketFee{
				RootAddress: tt.fields.RootAddress,
				AddFees:     tt.fields.AddFees,
				RemoveFees:  tt.fields.RemoveDenoms,
			}
			err := m.ValidateBasic()
			if len(tt.errMsg) != 0 {
				assert.ErrorContains(t, err, tt.errMsg)
			} else {
				assert.NilError(t, err)
			}
		})
	}
}
