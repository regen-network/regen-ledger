package server

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkmodule "github.com/cosmos/cosmos-sdk/types/module"

	"github.com/regen-network/regen-ledger/types/module"
)

// Module is the module type that all server modules must satisfy
type Module interface {
	module.TypeModule

	RegisterServices(Configurator)
}

type Configurator interface {
	sdkmodule.Configurator

	ModuleKey() RootModuleKey
	Marshaler() codec.Codec
	RequireServer(interface{})
	RegisterInvariantsHandler(registry RegisterInvariantsHandler)
	RegisterGenesisHandlers(module.InitGenesisHandler, module.ExportGenesisHandler)

	// Router() is temporarily added here to use in the group module.
	// TODO: remove once #225 addressed
	Router() sdk.Router

	RegisterWeightedOperationsHandler(WeightedOperationsHandler)
}

// LegacyRouteModule is the module type that a module must implement
// to support legacy sdk.Msg routing.
// This is currently used for the group module as part of #218.
type LegacyRouteModule interface {
	Route(Configurator) sdk.Route
}
