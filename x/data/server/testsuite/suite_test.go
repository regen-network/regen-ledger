package testsuite

import (
	"testing"

	"cosmossdk.io/log"
	storetypes "cosmossdk.io/store/types"
	"github.com/stretchr/testify/suite"

	addresscodec "github.com/cosmos/cosmos-sdk/codec/address"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdkmodules "github.com/cosmos/cosmos-sdk/types/module"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	disttypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	params "github.com/cosmos/cosmos-sdk/x/params/types/proposal"

	"github.com/regen-network/regen-ledger/types/v2/testutil/fixture"
	"github.com/regen-network/regen-ledger/x/data/v3"
	datamodule "github.com/regen-network/regen-ledger/x/data/v3/module"
)

func TestServer(t *testing.T) {
	ff := setup(t)
	s := NewIntegrationTestSuite(ff)
	suite.Run(t, s)
}

func setup(t *testing.T) fixture.Factory {
	ff := fixture.NewFixtureFactory(t, 8)
	baseApp := ff.BaseApp()
	cdc := ff.Codec()

	authtypes.RegisterInterfaces(cdc.InterfaceRegistry())
	params.RegisterInterfaces(cdc.InterfaceRegistry())

	authKey := storetypes.NewKVStoreKey(authtypes.StoreKey)
	bankKey := storetypes.NewKVStoreKey(banktypes.StoreKey)
	distKey := storetypes.NewKVStoreKey(disttypes.StoreKey)
	paramsKey := storetypes.NewKVStoreKey(paramstypes.StoreKey)
	dataKey := storetypes.NewKVStoreKey(data.ModuleName)
	tkey := storetypes.NewTransientStoreKey(paramstypes.TStoreKey)

	baseApp.MountStore(authKey, storetypes.StoreTypeIAVL)
	baseApp.MountStore(dataKey, storetypes.StoreTypeIAVL)
	baseApp.MountStore(bankKey, storetypes.StoreTypeIAVL)
	baseApp.MountStore(distKey, storetypes.StoreTypeIAVL)
	baseApp.MountStore(paramsKey, storetypes.StoreTypeIAVL)
	baseApp.MountStore(tkey, storetypes.StoreTypeTransient)

	maccPerms := map[string][]string{
		minttypes.ModuleName: {authtypes.Minter},
	}

	authority := authtypes.NewModuleAddress("gov")

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

	dataMod := datamodule.NewModule(dataKey, accountKeeper, bankKeeper, addresscodec.NewBech32Codec("regen"))
	dataMod.RegisterInterfaces(cdc.InterfaceRegistry())
	ff.SetModules([]sdkmodules.AppModule{dataMod})

	return ff
}
