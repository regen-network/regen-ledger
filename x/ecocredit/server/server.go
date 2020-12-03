package server

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/regen-network/regen-ledger/types/module/server"
	"github.com/regen-network/regen-ledger/x/bank"

	"github.com/regen-network/regen-ledger/orm"
	"github.com/regen-network/regen-ledger/x/ecocredit"
)

type serverImpl struct {
	key        server.RootModuleKey
	moduleAddr sdk.AccAddress

	// we use a single sequence to avoid having the same string/ID identifying a class and batch denom
	idSeq          orm.Sequence
	classInfoTable orm.NaturalKeyTable
	batchInfoTable orm.NaturalKeyTable

	bankMsgClient   bank.MsgClient
	bankQueryClient bank.QueryClient
}

func newServer(key server.RootModuleKey) serverImpl {
	s := serverImpl{key: key}
	s.moduleAddr = key.Address()

	s.idSeq = orm.NewSequence(key, IDSeqPrefix)

	classInfoTableBuilder := orm.NewNaturalKeyTableBuilder(ClassInfoTablePrefix, key, &ecocredit.ClassInfo{}, orm.Max255DynamicLengthIndexKeyCodec{})
	s.classInfoTable = classInfoTableBuilder.Build()

	batchInfoTableBuilder := orm.NewNaturalKeyTableBuilder(BatchInfoTablePrefix, key, &ecocredit.BatchInfo{}, orm.Max255DynamicLengthIndexKeyCodec{})
	s.batchInfoTable = batchInfoTableBuilder.Build()

	s.bankMsgClient = bank.NewMsgClient(key)
	s.bankQueryClient = bank.NewQueryClient(key)

	return s
}

func RegisterServices(configurator server.Configurator) {
	impl := newServer(configurator.ModuleKey())
	ecocredit.RegisterMsgServer(configurator.MsgServer(), impl)
	ecocredit.RegisterQueryServer(configurator.QueryServer(), impl)
	configurator.RequireServer((*bank.MsgServer)(nil))
}
