package server

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"

	"github.com/regen-network/regen-ledger/orm"
	regenmodule "github.com/regen-network/regen-ledger/types/module"
	"github.com/regen-network/regen-ledger/x/ecocredit"
)

const (
	TradableBalancePrefix  byte = 0x0
	TradableSupplyPrefix   byte = 0x1
	RetiredBalancePrefix   byte = 0x2
	RetiredSupplyPrefix    byte = 0x3
	IDSeqPrefix            byte = 0x4
	ClassInfoTablePrefix   byte = 0x5
	BatchInfoTablePrefix   byte = 0x6
	MaxDecimalPlacesPrefix byte = 0x7
)

type serverImpl struct {
	storeKey sdk.StoreKey

	// we use a single sequence to avoid having the same string/ID identifying a class and batch denom
	idSeq          orm.Sequence
	classInfoTable orm.NaturalKeyTable
	batchInfoTable orm.NaturalKeyTable
}

func newServer(storeKey sdk.StoreKey, cdc codec.BinaryMarshaler) serverImpl {
	s := serverImpl{storeKey: storeKey}

	s.idSeq = orm.NewSequence(storeKey, IDSeqPrefix)

	classInfoTableBuilder := orm.NewNaturalKeyTableBuilder(ClassInfoTablePrefix, storeKey, &ecocredit.ClassInfo{}, orm.Max255DynamicLengthIndexKeyCodec{}, cdc)
	s.classInfoTable = classInfoTableBuilder.Build()

	batchInfoTableBuilder := orm.NewNaturalKeyTableBuilder(BatchInfoTablePrefix, storeKey, &ecocredit.BatchInfo{}, orm.Max255DynamicLengthIndexKeyCodec{}, cdc)
	s.batchInfoTable = batchInfoTableBuilder.Build()

	return s
}

func RegisterServices(storeKey sdk.StoreKey, configurator module.Configurator) {
	cfg, ok := configurator.(regenmodule.Configurator)
	// We need regen configurator's BinaryMarshaler in order to
	// instantiate new table builders so panicking if it's not the case
	// until we use this upgraded configurator in the cosmos sdk
	if !ok {
		panic("configurator should implement regenmodule.Configurator")
	}
	impl := newServer(storeKey, cfg.BinaryMarshaler())
	ecocredit.RegisterMsgServer(configurator.MsgServer(), impl)
	ecocredit.RegisterQueryServer(configurator.QueryServer(), impl)
}
