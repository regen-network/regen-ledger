package module

import (
	"github.com/cosmos/cosmos-sdk/codec/types"
	regentypes "github.com/regen-network/regen-ledger/types"
)

type ModuleBase interface {
	RegisterTypes(registry types.InterfaceRegistry)
}

type ModuleMap map[string]ModuleBase

type ModuleID struct {
	ModuleName string
	Path       []byte
}

func (m ModuleID) Address() []byte {
	return regentypes.AddressHash(m.ModuleName, m.Path)
}
