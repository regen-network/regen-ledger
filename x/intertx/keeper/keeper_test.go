package keeper

import (
	"context"

	"github.com/golang/mock/gomock"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
	dbm "github.com/tendermint/tm-db"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/orm/model/ormtable"
	"github.com/cosmos/cosmos-sdk/orm/testing/ormtest"
	"github.com/cosmos/cosmos-sdk/store"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/regen-network/regen-ledger/x/intertx/keeper/mocks"
)

type baseSuite struct {
	t      gocuke.TestingT
	ctx    context.Context
	sdkCtx sdk.Context
	cdc    *codec.ProtoCodec
	k      Keeper
	cap    *mocks.MockCapabilityKeeper
	ica    *mocks.MockICAControllerKeeper
	addrs  []sdk.AccAddress
}

func setupBase(t gocuke.TestingT) *baseSuite {
	s := &baseSuite{t: t}

	// set up store
	db := dbm.NewMemDB()
	cms := store.NewCommitMultiStore(db)
	sk := storetypes.NewKVStoreKey("test")
	cms.MountStoreWithDB(sk, storetypes.StoreTypeIAVL, db)
	require.NoError(t, cms.LoadLatestVersion())

	// set up context
	ormCtx := ormtable.WrapContextDefault(ormtest.NewMemoryBackend())
	s.sdkCtx = sdk.NewContext(cms, tmproto.Header{}, false, log.NewNopLogger()).WithContext(ormCtx)
	s.ctx = sdk.WrapSDKContext(s.sdkCtx)

	ir := codectypes.NewInterfaceRegistry()
	s.cdc = codec.NewProtoCodec(ir)
	ctrl := gomock.NewController(t)
	s.cap = mocks.NewMockCapabilityKeeper(ctrl)
	s.ica = mocks.NewMockICAControllerKeeper(ctrl)
	s.k = NewKeeper(s.cdc, s.ica, s.cap)

	// set up addresses
	_, _, addr1 := testdata.KeyTestPubAddr()
	_, _, addr2 := testdata.KeyTestPubAddr()
	s.addrs = append(s.addrs, addr1)
	s.addrs = append(s.addrs, addr2)

	return s
}
