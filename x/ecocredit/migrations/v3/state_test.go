package v3_test

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/orm/model/ormdb"
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

	basketapi "github.com/regen-network/regen-ledger/api/regen/ecocredit/basket/v1"
	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
	v3 "github.com/regen-network/regen-ledger/x/ecocredit/migrations/v3"
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

	paramStore.WithKeyTable(core.ParamKeyTable())

	// initialize params
	paramStore.SetParamSet(sdkCtx, &core.Params{
		CreditClassFee:       sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(10))),
		BasketFee:            sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(10)), sdk.NewCoin("uregen", sdk.NewInt(2000000))),
		AllowedClassCreators: []string{},
		AllowlistEnabled:     true,
	})

	var params core.Params
	paramStore.GetParamSet(sdkCtx, &params)

	ormdb, err := ormdb.NewModuleDB(&ecocredit.ModuleSchema, ormdb.ModuleDBOptions{})
	assert.NilError(t, err)
	coreStore, err := api.NewStateStore(ormdb)
	assert.NilError(t, err)

	basketStore, err := basketapi.NewStateStore(ormdb)
	assert.NilError(t, err)

	assert.NilError(t, v3.MigrateState(sdkCtx, coreStore, basketStore, paramStore))

	// verify basket params migrated to orm table
	basketFees, err := basketStore.BasketFeesTable().Get(sdkCtx)
	assert.NilError(t, err)

	assert.Equal(t, len(basketFees.Fees), 2)
	assert.Equal(t, basketFees.Fees[0].Denom, sdk.DefaultBondDenom)
	assert.Equal(t, basketFees.Fees[0].Amount, "10")
	assert.Equal(t, basketFees.Fees[1].Denom, "uregen")
	assert.Equal(t, basketFees.Fees[1].Amount, "2000000")
}