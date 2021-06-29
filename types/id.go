package types

// TODO: rename to ModuleAccount
type ModuleID struct {
	// Name of a module used by router, used when registering a MsgServer and invoking
	// service methods.
	Name string
	// Key for identifying cross-module messages and module address derivation
	Key []byte
}
