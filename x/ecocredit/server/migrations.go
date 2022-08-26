package server

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	v3 "github.com/regen-network/regen-ledger/x/ecocredit/migrations/v3"
)

// Migrator is a struct for handling in-place store migrations.
type Migrator struct {
	keeper         Keeper
	legacySubspace paramtypes.Subspace
}

// NewMigrator returns a new Migrator.
func NewMigrator(keeper Keeper, legacySubspace paramtypes.Subspace) Migrator {
	return Migrator{
		keeper:         keeper,
		legacySubspace: legacySubspace,
	}
}

// Migrate2to3 migrates from version 2 to 3.
func (m Migrator) Migrate2to3(ctx sdk.Context) error {

	coreStore, basketStore, _ := m.keeper.GetStateStores()
	if err := v3.MigrateState(ctx, coreStore, basketStore, m.legacySubspace); err != nil {
		return err
	}

	return nil
}
