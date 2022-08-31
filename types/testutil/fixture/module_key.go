package fixture

type InvokerFactory func(callInfo CallInfo) (Invoker, error)

type CallInfo struct {
	Method string
	Caller ModuleID
}
