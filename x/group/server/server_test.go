package server_test

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/stretchr/testify/suite"

	"github.com/regen-network/regen-ledger/app"
	"github.com/regen-network/regen-ledger/testutil/server/configurator"
	"github.com/regen-network/regen-ledger/x/group/server"
	"github.com/regen-network/regen-ledger/x/group/server/testsuite"
	"github.com/regen-network/regen-ledger/x/group/types"
)

func TestServer(t *testing.T) {
	encodingConfig := app.MakeEncodingConfig()

	pKey, pTKey := sdk.NewKVStoreKey(paramstypes.StoreKey), sdk.NewTransientStoreKey(paramstypes.TStoreKey)
	paramSpace := paramstypes.NewSubspace(encodingConfig.Marshaler, encodingConfig.Amino, pKey, pTKey, types.DefaultParamspace)

	key := sdk.NewKVStoreKey(types.ModuleName)
	k := server.NewGroupKeeper(key, paramSpace, baseapp.NewRouter())

	addrs := configurator.MakeTestAddresses(2)
	cfg := configurator.NewFixture(t, []sdk.StoreKey{key}, addrs)
	server.RegisterServices(k, cfg)
	s := testsuite.NewIntegrationTestSuite(cfg)

	suite.Run(t, s)
}
