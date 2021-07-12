package types

// ModuleAcc represents an account managed by a module
type ModuleAcc struct {
	// Module name managing this account. Used by router, used when registering a MsgServer
	// and invoking service methods.
	Module string
	// Address of the module account generated using a derivation mechanism.
	Address []byte
}
