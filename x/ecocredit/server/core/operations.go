package core

import (
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

// WeightedOperations returns all the ecocredit module operations with their respective weights.
// TODO: sim refactor PR https://github.com/regen-network/regen-ledger/issues/728
func (s serverImpl) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	//key := s.storeKey.(servermodule.RootModuleKey)
	//queryClient := ecocredit.NewQueryClient(key)
	//
	//return simulation.WeightedOperations(
	//	simState.AppParams, simState.Cdc,
	//	s.accountKeeper, s.bankKeeper,
	//	queryClient,
	//)
	return nil
}
