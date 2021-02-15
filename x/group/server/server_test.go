package server_test

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/bank"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"

	"github.com/stretchr/testify/suite"

	"github.com/regen-network/regen-ledger/types/module"
	"github.com/regen-network/regen-ledger/types/module/server"
	groupmodule "github.com/regen-network/regen-ledger/x/group/module"
	"github.com/regen-network/regen-ledger/x/group/server/testsuite"
)

func TestServer(t *testing.T) {
	ff := server.NewFixtureFactory(t, 6)
	cdc := ff.Codec()
	// Setting up bank keeper
	banktypes.RegisterInterfaces(cdc.InterfaceRegistry())
	authtypes.RegisterInterfaces(cdc.InterfaceRegistry())

	paramsKey := sdk.NewKVStoreKey(paramstypes.StoreKey)
	authKey := sdk.NewKVStoreKey(authtypes.StoreKey)
	bankKey := sdk.NewKVStoreKey(banktypes.StoreKey)
	tkey := sdk.NewTransientStoreKey(paramstypes.TStoreKey)
	amino := codec.NewLegacyAmino()

	authSubspace := paramstypes.NewSubspace(cdc, amino, paramsKey, tkey, authtypes.ModuleName)
	bankSubspace := paramstypes.NewSubspace(cdc, amino, paramsKey, tkey, banktypes.ModuleName)

	accountKeeper := authkeeper.NewAccountKeeper(
		cdc, authKey, authSubspace, authtypes.ProtoBaseAccount, map[string][]string{},
	)

	bankKeeper := bankkeeper.NewBaseKeeper(
		cdc, bankKey, accountKeeper, bankSubspace, map[string]bool{},
	)

	baseApp := ff.BaseApp()

	baseApp.Router().AddRoute(sdk.NewRoute(banktypes.ModuleName, bank.NewHandler(bankKeeper)))
	baseApp.MountStore(tkey, sdk.StoreTypeTransient)
	baseApp.MountStore(paramsKey, sdk.StoreTypeIAVL)
	baseApp.MountStore(authKey, sdk.StoreTypeIAVL)
	baseApp.MountStore(bankKey, sdk.StoreTypeIAVL)

	ff.SetModules([]module.Module{groupmodule.Module{AccountKeeper: accountKeeper}})

	s := testsuite.NewIntegrationTestSuite(ff, accountKeeper, bankKeeper)

	suite.Run(t, s)
}
