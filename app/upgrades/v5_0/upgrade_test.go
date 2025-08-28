package v5_0_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	upgradetypes "cosmossdk.io/x/upgrade/types"

	"github.com/regen-network/regen-ledger/v7/app/testsuite"
	// "github.com/regen-network/regen-ledger/x/ecocredit/v4"
	// ecocreditv1 "github.com/regen-network/regen-ledger/x/ecocredit/v4/base/types/v1"
)

type UpgradeTestSuite struct {
	testsuite.UpgradeTestSuite
}

func TestUpgrade(t *testing.T) {
	suite.Run(t, new(UpgradeTestSuite))
}

const dummyUpgradeHeight = 5

func (suite *UpgradeTestSuite) TestV5Upgrade() {
	suite.Setup()

	suite.Ctx = suite.Ctx.WithBlockHeight(dummyUpgradeHeight - 1)
	suite.Ctx = suite.Ctx.WithChainID("testing-1")

	plan := upgradetypes.Plan{Name: "v5.0", Height: dummyUpgradeHeight}
	err := suite.App.UpgradeKeeper.ScheduleUpgrade(suite.Ctx, plan)
	suite.Require().NoError(err)
	_, err = suite.App.UpgradeKeeper.GetUpgradePlan(suite.Ctx)
	suite.Require().NoError(err)

	// force the app to have params so migration does not fail
	// ss, ok := suite.App.ParamsKeeper.GetSubspace(ecocredit.ModuleName)
	// assert.True(suite.T(), ok)
	// fees := sdk.NewCoins(sdk.NewInt64Coin("uregen", 10))
	// ss.SetParamSet(suite.Ctx, &ecocreditv1.Params{CreditClassFee: fees, BasketFee: fees})

	suite.Ctx = suite.Ctx.WithBlockHeight(dummyUpgradeHeight)
	suite.Require().NotPanics(func() {
		suite.App.BeginBlocker(suite.Ctx)
	})
}
