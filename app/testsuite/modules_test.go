package testsuite

import (
	ecocredit "github.com/regen-network/regen-ledger/x/ecocredit/client/testsuite"
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestEcocreditIntegration(t *testing.T) {
	cfg := DefaultConfig()

	suite.Run(t, ecocredit.NewIntegrationTestSuite(cfg))
}
