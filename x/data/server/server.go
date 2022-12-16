package server

import (
	"encoding/json"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/orm/model/ormdb"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/abci/types"

	api "github.com/regen-network/regen-ledger/api/v2/regen/data/v1"
	"github.com/regen-network/regen-ledger/types/v2/ormstore"
	"github.com/regen-network/regen-ledger/x/data/v2"
	"github.com/regen-network/regen-ledger/x/data/v2/server/hasher"
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
	stateStore    api.StateStore
	db            ormdb.ModuleDB
	bankKeeper    data.BankKeeper
	accountKeeper data.AccountKeeper
	iriHasher     hasher.Hasher
	iriPrefix     string
}

func (s serverImpl) QueryServer() data.QueryServer {
	return s
}

//nolint:revive
func NewServer(storeKey storetypes.StoreKey, ak data.AccountKeeper, bk data.BankKeeper, config data.Config) serverImpl {
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

	if config.IRIPrefix == "" {
		config.IRIPrefix = data.DefaultConfig().IRIPrefix
	}

	return serverImpl{
		stateStore:    stateStore,
		db:            db,
		bankKeeper:    bk,
		accountKeeper: ak,
		iriHasher:     hasher,
		iriPrefix:     config.IRIPrefix,
	}
}
