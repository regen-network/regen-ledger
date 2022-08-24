package v4_test

import (
	"fmt"
	"testing"

	"github.com/cosmos/cosmos-sdk/orm/model/ormtable"
	"github.com/cosmos/cosmos-sdk/orm/testing/ormtest"
	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/cosmos/cosmos-sdk/store"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	dbm "github.com/tendermint/tm-db"
	"gotest.tools/v3/assert"

	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

func TestMigrations(t *testing.T) {

	ecocreditKey := sdk.NewKVStoreKey("ecocredit")
	tecocreditKey := sdk.NewTransientStoreKey("transient_test")
	encCfg := simapp.MakeTestEncodingConfig()
	paramStore := paramtypes.NewSubspace(encCfg.Codec, encCfg.Amino, ecocreditKey, tecocreditKey, ecocredit.ModuleName)

	db := dbm.NewMemDB()
	cms := store.NewCommitMultiStore(db)
	cms.MountStoreWithDB(ecocreditKey, storetypes.StoreTypeIAVL, db)
	cms.MountStoreWithDB(tecocreditKey, storetypes.StoreTypeTransient, db)
	assert.NilError(t, cms.LoadLatestVersion())
	ormCtx := ormtable.WrapContextDefault(ormtest.NewMemoryBackend())
	sdkCtx := sdk.NewContext(cms, tmproto.Header{}, false, log.NewNopLogger()).WithContext(ormCtx)

	// store := sdkCtx.KVStore(ecocreditKey)

	paramStore.WithKeyTable(core.ParamKeyTable())

	paramStore.SetParamSet(sdkCtx, &core.Params{
		CreditClassFee:       sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(10))),
		BasketFee:            sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(10))),
		AllowedClassCreators: []string{},
		AllowlistEnabled:     true,
	})

	var params core.Params
	paramStore.GetParamSet(sdkCtx, &params)
	fmt.Println(params)

	assert.Equal(t, true, false)

}
