package server

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/regen-network/regen-ledger/x/axelarbridge"
)

func (s serverImpl) ExecBridgeEvent(ctx context.Context, req *axelarbridge.MsgExecBridgeEvent) (*axelarbridge.MsgExecBridgeEventResponse, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}

	event, err := s.stateStore.EventTable().Get(ctx, req.EventId)
	if err != nil {
		return nil, err
	}

	// TODO handler

	err = s.stateStore.EventTable().Delete(ctx, event)
	if err != nil {
		return nil, err
	}

	return &axelarbridge.MsgExecBridgeEventResponse{}, nil
}
