//go:build !nointegration

package testsuite

import (
	"testing"

	"github.com/stretchr/testify/suite"

	data "github.com/regen-network/regen-ledger/x/data/v2/client/testsuite"
	ecocredit "github.com/regen-network/regen-ledger/x/ecocredit/v3/client/testsuite"
)

func TestDataIntegration(t *testing.T) {
	cfg := DefaultConfig()
	suite.Run(t, data.NewIntegrationTestSuite(cfg))
}

func TestEcocreditIntegration(t *testing.T) {
	cfg := DefaultConfig()
	suite.Run(t, ecocredit.NewIntegrationTestSuite(cfg))
}
