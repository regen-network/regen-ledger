package server

import (
	"encoding/json"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/orm/types/ormjson"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/regen-network/regen-ledger/types"
)

// InitGenesis performs genesis initialization for the data module. It
// returns no validator updates.
func (s serverImpl) InitGenesis(ctx types.Context, cdc codec.Codec, data json.RawMessage) ([]abci.ValidatorUpdate, error) {
	jsonSource, err := ormjson.NewRawMessageSource(data)
	if err != nil {
		return nil, err
	}

	err = s.db.ImportJSON(ctx, jsonSource)
	if err != nil {
		return nil, err
	}

	return []abci.ValidatorUpdate{}, nil
}

// ExportGenesis will dump the data module state into a serializable GenesisState.
func (s serverImpl) ExportGenesis(ctx types.Context, cdc codec.Codec) (json.RawMessage, error) {
	jsonTarget := ormjson.NewRawMessageTarget()
	err := s.db.ExportJSON(ctx, jsonTarget)
	if err != nil {
		return nil, err
	}

	return jsonTarget.JSON()
}
