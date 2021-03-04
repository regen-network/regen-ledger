package app

import (
	"encoding/json"
	"fmt"
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
	return app.nm.WeightedOperations(simState, app.sm.Modules)
}

// CheckExportSimulation exports the app state and simulation parameters to JSON
// if the export paths are defined.
func CheckExportSimulation(
	app *RegenApp, config simulation.Config, params simulation.Params,
) error {
	if config.ExportStatePath != "" {
		fmt.Println("exporting app state...")
		exported, err := app.ExportAppStateAndValidators(false, nil)
		if err != nil {
			return err
		}

		if err := ioutil.WriteFile(config.ExportStatePath, []byte(exported.AppState), 0600); err != nil {
			return err
		}
	}

	if config.ExportParamsPath != "" {
		fmt.Println("exporting simulation params...")
		paramsBz, err := json.MarshalIndent(params, "", " ")
		if err != nil {
			return err
		}

		if err := ioutil.WriteFile(config.ExportParamsPath, paramsBz, 0600); err != nil {
			return err
		}
	}
	return nil
}
