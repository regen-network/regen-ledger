package types

// ModuleAcc represents an account managed by a module
type ModuleAcc struct {
	// Module name managing this account. Used by router, used when registering a MsgServer
	// and invoking service methods.
	Module string
	// Key for identifying cross-module messages and module address derivation
	// TODO: probably we don't need it?
	Key []byte
	// Address derived from other module account
	Address []byte
}
