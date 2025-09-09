package server

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/regen-network/regen-ledger/x/ecocredit/v4"
	basekeeper "github.com/regen-network/regen-ledger/x/ecocredit/v4/base/keeper"
)

// RegisterInvariants registers the ecocredit module invariants.
func (s serverImpl) RegisterInvariants(ir sdk.InvariantRegistry) {
	ir.RegisterRoute(ecocredit.ModuleName, "batch-supply", s.batchSupplyInvariant())
	s.basketKeeper.RegisterInvariants(ir)
}

func (s serverImpl) batchSupplyInvariant() sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		goCtx := ctx
		basketBalances, err := s.basketKeeper.GetBasketBalanceMap(goCtx)
		if err != nil {
			return err.Error(), true
		}

		msg, broken := basekeeper.BatchSupplyInvariant(goCtx, s.baseKeeper, basketBalances)
		return sdk.FormatInvariant(ecocredit.ModuleName, "batch-supply", msg), broken
	}
}
