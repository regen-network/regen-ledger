package server

import (
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"

	"github.com/regen-network/regen-ledger/x/ecocredit/simulation"
)

// WeightedOperations returns all the ecocredit module operations with their respective weights.
func (s serverImpl) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {

	// interfaceRegistry := types.NewInterfaceRegistry()
	// queryClient := ecocredit.NewQueryClient(s.key)
	return simulation.WeightedOperations(
		simState.AppParams, simState.Cdc,
		s.accountKeeper, s.bankKeeper,
	)
}
