package upgrade

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

type Keeper struct {
	storeKey          sdk.StoreKey
	cdc               *codec.Codec
	lastUpgradeHeight int64
	info              UpgradeInfo
	haveCachedInfo    bool
	willUpgrader      func(info UpgradeInfo)
	beforeShutdowner  func(info UpgradeInfo)
	doShutdowner      func()
}

type UpgradeInfo struct {
	// The height at which the upgrade must be performed
	Height int64 `json:"height"`

	// Any application specific upgrade info to be included on-chain
	// such as a git commit that validators could automatically upgrade to
	Memo string `json:"memo,omitempty"`
}

const (
	infoKey = "info"
)

func NewKeeper(storeKey sdk.StoreKey, cdc *codec.Codec, lastUpgradeHeight int64) Keeper {
	return Keeper{
		storeKey:          storeKey,
		cdc:               cdc,
		lastUpgradeHeight: lastUpgradeHeight,
		haveCachedInfo:    false,
	}
}

func (keeper Keeper) ScheduleUpgrade(ctx sdk.Context, info UpgradeInfo) {
	store := ctx.KVStore(keeper.storeKey)
	bz, err := keeper.cdc.MarshalBinaryBare(info)
	if err != nil {
		panic(err)
	}
	store.Set([]byte(infoKey), bz)
}

func (keeper Keeper) GetUpgradeInfo(ctx sdk.Context, info *UpgradeInfo) sdk.Error {
	store := ctx.KVStore(keeper.storeKey)
	bz := store.Get([]byte(infoKey))
	if bz == nil {
		return sdk.ErrUnknownRequest("Not found")
	}
	marshalErr := keeper.cdc.UnmarshalBinaryBare(bz, &info)
	if marshalErr != nil {
		return sdk.ErrUnknownRequest(marshalErr.Error())
	}
	return nil
}

func (keeper Keeper) SetWillUpgrader(ctx sdk.Context, willUpgrader func(info UpgradeInfo)) {
	keeper.willUpgrader = willUpgrader
}

func (keeper Keeper) SetBeforeShutdowner(ctx sdk.Context, beforeShutdowner func(info UpgradeInfo)) {
	keeper.beforeShutdowner = beforeShutdowner
}

func (keeper Keeper) SetDoShutdowner(ctx sdk.Context, doShutdowner func()) {
	keeper.doShutdowner = doShutdowner
}

func (keeper Keeper) BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock) {
	if !keeper.haveCachedInfo {
		err := keeper.GetUpgradeInfo(ctx, &keeper.info)
		if err != nil {
			return
		}
		keeper.haveCachedInfo = true
		if keeper.info.Height > keeper.lastUpgradeHeight {
			willUpgrader := keeper.willUpgrader
			if willUpgrader != nil {
				willUpgrader(keeper.info)
			}
		}
	}

	upgradeHeight := keeper.info.Height
	if upgradeHeight > keeper.lastUpgradeHeight && ctx.BlockHeight() >= upgradeHeight {
		beforeShutdowner := keeper.beforeShutdowner
		if beforeShutdowner != nil {
			beforeShutdowner(keeper.info)
		}

		doShutdowner := keeper.doShutdowner
		if doShutdowner != nil {
			doShutdowner()
		} else {
			panic(fmt.Sprintf("UPGRADE NEEDED at height %d: %s", upgradeHeight, keeper.info.Memo))
		}
	}
}
