package simulation

import (
	"github.com/cosmos/cosmos-sdk/orm/model/ormdb"
	"github.com/cosmos/cosmos-sdk/orm/model/ormtable"
	"github.com/cosmos/cosmos-sdk/types/module"
	dbm "github.com/tendermint/tm-db"

	api "github.com/regen-network/regen-ledger/api/regen/data/v1"
	"github.com/regen-network/regen-ledger/x/data/server"
)

// RandomizedGenState generates a random GenesisState for the data module.
func RandomizedGenState(simState *module.SimulationState) {

	db := dbm.NewMemDB()
	backend := ormtable.NewBackend(ormtable.BackendOptions{
		CommitmentStore: db,
		IndexStore:      db,
	})

	ormdb, err := ormdb.NewModuleDB(&server.ModuleSchema, ormdb.ModuleDBOptions{})
	if err != nil {
		panic(err)
	}

	ormCtx := ormtable.WrapContextDefault(backend)
	ss, err := api.NewStateStore(ormdb)
	if err != nil {
		panic(err)
	}

	manager := simState.Accounts[0]

	resolverId, err := ss.ResolverTable().InsertReturningID(ormCtx, &api.Resolver{
		Url:     "https://foo.bar",
		Manager: manager.Address,
	})
	if err != nil {
		panic(err)
	}

	err = ss.DataResolverTable().Insert(ormCtx, &api.DataResolver{
		ResolverId: resolverId,
	})
	if err != nil {
		panic(err)
	}

	ss.DataIDTable().Insert(ormCtx, &api.DataID{})

	ss.DataAnchorTable().Insert(ormCtx, &api.DataAnchor{})

	ss.DataAttestorTable().Insert(ormCtx, &api.DataAttestor{})

}
