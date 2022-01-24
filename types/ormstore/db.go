package ormstore

import (
	"context"

	"github.com/cosmos/cosmos-sdk/orm/model/ormdb"
	"github.com/cosmos/cosmos-sdk/orm/model/ormtable"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func NewStoreKeyDB(desc ormdb.ModuleSchema, key storetypes.StoreKey, options ormdb.ModuleDBOptions) (ormdb.ModuleDB, error) {
	getBackend := func(ctx context.Context) (ormtable.Backend, error) {
		sdkCtx := sdk.UnwrapSDKContext(ctx)
		store := sdkCtx.KVStore(key)
		wrapper := storeWrapper{store}
		return ormtable.NewBackend(ormtable.BackendOptions{
			CommitmentStore: wrapper,
			IndexStore:      wrapper,
		}), nil
	}
	options.GetBackend = getBackend
	options.GetReadBackend = func(ctx context.Context) (ormtable.ReadBackend, error) {
		return getBackend(ctx)
	}
	return ormdb.NewModuleDB(desc, options)
}
