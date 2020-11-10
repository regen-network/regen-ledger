package server

import (
	"context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/regen-network/regen-ledger/testutil/server"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	dbm "github.com/tendermint/tm-db"
	"testing"
)

type fixtureFactory struct {
	modules map[string]Module
	signers []sdk.AccAddress
}

var _ server.FixtureFactory = fixtureFactory{}

func (mm fixtureFactory) Setup() server.Fixture {
	db := dbm.NewMemDB()

	ms := store.NewCommitMultiStore(db)
	err := ms.LoadLatestVersion()
	require.NoError(mm.t, err)

	ctx := sdk.WrapSDKContext(sdk.NewContext(ms, tmproto.Header{}, false, log.NewNopLogger()))

	ir := types.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(ir)

	router := &router{handlers: map[string]handler{}}

	return fixture{
		router:  router,
		t:       nil,
		signers: nil,
		ctx:     nil,
	}
}

type fixture struct {
	router  *router
	t       *testing.T
	signers []sdk.AccAddress
	ctx     context.Context
}

var _
