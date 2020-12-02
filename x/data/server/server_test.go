package server_test

import (
	module2 "github.com/cosmos/cosmos-sdk/types/module"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/regen-network/regen-ledger/types/module/server"
	datamodule "github.com/regen-network/regen-ledger/x/data/module"
	"github.com/regen-network/regen-ledger/x/data/server/testsuite"
)

func TestServer(t *testing.T) {
	ff := server.NewFixtureFactory(t, 2, module2.NewBasicManager(
		datamodule.Module{},
	))

	s := testsuite.NewIntegrationTestSuite(ff)

	suite.Run(t, s)
}
