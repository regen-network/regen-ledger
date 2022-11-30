package v5_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	sdk "github.com/cosmos/cosmos-sdk/types"

	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/regen-network/regen-ledger/v5/app/testsuite"
	"github.com/regen-network/regen-ledger/x/ecocredit/v3"
	ecocreditv1 "github.com/regen-network/regen-ledger/x/ecocredit/v3/base/types/v1"
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
	plan := upgradetypes.Plan{Name: "v5", Height: dummyUpgradeHeight}
	err := suite.App.UpgradeKeeper.ScheduleUpgrade(suite.Ctx, plan)
	suite.Require().NoError(err)
	_, exists := suite.App.UpgradeKeeper.GetUpgradePlan(suite.Ctx)
	suite.Require().True(exists)

	// force the app to have params so migration does not fail
	ss, ok := suite.App.ParamsKeeper.GetSubspace(ecocredit.ModuleName)
	assert.True(suite.T(), ok)
	fees := sdk.NewCoins(sdk.NewInt64Coin("uregen", 10))
	ss.SetParamSet(suite.Ctx, &ecocreditv1.Params{CreditClassFee: fees, BasketFee: fees})

	suite.Ctx = suite.Ctx.WithBlockHeight(dummyUpgradeHeight)
	suite.Require().NotPanics(func() {
		beginBlockRequest := abci.RequestBeginBlock{}
		suite.App.BeginBlocker(suite.Ctx, beginBlockRequest)
	})
}
