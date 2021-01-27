package server

import (
	"github.com/regen-network/regen-ledger/types"
)

type ModuleKey interface {
	types.InvokerConn

	ModuleID() types.ModuleID
	Address() []byte
}

type InvokerFactory func(callInfo CallInfo) (types.Invoker, error)

type CallInfo struct {
	Method string
	Caller types.ModuleID
}
