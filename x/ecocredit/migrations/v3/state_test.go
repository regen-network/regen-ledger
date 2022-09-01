package v3_test

import (
	"testing"

	dbm "github.com/tendermint/tm-db"
	"gotest.tools/v3/assert"

	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/cosmos/cosmos-sdk/orm/model/ormdb"
	"github.com/cosmos/cosmos-sdk/orm/model/ormtable"
	"github.com/cosmos/cosmos-sdk/orm/testing/ormtest"
	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/cosmos/cosmos-sdk/store"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"

	basketapi "github.com/regen-network/regen-ledger/api/regen/ecocredit/basket/v1"
	baseapi "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	basetypes "github.com/regen-network/regen-ledger/x/ecocredit/base/types/v1"
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

	paramStore.WithKeyTable(basetypes.ParamKeyTable())

	creator1 := sdk.AccAddress("creator1")
	creator2 := sdk.AccAddress("creator2")

	// initialize params
	paramStore.SetParamSet(sdkCtx, &basetypes.Params{
		CreditClassFee:       sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(10)), sdk.NewCoin("uregen", sdk.NewInt(2000000))),
		BasketFee:            sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(10)), sdk.NewCoin("uregen", sdk.NewInt(2000000))),
		AllowedClassCreators: []string{creator1.String(), creator2.String()},
		AllowlistEnabled:     true,
	})

	var params basetypes.Params
	paramStore.GetParamSet(sdkCtx, &params)

	ormdb, err := ormdb.NewModuleDB(&ecocredit.ModuleSchema, ormdb.ModuleDBOptions{})
	assert.NilError(t, err)
	coreStore, err := baseapi.NewStateStore(ormdb)
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

	// verify core state migrated to orm table
	classFees, err := coreStore.ClassFeesTable().Get(sdkCtx)
	assert.NilError(t, err)

	assert.Equal(t, len(classFees.Fees), 2)
	assert.Equal(t, classFees.Fees[0].Denom, sdk.DefaultBondDenom)
	assert.Equal(t, classFees.Fees[0].Amount, "10")
	assert.Equal(t, classFees.Fees[1].Denom, "uregen")
	assert.Equal(t, classFees.Fees[1].Amount, "2000000")

	allowedListEnabled, err := coreStore.AllowListEnabledTable().Get(sdkCtx)
	assert.NilError(t, err)
	assert.Equal(t, allowedListEnabled.Enabled, true)

	itr, err := coreStore.AllowedClassCreatorTable().List(sdkCtx, baseapi.AllowedClassCreatorPrimaryKey{})
	assert.NilError(t, err)

	var expected []string
	for itr.Next() {
		val, err := itr.Value()
		assert.NilError(t, err)

		expected = append(expected, sdk.AccAddress(val.Address).String())
	}
	itr.Close()

	assert.DeepEqual(t, params.AllowedClassCreators, expected)
}
