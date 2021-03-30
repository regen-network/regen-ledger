package app

import (
	"encoding/json"
	"io/ioutil"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/types/simulation"
)

// SimulationOperations retrieves the simulation params from the provided file path
// and returns all the modules weighted operations
func SimulationOperations(app *RegenApp, cdc codec.JSONMarshaler, config simulation.Config) []simulation.WeightedOperation {
	simState := module.SimulationState{
		AppParams: make(simulation.AppParams),
		Cdc:       cdc,
	}

	if config.ParamsFile != "" {
		bz, err := ioutil.ReadFile(config.ParamsFile)
		if err != nil {
			panic(err)
		}

		err = json.Unmarshal(bz, &simState.AppParams)
		if err != nil {
			panic(err)
		}
	}

	simState.ParamChanges = app.SimulationManager().GenerateParamChanges(config.Seed)
	simState.Contents = app.SimulationManager().GetProposalContents(simState)
	return app.smm.WeightedOperations(simState, app.sm.Modules)
}
