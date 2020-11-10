package server

import (
	"context"
	"github.com/regen-network/regen-ledger/types/module"
	"google.golang.org/grpc"
)

type ModuleKey interface {
	grpc.ClientConnInterface

	ModuleID() module.ModuleID
	Address() []byte
}

type CallInfo struct {
	Method string
	Caller module.ModuleID
}

type Invoker func(callInfo CallInfo) func(ctx context.Context, request, response interface{}, opts ...interface{}) error
