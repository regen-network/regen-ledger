package ormutil

import (
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/cosmos/cosmos-sdk/orm/model/ormdb"
	"github.com/cosmos/cosmos-sdk/store"
	"github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	dbm "github.com/tendermint/tm-db"

	ecocreditv1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
)

func sdkContextForStoreKey(key *types.KVStoreKey) sdk.Context {
	db := dbm.NewMemDB()
	cms := store.NewCommitMultiStore(db)
	cms.MountStoreWithDB(key, sdk.StoreTypeIAVL, db)
	err := cms.LoadLatestVersion()
	if err != nil {
		panic(err)
	}
	return sdk.NewContext(cms, tmproto.Header{}, false, log.NewNopLogger())
}

func TestStoreKeyDB(t *testing.T) {
	storeKey := types.NewKVStoreKey("test")
	db, err := NewStoreKeyDB(
		ormdb.ModuleSchema{FileDescriptors: map[uint32]protoreflect.FileDescriptor{
			1: ecocreditv1.File_regen_ecocredit_v1_state_proto,
		}},
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
