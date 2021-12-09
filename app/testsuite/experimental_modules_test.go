//go:build experimental
// +build experimental

package testsuite

import (
	"testing"

	"github.com/stretchr/testify/suite"

	data "github.com/regen-network/regen-ledger/x/data/client/testsuite"
	group "github.com/regen-network/regen-ledger/x/group/client/testsuite"
)

func TestGroupIntegration(t *testing.T) {
	cfg := DefaultConfig()

	suite.Run(t, group.NewIntegrationTestSuite(cfg))
}

func TestDataIntegration(t *testing.T) {
	cfg := DefaultConfig()
	suite.Run(t, data.NewIntegrationTestSuite(cfg))
}
