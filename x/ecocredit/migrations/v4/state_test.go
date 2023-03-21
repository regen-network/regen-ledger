package v4_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	dbm "github.com/tendermint/tm-db"

	basketapi "github.com/regen-network/regen-ledger/api/v2/regen/ecocredit/basket/v1"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/cosmos/cosmos-sdk/orm/model/ormdb"
	"github.com/cosmos/cosmos-sdk/orm/model/ormtable"
	"github.com/cosmos/cosmos-sdk/orm/testing/ormtest"
	"github.com/cosmos/cosmos-sdk/store"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/regen-network/regen-ledger/x/ecocredit/v3"
	v4 "github.com/regen-network/regen-ledger/x/ecocredit/v3/migrations/v4"
)

func TestMainnetMigrateBatchMetadata(t *testing.T) {
	sdkCtx, basketStore := basketSetup(t)
	sdkCtx = sdkCtx.WithChainID("regen-1")

	curator := sdk.MustAccAddressFromBech32("regen1mrvlgpmrjn9s7r7ct69euqfgxjazjt2l5lzqcd")

	// http://mainnet.regen.network:1317/regen/ecocredit/basket/v1/baskets/eco.uC.NCT
	basket := &basketapi.Basket{
		BasketDenom:       "eco.uC.NCT",
		Name:              "NCT",
		DisableAutoRetire: false,
		CreditTypeAbbrev:  "C",
		DateCriteria:      nil,
		Exponent:          6,
		Curator:           curator,
	}

	// add basket to state
	require.NoError(t, basketStore.BasketTable().Insert(sdkCtx, basket))

	// execute state migrations
	require.NoError(t, v4.MigrateState(sdkCtx, basketStore))

	b, err := basketStore.BasketTable().GetByBasketDenom(sdkCtx, "eco.uC.NCT")
	require.NoError(t, err)

	require.Equal(t, "NCT", b.Name)
	require.Equal(t, true, b.DisableAutoRetire)
	require.Equal(t, "C", b.CreditTypeAbbrev)
	require.Equal(t, (*basketapi.DateCriteria)(nil), b.DateCriteria) // TODO
	require.Equal(t, uint32(6), b.Exponent)
	require.Equal(t, curator.Bytes(), b.Curator)
}

func basketSetup(t *testing.T) (sdk.Context, basketapi.StateStore) {
	key := sdk.NewKVStoreKey("basket")

	db := dbm.NewMemDB()
	cms := store.NewCommitMultiStore(db)
	cms.MountStoreWithDB(key, storetypes.StoreTypeIAVL, db)

	require.NoError(t, cms.LoadLatestVersion())

	ormCtx := ormtable.WrapContextDefault(ormtest.NewMemoryBackend())
	sdkCtx := sdk.NewContext(cms, tmproto.Header{}, false, log.NewNopLogger()).WithContext(ormCtx)

	modDB, err := ormdb.NewModuleDB(&ecocredit.ModuleSchema, ormdb.ModuleDBOptions{})
	require.NoError(t, err)

	basketStore, err := basketapi.NewStateStore(modDB)
	require.NoError(t, err)

	return sdkCtx, basketStore
}
