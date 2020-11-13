package server

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"

	"github.com/regen-network/regen-ledger/x/data"
)

type serverImpl struct {
	storeKey sdk.StoreKey
}

func newServer(storeKey sdk.StoreKey) serverImpl {
	return serverImpl{storeKey: storeKey}
}

func RegisterServices(storeKey sdk.StoreKey, configurator module.Configurator) {
	impl := newServer(storeKey)
	data.RegisterMsgServer(configurator.MsgServer(), impl)
	data.RegisterQueryServer(configurator.QueryServer(), impl)
}
