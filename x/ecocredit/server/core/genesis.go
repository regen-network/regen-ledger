package core

import (
	"encoding/json"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/orm/types/ormjson"
	regentypes "github.com/regen-network/regen-ledger/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

// InitGenesis performs genesis initialization for the ecocredit module. It
// returns no validator updates.
func (s serverImpl) InitGenesis(ctx regentypes.Context, cdc codec.Codec, data json.RawMessage) ([]abci.ValidatorUpdate, error) {
	goCtx := ctx.Context.Context()
	target, err := ormjson.NewRawMessageSource(data)
	if err != nil {
		return nil, err
	}
	return []abci.ValidatorUpdate{}, s.db.ImportJSON(goCtx, target)
}

// ExportGenesis will dump the ecocredit module state into a serializable GenesisState.
func (s serverImpl) ExportGenesis(ctx regentypes.Context, cdc codec.Codec) (json.RawMessage, error) {
	target := ormjson.NewRawMessageTarget()
	err := s.db.ExportJSON(ctx.Context.Context(), target)
	if err != nil {
		return nil, err
	}
	return target.JSON()
}
