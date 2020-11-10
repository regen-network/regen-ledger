package server

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdkmodule "github.com/cosmos/cosmos-sdk/types/module"
	"github.com/regen-network/regen-ledger/types/module"
)

type Module interface {
	module.ModuleBase

	RegisterServices(Configurator)
}

type Configurator interface {
	sdkmodule.Configurator

	ModuleKey() ModuleKey
	Marshaler() codec.Marshaler
}
