package v5

import (
	"testing"

	"github.com/stretchr/testify/require"
	dbm "github.com/tendermint/tm-db"

	"github.com/cosmos/cosmos-sdk/orm/model/ormdb"
	"github.com/cosmos/cosmos-sdk/orm/model/ormtable"
	"github.com/cosmos/cosmos-sdk/orm/testing/ormtest"
	"github.com/cosmos/cosmos-sdk/store"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	ecocreditv1 "github.com/regen-network/regen-ledger/api/v2/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/v3"
)

func TestMigrate(t *testing.T) {
	ctx, state := setup(t)

	// create some test classes
	clsKey1, err := state.ClassTable().InsertReturningID(ctx, &ecocreditv1.Class{
		Id: "C01",
	})
	require.NoError(t, err)

	clsKey2, err := state.ClassTable().InsertReturningID(ctx, &ecocreditv1.Class{
		Id: "BIO01",
	})
	require.NoError(t, err)
	// create some test projects
	projKey1, err := state.ProjectTable().InsertReturningID(ctx, &ecocreditv1.Project{
		Id:       "C01-001",
		ClassKey: clsKey1,
	})
	require.NoError(t, err)

	projKey2, err := state.ProjectTable().InsertReturningID(ctx, &ecocreditv1.Project{
		Id:       "C01-002",
		ClassKey: clsKey1,
	})
	require.NoError(t, err)

	projKey3, err := state.ProjectTable().InsertReturningID(ctx, &ecocreditv1.Project{
		Id:       "BIO01-001",
		ClassKey: clsKey2,
	})
	require.NoError(t, err)
	// create some test batches
	batchKey1, err := state.BatchTable().InsertReturningID(ctx, &ecocreditv1.Batch{
		Denom:      "C01-001-001",
		ProjectKey: projKey1,
	})
	require.NoError(t, err)

	batchKey2, err := state.BatchTable().InsertReturningID(ctx, &ecocreditv1.Batch{
		Denom:      "C01-001-002",
		ProjectKey: projKey1,
	})
	require.NoError(t, err)

	batchKey3, err := state.BatchTable().InsertReturningID(ctx, &ecocreditv1.Batch{
		Denom:      "C01-002-001",
		ProjectKey: projKey2,
	})
	require.NoError(t, err)

	batchKey4, err := state.BatchTable().InsertReturningID(ctx, &ecocreditv1.Batch{
		Denom:      "BIO01-001-001",
		ProjectKey: projKey3,
	})
	require.NoError(t, err)

	batchKey5, err := state.BatchTable().InsertReturningID(ctx, &ecocreditv1.Batch{
		Denom:      "BIO01-001-002",
		ProjectKey: projKey3,
	})
	require.NoError(t, err)

	// run the migration
	require.NoError(t, MigrateState(ctx, state))

	// check that the class keys on the batches are set correctly
	batch1, err := state.BatchTable().Get(ctx, batchKey1)
	require.NoError(t, err)
	require.Equal(t, clsKey1, batch1.ClassKey)

	batch2, err := state.BatchTable().Get(ctx, batchKey2)
	require.NoError(t, err)
	require.Equal(t, clsKey1, batch2.ClassKey)

	batch3, err := state.BatchTable().Get(ctx, batchKey3)
	require.NoError(t, err)
	require.Equal(t, clsKey1, batch3.ClassKey)

	batch4, err := state.BatchTable().Get(ctx, batchKey4)
	require.NoError(t, err)
	require.Equal(t, clsKey2, batch4.ClassKey)

	batch5, err := state.BatchTable().Get(ctx, batchKey5)
	require.NoError(t, err)
	require.Equal(t, clsKey2, batch5.ClassKey)

	// check that enrollment entries are created for all project class relationships
	enrollment1, err := state.ProjectEnrollmentTable().Get(ctx, projKey1, clsKey1)
	require.NoError(t, err)
	require.Equal(t, projKey1, enrollment1.ProjectKey)
	require.Equal(t, clsKey1, enrollment1.ClassKey)

	enrollment2, err := state.ProjectEnrollmentTable().Get(ctx, projKey2, clsKey1)
	require.NoError(t, err)
	require.Equal(t, projKey2, enrollment2.ProjectKey)
	require.Equal(t, clsKey1, enrollment2.ClassKey)

	enrollment3, err := state.ProjectEnrollmentTable().Get(ctx, projKey3, clsKey2)
	require.NoError(t, err)
	require.Equal(t, projKey3, enrollment3.ProjectKey)

}

func setup(t *testing.T) (sdk.Context, ecocreditv1.StateStore) {
	ecocreditKey := sdk.NewKVStoreKey("ecocredit")

	db := dbm.NewMemDB()
	cms := store.NewCommitMultiStore(db)
	cms.MountStoreWithDB(ecocreditKey, storetypes.StoreTypeIAVL, db)

	require.NoError(t, cms.LoadLatestVersion())

	ormCtx := ormtable.WrapContextDefault(ormtest.NewMemoryBackend())
	sdkCtx := sdk.NewContext(cms, tmproto.Header{}, false, log.NewNopLogger()).WithContext(ormCtx)

	modDB, err := ormdb.NewModuleDB(&ecocredit.ModuleSchema, ormdb.ModuleDBOptions{})
	require.NoError(t, err)

	baseStore, err := ecocreditv1.NewStateStore(modDB)
	require.NoError(t, err)

	return sdkCtx, baseStore
}
