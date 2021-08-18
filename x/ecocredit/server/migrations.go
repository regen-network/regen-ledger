package server

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	v2 "github.com/regen-network/regen-ledger/x/ecocredit/migrations/v2"
)

// Migrator is a struct for handling in-place store migrations.
type Migrator struct {
	server serverImpl
}

// Migrate0to1 migrates from version 1 to 2.
func (m Migrator) Migrate0to1(ctx sdk.Context) error {
	return v2.UpgradeEco(m.server.paramSpace, ctx)
}
