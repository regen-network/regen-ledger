package module

import (
	"github.com/cosmos/cosmos-sdk/codec/types"
	regentypes "github.com/regen-network/regen-ledger/types"
)

type ModuleBase interface {
	Name() string
	RegisterInterfaces(types.InterfaceRegistry)
}

type Modules []ModuleBase

func (mm Modules) RegisterInterfaces(registry types.InterfaceRegistry) {
	for _, m := range mm {
		m.RegisterInterfaces(registry)
	}
}

type ModuleID struct {
	ModuleName string
	Path       []byte
}

func (m ModuleID) Address() []byte {
	return regentypes.AddressHash(m.ModuleName, m.Path)
}
