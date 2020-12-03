package server

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/regen-network/regen-ledger/types"
	"google.golang.org/grpc"
)

type ModuleKey interface {
	grpc.ClientConnInterface

	ModuleID() types.ModuleID
	Address() sdk.AccAddress
}

type CallInfo struct {
	Method string
	Caller types.ModuleID
}

type InvokerFactory func(callInfo CallInfo) (Invoker, error)

type Invoker func(ctx context.Context, request, response interface{}, opts ...interface{}) error
