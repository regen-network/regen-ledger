package core

import (
	"testing"

	"gotest.tools/v3/assert"

	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func TestMsgUpdateAllowedCreditClassCreators_ValidateBasic(t *testing.T) {
	type fields struct {
		RootAddress    string
		AddCreators    []string
		RemoveCreators []string
	}
	_, _, addr := testdata.KeyTestPubAddr()
	tests := []struct {
		name   string
		fields fields
		errMsg string
	}{
		{
			name: "valid",
			fields: fields{
				RootAddress:    addr.String(),
				AddCreators:    []string{addr.String()},
				RemoveCreators: []string{addr.String()},
			},
		},
		{
			name: "invalid root addr",
			fields: fields{
				RootAddress:    "foobar",
				AddCreators:    []string{addr.String()},
				RemoveCreators: []string{addr.String()},
			},
			errMsg: sdkerrors.ErrInvalidAddress.Error(),
		},
		{
			name: "none specified",
			fields: fields{
				RootAddress: addr.String(),
			},
			errMsg: "must specify creators to add and/or remove",
		},
		{
			name: "invalid add creator",
			fields: fields{
				RootAddress: addr.String(),
				AddCreators: []string{"foo"},
			},
			errMsg: sdkerrors.ErrInvalidAddress.Error(),
		},
		{
			name: "invalid remove creator",
			fields: fields{
				RootAddress:    addr.String(),
				RemoveCreators: []string{"foo"},
			},
			errMsg: sdkerrors.ErrInvalidAddress.Error(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := MsgUpdateAllowedCreditClassCreators{
				RootAddress:    tt.fields.RootAddress,
				AddCreators:    tt.fields.AddCreators,
				RemoveCreators: tt.fields.RemoveCreators,
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
