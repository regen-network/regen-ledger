//go:build !nointegration

package testsuite

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/cosmos/cosmos-sdk/testutil/network"

	data "github.com/regen-network/regen-ledger/x/data/v3/client/testsuite"
	ecocredit "github.com/regen-network/regen-ledger/x/ecocredit/v3/client/testsuite"
)

func TestDataIntegration(t *testing.T) {
	cfg := network.DefaultConfig(NewTestNetworkFixture)
	suite.Run(t, data.NewIntegrationTestSuite(cfg))
}

func TestEcocreditIntegration(t *testing.T) {
	cfg := network.DefaultConfig(NewTestNetworkFixture)
	suite.Run(t, ecocredit.NewIntegrationTestSuite(cfg))
}
