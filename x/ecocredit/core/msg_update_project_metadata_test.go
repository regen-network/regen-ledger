package core

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/tendermint/tendermint/libs/rand"
	"gotest.tools/v3/assert"

	"github.com/regen-network/regen-ledger/x/ecocredit"
)

func TestMsgUpdateProjectMetadata_ValidateBasic(t *testing.T) {
	addr := sdk.AccAddress("addr1").String()
	type fields struct {
		Admin       string
		NewMetadata string
		ProjectId   string
	}
	tests := []struct {
		name   string
		fields fields
		errMsg string
	}{
		{
			name: "valid",
			fields: fields{
				Admin:       addr,
				NewMetadata: "new meta data",
				ProjectId:   "C01-001",
			},
		},
		{
			name: "invalid admin addr",
			fields: fields{
				Admin: "foo",
			},
			errMsg: sdkerrors.ErrInvalidAddress.Error(),
		},
		{
			name: "metadata too long",
			fields: fields{
				Admin:       addr,
				NewMetadata: rand.Str(MaxMetadataLength + 1),
			},
			errMsg: ecocredit.ErrMaxLimit.Error(),
		},
		{
			name: "invalid project id",
			fields: fields{
				Admin:       addr,
				NewMetadata: "new metadata",
				ProjectId:   "001",
			},
			errMsg: "invalid project id",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := MsgUpdateProjectMetadata{
				Admin:       tt.fields.Admin,
				NewMetadata: tt.fields.NewMetadata,
				ProjectId:   tt.fields.ProjectId,
			}
			if len(tt.errMsg) == 0 {
				assert.NilError(t, m.ValidateBasic())
			} else {
				assert.ErrorContains(t, m.ValidateBasic(), tt.errMsg)
			}
		})
	}
}
