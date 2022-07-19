//go:build !experimental
// +build !experimental

// DONTCOVER

package app

import (
	"fmt"
	"sort"

	"github.com/CosmWasm/wasmd/x/wasm"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/auth/ante"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	vestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	distrclient "github.com/cosmos/cosmos-sdk/x/distribution/client"
	"github.com/cosmos/cosmos-sdk/x/gov"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	paramsclient "github.com/cosmos/cosmos-sdk/x/params/client"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	upgradeclient "github.com/cosmos/cosmos-sdk/x/upgrade/client"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"

	"github.com/regen-network/regen-ledger/types/module/server"
	"github.com/regen-network/regen-ledger/x/data"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	ecocreditcore "github.com/regen-network/regen-ledger/x/ecocredit/client/core"
	"github.com/regen-network/regen-ledger/x/ecocredit/client/marketplace"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
	ecocreditmodule "github.com/regen-network/regen-ledger/x/ecocredit/module"
)

func setCustomModuleBasics() []module.AppModuleBasic {
	return []module.AppModuleBasic{
		gov.NewAppModuleBasic(
			paramsclient.ProposalHandler, distrclient.ProposalHandler,
			upgradeclient.ProposalHandler, upgradeclient.CancelProposalHandler,
			ecocreditcore.CreditTypeProposalHandler, marketplace.AllowDenomProposalHandler,
		),
	}
}

// setCustomModules registers new modules with the server module manager.
// It does nothing here and returns an empty manager since we're not using experimental mode.
func setCustomModules(app *RegenApp, interfaceRegistry types.InterfaceRegistry) *server.Manager {
	return server.NewManager(app.BaseApp, codec.NewProtoCodec(interfaceRegistry))
}
func setCustomKVStoreKeys() []string {
	return []string{}
}

func setCustomMaccPerms() map[string][]string {
	return map[string][]string{}
}

func setCustomOrderBeginBlocker() []string {
	return []string{}
}

func setCustomOrderEndBlocker() []string {
	return []string{}
}

func (app *RegenApp) registerUpgradeHandlers() {
	const upgradeName = "v4.0.0"
	const mainnet = "regen-1"
	const redwood = "regen-redwood-1"

	app.UpgradeKeeper.SetUpgradeHandler(upgradeName, func(ctx sdk.Context, plan upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
		chainId := ctx.ChainID()

		// run state migrations for sdk modules
		toVersion, err := app.mm.RunMigrations(ctx, app.configurator, fromVM)
		if err != nil {
			return nil, err
		}

		// run state migrations for x/ecocredit module
		if err := app.smm.RunMigrations(ctx, app.AppCodec()); err != nil {
			return nil, err
		}
		toVersion[ecocredit.ModuleName] = ecocreditmodule.Module{}.ConsensusVersion()

		// update x/ecocredit basket fee param (the basket fee param key has changed but the
		// value will be the same value as is on regen-1 and regen-redwood-1 at the time of
		// the upgrade; the value will be reset to empty for unsupported chains)
		ecocreditSubspace, _ := app.ParamsKeeper.GetSubspace(ecocredit.ModuleName)
		if chainId == mainnet {
			ecocreditSubspace.Set(ctx, core.KeyBasketFee, sdk.NewCoins(sdk.NewInt64Coin("uregen", 1e9)))
		} else if chainId == redwood {
			ecocreditSubspace.Set(ctx, core.KeyBasketFee, sdk.NewCoins(sdk.NewInt64Coin("uregen", 2e7)))
		} else {
			ecocreditSubspace.Set(ctx, core.KeyBasketFee, sdk.Coins{})
		}

		// recover funds for community member (regen-1 governance proposal #11)
		if chainId == mainnet {
			if err := recoverFunds(ctx, app.AccountKeeper, app.BankKeeper); err != nil {
				return nil, err
			}
		}

		// add name and symbol to regen denom metadata
		if chainId == mainnet || chainId == redwood {
			if err := migrateDenomMetadata(ctx, app.BankKeeper); err != nil {
				return nil, err
			}
		}

		// update denom unit order for all tokens (fixes basket token order)
		migrateDenomUnits(ctx, app.BankKeeper)

		return toVersion, nil
	})

	upgradeInfo, err := app.UpgradeKeeper.ReadUpgradeInfoFromDisk()
	if err != nil {
		panic(err)
	}

	if upgradeInfo.Name == upgradeName && !app.UpgradeKeeper.IsSkipHeight(upgradeInfo.Height) {
		storeUpgrades := storetypes.StoreUpgrades{
			Added: []string{
				data.ModuleName,
			},
		}

		// configure store loader that checks if version == upgradeHeight and applies store upgrades
		app.SetStoreLoader(upgradetypes.UpgradeStoreLoader(upgradeInfo.Height, &storeUpgrades))
	}
}

