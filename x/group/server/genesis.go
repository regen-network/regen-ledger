package server

import (
	"encoding/json"

	"github.com/cosmos/cosmos-sdk/codec"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/x/group"
)

func (s serverImpl) InitGenesis(ctx types.Context, cdc codec.JSONMarshaler, data json.RawMessage) []abci.ValidatorUpdate {
	var genesisState group.GenesisState
	cdc.MustUnmarshalJSON(data, &genesisState)

	// TODO
	return []abci.ValidatorUpdate{}
}

func (s serverImpl) ExportGenesis(ctx types.Context, cdc codec.JSONMarshaler) json.RawMessage {
	// TODO
	return nil
}
