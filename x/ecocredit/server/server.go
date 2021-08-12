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
)

type serverImpl struct {
	storeKey sdk.StoreKey

	paramSpace paramtypes.Subspace
	bankKeeper ecocredit.BankKeeper

	// Store sequence numbers per credit type
	creditTypeSeqTable orm.PrimaryKeyTable

	classInfoTable orm.PrimaryKeyTable
	batchInfoTable orm.PrimaryKeyTable
}

func newServer(storeKey sdk.StoreKey, paramSpace paramtypes.Subspace, bankKeeper ecocredit.BankKeeper, cdc codec.Codec) serverImpl {
	s := serverImpl{
		storeKey:   storeKey,
		paramSpace: paramSpace,
		bankKeeper: bankKeeper,
	}

	creditTypeSeqTable := orm.NewPrimaryKeyTableBuilder(CreditTypeSeqTablePrefix, storeKey, &ecocredit.CreditTypeSeq{}, orm.Max255DynamicLengthIndexKeyCodec{}, cdc)
	s.creditTypeSeqTable = creditTypeSeqTable.Build()

	classInfoTableBuilder := orm.NewPrimaryKeyTableBuilder(ClassInfoTablePrefix, storeKey, &ecocredit.ClassInfo{}, orm.Max255DynamicLengthIndexKeyCodec{}, cdc)
	s.classInfoTable = classInfoTableBuilder.Build()

	batchInfoTableBuilder := orm.NewPrimaryKeyTableBuilder(BatchInfoTablePrefix, storeKey, &ecocredit.BatchInfo{}, orm.Max255DynamicLengthIndexKeyCodec{}, cdc)
	s.batchInfoTable = batchInfoTableBuilder.Build()

	return s
}

func RegisterServices(configurator server.Configurator, paramSpace paramtypes.Subspace, bankKeeper ecocredit.BankKeeper) {
	impl := newServer(configurator.ModuleKey(), paramSpace, bankKeeper, configurator.Marshaler())
	ecocredit.RegisterMsgServer(configurator.MsgServer(), impl)
	ecocredit.RegisterQueryServer(configurator.QueryServer(), impl)
	configurator.RegisterGenesisHandlers(impl.InitGenesis, impl.ExportGenesis)
	configurator.RegisterInvariantsHandler(impl.RegisterInvariants)
}
