package upgrade

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

type Keeper struct {
	storeKey        sdk.StoreKey
	cdc             *codec.Codec
	plan            UpgradePlan
	haveCachedInfo  bool
	doShutdowner    func(sdk.Context, UpgradePlan)
	upgradeHandlers map[string]UpgradeHandler
}

const (
	planKey = "plan"
)

func NewKeeper(storeKey sdk.StoreKey, cdc *codec.Codec) Keeper {
	return Keeper{
		storeKey:        storeKey,
		cdc:             cdc,
		haveCachedInfo:  false,
		upgradeHandlers: map[string]UpgradeHandler{},
	}
}

// Sets an upgrade handler for the upgrade specified by name. This handler will be called when the upgrade
// with this name is applied. In order for an upgrade with the given name to proceed, a handler for this upgrade
// must be set even if it is a no-op function.
func (keeper Keeper) SetUpgradeHandler(name string, upgradeHandler UpgradeHandler) {
	keeper.upgradeHandlers[name] = upgradeHandler
}

func (keeper Keeper) ScheduleUpgrade(ctx sdk.Context, plan UpgradePlan) sdk.Error {
	err := plan.ValidateBasic()
	if err != nil {
		return err
	}
	if !plan.Time.IsZero() {
		if !plan.Time.After(ctx.BlockHeader().Time) {
			return sdk.ErrUnknownRequest("Upgrade cannot be scheduled in the past")
		}
	} else {
		if plan.Height <= ctx.BlockHeight() {
			return sdk.ErrUnknownRequest("Upgrade cannot be scheduled in the past")
		}
	}
	store := ctx.KVStore(keeper.storeKey)
	if store.Has(upgradeDoneKey(plan.Name)) {
		return sdk.ErrUnknownRequest(fmt.Sprintf("Upgrade with name %s has already been completed", plan.Name))
	}
	bz := keeper.cdc.MustMarshalBinaryBare(plan)
	store.Set([]byte(planKey), bz)
	return nil
}

func (keeper Keeper) ClearUpgradePlan(ctx sdk.Context) {
	store := ctx.KVStore(keeper.storeKey)
	store.Delete([]byte(planKey))
}

func (plan UpgradePlan) ValidateBasic() sdk.Error {
	if len(plan.Name) == 0 {
		return sdk.ErrUnknownRequest("Name cannot be empty")

	}
	return nil
}

func (keeper Keeper) GetUpgradeInfo(ctx sdk.Context, plan *UpgradePlan) sdk.Error {
	store := ctx.KVStore(keeper.storeKey)
	bz := store.Get([]byte(planKey))
	if bz == nil {
		return sdk.ErrUnknownRequest("Not found")
	}
	keeper.cdc.MustUnmarshalBinaryBare(bz, &plan)
	return nil
}

func (keeper *Keeper) SetDoShutdowner(doShutdowner func(ctx sdk.Context, plan UpgradePlan)) {
	keeper.doShutdowner = doShutdowner
}

func upgradeDoneKey(name string) []byte {
	return []byte(fmt.Sprintf("done/%s", name))
}

func (keeper *Keeper) BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock) {
	blockTime := ctx.BlockHeader().Time
	blockHeight := ctx.BlockHeight()

	if !keeper.haveCachedInfo {
		err := keeper.GetUpgradeInfo(ctx, &keeper.plan)
		if err != nil {
			return
		}
		keeper.haveCachedInfo = true
	}

	if keeper.haveCachedInfo {
		upgradeTime := keeper.plan.Time
		upgradeHeight := keeper.plan.Height
		if (!upgradeTime.IsZero() && !blockTime.Before(upgradeTime)) || upgradeHeight <= blockHeight {
			handler, ok := keeper.upgradeHandlers[keeper.plan.Name]
			if ok {
				// We have an upgrade handler for this upgrade name, so apply the upgrade
				ctx.Logger().Info(fmt.Sprintf("Applying upgrade \"%s\" at height %d", keeper.plan.Name, blockHeight))
				handler(ctx, keeper.plan)
				keeper.ClearUpgradePlan(ctx)
				// Mark this upgrade name as being done so the name can't be reused accidentally
				store := ctx.KVStore(keeper.storeKey)
				store.Set(upgradeDoneKey(keeper.plan.Name), []byte("1"))
			} else {
				// We don't have an upgrade handler for this upgrade name, meaning this software is out of date so shutdown
				ctx.Logger().Error(fmt.Sprintf("UPGRADE \"%s\" NEEDED needed at height %d: %s", keeper.plan.Name, blockHeight, keeper.plan.Memo))
				doShutdowner := keeper.doShutdowner
				if doShutdowner != nil {
					doShutdowner(ctx, keeper.plan)
				} else {
					panic("UPGRADE REQUIRED!")
				}
			}
		}
	}
}
