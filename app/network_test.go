package app_test

import (
	"testing"

	"github.com/regen-network/regen-ledger/testutil/network"
	ecocreditsuite "github.com/regen-network/regen-ledger/x/ecocredit/server/testsuite"
	"github.com/stretchr/testify/suite"
)

func TestNetwork(t *testing.T) {
	cfg := network.DefaultConfig()
	cfg.NumValidators = 6
	ff := network.NewFixtureFactory(t, cfg)
	s1 := ecocreditsuite.NewIntegrationTestSuite(ff)
	suite.Run(t, s1)
}
