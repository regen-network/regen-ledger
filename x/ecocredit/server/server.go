package server

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/orm/model/ormdb"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	basketv1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/basket/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/server/basket"
	"github.com/regen-network/regen-ledger/x/ecocredit/server/ormutil"
	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/regen-network/regen-ledger/orm"
	"github.com/regen-network/regen-ledger/types/module/server"
	"github.com/regen-network/regen-ledger/x/ecocredit"
)

const (
	TradableBalancePrefix    byte = 0x0
	TradableSupplyPrefix     byte = 0x1
	RetiredBalancePrefix     byte = 0x2
	RetiredSupplyPrefix      byte = 0x3
	CreditTypeSeqTablePrefix byte = 0x4
	ClassInfoTablePrefix     byte = 0x5
	BatchInfoTablePrefix     byte = 0x6
	ORMPrefix                byte = 0x7
)

var ModuleSchema = ormdb.ModuleSchema{
	FileDescriptors: map[uint32]protoreflect.FileDescriptor{
		1: basketv1.File_regen_ecocredit_basket_v1_state_proto,
	},
	Prefix: []byte{ORMPrefix},
}

type serverImpl struct {
	storeKey sdk.StoreKey

	paramSpace    paramtypes.Subspace
	bankKeeper    ecocredit.BankKeeper
	accountKeeper ecocredit.AccountKeeper

	// Store sequence numbers per credit type
	creditTypeSeqTable orm.PrimaryKeyTable

	classInfoTable orm.PrimaryKeyTable
	batchInfoTable orm.PrimaryKeyTable

	db ormdb.ModuleDB
}

func newServer(storeKey sdk.StoreKey, paramSpace paramtypes.Subspace,
	accountKeeper ecocredit.AccountKeeper, bankKeeper ecocredit.BankKeeper, cdc codec.Codec) serverImpl {
	s := serverImpl{
		storeKey:      storeKey,
		paramSpace:    paramSpace,
		bankKeeper:    bankKeeper,
		accountKeeper: accountKeeper,
	}

	creditTypeSeqTable, err := orm.NewPrimaryKeyTableBuilder(CreditTypeSeqTablePrefix, storeKey, &ecocredit.CreditTypeSeq{}, cdc)
	if err != nil {
		panic(err.Error())
	}
	s.creditTypeSeqTable = creditTypeSeqTable.Build()

	classInfoTableBuilder, err := orm.NewPrimaryKeyTableBuilder(ClassInfoTablePrefix, storeKey, &ecocredit.ClassInfo{}, cdc)
	if err != nil {
		panic(err.Error())
	}
	s.classInfoTable = classInfoTableBuilder.Build()

	batchInfoTableBuilder, err := orm.NewPrimaryKeyTableBuilder(BatchInfoTablePrefix, storeKey, &ecocredit.BatchInfo{}, cdc)
	if err != nil {
		panic(err.Error())
	}
	s.batchInfoTable = batchInfoTableBuilder.Build()

	s.db, err = ormutil.NewStoreKeyDB(ModuleSchema, storeKey, ormdb.ModuleDBOptions{})
	if err != nil {
		panic(err)
	}

	return s
}

func RegisterServices(configurator server.Configurator, paramSpace paramtypes.Subspace, accountKeeper ecocredit.AccountKeeper,
	bankKeeper ecocredit.BankKeeper) {
	impl := newServer(configurator.ModuleKey(), paramSpace, accountKeeper, bankKeeper, configurator.Marshaler())
	ecocredit.RegisterMsgServer(configurator.MsgServer(), impl)
	ecocredit.RegisterQueryServer(configurator.QueryServer(), impl)
	configurator.RegisterGenesisHandlers(impl.InitGenesis, impl.ExportGenesis)
	configurator.RegisterWeightedOperationsHandler(impl.WeightedOperations)
	configurator.RegisterInvariantsHandler(impl.RegisterInvariants)

	db, err := ormutil.NewStoreKeyDB(ModuleSchema, configurator.ModuleKey(), ormdb.ModuleDBOptions{})
	if err != nil {
		panic(err)
	}

	_ = basket.NewKeeper(db, impl, bankKeeper)
	// TODO Msg and Query server registration
}
