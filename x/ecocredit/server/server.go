package server

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/modules/incubator/orm"
	_ "github.com/cosmos/modules/incubator/orm"

	"github.com/regen-network/regen-ledger/x/ecocredit"
)

const (
	TradeableBalancePrefix byte = 0x0
	TradeableSupplyPrefix  byte = 0x1
	RetiredBalancePrefix   byte = 0x2
	RetiredSupplyPrefix    byte = 0x3
	ClassInfoSeqPrefix     byte = 0x4
	ClassInfoTablePrefix   byte = 0x5
	BatchInfoSeqPrefix     byte = 0x6
	BatchInfoTablePrefix   byte = 0x7
)

type serverImpl struct {
	storeKey sdk.StoreKey

	classInfoSeq   orm.Sequence
	classInfoTable orm.NaturalKeyTable

	batchInfoSeq   orm.Sequence
	batchInfoTable orm.NaturalKeyTable
}

type Server interface {
	ecocredit.MsgServer
	ecocredit.QueryServer
}

func NewServer(storeKey sdk.StoreKey) Server {
	s := serverImpl{storeKey: storeKey}

	s.classInfoSeq = orm.NewSequence(storeKey, ClassInfoSeqPrefix)
	classInfoTableBuilder := orm.NewNaturalKeyTableBuilder(ClassInfoTablePrefix, storeKey, &ecocredit.ClassInfo{}, orm.Max255DynamicLengthIndexKeyCodec{})
	s.classInfoTable = classInfoTableBuilder.Build()

	s.batchInfoSeq = orm.NewSequence(storeKey, BatchInfoSeqPrefix)
	batchInfoTableBuilder := orm.NewNaturalKeyTableBuilder(BatchInfoTablePrefix, storeKey, &ecocredit.BatchInfo{}, orm.Max255DynamicLengthIndexKeyCodec{})
	s.classInfoTable = batchInfoTableBuilder.Build()

	return s
}

func TradeableBalanceKey(acc string, batchDenom string) []byte {
	key := []byte{TradeableBalancePrefix}
	str := fmt.Sprintf("%s|%s", acc, batchDenom)
	return append(key, str...)
}

func TradeableSupplyKey(batchDenom string) []byte {
	key := []byte{TradeableSupplyPrefix}
	return append(key, batchDenom...)
}

func RetiredBalanceKey(acc string, batchDenom string) []byte {
	key := []byte{RetiredBalancePrefix}
	str := fmt.Sprintf("%s|%s", acc, batchDenom)
	return append(key, str...)
}

func RetiredSupplyKey(batchDenom string) []byte {
	key := []byte{RetiredSupplyPrefix}
	return append(key, batchDenom...)
}
