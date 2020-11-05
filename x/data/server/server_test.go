package server

import (
	"github.com/regen-network/regen-ledger/testutil/server/configurator"
	"github.com/regen-network/regen-ledger/x/data/server/testsuite"
	"testing"

	"github.com/stretchr/testify/suite"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func TestServer(t *testing.T) {
	key := sdk.NewKVStoreKey("data")
	addrs := configurator.MakeTestAddresses(2)
	configuratorFixture := configurator.NewFixture(t, []sdk.StoreKey{key}, addrs)
	RegisterServices(key, configuratorFixture)
	s := testsuite.NewIntegrationTestSuite(configuratorFixture)

	suite.Run(t, s)
}
