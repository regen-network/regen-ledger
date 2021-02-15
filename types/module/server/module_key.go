package server

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/regen-network/regen-ledger/types"
)

type ModuleKey interface {
	types.InvokerConn

	ModuleID() types.ModuleID
	Address() sdk.AccAddress
}

type InvokerFactory func(callInfo CallInfo) (types.Invoker, error)

type CallInfo struct {
	Method string
	Caller types.ModuleID
}
