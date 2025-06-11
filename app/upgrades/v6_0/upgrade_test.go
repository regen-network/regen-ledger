package v6_0_test

import (
	"testing"

	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/stretchr/testify/suite"

	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
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
	suite.Require().True(exists)

	suite.Ctx = suite.Ctx.WithBlockHeight(upgradeHeight)

	suite.Require().NotPanics(func() {
		beginBockRequest := abci.RequestBeginBlock{}
		suite.App.BeginBlocker(suite.Ctx, beginBockRequest)
	})
}
