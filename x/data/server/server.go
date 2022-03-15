package server

import (
	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/cosmos/cosmos-sdk/orm/model/ormdb"
	sdk "github.com/cosmos/cosmos-sdk/types"

	api "github.com/regen-network/regen-ledger/api/regen/data/v1"
	servermodule "github.com/regen-network/regen-ledger/types/module/server"
	"github.com/regen-network/regen-ledger/types/ormstore"
	"github.com/regen-network/regen-ledger/x/data"
	"github.com/regen-network/regen-ledger/x/data/server/lookup"
)

var ModuleSchema = ormdb.ModuleSchema{
	FileDescriptors: map[uint32]protoreflect.FileDescriptor{
		1: api.File_regen_data_v1_state_proto,
	},
	Prefix: []byte{ORMStatePrefix},
}

type serverImpl struct {
	storeKey   sdk.StoreKey
	iriIDTable lookup.Table
	stateStore api.StateStore
}

func newServer(storeKey sdk.StoreKey) serverImpl {
	tbl, err := lookup.NewTable([]byte{IriIDTablePrefix})
	if err != nil {
		panic(err)
	}

	db, err := ormstore.NewStoreKeyDB(ModuleSchema, storeKey, ormdb.ModuleDBOptions{})
	if err != nil {
		panic(err)
	}

	stateStore, err := api.NewStateStore(db)
	if err != nil {
		panic(err)
	}

	return serverImpl{
		storeKey:   storeKey,
		iriIDTable: tbl,
		stateStore: stateStore,
	}
}

func RegisterServices(configurator servermodule.Configurator) {
	impl := newServer(configurator.ModuleKey())
	data.RegisterMsgServer(configurator.MsgServer(), impl)
	data.RegisterQueryServer(configurator.QueryServer(), impl)
}
