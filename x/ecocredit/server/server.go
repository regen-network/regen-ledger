package server

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"

	"github.com/regen-network/regen-ledger/orm"
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

func newServer(storeKey sdk.StoreKey) serverImpl {
	s := serverImpl{storeKey: storeKey}

	s.idSeq = orm.NewSequence(storeKey, IDSeqPrefix)

	classInfoTableBuilder := orm.NewNaturalKeyTableBuilder(ClassInfoTablePrefix, storeKey, &ecocredit.ClassInfo{}, orm.Max255DynamicLengthIndexKeyCodec{})
	s.classInfoTable = classInfoTableBuilder.Build()

	batchInfoTableBuilder := orm.NewNaturalKeyTableBuilder(BatchInfoTablePrefix, storeKey, &ecocredit.BatchInfo{}, orm.Max255DynamicLengthIndexKeyCodec{})
	s.batchInfoTable = batchInfoTableBuilder.Build()

	return s
}

func RegisterServices(storeKey sdk.StoreKey, configurator module.Configurator) {
	impl := newServer(storeKey)
	ecocredit.RegisterMsgServer(configurator.MsgServer(), impl)
	ecocredit.RegisterQueryServer(configurator.QueryServer(), impl)
}
