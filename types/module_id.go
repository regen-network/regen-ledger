package types

type ModuleID interface {
	Address() []byte
}

type RootModuleID string

type DerivedModuleID struct {
	ModuleName string
	Path       []byte
}
