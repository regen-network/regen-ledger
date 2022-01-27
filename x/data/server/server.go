package server

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	datav1alpha2 "github.com/regen-network/regen-ledger/api/regen/data/v1alpha2"
	servermodule "github.com/regen-network/regen-ledger/types/module/server"
	"github.com/regen-network/regen-ledger/x/data"
	"github.com/regen-network/regen-ledger/x/data/server/lookup"
)

type serverImpl struct {
	storeKey   sdk.StoreKey
	iriIDTable lookup.Table
	stateStore datav1alpha2.StateStore
}

func newServer(storeKey sdk.StoreKey) serverImpl {
	tbl, err := lookup.NewTable([]byte{IriIDTablePrefix})
	if err != nil {
		panic(err)
	}

	return serverImpl{
		storeKey:   storeKey,
		iriIDTable: tbl,
	}
}

func RegisterServices(configurator servermodule.Configurator) {
	impl := newServer(configurator.ModuleKey())
	data.RegisterMsgServer(configurator.MsgServer(), impl)
	data.RegisterQueryServer(configurator.QueryServer(), impl)
}
