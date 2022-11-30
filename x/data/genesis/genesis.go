package genesis

import (
	"encoding/json"

	dbm "github.com/tendermint/tm-db"
	"google.golang.org/protobuf/proto"

	"github.com/cosmos/cosmos-sdk/orm/model/ormdb"
	"github.com/cosmos/cosmos-sdk/orm/model/ormtable"
	"github.com/cosmos/cosmos-sdk/orm/types/ormjson"

	api "github.com/regen-network/regen-ledger/api/v2/regen/data/v1"
	"github.com/regen-network/regen-ledger/types/v2/ormutil"
	"github.com/regen-network/regen-ledger/x/data/v2"
)

// ValidateGenesis performs basic validation of genesis state.
func ValidateGenesis(jsonData json.RawMessage) error {
	db := dbm.NewMemDB()
	backend := ormtable.NewBackend(ormtable.BackendOptions{
		CommitmentStore: db,
		IndexStore:      db,
	})

	moduleDB, err := ormdb.NewModuleDB(&data.ModuleSchema, ormdb.ModuleDBOptions{
		JSONValidator: validateMsg,
	})
	if err != nil {
		return err
	}

	ormCtx := ormtable.WrapContextDefault(backend)

	jsonSource, err := ormjson.NewRawMessageSource(jsonData)
	if err != nil {
		return err
	}

	err = moduleDB.ImportJSON(ormCtx, jsonSource)
	if err != nil {
		return err
	}

	if err := moduleDB.ValidateJSON(jsonSource); err != nil {
		return err
	}

	return nil
}

func validateMsg(m proto.Message) error {
	switch m.(type) {
	case *api.DataID:
		msg := &data.DataID{}
		if err := ormutil.PulsarToGogoSlow(m, msg); err != nil {
			return err
		}
		return msg.Validate()
	case *api.DataAnchor:
		msg := &data.DataAnchor{}
		if err := ormutil.PulsarToGogoSlow(m, msg); err != nil {
			return err
		}
		return msg.Validate()
	case *api.DataAttestor:
		msg := &data.DataAttestor{}
		if err := ormutil.PulsarToGogoSlow(m, msg); err != nil {
			return err
		}
		return msg.Validate()
	case *api.Resolver:
		msg := &data.Resolver{}
		if err := ormutil.PulsarToGogoSlow(m, msg); err != nil {
			return err
		}
		return msg.Validate()
	case *api.DataResolver:
		msg := &data.DataResolver{}
		if err := ormutil.PulsarToGogoSlow(m, msg); err != nil {
			return err
		}
		return msg.Validate()
	}

	return nil
}
