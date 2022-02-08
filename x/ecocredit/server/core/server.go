package core

import (
	"github.com/cosmos/cosmos-sdk/orm/model/ormdb"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	ecocreditv1beta1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1beta1"
	"github.com/regen-network/regen-ledger/types/module/server"
	"github.com/regen-network/regen-ledger/types/ormstore"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/regen-network/regen-ledger/x/ecocredit/v1beta1"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type serverImpl struct {
	storeKey sdk.StoreKey

	paramSpace    paramtypes.Subspace
	bankKeeper    ecocredit.BankKeeper
	accountKeeper ecocredit.AccountKeeper

	db ormdb.ModuleDB

	stateStore ecocreditv1beta1.StateStore
}

func newServer(storeKey sdk.StoreKey, paramSpace paramtypes.Subspace,
	accountKeeper ecocredit.AccountKeeper, bankKeeper ecocredit.BankKeeper) serverImpl {

	s := serverImpl{
		storeKey:      storeKey,
		paramSpace:    paramSpace,
		bankKeeper:    bankKeeper,
		accountKeeper: accountKeeper,
	}

	ecocreditSchema := ormdb.ModuleSchema{
		FileDescriptors: map[uint32]protoreflect.FileDescriptor{1: ecocreditv1beta1.File_regen_ecocredit_v1beta1_state_proto},
		Prefix:          nil,
	}

	db, err := ormstore.NewStoreKeyDB(ecocreditSchema, storeKey, ormdb.ModuleDBOptions{})
	if err != nil {
		panic(err)
	}
	s.db = db

	stateStore, err := ecocreditv1beta1.NewStateStore(db)
	if err != nil {
		panic(err)
	}
	s.stateStore = stateStore

	return s
}

func RegisterServices(configurator server.Configurator, paramSpace paramtypes.Subspace, accountKeeper ecocredit.AccountKeeper,
	bankKeeper ecocredit.BankKeeper) v1beta1.MsgServer {
	impl := newServer(configurator.ModuleKey(), paramSpace, accountKeeper, bankKeeper)
	v1beta1.RegisterMsgServer(configurator.MsgServer(), impl)
	v1beta1.RegisterQueryServer(configurator.QueryServer(), impl)
	configurator.RegisterGenesisHandlers(impl.InitGenesis, impl.ExportGenesis)
	configurator.RegisterWeightedOperationsHandler(impl.WeightedOperations)
	configurator.RegisterInvariantsHandler(impl.RegisterInvariants)
	return impl
}
