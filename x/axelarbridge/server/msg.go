package server

import (
	"context"

	api "github.com/regen-network/regen-ledger/api/axelar/bridge/v1"
	"github.com/regen-network/regen-ledger/x/axelarbridge"
)

func (s serverImpl) ExecBridgeEvent(ctx context.Context, req *axelarbridge.MsgExecBridgeEvent) (*axelarbridge.MsgExecBridgeEventResponse, error) {
	panic("todo")
}

func (s serverImpl) RecordBridgeEvent(ctx context.Context, req *axelarbridge.MsgRecordBridgeEvent) (*axelarbridge.MsgRecordBridgeEventResponse, error) {
	id, err := s.stateStore.EventTable().InsertReturningID(ctx, &api.Event{
		SenderAddress: req.SenderAddress,
		DestAddress:   req.DestAddress,
		Payload:       req.Payload,
	})
	if err != nil {
		return nil, err
	}
	return &axelarbridge.MsgRecordBridgeEventResponse{id}, nil

}

func (s serverImpl) SendBridgeEvent(ctx context.Context, req *axelarbridge.MsgSendBridgeEvent) (*axelarbridge.MsgSendBridgeEventResponse, error) {
	panic("todo")
}
