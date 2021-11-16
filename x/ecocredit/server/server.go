package server

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"

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
	SellOrderTablePrefix     byte = 0x7
	SellOrderTableSeqPrefix  byte = 0x8
	BuyOrderTablePrefix      byte = 0x9
	BuyOrderTableSeqPrefix   byte = 0x10
	AskDenomTablePrefix      byte = 0x11
)

type serverImpl struct {
	storeKey sdk.StoreKey

	paramSpace    paramtypes.Subspace
	bankKeeper    ecocredit.BankKeeper
	accountKeeper ecocredit.AccountKeeper

	// Store sequence numbers per credit type
	creditTypeSeqTable orm.PrimaryKeyTable

	classInfoTable orm.PrimaryKeyTable
	batchInfoTable orm.PrimaryKeyTable
	sellOrderTable orm.AutoUInt64Table
	buyOrderTable orm.AutoUInt64Table
	askDenomTable orm.PrimaryKeyTable
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

	sellOrderTableBuilder, err := orm.NewAutoUInt64TableBuilder(SellOrderTablePrefix, SellOrderTableSeqPrefix, storeKey, &ecocredit.SellOrder{}, cdc)
	if err != nil {
		panic(err.Error())
	}
	s.sellOrderTable = sellOrderTableBuilder.Build()

	buyOrderTableBuilder, err := orm.NewAutoUInt64TableBuilder(BuyOrderTablePrefix, BuyOrderTableSeqPrefix, storeKey, &ecocredit.BuyOrder{}, cdc)
	if err != nil {
		panic(err.Error())
	}
	s.buyOrderTable = buyOrderTableBuilder.Build()

	askDenomTableBuilder, err := orm.NewPrimaryKeyTableBuilder(AskDenomTablePrefix, storeKey, &ecocredit.AskDenom{}, cdc)
	if err != nil {
		panic(err.Error())
	}
	s.askDenomTable = askDenomTableBuilder.Build()

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
}
