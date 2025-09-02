package tests

import (
	"cosmossdk.io/log"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/testutil/testdata"
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
	ecocredittypes "github.com/regen-network/regen-ledger/x/ecocredit/v4"
	basetypes "github.com/regen-network/regen-ledger/x/ecocredit/v4/base/types/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/v4/basket"
	"github.com/regen-network/regen-ledger/x/ecocredit/v4/marketplace"
	markettypes "github.com/regen-network/regen-ledger/x/ecocredit/v4/marketplace/types/v1"
	ecocredit "github.com/regen-network/regen-ledger/x/ecocredit/v4/module"
)

type ecocreditServer struct {
	basetypes.MsgClient
	basetypes.QueryClient
}

type marketServer struct {
	markettypes.MsgClient
	markettypes.QueryClient
}

func NewEcocreditModule(ff fixture.Factory) *ecocredit.Module {
	baseApp := ff.BaseApp()
	cdc := ff.Codec()
	amino := codec.NewLegacyAmino()

	authtypes.RegisterInterfaces(cdc.InterfaceRegistry())
	params.RegisterInterfaces(cdc.InterfaceRegistry())

	authKey := storetypes.NewKVStoreKey(authtypes.StoreKey)
	ecocreditKey := storetypes.NewKVStoreKey(ecocredittypes.ModuleName)
	bankKey := storetypes.NewKVStoreKey(banktypes.StoreKey)
	distKey := storetypes.NewKVStoreKey(disttypes.StoreKey)
	paramsKey := storetypes.NewKVStoreKey(paramstypes.StoreKey)
	tkey := storetypes.NewTransientStoreKey(paramstypes.TStoreKey)

	baseApp.MountStore(authKey, storetypes.StoreTypeIAVL)
	baseApp.MountStore(ecocreditKey, storetypes.StoreTypeIAVL)
	baseApp.MountStore(bankKey, storetypes.StoreTypeIAVL)
	baseApp.MountStore(distKey, storetypes.StoreTypeIAVL)
	baseApp.MountStore(paramsKey, storetypes.StoreTypeIAVL)
	baseApp.MountStore(tkey, storetypes.StoreTypeTransient)

	ecocreditSubspace := paramstypes.NewSubspace(cdc, amino, paramsKey, tkey, ecocredittypes.ModuleName)

	maccPerms := map[string][]string{
		minttypes.ModuleName:       {authtypes.Minter},
		ecocredittypes.ModuleName:  {authtypes.Burner},
		basket.BasketSubModuleName: {authtypes.Burner, authtypes.Minter},
		marketplace.FeePoolName:    {authtypes.Burner},
	}

	govAddr := authtypes.NewModuleAddress(govtypes.ModuleName).String()
	accountKeeper := authkeeper.NewAccountKeeper(
		cdc,
		runtime.NewKVStoreService(authKey),
		authtypes.ProtoBaseAccount,
		maccPerms,
		addresscodec.NewBech32Codec("regen"),
		"regen",
		govAddr,
	)

	bankKeeper := bankkeeper.NewBaseKeeper(cdc, runtime.NewKVStoreService(bankKey), accountKeeper, nil, govAddr, log.NewNopLogger())

	_, _, addr := testdata.KeyTestPubAddr()
	ecocreditModule := ecocredit.NewModule(ecocreditKey, addr, accountKeeper, bankKeeper, ecocreditSubspace, nil)
	ecocreditModule.RegisterInterfaces(cdc.InterfaceRegistry())
	return ecocreditModule
}
