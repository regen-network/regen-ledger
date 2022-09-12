//go:build !nointegration

package testsuite

import (
	"testing"

	"github.com/stretchr/testify/suite"

	ecocredit "github.com/regen-network/regen-ledger/x/ecocredit/client/testsuite"
)

// func TestDataIntegration(t *testing.T) {
// 	cfg := DefaultConfig()
// 	suite.Run(t, data.NewIntegrationTestSuite(cfg))
// }

func TestEcocreditIntegration(t *testing.T) {
	cfg := DefaultConfig()
	suite.Run(t, ecocredit.NewIntegrationTestSuite(cfg))
}
