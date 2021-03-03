package simulation

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/regen-network/regen-ledger/types/module/server"
)

// AppSimulationManager defines a simulation manager that provides the high level utility
// for managing and executing simulation functionalities for a group of modules
type AppSimulationManager struct {
	Modules       []module.AppModuleSimulation // array of app modules; we use an array for deterministic simulation tests
	StoreDecoders sdk.StoreDecoderRegistry     // functions to decode the key-value pairs from each module's store
	manager       *server.Manager
}

// NewAppSimulationManager creates a new SimulationManager object
// CONTRACT: All the modules provided must be also registered on the module Manager
func NewAppSimulationManager(serverManager *server.Manager, modules ...module.AppModuleSimulation) *AppSimulationManager {
	return &AppSimulationManager{
		Modules:       modules,
		StoreDecoders: make(sdk.StoreDecoderRegistry),
		manager:       serverManager,
	}
}

// GetProposalContents returns each module's proposal content generator function
// with their default operation weight and key.
func (sm *AppSimulationManager) GetProposalContents(simState module.SimulationState) []simulation.WeightedProposalContent {
	wContents := make([]simulation.WeightedProposalContent, 0, len(sm.Modules))
	for _, module := range sm.Modules {
		wContents = append(wContents, module.ProposalContents(simState)...)
	}

	return wContents
}

// RegisterStoreDecoders registers each of the modules' store decoders into a map
func (sm *AppSimulationManager) RegisterStoreDecoders() {
	for _, module := range sm.Modules {
		module.RegisterStoreDecoder(sm.StoreDecoders)
	}
}

// GenerateGenesisStates generates a randomized GenesisState for each of the
// registered modules
func (sm *AppSimulationManager) GenerateGenesisStates(simState *module.SimulationState) {
	for _, module := range sm.Modules {
		module.GenerateGenesisState(simState)
	}
}

// GenerateParamChanges generates randomized contents for creating params change
// proposal transactions
func (sm *AppSimulationManager) GenerateParamChanges(seed int64) (paramChanges []simulation.ParamChange) {
	r := rand.New(rand.NewSource(seed))
	for _, module := range sm.Modules {
		paramChanges = append(paramChanges, module.RandomizedParams(r)...)
	}

	return
}

// WeightedOperations returns all the modules' weighted operations of an application
func (sm *AppSimulationManager) WeightedOperations(simState module.SimulationState) []simulation.WeightedOperation {
	wOps := make([]simulation.WeightedOperation, 0, len(sm.Modules))
	modules := sm.manager.GetWeightedOperationsHandlers()
	for _, module := range modules {
		wOps = append(wOps, module(simState)...)
	}

	return wOps
}
