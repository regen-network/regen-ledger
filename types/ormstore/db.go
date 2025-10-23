package ormstore

import (
	"context"

	storetypes "cosmossdk.io/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	ormv1alpha1 "github.com/regen-network/regen-ledger/api/v2/regen/orm/v1alpha1"
	"github.com/regen-network/regen-ledger/orm/model/ormdb"
	"github.com/regen-network/regen-ledger/orm/model/ormtable"
)

// NewStoreKeyDB creates an ormdb.ModuleDB from an ormdb.ModuleDB and a StoreKey.
// It is an interim solution for using the ORM in existing Cosmos SDK modules
// before fuller integration has been done.
func NewStoreKeyDB(desc *ormv1alpha1.ModuleSchemaDescriptor, key storetypes.StoreKey, options ormdb.ModuleDBOptions) (ormdb.ModuleDB, error) {
	backEndResolver := func(_ ormv1alpha1.StorageType) (ormtable.BackendResolver, error) {
		getBackend := func(ctx context.Context) (ormtable.ReadBackend, error) {
			sdkCtx := sdk.UnwrapSDKContext(ctx)
			store := sdkCtx.KVStore(key)
			wrapper := storeWrapper{store}
			return ormtable.NewBackend(ormtable.BackendOptions{
				CommitmentStore: wrapper,
				IndexStore:      wrapper,
			}), nil
		}
		return getBackend, nil
	}
	options.GetBackendResolver = backEndResolver
	return ormdb.NewModuleDB(desc, options)
}
