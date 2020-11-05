package server

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"

	"github.com/regen-network/regen-ledger/testutil/server/configurator"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/regen-network/regen-ledger/x/ecocredit/server/testsuite"
)

func registerServers(key sdk.StoreKey, cfg configurator.Fixture) {
	srv := NewServer(key)
	ecocredit.RegisterMsgServer(cfg.MsgServer(), srv)
	ecocredit.RegisterQueryServer(cfg.QueryServer(), srv)
}

func TestServer(t *testing.T) {
	key := sdk.NewKVStoreKey(ecocredit.ModuleName)
	addrs := configurator.MakeTestAddresses(6)
	cfg := configurator.NewFixture(t, []sdk.StoreKey{key}, addrs)
	registerServers(key, cfg)
	s := testsuite.NewIntegrationTestSuite(cfg)

	suite.Run(t, s)
}
