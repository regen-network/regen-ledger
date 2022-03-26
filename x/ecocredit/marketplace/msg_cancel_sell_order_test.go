package marketplace

import (
	"testing"

	"gotest.tools/v3/assert"

	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func TestMsgCancelSellOrder_ValidateBasic(t *testing.T) {
	type fields struct {
		Seller      string
		SellOrderId uint64
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
				Seller:      addr.String(),
				SellOrderId: 1,
			},
		},
		{
			name: "bad seller",
			fields: fields{
				Seller:      "foo",
				SellOrderId: 1,
			},
			errMsg: sdkerrors.ErrInvalidAddress.Error(),
		},
		{
			name: "bad sell order id",
			fields: fields{
				Seller:      addr.String(),
				SellOrderId: 0,
			},
			errMsg: "0 is not a valid sell order id",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MsgCancelSellOrder{
				Seller:      tt.fields.Seller,
				SellOrderId: tt.fields.SellOrderId,
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
