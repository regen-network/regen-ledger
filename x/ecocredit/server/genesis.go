package server

import (
	"bytes"
	"encoding/json"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gogo/protobuf/jsonpb"
	"github.com/gogo/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/orm/types/ormjson"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

// InitGenesis performs genesis initialization for the ecocredit module. It
// returns no validator updates.
func (s serverImpl) InitGenesis(ctx sdk.Context, cdc codec.JSONCodec, data json.RawMessage) ([]abci.ValidatorUpdate, error) {
	jsonSource, err := ormjson.NewRawMessageSource(data)
	if err != nil {
		return nil, err
	}

	err = s.db.ImportJSON(sdk.WrapSDKContext(ctx), jsonSource)
	if err != nil {
		return nil, err
	}

	var params core.Params
	r, err := jsonSource.OpenReader(protoreflect.FullName(proto.MessageName(&params)))
	if err != nil {
		return nil, err
	}

	if r == nil { // r is nil when theres no table data, so we can just unmarshal the data given
		bz := bytes.NewBuffer(data)
		err = (&jsonpb.Unmarshaler{AllowUnknownFields: true}).Unmarshal(bz, &params)
		if err != nil {
			return nil, err
		}
	} else { // r is not nil, so there is table data and we can just use r.
		err = (&jsonpb.Unmarshaler{AllowUnknownFields: true}).Unmarshal(r, &params)
		if err != nil {
			return nil, err
		}
	}

	s.legacySubspace.SetParamSet(ctx, &params)

	return []abci.ValidatorUpdate{}, nil
}

// ExportGenesis will dump the ecocredit module state into a serializable GenesisState.
func (s serverImpl) ExportGenesis(ctx sdk.Context, cdc codec.JSONCodec) (json.RawMessage, error) {
	// Get Params from the store and put them in the genesis state
	var params core.Params
	s.legacySubspace.GetParamSet(ctx, &params)

	jsonTarget := ormjson.NewRawMessageTarget()
	err := s.db.ExportJSON(sdk.WrapSDKContext(ctx), jsonTarget)
	if err != nil {
		return nil, err
	}

	return jsonTarget.JSON()
}
