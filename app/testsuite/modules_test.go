package testsuite

import (
	"testing"

	"github.com/stretchr/testify/suite"

	data "github.com/regen-network/regen-ledger/x/data/client/testsuite"
	"github.com/regen-network/regen-ledger/x/ecocredit/client/testsuite"
)

func TestEcocreditIntegration(t *testing.T) {
	cfg := DefaultConfig()
	testsuite.RunCLITests(t, cfg)
}

func TestDataIntegration(t *testing.T) {
	cfg := DefaultConfig()
	suite.Run(t, data.NewIntegrationTestSuite(cfg))
}
