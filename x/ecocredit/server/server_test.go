package server_test

import (
	"github.com/cosmos/cosmos-sdk/types/module"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/regen-network/regen-ledger/types/module/server"
	ecocreditmodule "github.com/regen-network/regen-ledger/x/ecocredit/module"
	"github.com/regen-network/regen-ledger/x/ecocredit/server/testsuite"
)

func TestServer(t *testing.T) {
	ff := server.NewFixtureFactory(t, 6, module.NewBasicManager(
		ecocreditmodule.Module{},
	))

	s := testsuite.NewIntegrationTestSuite(ff)

	suite.Run(t, s)
}
