package v7_0_test

import (
	"fmt"
	"testing"

	"cosmossdk.io/core/header"
	"github.com/cometbft/cometbft/abci/types"
	"github.com/stretchr/testify/suite"

	upgradetypes "cosmossdk.io/x/upgrade/types"
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"

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

	plan := upgradetypes.Plan{Name: "v7_0", Height: upgradeHeight}
	err := suite.App.UpgradeKeeper.ScheduleUpgrade(suite.Ctx, plan)
	suite.Require().NoError(err)

	_, err = suite.App.UpgradeKeeper.GetUpgradePlan(suite.Ctx)
	suite.Require().NoError(err)

	params := suite.App.WasmKeeper.GetParams(suite.Ctx)
	fmt.Println(params)
	suite.Require().Len(params.CodeUploadAccess.Addresses, 0)
	suite.Require().Equal(params.CodeUploadAccess.Permission, wasmtypes.AccessTypeNobody)
	suite.Require().Equal(params.CodeUploadAccess.Permission, wasmtypes.AccessTypeNobody)

	ctx := suite.Ctx.WithBlockHeight(upgradeHeight).WithHeaderInfo(header.Info{Height: upgradeHeight})
	suite.Require().NotPanics(func() {
		_, err := suite.App.PreBlocker(ctx, &types.RequestFinalizeBlock{
			Height: upgradeHeight,
		})
		suite.Require().NoError(err)
	})

	updatedParams := suite.App.WasmKeeper.GetParams(ctx)
	suite.Require().Len(updatedParams.CodeUploadAccess.Addresses, 1)
	suite.Require().Equal(updatedParams.CodeUploadAccess.Addresses[0], suite.App.GovKeeper.GetAuthority())
	suite.Require().Equal(updatedParams.InstantiateDefaultPermission, wasmtypes.AccessTypeAnyOfAddresses)
	suite.Require().Equal(updatedParams.CodeUploadAccess.Permission, wasmtypes.AccessTypeAnyOfAddresses)
}
