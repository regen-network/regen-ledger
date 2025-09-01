package server

import (
	"encoding/json"

	"github.com/cometbft/cometbft/abci/types"

	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/regen-network/regen-ledger/orm/model/ormdb"

	api "github.com/regen-network/regen-ledger/api/v2/regen/data/v1"
	"github.com/regen-network/regen-ledger/types/v2/ormstore"
	"github.com/regen-network/regen-ledger/x/data/v3"
	"github.com/regen-network/regen-ledger/x/data/v3/server/hasher"
)

var _ data.MsgServer = serverImpl{}
var _ data.QueryServer = serverImpl{}

var _ Keeper = serverImpl{}

type Keeper interface {
	InitGenesis(ctx sdk.Context, cdc codec.JSONCodec, data json.RawMessage) ([]types.ValidatorUpdate, error)
	ExportGenesis(ctx sdk.Context, cdc codec.JSONCodec) (json.RawMessage, error)
	QueryServer() data.QueryServer
}

type serverImpl struct {
	iriHasher     hasher.Hasher
	stateStore    api.StateStore
	db            ormdb.ModuleDB
	bankKeeper    data.BankKeeper
	accountKeeper data.AccountKeeper
}

func (s serverImpl) QueryServer() data.QueryServer {
	return s
}

//nolint:revive
func NewServer(storeKey storetypes.StoreKey, ak data.AccountKeeper, bk data.BankKeeper) serverImpl {
	hasher, err := hasher.NewHasher()
	if err != nil {
		panic(err)
	}

	db, err := ormstore.NewStoreKeyDB(&data.ModuleSchema, storeKey, ormdb.ModuleDBOptions{})
	if err != nil {
		panic(err)
	}

	stateStore, err := api.NewStateStore(db)
	if err != nil {
		panic(err)
	}

	return serverImpl{
		iriHasher:     hasher,
		stateStore:    stateStore,
		db:            db,
		bankKeeper:    bk,
		accountKeeper: ak,
	}
}
