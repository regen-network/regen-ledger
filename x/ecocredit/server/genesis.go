package server

import (
	"encoding/json"

	"github.com/cosmos/cosmos-sdk/codec"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/x/ecocredit"
)


// InitGenesis performs genesis initialization for the ecocredit module. It
// returns no validator updates.
func (s serverImpl) InitGenesis(ctx types.Context, cdc codec.JSONCodec, data json.RawMessage) ([]abci.ValidatorUpdate, error) {
	var genesisState ecocredit.GenesisState
	cdc.MustUnmarshalJSON(data, &genesisState)
	s.paramSpace.SetParamSet(ctx.Context, &genesisState.Params)
	return []abci.ValidatorUpdate{}, nil
}

func (s serverImpl) ExportGenesis(ctx types.Context, cdc codec.JSONCodec) (json.RawMessage, error) {
	// Get Params from the store and put them in the genesis state
	var params ecocredit.Params
	s.paramSpace.GetParamSet(ctx.Context, &params)

	gs := &ecocredit.GenesisState{
		Params: params,
	}
	return cdc.MustMarshalJSON(gs), nil
}
