package server_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/regen-network/regen-ledger/types/module"
	"github.com/regen-network/regen-ledger/types/module/server"
	ecocreditmodule "github.com/regen-network/regen-ledger/x/ecocredit/module"
	"github.com/regen-network/regen-ledger/x/ecocredit/server/testsuite"
)

func TestServer(t *testing.T) {
	ff := server.NewFixtureFactory(t, 6)
	ff.SetModules([]module.Module{ecocreditmodule.Module{}})
	s := testsuite.NewIntegrationTestSuite(ff)
	suite.Run(t, s)
}
