package types

type ModuleID struct {
	// Name of a module used by router, used when registering a MsgServer and invoking
	// gRPC services.
	Name string
	// Address of a module
	Address []byte
}
