package core

import (
	"context"
	gomath "math"

	"github.com/cosmos/cosmos-sdk/orm/model/ormlist"
	"github.com/cosmos/cosmos-sdk/types/query"
	ecocreditv1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/types/ormutil"
)

// import (
// 	sdk "github.com/cosmos/cosmos-sdk/types"

// 	"github.com/regen-network/regen-ledger/x/ecocredit"
// )

// func (k Keeper) RegisterInvariants(ir sdk.InvariantRegistry) {
// 	ir.RegisterRoute(ecocredit.ModuleName, "tradable-supply", k.tradableSupplyInvariant())
// 	ir.RegisterRoute(ecocredit.ModuleName, "retired-supply", k.retiredSupplyInvariant())
// }

// func (k Keeper) tradableSupplyInvariant() sdk.Invariant {
// 	k.
// 	return func(ctx sdk.Context) (string, bool) {
// 		store := ctx.KVStore(s.storeKey)
// 		goCtx := sdk.WrapSDKContext(ctx)
// 		basketBalances := k.
// 		return tradableSupplyInvariant(store, basketBalances)
// 	}
// }

// func (k Keeper) retiredSupplyInvariant() sdk.Invariant {
// 	return func(ctx sdk.Context) (string, bool) {
// 	}
// }

func (k Keeper) BatchBalanceIterator(ctx context.Context) (ecocreditv1.BatchBalanceIterator, error) {
	pulsarPageReq, err := ormutil.GogoPageReqToPulsarPageReq(&query.PageRequest{
		Offset: 0,
		Limit:  gomath.MaxUint64,
	})
	if err != nil {
		return ecocreditv1.BatchBalanceIterator{}, err
	}

	return k.stateStore.BatchBalanceTable().List(ctx, ecocreditv1.BatchBalancePrimaryKey{}, ormlist.Paginate(pulsarPageReq))
}

func (k Keeper) BatchSupplyIterator(ctx context.Context) (ecocreditv1.BatchSupplyIterator, error) {
	pulsarPageReq, err := ormutil.GogoPageReqToPulsarPageReq(&query.PageRequest{
		Offset: 0,
		Limit:  gomath.MaxUint64,
	})
	if err != nil {
		return ecocreditv1.BatchSupplyIterator{}, err
	}

	return k.stateStore.BatchSupplyTable().List(ctx, ecocreditv1.BatchSupplyPrimaryKey{}, ormlist.Paginate(pulsarPageReq))
}
