package server

import (
	"context"

	"github.com/regen-network/regen-ledger/x/axelarbridge"
)

func (s serverImpl) ExecBridgeEvent(ctx context.Context, req *axelarbridge.MsgExecBridgeEvent) (*axelarbridge.MsgExecBridgeEventResponse, error) {
	panic("todo")
}

func (s serverImpl) RecordBridgeEvent(ctx context.Context, req *axelarbridge.MsgRecordBridgeEvent) (*axelarbridge.MsgRecordBridgeEventResponse, error) {
	panic("todo")
}

func (s serverImpl) SendBridgeEvent(ctx context.Context, req *axelarbridge.MsgSendBridgeEvent) (*axelarbridge.MsgSendBridgeEventResponse, error) {
	panic("todo")
}
