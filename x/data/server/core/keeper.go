package core

import (
	"fmt"

	api "github.com/regen-network/regen-ledger/api/regen/data/v1"
	"github.com/regen-network/regen-ledger/x/data/server"
	"github.com/regen-network/regen-ledger/x/data/server/lookup"

	"github.com/cosmos/cosmos-sdk/orm/model/ormdb"
)

// var _ data.MsgServer = Keeper{}
// var _ data.QueryServer = Keeper{}

type Keeper struct {
	stateStore api.StateStore
	iriIDTable lookup.Table
}

func NewKeeper(db ormdb.ModuleDB) Keeper {
	ss, err := api.NewStateStore(db)
	if err != nil {
		panic(fmt.Sprintf("failed to create state tables for data module: %s", err.Error()))
	}
	tbl, err := lookup.NewTable([]byte{server.IriIDTablePrefix})
	if err != nil {
		panic(err)
	}
	return Keeper{
		stateStore: ss,
		iriIDTable: tbl,
	}
}
