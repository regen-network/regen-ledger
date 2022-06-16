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
	t.Parallel()
	validStart, validEnd := time.Now(), time.Date(2022, time.June, 25, 0, 0, 0, 0, time.Local)
	validOriginTx := OriginTx{
		Id:     "0x12345",
		Source: "polygon",
	}
	_, _, accAddr := testdata.KeyTestPubAddr()
	addr := accAddr.String()

	validBatch := MsgBridgeReceive_Batch{
		Recipient: addr,
		Amount:    "10.5",
		StartDate: &validStart,
		EndDate:   &validEnd,
		Metadata:  "some metadata",
	}
	validProject := MsgBridgeReceive_Project{
		ReferenceId:  "VCS-001",
		Jurisdiction: "US-KY",
		Metadata:     "some metadata",
	}
	validMsg := MsgBridgeReceive{
		Issuer:   addr,
		Batch:    &validBatch,
		Project:  &validProject,
		OriginTx: &validOriginTx,
		ClassId:  "C01",
		Note:     "some note",
	}

	resetMsg := func() MsgBridgeReceive {
		msg := validMsg
		msg.Batch = &validBatch
		msg.Project = &validProject
		msg.OriginTx = &validOriginTx
		return msg
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
				validMsg.Issuer = "0xfoobar"
				return validMsg
			},
			errMsg: sdkerrors.ErrInvalidAddress.Error(),
		},
		{
			name: "invalid: null batch",
			getMsg: func(validMsg MsgBridgeReceive) MsgBridgeReceive {
				validMsg.Batch = nil
				return validMsg
			},
			errMsg: sdkerrors.ErrInvalidRequest.Wrapf("batch cannot be empty").Error(),
		},
		{
			name: "invalid: recipient",
			getMsg: func(validMsg MsgBridgeReceive) MsgBridgeReceive {
				validMsg.Batch.Recipient = "0xfoobar"
				return validMsg
			},
			errMsg: sdkerrors.ErrInvalidAddress.Error(),
		},
		{
			name: "invalid: decimal",
			getMsg: func(validMsg MsgBridgeReceive) MsgBridgeReceive {
				validMsg.Batch.Amount = "-1"
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
			name: "invalid: null project",
			getMsg: func(validMsg MsgBridgeReceive) MsgBridgeReceive {
				validMsg.Project = nil
				return validMsg
			},
			errMsg: sdkerrors.ErrInvalidRequest.Wrapf("project cannot be empty").Error(),
		},
		{
			name: "invalid: no project reference id",
			getMsg: func(validMsg MsgBridgeReceive) MsgBridgeReceive {
				validMsg.Project.ReferenceId = ""
				return validMsg
			},
			errMsg: sdkerrors.ErrInvalidRequest.Wrap("project_ref_id is required").Error(),
		},
		{
			name: "invalid: jurisdiction",
			getMsg: func(validMsg MsgBridgeReceive) MsgBridgeReceive {
				validMsg.Project.Jurisdiction = ""
				return validMsg
			},
			errMsg: "invalid jurisdiction",
		},
		{
			name: "invalid: no start date",
			getMsg: func(validMsg MsgBridgeReceive) MsgBridgeReceive {
				validMsg.Batch.StartDate = nil
				return validMsg
			},
			errMsg: sdkerrors.ErrInvalidRequest.Wrap("start_date is required").Error(),
		},
		{
			name: "invalid: no end date",
			getMsg: func(validMsg MsgBridgeReceive) MsgBridgeReceive {
				validMsg.Batch.EndDate = nil
				return validMsg
			},
			errMsg: sdkerrors.ErrInvalidRequest.Wrap("end_date is required").Error(),
		},
		{
			name: "invalid: start date after end date",
			getMsg: func(validMsg MsgBridgeReceive) MsgBridgeReceive {
				before := time.Date(2020, 0, 0, 0, 0, 0, 0, time.Local)
				after := time.Date(2021, 0, 0, 0, 0, 0, 0, time.Local)
				validMsg.Batch.StartDate = &after
				validMsg.Batch.EndDate = &before
				return validMsg
			},
			errMsg: sdkerrors.ErrInvalidRequest.Wrap("start_date must be a time before end_date").Error(),
		},
		{
			name: "invalid: project metadata length",
			getMsg: func(validMsg MsgBridgeReceive) MsgBridgeReceive {
				validMsg.Project.Metadata = strings.Repeat("X", MaxMetadataLength+1)
				return validMsg
			},
			errMsg: sdkerrors.ErrInvalidRequest.Wrapf("project_metadata length (%d) exceeds max metadata length: %d", MaxMetadataLength+1, MaxMetadataLength).Error(),
		},
		{
			name: "invalid: batch metadata length",
			getMsg: func(validMsg MsgBridgeReceive) MsgBridgeReceive {
				validMsg.Batch.Metadata = strings.Repeat("X", MaxMetadataLength+1)
				return validMsg
			},
			errMsg: sdkerrors.ErrInvalidRequest.Wrapf("batch_metadata length (%d) exceeds max metadata length: %d", MaxMetadataLength+1, MaxMetadataLength).Error(),
		},
		{
			name: "invalid: note length",
			getMsg: func(validMsg MsgBridgeReceive) MsgBridgeReceive {
				validMsg.Note = strings.Repeat("X", MaxNoteLength+1)
				return validMsg
			},
			errMsg: sdkerrors.ErrInvalidRequest.Wrapf("note length (%d) exceeds max length: %d", MaxNoteLength+1, MaxNoteLength).Error(),
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
			t.Parallel()
			// reset the valid message, cleaning the pointer values
			vMsg := resetMsg()
			msg := tt.getMsg(vMsg)
			err := msg.ValidateBasic()

			if len(tt.errMsg) == 0 {
				require.NoError(t, err)
			} else {
				require.ErrorContains(t, err, tt.errMsg)
			}
		})
	}
}
