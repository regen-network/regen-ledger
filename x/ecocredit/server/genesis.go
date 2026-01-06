package server

import (
	"encoding/json"

	abci "github.com/cometbft/cometbft/abci/types"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/regen-network/regen-ledger/api/v2/orm/types/ormjson"
)

// InitGenesis performs genesis initialization for the ecocredit module. It
// returns no validator updates.
func (s serverImpl) InitGenesis(ctx sdk.Context, _ codec.JSONCodec, data json.RawMessage) ([]abci.ValidatorUpdate, error) {
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

// ExportGenesis will dump the ecocredit module state into a serializable GenesisState.
func (s serverImpl) ExportGenesis(ctx sdk.Context, _ codec.JSONCodec) (json.RawMessage, error) {
	jsonTarget := ormjson.NewRawMessageTarget()
	err := s.db.ExportJSON(ctx, jsonTarget)
	if err != nil {
		return nil, err
	}

	return jsonTarget.JSON()
}
