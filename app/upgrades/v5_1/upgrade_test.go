package v5_1_test

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

func (suite *UpgradeTestSuite) TestUpgrade() {
	suite.Setup()

	suite.Ctx = suite.Ctx.WithBlockHeight(upgradeHeight - 1)
	suite.Ctx = suite.Ctx.WithChainID("regen-1")

	plan := upgradetypes.Plan{Name: "v5.1", Height: upgradeHeight}
	err := suite.App.UpgradeKeeper.ScheduleUpgrade(suite.Ctx, plan)
	suite.Require().NoError(err)

	_, err = suite.App.UpgradeKeeper.GetUpgradePlan(suite.Ctx)
	suite.Require().NoError(err)

	suite.Ctx = suite.Ctx.WithBlockHeight(upgradeHeight)

	suite.Require().NotPanics(func() {
		suite.App.BeginBlocker(suite.Ctx)
	})
}
