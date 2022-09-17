package simulation

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"

	dbm "github.com/tendermint/tm-db"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/cosmos/cosmos-sdk/orm/model/ormdb"
	"github.com/cosmos/cosmos-sdk/orm/model/ormtable"
	"github.com/cosmos/cosmos-sdk/orm/types/ormjson"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"

	api "github.com/regen-network/regen-ledger/api/regen/data/v1"
	"github.com/regen-network/regen-ledger/x/data"
	"github.com/regen-network/regen-ledger/x/data/server/hasher"
)

// RandomizedGenState generates a random GenesisState for the data module.
func RandomizedGenState(simState *module.SimulationState) {
	r := simState.Rand

	db := dbm.NewMemDB()
	backend := ormtable.NewBackend(ormtable.BackendOptions{
		CommitmentStore: db,
	})

	ormdb, err := ormdb.NewModuleDB(&data.ModuleSchema, ormdb.ModuleDBOptions{})
	if err != nil {
		panic(err)
	}

	ormCtx := ormtable.WrapContextDefault(backend)
	ss, err := api.NewStateStore(ormdb)
	if err != nil {
		panic(err)
	}

	if err := generateGenesisState(ormCtx, r, ss, simState); err != nil {
		panic(err)
	}

	jsonTarget := ormjson.NewRawMessageTarget()
	if err := ormdb.ExportJSON(ormCtx, jsonTarget); err != nil {
		panic(err)
	}

	rawJSON, err := jsonTarget.JSON()
	if err != nil {
		panic(err)
	}

	bz, err := json.Marshal(rawJSON)
	if err != nil {
		panic(err)
	}

	simState.GenState[data.ModuleName] = bz
}

func generateGenesisState(ormCtx context.Context, r *rand.Rand, ss api.StateStore,
	simState *module.SimulationState) error {
	hasher, err := hasher.NewHasher()
	if err != nil {
		return err
	}

	for i := 0; i < 5; i++ {
		contentHash, err := getContentHash(r)
		if err != nil {
			return err
		}

		iri, err := contentHash.ToIRI()
		if err != nil {
			return err
		}

		id := hasher.CreateID([]byte(iri), i)
		if err := ss.DataIDTable().Insert(ormCtx, &api.DataID{
			Iri: iri,
			Id:  id,
		}); err != nil {
			return err
		}

		if err := ss.DataAnchorTable().Insert(ormCtx, &api.DataAnchor{
			Id:        id,
			Timestamp: timestamppb.New(simtypes.RandTimestamp(r).UTC()),
		}); err != nil {
			return err
		}

		attestor, _ := simtypes.RandomAcc(r, simState.Accounts)
		if err := ss.DataAttestorTable().Insert(ormCtx, &api.DataAttestor{
			Id:        id,
			Attestor:  attestor.Address,
			Timestamp: timestamppb.New(simtypes.RandTimestamp(r).UTC()),
		}); err != nil {
			return err
		}

		domain := simtypes.RandStringOfLength(r, 3)
		manager, _ := simtypes.RandomAcc(r, simState.Accounts)
		resolverID, err := ss.ResolverTable().InsertReturningID(ormCtx, &api.Resolver{
			Url:     fmt.Sprintf("https://%s.foo", domain),
			Manager: manager.Address,
		})
		if err != nil {
			return err
		}

		err = ss.DataResolverTable().Insert(ormCtx, &api.DataResolver{
			ResolverId: resolverID,
			Id:         id,
		})
		if err != nil {
			return err
		}
	}

	return nil
}
