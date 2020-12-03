package module

import (
	"github.com/cosmos/cosmos-sdk/codec/types"
)

type Module interface {
	Name() string
	RegisterInterfaces(types.InterfaceRegistry)
}
