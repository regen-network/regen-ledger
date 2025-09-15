package v6_0_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	upgradetypes "cosmossdk.io/x/upgrade/types"

	"github.com/regen-network/regen-ledger/v7/app/testsuite"
)

type UpgradeTestSuite struct {
	testsuite.UpgradeTestSuite
}

func TestUpgrade(t *testing.T) {
	suite.Run(t, new(UpgradeTestSuite))
}

const upgradeHeight = 5

func (suite *UpgradeTestSuite) TestUpgrade_Upgrade_OK() {
	suite.Setup()

	suite.Ctx = suite.Ctx.WithBlockHeight(upgradeHeight - 1)
	suite.Ctx = suite.Ctx.WithChainID("regen-2")

	plan := upgradetypes.Plan{Name: "v6.0", Height: upgradeHeight}

	err := suite.App.UpgradeKeeper.ScheduleUpgrade(suite.Ctx, plan)
	suite.Require().NoError(err)

	_, exists := suite.App.UpgradeKeeper.GetUpgradePlan(suite.Ctx)
	suite.Require().NoError(exists)

	suite.Ctx = suite.Ctx.WithBlockHeight(upgradeHeight)

	suite.Require().NotPanics(func() {
		suite.App.BeginBlocker(suite.Ctx)
	})
}
