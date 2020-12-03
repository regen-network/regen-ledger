package server

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	sdkmodule "github.com/cosmos/cosmos-sdk/types/module"
	"github.com/regen-network/regen-ledger/types/module"
)

type Module interface {
	module.Module

	RegisterInterfaces(registry types.InterfaceRegistry)
	RegisterServices(Configurator)
}

type Configurator interface {
	sdkmodule.Configurator

	ModuleKey() RootModuleKey
	BinaryMarshaler() codec.BinaryMarshaler
	RequireServer(interface{})
}
