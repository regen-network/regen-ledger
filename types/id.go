package types

type ModuleID struct {
	ModuleName string
	Path       []byte
}

func (m ModuleID) Address() []byte {
	return AddressHash(m.ModuleName, m.Path)
}
