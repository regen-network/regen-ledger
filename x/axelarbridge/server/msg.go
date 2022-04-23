package server

import (
	"context"

	"github.com/regen-network/regen-ledger/x/axelarbridge"
)

func (s serverImpl) ExecBrigeEvent(ctx context.Context, req *axelarbridge.MsgExecBrigeEvent) (*axelarbridge.MsgExecBrigeEventResponse, error) {
	panic("todo")
}

func (s serverImpl) RecordBridgeEvent(ctx context.Context, req *axelarbridge.MsgRecordBridgeEvent) (*axelarbridge.MsgRecordBridgeEventResponse, error) {
	panic("todo")
}

func (s serverImpl) SendBridgeMessage(ctx context.Context, req *axelarbridge.MsgSendBridgeMessage) (*axelarbridge.MsgSendBridgeMessageResponse, error) {
	panic("todo")
}
