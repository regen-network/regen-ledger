package basket

import (
	"fmt"
	"testing"

	"gotest.tools/v3/assert"

	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func TestMsgUpdateBasketFeeRequest_ValidateBasic(t *testing.T) {
	type fields struct {
		RootAddress  string
		AddFees      []*MsgUpdateBasketFeeRequest_Fee
		RemoveDenoms []string
	}
	_, _, validAddr := testdata.KeyTestPubAddr()
	tests := []struct {
		name   string
		fields fields
		errMsg string
	}{
		{name: "valid",
			fields: fields{
				RootAddress:  validAddr.String(),
				AddFees:      []*MsgUpdateBasketFeeRequest_Fee{{Denom: "foo", Amount: "10"}},
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
				AddFees:      []*MsgUpdateBasketFeeRequest_Fee{{Denom: "al;skdjg;l"}},
				RemoveDenoms: nil,
			},
			errMsg: "invalid denom",
		},
		{
			name: "no bad AddFees amount",
			fields: fields{
				RootAddress:  validAddr.String(),
				AddFees:      []*MsgUpdateBasketFeeRequest_Fee{{Denom: "foo", Amount: "19.fee"}},
				RemoveDenoms: nil,
			},
			errMsg: fmt.Sprintf("could not convert 19.fee to %T", sdk.Int{}),
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
			m := &MsgUpdateBasketFeeRequest{
				RootAddress:  tt.fields.RootAddress,
				AddFees:      tt.fields.AddFees,
				RemoveDenoms: tt.fields.RemoveDenoms,
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
