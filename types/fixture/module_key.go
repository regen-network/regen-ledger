package fixture

import (
	"github.com/regen-network/regen-ledger/types"
)

type InvokerFactory func(callInfo CallInfo) (types.Invoker, error)

type CallInfo struct {
	Method string
	Caller types.ModuleID
}
