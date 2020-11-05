package server

import (
	"testing"

	"github.com/stretchr/testify/suite"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/regen-network/regen-ledger/testutil/server/configurator"
	"github.com/regen-network/regen-ledger/x/ecocredit/server/testsuite"
)

func TestServer(t *testing.T) {
	key := sdk.NewKVStoreKey("ecocredit")
	addrs := configurator.MakeTestAddresses(6)
	configuratorFixture := configurator.NewFixture(t, []sdk.StoreKey{key}, addrs)
	RegisterServices(key, configuratorFixture)
	s := testsuite.NewIntegrationTestSuite(configuratorFixture)

	suite.Run(t, s)
}
