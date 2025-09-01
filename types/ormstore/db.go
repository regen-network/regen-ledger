package ormstore

import (
	ormv1alpha1 "github.com/regen-network/regen-ledger/api/v2/regen/orm/v1alpha1"

	storetypes "cosmossdk.io/store/types"
	"github.com/regen-network/regen-ledger/orm/model/ormdb"
)

// NewStoreKeyDB creates an ormdb.ModuleDB from an ormdb.ModuleDB and a StoreKey.
// It is an interim solution for using the ORM in existing Cosmos SDK modules
// before fuller integration has been done.
func NewStoreKeyDB(desc *ormv1alpha1.ModuleSchemaDescriptor, key storetypes.StoreKey, options ormdb.ModuleDBOptions) (ormdb.ModuleDB, error) {

	return ormdb.NewModuleDB(desc, options)
}
