package module

import (
	"github.com/cosmos/cosmos-sdk/codec/types"
)

type ModuleBase interface {
	Name() string
	RegisterInterfaces(types.InterfaceRegistry)
}
