package server

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	api "github.com/regen-network/regen-ledger/api/axelar/bridge/v1"
	"github.com/regen-network/regen-ledger/x/axelarbridge"
)

func (s serverImpl) RecordBridgeEvent(ctx context.Context, req *axelarbridge.MsgRecordBridgeEvent) (*axelarbridge.MsgRecordBridgeEventResponse, error) {
	if _, ok := s.handlers[req.Handler]; !ok {
		return nil, sdkerrors.ErrInvalidRequest.Wrap("Undefined handler: " + req.Handler)
	}
	id, err := s.stateStore.EventTable().InsertReturningID(ctx, &api.Event{
		SrcChain: req.SrcChain,
		SrcTxId:  req.SrcTxId,
		Sender:   req.Sender,
		Handler:  req.Handler,
		Payload:  req.Payload,
	})
	if err != nil {
		return nil, err
	}
	return &axelarbridge.MsgRecordBridgeEventResponse{id}, nil

}

func (s serverImpl) SendBridgeEvent(ctx context.Context, req *axelarbridge.MsgSendBridgeEvent) (*axelarbridge.MsgSendBridgeEventResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	err := sdkCtx.EventManager().EmitTypedEvent(&axelarbridge.SendBridgeEvent{
		Sender:  req.Sender,
		Handler: req.Handler,
		Payload: req.Payload,
	})
	return &axelarbridge.MsgSendBridgeEventResponse{}, err
}
