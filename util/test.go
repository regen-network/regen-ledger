package util

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	dbm "github.com/tendermint/tm-db"
)

type TestHarness struct {
	suite.Suite
	Ctx   sdk.Context
	Cms   store.CommitMultiStore
	Cdc   *codec.LegacyAmino
	Db    *dbm.MemDB
	Addr1 sdk.AccAddress
	Addr2 sdk.AccAddress
}

func (s *TestHarness) Setup() {
	s.Db = dbm.NewMemDB()
	s.Cms = store.NewCommitMultiStore(s.Db)
	s.Cdc = codec.NewLegacyAmino()
	s.Ctx = sdk.NewContext(s.Cms, tmproto.Header{}, false, log.NewNopLogger())
	s.Addr1 = sdk.AccAddress{0, 1, 2, 3, 4, 5, 6, 7, 8}
	s.Addr2 = sdk.AccAddress{1, 2, 3, 4, 5, 6, 7, 8, 9}
}
