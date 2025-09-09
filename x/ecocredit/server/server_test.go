package server_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"cosmossdk.io/log"
	storetypes "cosmossdk.io/store/types"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdkmodule "github.com/cosmos/cosmos-sdk/types/module"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	disttypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	params "github.com/cosmos/cosmos-sdk/x/params/types/proposal"

	addresscodec "github.com/cosmos/cosmos-sdk/codec/address"
	"github.com/regen-network/regen-ledger/types/v2/testutil/fixture"
	"github.com/regen-network/regen-ledger/x/ecocredit/v4"
	"github.com/regen-network/regen-ledger/x/ecocredit/v4/basket"
	"github.com/regen-network/regen-ledger/x/ecocredit/v4/marketplace"
	"github.com/regen-network/regen-ledger/x/ecocredit/v4/module"
	"github.com/regen-network/regen-ledger/x/ecocredit/v4/server/testsuite"
)

func TestServer(t *testing.T) {
	ff, bankKeeper, accountKeeper := setup(t)
	s := testsuite.NewIntegrationTestSuite(ff, bankKeeper, accountKeeper)
	suite.Run(t, s)
}

func TestGenesis(t *testing.T) {
	ff, bankKeeper, _ := setup(t)
	s := testsuite.NewGenesisTestSuite(ff, bankKeeper)
	suite.Run(t, s)
}

func setup(t *testing.T) (fixture.Factory, bankkeeper.BaseKeeper, authkeeper.AccountKeeper) {
	ff := fixture.NewFixtureFactory(t, 8)
	baseApp := ff.BaseApp()
	cdc := ff.Codec()
	amino := codec.NewLegacyAmino()

	authtypes.RegisterInterfaces(cdc.InterfaceRegistry())
	params.RegisterInterfaces(cdc.InterfaceRegistry())

	authKey := storetypes.NewKVStoreKey(authtypes.StoreKey)
	bankKey := storetypes.NewKVStoreKey(banktypes.StoreKey)
	distKey := storetypes.NewKVStoreKey(disttypes.StoreKey)
	paramsKey := storetypes.NewKVStoreKey(paramstypes.StoreKey)
	ecoKey := storetypes.NewKVStoreKey(ecocredit.ModuleName)
	tkey := storetypes.NewTransientStoreKey(paramstypes.TStoreKey)

	baseApp.MountStore(authKey, storetypes.StoreTypeIAVL)
	baseApp.MountStore(ecoKey, storetypes.StoreTypeIAVL)
	baseApp.MountStore(bankKey, storetypes.StoreTypeIAVL)
	baseApp.MountStore(distKey, storetypes.StoreTypeIAVL)
	baseApp.MountStore(paramsKey, storetypes.StoreTypeIAVL)
	baseApp.MountStore(tkey, storetypes.StoreTypeTransient)

	ecocreditSubspace := paramstypes.NewSubspace(cdc, amino, paramsKey, tkey, ecocredit.ModuleName)

	maccPerms := map[string][]string{
		minttypes.ModuleName:       {authtypes.Minter},
		ecocredit.ModuleName:       {authtypes.Burner},
		basket.BasketSubModuleName: {authtypes.Burner, authtypes.Minter},
		marketplace.FeePoolName:    {authtypes.Burner},
	}

	authority := authtypes.NewModuleAddress(govtypes.ModuleName)
	accountKeeper := authkeeper.NewAccountKeeper(
		cdc,
		runtime.NewKVStoreService(authKey),
		authtypes.ProtoBaseAccount,
		maccPerms,
		addresscodec.NewBech32Codec("regen"),
		"regen",
		authority.String(),
	)

	bankKeeper := bankkeeper.NewBaseKeeper(cdc, runtime.NewKVStoreService(bankKey), accountKeeper, nil, authority.String(), log.NewNopLogger())

	ecocreditModule := module.NewModule(ecoKey, authority, accountKeeper, bankKeeper, ecocreditSubspace, nil)
	ecocreditModule.RegisterInterfaces(cdc.InterfaceRegistry())
	ff.SetModules([]sdkmodule.AppModule{ecocreditModule})

	return ff, bankKeeper, accountKeeper
}
