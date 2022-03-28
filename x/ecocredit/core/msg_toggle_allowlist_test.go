package core

import (
	"testing"

	"gotest.tools/v3/assert"

	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func TestMsgToggleAllowList_ValidateBasic(t *testing.T) {
	type fields struct {
		RootAddress string
		Toggle      bool
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
				RootAddress: addr.String(),
				Toggle:      false,
			},
		},
		{
			name: "invalid",
			fields: fields{
				RootAddress: "foo",
				Toggle:      false,
			},
			errMsg: sdkerrors.ErrInvalidAddress.Error(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := MsgToggleAllowList{
				RootAddress: tt.fields.RootAddress,
				Toggle:      tt.fields.Toggle,
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
