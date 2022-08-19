package server

import (
	"encoding/json"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/tendermint/tendermint/abci/types"
)

var _ Keeper = serverImpl{}

// Keeper defines a set of methods the ecocredit module exposes.
type Keeper interface {
	PruneOrders(ctx sdk.Context) error
	RegisterInvariants(sdk.InvariantRegistry)
	InitGenesis(ctx sdk.Context, cdc codec.JSONCodec, data json.RawMessage) ([]types.ValidatorUpdate, error)
	ExportGenesis(ctx sdk.Context, cdc codec.JSONCodec) (json.RawMessage, error)
	WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation
}
