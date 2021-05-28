// +build experimental
// +build norace

package testsuite_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/regen-network/regen-ledger/app/testsuite"
	group "github.com/regen-network/regen-ledger/x/group/client/testsuite"
)

func TestModules(t *testing.T) {
	cfg := testsuite.DefaultConfig()

	suite.Run(t, group.NewIntegrationTestSuite(cfg))
}
