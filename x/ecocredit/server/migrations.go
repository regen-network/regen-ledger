package server

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"

	ecocreditv1 "github.com/regen-network/regen-ledger/api/v2/regen/ecocredit/v1"
	v3 "github.com/regen-network/regen-ledger/x/ecocredit/v3/migrations/v3"
	v4 "github.com/regen-network/regen-ledger/x/ecocredit/v3/migrations/v4"
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

	baseStore, basketStore, _ := m.keeper.GetStateStores()
	if err := v3.MigrateState(ctx, baseStore, basketStore, m.legacySubspace); err != nil {
		return err
	}

	// add polygon to the allowed bridge chain table, as this was a hard coded requirement previously.
	err := baseStore.AllowedBridgeChainTable().Insert(ctx, &ecocreditv1.AllowedBridgeChain{ChainName: "polygon"})
	if err != nil {
		return err
	}

	return nil
}

// Migrate3to4 migrates from version 3 to 4.
func (m Migrator) Migrate3to4(ctx sdk.Context) error {
	baseStore, basketStore, _ := m.keeper.GetStateStores()
	return v4.MigrateState(ctx, baseStore, basketStore)
}
