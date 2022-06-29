package server

import (
	"github.com/cosmos/cosmos-sdk/orm/model/ormdb"
	sdk "github.com/cosmos/cosmos-sdk/types"

	api "github.com/regen-network/regen-ledger/api/regen/data/v1"
	servermodule "github.com/regen-network/regen-ledger/types/module/server"
	"github.com/regen-network/regen-ledger/types/ormstore"
	"github.com/regen-network/regen-ledger/x/data"
	"github.com/regen-network/regen-ledger/x/data/server/hasher"
)

var _ data.MsgServer = serverImpl{}
var _ data.QueryServer = serverImpl{}

type serverImpl struct {
	storeKey      sdk.StoreKey
	iriHasher     hasher.Hasher
	stateStore    api.StateStore
	db            ormdb.ModuleDB
	bankKeeper    data.BankKeeper
	accountKeeper data.AccountKeeper
}

func newServer(storeKey sdk.StoreKey, ak data.AccountKeeper, bk data.BankKeeper) serverImpl {
	hasher, err := hasher.NewHasher()
	if err != nil {
		panic(err)
	}

	db, err := ormstore.NewStoreKeyDB(&data.ModuleSchema, storeKey, ormdb.ModuleDBOptions{})
	if err != nil {
		panic(err)
	}

	stateStore, err := api.NewStateStore(db)
	if err != nil {
		panic(err)
	}

	return serverImpl{
		storeKey:      storeKey,
		iriHasher:     hasher,
		stateStore:    stateStore,
		db:            db,
		bankKeeper:    bk,
		accountKeeper: ak,
	}
}

func RegisterServices(configurator servermodule.Configurator, ak data.AccountKeeper, bk data.BankKeeper) {
	impl := newServer(configurator.ModuleKey(), ak, bk)
	data.RegisterMsgServer(configurator.MsgServer(), impl)
	data.RegisterQueryServer(configurator.QueryServer(), impl)

	configurator.RegisterGenesisHandlers(impl.InitGenesis, impl.ExportGenesis)
	configurator.RegisterWeightedOperationsHandler(impl.WeightedOperations)
}
