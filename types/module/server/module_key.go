package server

import (
	"context"

	"github.com/regen-network/regen-ledger/types"

	"google.golang.org/grpc"
)

type ModuleKey interface {
	grpc.ClientConnInterface

	ModuleID() types.ModuleID
	Address() []byte
}

type CallInfo struct {
	Method string
	Caller types.ModuleID
}

type InvokerFactory func(callInfo CallInfo) (Invoker, error)

type Invoker func(ctx context.Context, request, response interface{}, opts ...interface{}) error
