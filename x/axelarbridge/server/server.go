package server

import (
	ormv1alpha1 "github.com/cosmos/cosmos-sdk/api/cosmos/orm/v1alpha1"
	"github.com/cosmos/cosmos-sdk/orm/model/ormdb"
	sdk "github.com/cosmos/cosmos-sdk/types"

	api "github.com/regen-network/regen-ledger/api/axelar/bridge/v1"
	servermodule "github.com/regen-network/regen-ledger/types/module/server"
	"github.com/regen-network/regen-ledger/types/ormstore"
	"github.com/regen-network/regen-ledger/x/axelarbridge"
)

const (
	ORMPrefix byte = iota
)

var _ axelarbridge.MsgServer = serverImpl{}

var ModuleSchema = ormv1alpha1.ModuleSchemaDescriptor{
	SchemaFile: []*ormv1alpha1.ModuleSchemaDescriptor_FileEntry{
		{Id: 1, ProtoFileName: api.File_axelar_bridge_v1_state_proto.Path(), StorageType: ormv1alpha1.StorageType_STORAGE_TYPE_DEFAULT_UNSPECIFIED},
	},
	Prefix: []byte{ORMPrefix},
}

type serverImpl struct {
	storeKey   sdk.StoreKey
	stateStore api.StateStore
	db         ormdb.ModuleDB
}

func newServer(storeKey sdk.StoreKey) serverImpl {
	db, err := ormstore.NewStoreKeyDB(&ModuleSchema, storeKey, ormdb.ModuleDBOptions{})
	if err != nil {
		panic(err)
	}

	stateStore, err := api.NewStateStore(db)
	if err != nil {
		panic(err)
	}

	return serverImpl{
		storeKey:   storeKey,
		stateStore: stateStore,
		db:         nil, // db,
	}
}

func RegisterServices(configurator servermodule.Configurator) {
	impl := newServer(configurator.ModuleKey())
	axelarbridge.RegisterMsgServer(configurator.MsgServer(), impl)
}
