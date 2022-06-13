package core

import (
	"strings"
	"testing"
	"time"

	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgBridgeReceive_ValidateBasic(t *testing.T) {

	validStart, validEnd := time.Now(), time.Date(2022, time.June, 25, 0, 0, 0, 0, time.Local)
	validOriginTx := OriginTx{
		Id:     "0x12345",
		Source: "polygon",
	}
	_, _, accAddr := testdata.KeyTestPubAddr()
	addr := accAddr.String()

	validMsg := MsgBridgeReceive{
		ServiceAddress:      addr,
		Recipient:           addr,
		Amount:              "10.5",
		OriginTx:            &validOriginTx,
		ProjectRefId:        "toucan",
		ProjectJurisdiction: "US",
		StartDate:           &validStart,
		EndDate:             &validEnd,
		ProjectMetadata:     "some metadata",
		BatchMetadata:       "some metadata",
		Note:                "some note",
		ClassId:             "C01",
	}

	tests := []struct {
		name   string
		getMsg func(validMsg MsgBridgeReceive) MsgBridgeReceive
		errMsg string
	}{
		{
			name: "valid case",
			getMsg: func(validMsg MsgBridgeReceive) MsgBridgeReceive {
				return validMsg
			},
		},
		{
			name: "invalid: service address",
			getMsg: func(validMsg MsgBridgeReceive) MsgBridgeReceive {
				validMsg.ServiceAddress = "0xfoobar"
				return validMsg
			},
			errMsg: sdkerrors.ErrInvalidAddress.Error(),
		},
		{
			name: "invalid: recipient",
			getMsg: func(validMsg MsgBridgeReceive) MsgBridgeReceive {
				validMsg.Recipient = "0xfoobar"
				return validMsg
			},
			errMsg: sdkerrors.ErrInvalidAddress.Error(),
		},
		{
			name: "invalid: decimal",
			getMsg: func(validMsg MsgBridgeReceive) MsgBridgeReceive {
				validMsg.Amount = "-1"
				return validMsg
			},
			errMsg: "expected a positive decimal",
		},
		{
			name: "invalid: nil origin tx",
			getMsg: func(validMsg MsgBridgeReceive) MsgBridgeReceive {
				validMsg.OriginTx = nil
				return validMsg
			},
			errMsg: sdkerrors.ErrInvalidRequest.Wrap("origin_tx is required").Error(),
		},
		{
			name: "invalid: empty origin tx",
			getMsg: func(validMsg MsgBridgeReceive) MsgBridgeReceive {
				validMsg.OriginTx = &OriginTx{}
				return validMsg
			},
			errMsg: sdkerrors.ErrInvalidRequest.Wrap("invalid OriginTx: no id").Error(),
		},
		{
			name: "invalid: no project reference id",
			getMsg: func(validMsg MsgBridgeReceive) MsgBridgeReceive {
				validMsg.ProjectRefId = ""
				return validMsg
			},
			errMsg: sdkerrors.ErrInvalidRequest.Wrap("project_ref_id is required").Error(),
		},
		{
			name: "invalid: jurisdiction",
			getMsg: func(validMsg MsgBridgeReceive) MsgBridgeReceive {
				validMsg.ProjectJurisdiction = ""
				return validMsg
			},
			errMsg: "invalid jurisdiction",
		},
		{
			name: "invalid: no start date",
			getMsg: func(validMsg MsgBridgeReceive) MsgBridgeReceive {
				validMsg.StartDate = nil
				return validMsg
			},
			errMsg: sdkerrors.ErrInvalidRequest.Wrap("start_date is required").Error(),
		},
		{
			name: "invalid: no end date",
			getMsg: func(validMsg MsgBridgeReceive) MsgBridgeReceive {
				validMsg.EndDate = nil
				return validMsg
			},
			errMsg: sdkerrors.ErrInvalidRequest.Wrap("end_date is required").Error(),
		},
		{
			name: "invalid: start date after end date",
			getMsg: func(validMsg MsgBridgeReceive) MsgBridgeReceive {
				before := time.Date(2020, 0, 0, 0, 0, 0, 0, time.Local)
				after := time.Date(2021, 0, 0, 0, 0, 0, 0, time.Local)
				validMsg.StartDate = &after
				validMsg.EndDate = &before
				return validMsg
			},
			errMsg: sdkerrors.ErrInvalidRequest.Wrap("start_date must be a time before end_date").Error(),
		},
		{
			name: "invalid: project metadata length",
			getMsg: func(validMsg MsgBridgeReceive) MsgBridgeReceive {
				validMsg.ProjectMetadata = strings.Repeat("X", MaxMetadataLength+1)
				return validMsg
			},
			errMsg: sdkerrors.ErrInvalidRequest.Wrapf("project_metadata length (%d) exceeds max metadata length: %d", MaxMetadataLength+1, MaxMetadataLength).Error(),
		},
		{
			name: "invalid: batch metadata length",
			getMsg: func(validMsg MsgBridgeReceive) MsgBridgeReceive {
				validMsg.BatchMetadata = strings.Repeat("X", MaxMetadataLength+1)
				return validMsg
			},
			errMsg: sdkerrors.ErrInvalidRequest.Wrapf("batch_metadata length (%d) exceeds max metadata length: %d", MaxMetadataLength+1, MaxMetadataLength).Error(),
		},
		{
			name: "invalid: note length",
			getMsg: func(validMsg MsgBridgeReceive) MsgBridgeReceive {
				validMsg.Note = strings.Repeat("X", MaxMetadataLength+1)
				return validMsg
			},
			errMsg: sdkerrors.ErrInvalidRequest.Wrapf("note length (%d) exceeds max length: %d", MaxMetadataLength+1, MaxMetadataLength).Error(),
		},
		{
			name: "invalid: class Id",
			getMsg: func(validMsg MsgBridgeReceive) MsgBridgeReceive {
				validMsg.ClassId = "Foobar12345"
				return validMsg
			},
			errMsg: "class ID didn't match the format",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			msg := tt.getMsg(validMsg)
			err := msg.ValidateBasic()

			if len(tt.errMsg) == 0 {
				require.NoError(t, err)
			} else {
				require.ErrorContains(t, err, tt.errMsg)
			}
		})
	}
}
