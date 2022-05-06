package core

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"gotest.tools/v3/assert"
)

func TestMsgUpdateProjectAdmin_ValidateBasic(t *testing.T) {
	addr := sdk.AccAddress("addr1").String()
	addr2 := sdk.AccAddress("addr2").String()
	type fields struct {
		Admin     string
		NewAdmin  string
		ProjectId string
	}
	tests := []struct {
		name   string
		fields fields
		errMsg string
	}{
		{
			name: "valid",
			fields: fields{
				Admin:     addr,
				NewAdmin:  addr2,
				ProjectId: "C01-001",
			},
		},
		{
			name: "invalid admin",
			fields: fields{
				Admin: "foo",
			},
			errMsg: sdkerrors.ErrInvalidAddress.Error(),
		},
		{
			name: "invalid new admin",
			fields: fields{
				Admin:    addr,
				NewAdmin: "foo",
			},
			errMsg: sdkerrors.ErrInvalidAddress.Error(),
		},
		{
			name: "cannot have same address",
			fields: fields{
				Admin:    addr,
				NewAdmin: addr,
			},
			errMsg: sdkerrors.ErrInvalidRequest.Wrap("new_admin and admin addresses cannot be the same").Error(),
		},
		{
			name: "invalid project id",
			fields: fields{
				Admin:     addr,
				NewAdmin:  addr2,
				ProjectId: "001",
			},
			errMsg: "invalid project id",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := MsgUpdateProjectAdmin{
				Admin:     tt.fields.Admin,
				NewAdmin:  tt.fields.NewAdmin,
				ProjectId: tt.fields.ProjectId,
			}
			if len(tt.errMsg) == 0 {
				assert.NilError(t, m.ValidateBasic())
			} else {
				assert.ErrorContains(t, m.ValidateBasic(), tt.errMsg)
			}
		})
	}
}
