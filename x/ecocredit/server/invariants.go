package server

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/regen-network/regen-ledger/x/ecocredit"
	basekeeper "github.com/regen-network/regen-ledger/x/ecocredit/base/keeper"
)

// RegisterInvariants registers the ecocredit module invariants.
func (s serverImpl) RegisterInvariants(ir sdk.InvariantRegistry) {
	ir.RegisterRoute(ecocredit.ModuleName, "batch-supply", s.batchSupplyInvariant())
	s.BasketKeeper.RegisterInvariants(ir)
}

func (s serverImpl) batchSupplyInvariant() sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		goCtx := sdk.WrapSDKContext(ctx)
		basketBalances, err := s.BasketKeeper.GetBasketBalanceMap(goCtx)
		if err != nil {
			return err.Error(), true
		}

		msg, broken := basekeeper.BatchSupplyInvariant(goCtx, s.BaseKeeper, basketBalances)
		return sdk.FormatInvariant(ecocredit.ModuleName, "batch-supply", msg), broken
	}
}
