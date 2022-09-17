package v3_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	dbm "github.com/tendermint/tm-db"

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
	require.NoError(t, cms.LoadLatestVersion())
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
	require.NoError(t, err)
	baseStore, err := baseapi.NewStateStore(ormdb)
	require.NoError(t, err)

	basketStore, err := basketapi.NewStateStore(ormdb)
	require.NoError(t, err)

	require.NoError(t, v3.MigrateState(sdkCtx, baseStore, basketStore, paramStore))

	// verify basket params migrated to orm table
	basketFee, err := basketStore.BasketFeeTable().Get(sdkCtx)
	require.NoError(t, err)

	require.NotEmpty(t, basketFee.Fee)
	require.Equal(t, basketFee.Fee.Denom, sdk.DefaultBondDenom)
	require.Equal(t, basketFee.Fee.Amount, "10")

	// verify core state migrated to orm table
	classFee, err := baseStore.ClassFeeTable().Get(sdkCtx)
	require.NoError(t, err)

	require.NotEmpty(t, classFee.Fee)
	require.Equal(t, classFee.Fee.Denom, sdk.DefaultBondDenom)
	require.Equal(t, classFee.Fee.Amount, "10")

	allowlist, err := baseStore.ClassCreatorAllowlistTable().Get(sdkCtx)
	require.NoError(t, err)
	require.Equal(t, allowlist.Enabled, true)

	itr, err := baseStore.AllowedClassCreatorTable().List(sdkCtx, baseapi.AllowedClassCreatorPrimaryKey{})
	require.NoError(t, err)

	var expected []string
	for itr.Next() {
		val, err := itr.Value()
		require.NoError(t, err)

		expected = append(expected, sdk.AccAddress(val.Address).String())
	}
	itr.Close()

	require.Equal(t, params.AllowedClassCreators, expected)
}
