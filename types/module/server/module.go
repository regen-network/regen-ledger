package server

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkmodule "github.com/cosmos/cosmos-sdk/types/module"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/regen-network/regen-ledger/types/module"
)

// Module is the module type that all server modules must satisfy
type Module interface {
	module.TypeModule

	RegisterServices(Configurator)
}

// BeginBlockerModule is a module exposing begin blocker for server module manager
type BeginBlockerModule interface {
	BeginBlock(sdk.Context, abci.RequestBeginBlock)
}

// EndBlockerModule is a module exposing end blocker for server module manager
type EndBlockerModule interface {
	EndBlock(sdk.Context, abci.RequestEndBlock) []abci.ValidatorUpdate
}

type Configurator interface {
	sdkmodule.Configurator

	ModuleKey() RootModuleKey
	Marshaler() codec.Codec
	RequireServer(interface{})
	RegisterInvariantsHandler(registry RegisterInvariantsHandler)
	RegisterGenesisHandlers(module.InitGenesisHandler, module.ExportGenesisHandler)
	RegisterWeightedOperationsHandler(WeightedOperationsHandler)
	RegisterMigrationHandler(MigrationHandler)
}
