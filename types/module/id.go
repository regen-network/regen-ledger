package module

import "github.com/regen-network/regen-ledger/types"

type ModuleID struct {
	ModuleName string
	Path       []byte
}

func (m ModuleID) Address() []byte {
	return types.AddressHash(m.ModuleName, m.Path)
}
