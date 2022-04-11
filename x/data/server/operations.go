package server

import (
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	servermodule "github.com/regen-network/regen-ledger/types/module/server"

	"github.com/regen-network/regen-ledger/x/data"
	"github.com/regen-network/regen-ledger/x/data/simulation"
)

// WeightedOperations returns all the ecocredit module operations with their respective weights.
func (s serverImpl) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	key := s.storeKey.(servermodule.RootModuleKey)
	queryClient := data.NewQueryClient(key)

	return simulation.WeightedOperations(
		simState.AppParams, simState.Cdc,
		s.accountKeeper, s.bankKeeper,
		queryClient,
	)
}
