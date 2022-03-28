package core

import (
	"testing"

	"gotest.tools/v3/assert"

	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func TestMsgUpdateCreditClassFeeRequest_ValidateBasic(t *testing.T) {
	type fields struct {
		RootAddress  string
		AddFees      []*sdk.Coin
		RemoveDenoms []string
	}
	_, _, validAddr := testdata.KeyTestPubAddr()
	tests := []struct {
		name   string
		fields fields
		errMsg string
	}{
		{
			name: "valid",
			fields: fields{
				RootAddress:  validAddr.String(),
				AddFees:      []*sdk.Coin{{Denom: "foo", Amount: sdk.NewInt(10)}},
				RemoveDenoms: []string{"foo"},
			},
		},
		{
			name: "invalid addr",
			fields: fields{
				RootAddress:  "foo",
				AddFees:      []*sdk.Coin{{Denom: "foo", Amount: sdk.NewInt(10)}},
				RemoveDenoms: []string{"foo"},
			},
			errMsg: sdkerrors.ErrInvalidAddress.Error(),
		},
		{
			name: "none specified",
			fields: fields{
				RootAddress: validAddr.String(),
			},
			errMsg: "at least one of add_fees or remove_denoms must be specified",
		},
		{
			name: "bad add denom",
			fields: fields{
				RootAddress: validAddr.String(),
				AddFees:     []*sdk.Coin{{Denom: ".,zxcjvn", Amount: sdk.NewInt(10)}},
			},
			errMsg: "invalid denom",
		},
		{
			name: "bad remove denom",
			fields: fields{
				RootAddress:  validAddr.String(),
				RemoveDenoms: []string{"!32.10!"},
			},
			errMsg: "invalid denom",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MsgUpdateCreditClassFee{
				RootAddress:  tt.fields.RootAddress,
				AddFees:      tt.fields.AddFees,
				RemoveDenoms: tt.fields.RemoveDenoms,
			}
			err := m.ValidateBasic()
			if len(tt.errMsg) == 0 {
				assert.NilError(t, err)
			} else {
				assert.ErrorContains(t, err, tt.errMsg)
			}
		})
	}
}
