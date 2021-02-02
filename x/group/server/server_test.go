package server_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/regen-network/regen-ledger/types/module"
	"github.com/regen-network/regen-ledger/types/module/server"
	groupmodule "github.com/regen-network/regen-ledger/x/group/module"
	"github.com/regen-network/regen-ledger/x/group/server/testsuite"
)

func TestServer(t *testing.T) {
	ff := server.NewFixtureFactory(t, 6, []module.Module{
		groupmodule.Module{},
	})

	s := testsuite.NewIntegrationTestSuite(ff)

	suite.Run(t, s)
}
