package testsuite

import (
	"testing"

	"github.com/stretchr/testify/suite"

	data "github.com/regen-network/regen-ledger/x/data/client/testsuite"
)

func TestEcocreditIntegration(t *testing.T) {
	/*cfg := DefaultConfig()
	TODO: uncomment after CLI v1 integration https://github.com/regen-network/regen-ledger/issues/876
	ecocredit.RunCLITests(t, cfg)*/
}

func TestDataIntegration(t *testing.T) {
	cfg := DefaultConfig()
	suite.Run(t, data.NewIntegrationTestSuite(cfg))
}
