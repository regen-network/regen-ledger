package server

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
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

	handler, ok := s.handlers[event.Handler]
	if !ok {
		return nil, sdkerrors.ErrInvalidRequest.Wrap("Undefined handler: " + event.Handler)
	}
	if err := handler(ctx, *event); err != nil {
		return nil, err
	}

	err = s.stateStore.EventTable().Delete(ctx, event)
	if err != nil {
		return nil, err
	}

	return &axelarbridge.MsgExecBridgeEventResponse{}, nil
}
