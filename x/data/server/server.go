package server

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/regen-network/regen-ledger/orm"
	servermodule "github.com/regen-network/regen-ledger/types/module/server"
	"github.com/regen-network/regen-ledger/x/data"
)

type serverImpl struct {
	storeKey sdk.StoreKey
	idSeq    orm.Sequence
}

func newServer(storeKey sdk.StoreKey) serverImpl {
	impl := serverImpl{storeKey: storeKey}
	impl.idSeq = orm.NewSequence(storeKey, IdSeqPrefix)
	return impl
}

func RegisterServices(configurator servermodule.Configurator) {
	impl := newServer(configurator.ModuleKey())
	data.RegisterMsgServer(configurator.MsgServer(), impl)
	data.RegisterQueryServer(configurator.QueryServer(), impl)
}
