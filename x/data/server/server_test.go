package server_test

import (
	"github.com/regen-network/regen-ledger/types/module"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/regen-network/regen-ledger/types/module/server"
	datamodule "github.com/regen-network/regen-ledger/x/data/module"
	"github.com/regen-network/regen-ledger/x/data/server/testsuite"
)

func TestServer(t *testing.T) {
	ff := server.NewFixtureFactory(t, 2, module.Modules{
		datamodule.Module{},
	})

	s := testsuite.NewIntegrationTestSuite(ff)

	suite.Run(t, s)
}
