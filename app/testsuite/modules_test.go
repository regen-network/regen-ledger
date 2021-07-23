// +build experimental

package testsuite

import (
	"testing"

	"github.com/stretchr/testify/suite"

	ecocredit "github.com/regen-network/regen-ledger/x/ecocredit/client/testsuite"
	group "github.com/regen-network/regen-ledger/x/group/client/testsuite"
)

func TestEcocreditIntegration(t *testing.T) {
	cfg := DefaultConfig()

	suite.Run(t, ecocredit.NewIntegrationTestSuite(cfg))
}

func TestGroupIntegration(t *testing.T) {
	cfg := DefaultConfig()

	suite.Run(t, group.NewIntegrationTestSuite(cfg))
}