// migrateDenomMetadata adds missing denom metadata to the REGEN token
func migrateDenomMetadata(ctx sdk.Context, bk bankkeeper.Keeper) error {
	denom := "uregen"
	metadata, found := bk.GetDenomMetaData(ctx, denom)
	if !found {
		return fmt.Errorf("no metadata found for %s denom", denom)
	}

	metadata.Name = "Regen"
	metadata.Symbol = "REGEN"
	bk.SetDenomMetaData(ctx, metadata)

	return nil
}

// migrateDenomUnits update basket metadata denom units list in ascending order
func migrateDenomUnits(ctx sdk.Context, bk bankkeeper.Keeper) {
	metadataList := make([]banktypes.Metadata, 0)
	bk.IterateAllDenomMetaData(ctx, func(m banktypes.Metadata) bool {
		metadata := m
		denomUnits := metadata.DenomUnits
		sort.Slice(denomUnits, func(i, j int) bool {
			return denomUnits[i].Exponent < denomUnits[j].Exponent
		})

		metadata.DenomUnits = denomUnits
		metadataList = append(metadataList, metadata)
		return false
	})

	for _, metadata := range metadataList {
		bk.SetDenomMetaData(ctx, metadata)
	}
}

func recoverFunds(ctx sdk.Context, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper) error {
	// address with funds inaccessible
	lostAddr, err := sdk.AccAddressFromBech32("regen1c3lpjaq0ytdtsrnjqzmtj3hceavl8fe2vtkj7f")
	if err != nil {
		return err
	}

	// address that the community member has access to
	newAddr, err := sdk.AccAddressFromBech32("regen14tpuqrwf95evu3ejm9z7dn20ttcyzqy3jjpfv4")
	if err != nil {
		return err
	}

	lostAccount := ak.GetAccount(ctx, lostAddr)
	if lostAccount == nil {
		return fmt.Errorf("%s account not found", lostAccount.GetAddress().String())
	}

	newAccount := ak.GetAccount(ctx, newAddr)
	if newAccount == nil {
		return fmt.Errorf("%s account not found", newAccount.GetAddress().String())
	}

	va, ok := lostAccount.(*vestingtypes.PeriodicVestingAccount)
	if !ok {
		return fmt.Errorf("%s is not a vesting account", lostAddr)
	}

	vestingPeriods := va.VestingPeriods
	// unlock vesting tokens
	newVestingPeriods := make([]vestingtypes.Period, len(va.VestingPeriods))
	for i, vp := range va.VestingPeriods {
		vp.Length = 0
		newVestingPeriods[i] = vp
	}
	va.VestingPeriods = newVestingPeriods
	ak.SetAccount(ctx, va)

	// send spendable balance from lost account to new account
	spendable := bk.SpendableCoins(ctx, lostAccount.GetAddress())
	if err := bk.SendCoins(ctx, lostAccount.GetAddress(), newAccount.GetAddress(), spendable); err != nil {
		return err
	}

	newPVA := vestingtypes.NewPeriodicVestingAccount(
		authtypes.NewBaseAccount(newAccount.GetAddress(), newAccount.GetPubKey(), newAccount.GetAccountNumber(), newAccount.GetSequence()),
		va.OriginalVesting, va.StartTime, vestingPeriods,
	)
	ak.SetAccount(ctx, newPVA)

	// delete old account
	ak.RemoveAccount(ctx, lostAccount)

	return nil
}

func (app *RegenApp) setCustomAnteHandler(encCfg simappparams.EncodingConfig, wasmKey *sdk.KVStoreKey, _ *wasm.Config) (sdk.AnteHandler, error) {
	return ante.NewAnteHandler(
		ante.HandlerOptions{
			AccountKeeper:   app.AccountKeeper,
			BankKeeper:      app.BankKeeper,
			SignModeHandler: encCfg.TxConfig.SignModeHandler(),
			FeegrantKeeper:  app.FeeGrantKeeper,
			SigGasConsumer:  ante.DefaultSigVerificationGasConsumer,
		},
	)
}

func (app *RegenApp) setCustomModuleManager() []module.AppModule {
	return []module.AppModule{}
}

func (app *RegenApp) setCustomKeepers(_ *baseapp.BaseApp, keys map[string]*sdk.KVStoreKey, appCodec codec.Codec, _ govtypes.Router, _ string,
	_ servertypes.AppOptions,
	_ []wasm.Option) {
}

func setCustomOrderInitGenesis() []string {
	return []string{}
}

func (app *RegenApp) setCustomSimulationManager() []module.AppModuleSimulation {
	return []module.AppModuleSimulation{}
}

func initCustomParamsKeeper(_ *paramskeeper.Keeper) {}

func (app *RegenApp) initializeCustomScopedKeepers() {}
