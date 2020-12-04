package types

type ModuleID struct {
	ModuleName string
	Path       []byte
}

func (m ModuleID) Address() []byte {
	return AddressHash(m.ModuleName, m.Path)
}

func RootModuleID(moduleName string) ModuleID {
	return ModuleID{ModuleName: moduleName}
}
