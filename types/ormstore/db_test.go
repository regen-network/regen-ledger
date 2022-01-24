package ormstore

import (
	"testing"

	ecocreditv1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"

	"github.com/cosmos/cosmos-sdk/store"
	storetypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	dbm "github.com/tendermint/tm-db"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/cosmos/cosmos-sdk/orm/model/ormdb"
	"github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func sdkContextForStoreKey(key *types.KVStoreKey) sdk.Context {
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
	storeKey := types.NewKVStoreKey("test")
	db, err := NewStoreKeyDB(
		ormdb.ModuleSchema{FileDescriptors: map[uint32]protoreflect.FileDescriptor{
			1: ecocreditv1.File_regen_ecocredit_v1_state_proto,
		}},
		storeKey,
		ormdb.ModuleDBOptions{},
	)
	require.NoError(t, err)
	ctx := sdkContextForStoreKey(storeKey)

}
