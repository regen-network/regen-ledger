package v5_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/regen-network/regen-ledger/v4/app/testutil"
)

type UpgradeTestSuite struct {
	testutil.KeeperTestHelper
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(UpgradeTestSuite))
}

const dummyUpgradeHeight = 5

func (suite *UpgradeTestSuite) TestV5Upgrade() {
	suite.Setup()

	suite.Ctx = suite.Ctx.WithBlockHeight(dummyUpgradeHeight - 1)
	plan := upgradetypes.Plan{Name: "v5", Height: dummyUpgradeHeight}
	err := suite.App.UpgradeKeeper.ScheduleUpgrade(suite.Ctx, plan)
	suite.Require().NoError(err)
	_, exists := suite.App.UpgradeKeeper.GetUpgradePlan(suite.Ctx)
	suite.Require().True(exists)

	suite.Ctx = suite.Ctx.WithBlockHeight(dummyUpgradeHeight)
	suite.Require().NotPanics(func() {
		beginBlockRequest := abci.RequestBeginBlock{}
		suite.App.BeginBlocker(suite.Ctx, beginBlockRequest)
	})
}
