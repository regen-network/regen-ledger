package server

import (
	"encoding/json"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/pkg/errors"
	bond "github.com/regen-network/regen-ledger/v2/x/bond"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/regen-network/regen-ledger/types"
)

// InitGenesis performs genesis initialization for the bond module. It returns no validator updates.
func (s serverImpl) InitGenesis(ctx types.Context, cdc codec.Codec, data json.RawMessage) ([]abci.ValidatorUpdate, error) {
	var genesisState bond.GenesisState
	cdc.MustUnmarshalJSON(data, &genesisState)

	s.paramSpace.SetParamSet(ctx.Context, &genesisState.Params)

	if err := s.bondInfoTable.Import(ctx, genesisState.BondInfo, 0); err != nil {
		return nil, errors.Wrap(err, "sequences")
	}

	return []abci.ValidatorUpdate{}, nil
}

// ExportGenesis will dump the bond module state into a serializable GenesisState.
func (s serverImpl) ExportGenesis(ctx types.Context, cdc codec.Codec) (json.RawMessage, error) {
	// Get Params from the store and put them in the genesis state
	var params bond.Params
	s.paramSpace.GetParamSet(ctx.Context, &params)

	var bondInfo []*bond.BondInfo
	if _, err := s.bondInfoTable.Export(ctx, &bondInfo); err != nil {
		return nil, errors.Wrap(err, "project-info")
	}

	gs := &bond.GenesisState{
		Params:   params,
		BondInfo: bondInfo,
	}

	return cdc.MustMarshalJSON(gs), nil
}
