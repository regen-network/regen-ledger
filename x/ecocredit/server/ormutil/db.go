package ormutil

import (
	"context"

	"github.com/regen-network/regen-ledger/types"

	"github.com/cosmos/cosmos-sdk/orm/model/ormdb"
	"github.com/cosmos/cosmos-sdk/orm/model/ormtable"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
)

// NewStoreKeyDB creates an ormdb.ModuleDB from an ormdb.ModuleDB and a StoreKey.
// It is an interim solution for using the ORM in existing Cosmos SDK modules
// before fuller integration has been done.
func NewStoreKeyDB(desc ormdb.ModuleSchema, key storetypes.StoreKey, options ormdb.ModuleDBOptions) (ormdb.ModuleDB, error) {
	getBackend := func(ctx context.Context) (ormtable.Backend, error) {
		sdkCtx := types.UnwrapSDKContext(ctx)
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
