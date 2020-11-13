package server

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"

	"github.com/regen-network/regen-ledger/testutil/server/configurator"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/regen-network/regen-ledger/x/ecocredit/server/testsuite"
)

func TestServer(t *testing.T) {
	key := sdk.NewKVStoreKey(ecocredit.ModuleName)
	addrs := configurator.MakeTestAddresses(6)
	cfg := configurator.NewFixture(t, []sdk.StoreKey{key}, addrs)
	RegisterServices(key, cfg)
	s := testsuite.NewIntegrationTestSuite(cfg)

	suite.Run(t, s)
}
