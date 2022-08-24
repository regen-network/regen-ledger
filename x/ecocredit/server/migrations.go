package server

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	v4 "github.com/regen-network/regen-ledger/x/ecocredit/migrations/v4"
)

func (s serverImpl) RunMigrations(ctx sdk.Context, cdc codec.Codec) error {

	if err := v4.MigrateState(ctx, s.storeKey, cdc, s.stateStore, s.basketStore, s.legacySubspace); err != nil {
		return err
	}

	return nil
}
