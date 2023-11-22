package ormstore

import (
	"testing"

	"github.com/stretchr/testify/require"
	dbm "github.com/tendermint/tm-db"

	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	ormv1alpha1 "github.com/cosmos/cosmos-sdk/api/cosmos/orm/v1alpha1"
	"github.com/cosmos/cosmos-sdk/orm/model/ormdb"
	"github.com/cosmos/cosmos-sdk/store"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	ecocreditv1 "github.com/regen-network/regen-ledger/api/v2/regen/ecocredit/v1"
)

func sdkContextForStoreKey(key *storetypes.KVStoreKey) sdk.Context {
	db := dbm.NewMemDB()
	cms := store.NewCommitMultiStore(db)
	cms.MountStoreWithDB(key, storetypes.StoreTypeIAVL, db)
	err := cms.LoadLatestVersion()
	if err != nil {
		panic(err)
	}
	return sdk.NewContext(cms, tmproto.Header{}, false, log.NewNopLogger())
}

func TestStoreKeyDB(t *testing.T) {
	storeKey := storetypes.NewKVStoreKey("test")
	db, err := NewStoreKeyDB(
		&ormv1alpha1.ModuleSchemaDescriptor{
			SchemaFile: []*ormv1alpha1.ModuleSchemaDescriptor_FileEntry{
				{Id: 1, ProtoFileName: ecocreditv1.File_regen_ecocredit_v1_state_proto.Path()},
			},
			Prefix: nil,
		},
		storeKey,
		ormdb.ModuleDBOptions{},
	)
	require.NoError(t, err)
	sdkCtx := sdkContextForStoreKey(storeKey)
	ctx := sdk.WrapSDKContext(sdkCtx)

	creditTypeTable := db.GetTable(&ecocreditv1.CreditType{})
	require.NotNil(t, creditTypeTable)

	require.NoError(t, creditTypeTable.Save(ctx, &ecocreditv1.CreditType{
		Name:         "carbon",
		Abbreviation: "C",
		Unit:         "tons of co2e",
		Precision:    6,
	}))

	creditType := &ecocreditv1.CreditType{
		Abbreviation: "C",
	}
	found, err := creditTypeTable.Get(ctx, creditType)
	require.NoError(t, err)
	require.True(t, found)
	require.Equal(t, "C", creditType.Abbreviation)
	require.Equal(t, "carbon", creditType.Name)
	require.Equal(t, "tons of co2e", creditType.Unit)
	require.Equal(t, uint32(6), creditType.Precision)
}
