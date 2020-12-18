package server_test

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/bank"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/stretchr/testify/suite"

	"github.com/regen-network/regen-ledger/app"
	"github.com/regen-network/regen-ledger/testutil/server/configurator"
	"github.com/regen-network/regen-ledger/x/group"
	"github.com/regen-network/regen-ledger/x/group/server"
	"github.com/regen-network/regen-ledger/x/group/server/testsuite"
)

func TestServer(t *testing.T) {
	encodingConfig := app.MakeEncodingConfig()
	cdc := encodingConfig.Marshaler

	paramsKey := sdk.NewKVStoreKey(paramstypes.StoreKey)
	groupKey := sdk.NewKVStoreKey(group.StoreKey)
	authKey := sdk.NewKVStoreKey(authtypes.StoreKey)
	bankKey := sdk.NewKVStoreKey(banktypes.StoreKey)
	tkey := sdk.NewTransientStoreKey(paramstypes.TStoreKey)

	groupSubspace := paramstypes.NewSubspace(cdc, encodingConfig.Amino, paramsKey, tkey, group.DefaultParamspace)
	authSubspace := paramstypes.NewSubspace(cdc, encodingConfig.Amino, paramsKey, tkey, authtypes.ModuleName)
	bankSubspace := paramstypes.NewSubspace(cdc, encodingConfig.Amino, paramsKey, tkey, banktypes.ModuleName)

	router := baseapp.NewRouter()
	accountKeeper := authkeeper.NewAccountKeeper(
		cdc, authKey, authSubspace, authtypes.ProtoBaseAccount, map[string][]string{},
	)
	bankKeeper := bankkeeper.NewBaseKeeper(
		cdc, bankKey, accountKeeper, bankSubspace, map[string]bool{},
	)
	router.AddRoute(sdk.NewRoute(banktypes.ModuleName, bank.NewHandler(bankKeeper)))

	addrs := configurator.MakeTestAddresses(4)
	cfg := configurator.NewFixture(t, []sdk.StoreKey{paramsKey, tkey, groupKey, authKey, bankKey}, addrs, cdc)

	server.RegisterServices(groupKey, groupSubspace, router, cfg)
	s := testsuite.NewIntegrationTestSuite(cfg, groupSubspace, bankKeeper, router)

	suite.Run(t, s)
}
