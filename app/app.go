package app

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/cosmos/cosmos-sdk/x/auth/genaccounts"
	"github.com/cosmos/cosmos-sdk/x/crisis"
	distr "github.com/cosmos/cosmos-sdk/x/distribution"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	"github.com/cosmos/cosmos-sdk/x/gov"
	"github.com/cosmos/cosmos-sdk/x/mint"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/cosmos/cosmos-sdk/x/slashing"
	"github.com/cosmos/cosmos-sdk/x/staking"
	"github.com/cosmos/cosmos-sdk/x/upgrade"
	"github.com/regen-network/regen-ledger/index/postgresql"
	"github.com/regen-network/regen-ledger/x/geo"
	"github.com/spf13/viper"
	"github.com/tendermint/tendermint/libs/cli"
	"github.com/tendermint/tendermint/libs/log"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	//"os"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/bank"

	bam "github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
	cmn "github.com/tendermint/tendermint/libs/common"
	dbm "github.com/tendermint/tendermint/libs/db"
)

const (
	appName = "xrn"
)

var (
	// default home directories for xrncli
	DefaultCLIHome = os.ExpandEnv("$HOME/.xrncli")

	// default home directories for xrnd
	DefaultNodeHome = os.ExpandEnv("$HOME/.xrnd")

	// The ModuleBasicManager is in charge of setting up basic,
	// non-dependant module elements, such as codec registration
	// and genesis verification.
	ModuleBasics sdk.ModuleBasicManager
)

func init() {
	ModuleBasics = sdk.NewModuleBasicManager(
		genaccounts.AppModuleBasic{},
		genutil.AppModuleBasic{},
		auth.AppModuleBasic{},
		bank.AppModuleBasic{},
		staking.AppModuleBasic{},
		mint.AppModuleBasic{},
		distr.AppModuleBasic{},
		gov.AppModuleBasic{},
		params.AppModuleBasic{},
		crisis.AppModuleBasic{},
		slashing.AppModuleBasic{},
		geo.AppModuleBasic{},
		upgrade.AppModuleBasic{},
	)
}

// custom tx codec
func MakeCodec() *codec.Codec {
	var cdc = codec.New()
	ModuleBasics.RegisterCodec(cdc)
	sdk.RegisterCodec(cdc)
	codec.RegisterCrypto(cdc)
	return cdc
}

// Extended ABCI application
type XrnApp struct {
	*bam.BaseApp
	cdc *codec.Codec

	invCheckPeriod uint

	// keys to access the substores
	keyMain          *sdk.KVStoreKey
	keyAccount       *sdk.KVStoreKey
	keyStaking       *sdk.KVStoreKey
	tkeyStaking      *sdk.TransientStoreKey
	keySlashing      *sdk.KVStoreKey
	keyMint          *sdk.KVStoreKey
	keyDistr         *sdk.KVStoreKey
	tkeyDistr        *sdk.TransientStoreKey
	keyGov           *sdk.KVStoreKey
	keyFeeCollection *sdk.KVStoreKey
	keyParams        *sdk.KVStoreKey
	tkeyParams       *sdk.TransientStoreKey
	upgradeStoreKey    *sdk.KVStoreKey
	//dataStoreKey       *sdk.KVStoreKey
	//schemaStoreKey     *sdk.KVStoreKey
	//espStoreKey        *sdk.KVStoreKey
	geoStoreKey *sdk.KVStoreKey
	//agentStoreKey      *sdk.KVStoreKey
	//proposalStoreKey   *sdk.KVStoreKey
	//consortiumStoreKey *sdk.KVStoreKey

	// keepers
	accountKeeper       auth.AccountKeeper
	feeCollectionKeeper auth.FeeCollectionKeeper
	bankKeeper          bank.Keeper
	stakingKeeper       staking.Keeper
	slashingKeeper      slashing.Keeper
	mintKeeper          mint.Keeper
	distrKeeper         distr.Keeper
	govKeeper           gov.Keeper
	crisisKeeper        crisis.Keeper
	paramsKeeper        params.Keeper
	upgradeKeeper       upgrade.Keeper
	//dataKeeper          data.Keeper
	//schemaKeeper        schema.Keeper
	//espKeeper           esp.Keeper
	geoKeeper geo.Keeper
	//agentKeeper         group.Keeper
	//proposalKeeper      proposal.Keeper
	//consortiumKeeper    consortium.Keeper

	// the module manager
	mm *sdk.ModuleManager

	txDecoder sdk.TxDecoder
	pgIndexer postgresql.Indexer
}

func NewXrnApp(logger log.Logger, db dbm.DB, traceStore io.Writer, loadLatest bool,
	invCheckPeriod uint, baseAppOptions ...func(*bam.BaseApp)) *XrnApp {
	config := sdk.GetConfig()
	config.Seal()

	cdc := MakeCodec()

	txDecoder := auth.DefaultTxDecoder(cdc)
	bApp := bam.NewBaseApp(appName, logger, db, txDecoder, baseAppOptions...)
	bApp.SetCommitMultiStoreTracer(traceStore)
	bApp.SetAppVersion(version.Version)

	var app = &XrnApp{
		BaseApp:          bApp,
		cdc:              cdc,
		invCheckPeriod:   invCheckPeriod,
		keyMain:          sdk.NewKVStoreKey("main"),
		keyAccount:       sdk.NewKVStoreKey("acc"),
		keyStaking:       sdk.NewKVStoreKey(staking.StoreKey),
		tkeyStaking:      sdk.NewTransientStoreKey(staking.TStoreKey),
		keyMint:          sdk.NewKVStoreKey(mint.StoreKey),
		keyDistr:         sdk.NewKVStoreKey(distr.StoreKey),
		tkeyDistr:        sdk.NewTransientStoreKey(distr.TStoreKey),
		keySlashing:      sdk.NewKVStoreKey(slashing.StoreKey),
		keyGov:           sdk.NewKVStoreKey(gov.StoreKey),
		keyFeeCollection: sdk.NewKVStoreKey(auth.FeeStoreKey),
		keyParams:        sdk.NewKVStoreKey(params.StoreKey),
		tkeyParams:       sdk.NewTransientStoreKey(params.TStoreKey),
		upgradeStoreKey:    sdk.NewKVStoreKey(upgrade.StoreKey),
		//dataStoreKey:       sdk.NewKVStoreKey("data"),
		//schemaStoreKey:     sdk.NewKVStoreKey("schema"),
		//espStoreKey:        sdk.NewKVStoreKey("esp"),
		geoStoreKey: sdk.NewKVStoreKey("geo"),
		//agentStoreKey:      sdk.NewKVStoreKey("group"),
		//proposalStoreKey:   sdk.NewKVStoreKey("proposal"),
		txDecoder: txDecoder,
	}

	// init params keeper and subspaces
	app.paramsKeeper = params.NewKeeper(app.cdc, app.keyParams, app.tkeyParams, params.DefaultCodespace)
	authSubspace := app.paramsKeeper.Subspace(auth.DefaultParamspace)
	bankSubspace := app.paramsKeeper.Subspace(bank.DefaultParamspace)
	stakingSubspace := app.paramsKeeper.Subspace(staking.DefaultParamspace)
	mintSubspace := app.paramsKeeper.Subspace(mint.DefaultParamspace)
	distrSubspace := app.paramsKeeper.Subspace(distr.DefaultParamspace)
	slashingSubspace := app.paramsKeeper.Subspace(slashing.DefaultParamspace)
	govSubspace := app.paramsKeeper.Subspace(gov.DefaultParamspace)
	crisisSubspace := app.paramsKeeper.Subspace(crisis.DefaultParamspace)

	// add keepers
	app.accountKeeper = auth.NewAccountKeeper(app.cdc, app.keyAccount, authSubspace, auth.ProtoBaseAccount)
	app.bankKeeper = bank.NewBaseKeeper(app.accountKeeper, bankSubspace, bank.DefaultCodespace)
	app.feeCollectionKeeper = auth.NewFeeCollectionKeeper(app.cdc, app.keyFeeCollection)
	stakingKeeper := staking.NewKeeper(app.cdc, app.keyStaking, app.tkeyStaking, app.bankKeeper,
		stakingSubspace, staking.DefaultCodespace)
	app.mintKeeper = mint.NewKeeper(app.cdc, app.keyMint, mintSubspace, &stakingKeeper, app.feeCollectionKeeper)
	app.distrKeeper = distr.NewKeeper(app.cdc, app.keyDistr, distrSubspace, app.bankKeeper, &stakingKeeper,
		app.feeCollectionKeeper, distr.DefaultCodespace)
	app.slashingKeeper = slashing.NewKeeper(app.cdc, app.keySlashing, &stakingKeeper,
		slashingSubspace, slashing.DefaultCodespace)
	app.crisisKeeper = crisis.NewKeeper(crisisSubspace, invCheckPeriod, app.distrKeeper,
		app.bankKeeper, app.feeCollectionKeeper)
	app.upgradeKeeper = upgrade.NewKeeper(app.upgradeStoreKey, app.cdc)
	app.upgradeKeeper.SetUpgradeHandler("U1", func(ctx sdk.Context, plan upgrade.Plan) {
		// upgrade app state
	})

	// register the proposal types
	govRouter := gov.NewRouter()
	govRouter.AddRoute(gov.RouterKey, gov.ProposalHandler).
		AddRoute(params.RouterKey, params.NewParamChangeProposalHandler(app.paramsKeeper)).
		AddRoute(distr.RouterKey, distr.NewCommunityPoolSpendProposalHandler(app.distrKeeper)).
		AddRoute(upgrade.RouterKey, upgrade.NewSoftwareUpgradeProposalHandler(app.upgradeKeeper))
	app.govKeeper = gov.NewKeeper(app.cdc, app.keyGov, app.paramsKeeper, govSubspace,
		app.bankKeeper, &stakingKeeper, gov.DefaultCodespace, govRouter)

	// register the staking hooks
	// NOTE: stakingKeeper above is passed by reference, so that it will contain these hooks
	app.stakingKeeper = *stakingKeeper.SetHooks(
		staking.NewMultiStakingHooks(app.distrKeeper.Hooks(), app.slashingKeeper.Hooks()))

	//app.schemaKeeper = schema.NewKeeper(app.schemaStoreKey, cdc)
	//
	//app.dataKeeper = data.NewKeeper(app.dataStoreKey, app.schemaKeeper, cdc)
	//
	//app.agentKeeper = group.NewKeeper(app.agentStoreKey, cdc, app.accountKeeper)

	app.geoKeeper = geo.NewKeeper(app.geoStoreKey, cdc, app.pgIndexer)

	//app.upgradeKeeper = upgrade.NewKeeper(app.upgradeStoreKey, cdc)
	//app.upgradeKeeper.SetDoShutdowner(app.shutdownOnUpgrade)

	//proposalRouter := proposal.NewRouter().
	//	AddRoute("esp", app.espKeeper).
	//	AddRoute("consortium", app.consortiumKeeper)
	//
	//app.proposalKeeper = proposal.NewKeeper(app.proposalStoreKey, proposalRouter, cdc)

	app.mm = sdk.NewModuleManager(
		genaccounts.NewAppModule(app.accountKeeper),
		genutil.NewAppModule(app.accountKeeper, app.stakingKeeper, app.BaseApp.DeliverTx),
		auth.NewAppModule(app.accountKeeper, app.feeCollectionKeeper),
		bank.NewAppModule(app.bankKeeper, app.accountKeeper),
		crisis.NewAppModule(app.crisisKeeper, app.Logger()),
		distr.NewAppModule(app.distrKeeper),
		gov.NewAppModule(app.govKeeper),
		mint.NewAppModule(app.mintKeeper),
		slashing.NewAppModule(app.slashingKeeper, app.stakingKeeper),
		staking.NewAppModule(app.stakingKeeper, app.feeCollectionKeeper, app.distrKeeper, app.accountKeeper),
		geo.NewAppModule(app.geoKeeper),
		upgrade.NewAppModule(app.upgradeKeeper),
	)

	// During begin block slashing happens after distr.BeginBlocker so that
	// there is nothing left over in the validator fee pool, so as to keep the
	// CanWithdrawInvariant invariant.
	app.mm.SetOrderBeginBlockers(upgrade.ModuleName, mint.ModuleName, distr.ModuleName, slashing.ModuleName)

	app.mm.SetOrderEndBlockers(gov.ModuleName, staking.ModuleName)

	// genutils must occur after staking so that pools are properly
	// initialized with tokens from genesis accounts.
	app.mm.SetOrderInitGenesis(genaccounts.ModuleName, distr.ModuleName,
		staking.ModuleName, auth.ModuleName, bank.ModuleName, slashing.ModuleName,
		gov.ModuleName, mint.ModuleName, crisis.ModuleName, genutil.ModuleName, geo.ModuleName)

	app.mm.RegisterInvariants(&app.crisisKeeper)
	app.mm.RegisterRoutes(app.Router(), app.QueryRouter())

	app.MountStores(app.keyMain, app.keyAccount, app.keyStaking, app.keyMint,
		app.keyDistr, app.keySlashing, app.keyGov, app.keyFeeCollection,
		app.keyParams, app.tkeyParams, app.tkeyStaking, app.tkeyDistr,
		app.upgradeStoreKey,
		app.geoStoreKey,
		//app.schemaStoreKey, app.dataStoreKey,
		//app.espStoreKey, app.geoStoreKey, app.agentStoreKey,
		//app.proposalStoreKey,
	)

	// initialize BaseApp
	app.SetInitChainer(app.InitChainer)
	app.SetBeginBlocker(app.BeginBlocker)
	app.SetAnteHandler(auth.NewAnteHandler(app.accountKeeper, app.feeCollectionKeeper, auth.DefaultSigVerificationGasConsumer))
	app.SetEndBlocker(app.EndBlocker)

	if loadLatest {
		err := app.LoadLatestVersion(app.keyMain)
		if err != nil {
			cmn.Exit(err.Error())
		}
	}
	return app
}

// application updates every begin block
func (app *XrnApp) BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock) abci.ResponseBeginBlock {
	return app.mm.BeginBlock(ctx, req)
}

// application updates every end block
func (app *XrnApp) EndBlocker(ctx sdk.Context, req abci.RequestEndBlock) abci.ResponseEndBlock {
	return app.mm.EndBlock(ctx, req)
}

// application update at chain initialization
func (app *XrnApp) InitChainer(ctx sdk.Context, req abci.RequestInitChain) abci.ResponseInitChain {
	var genesisState GenesisState
	app.cdc.MustUnmarshalJSON(req.AppStateBytes, &genesisState)
	return app.mm.InitGenesis(ctx, genesisState)
}

// load a particular height
func (app *XrnApp) LoadHeight(height int64) error {
	return app.LoadVersion(height, app.keyMain)
}

func (app *XrnApp) ConfigurePostgreSQLIndexer(postgresUrl string) {
	pgIndexer, err := postgresql.NewIndexer(postgresUrl, app.txDecoder)
	if err == nil {
		pgIndexer.AddMigrations("geo", geo.PostgresMigrations)
		app.pgIndexer = pgIndexer
		app.Logger().Info("Started PostgreSQL Indexer")
	} else {
		app.Logger().Error("Error Starting PostgreSQL Indexer", err)
	}
}

func (app *XrnApp) shutdownOnUpgrade(ctx sdk.Context, plan upgrade.Plan) {
	if len(plan.Info) != 0 {
		home := viper.GetString(cli.HomeFlag)
		_ = ioutil.WriteFile(filepath.Join(home, "data", "upgrade-info"), []byte(plan.Info), 0644)
	}
	ctx.Logger().Error(fmt.Sprintf("UPGRADE \"%s\" NEEDED needed at height %d: %s", plan.Name, ctx.BlockHeight(), plan.Info))
	os.Exit(1)
}

func (app *XrnApp) InitChain(req abci.RequestInitChain) (res abci.ResponseInitChain) {
	res = app.BaseApp.InitChain(req)
	if app.pgIndexer != nil {
		app.pgIndexer.OnInitChain(req, res)
	}
	return res
}

func (app *XrnApp) BeginBlock(req abci.RequestBeginBlock) (res abci.ResponseBeginBlock) {
	res = app.BaseApp.BeginBlock(req)
	if app.pgIndexer != nil {
		app.pgIndexer.OnBeginBlock(req, res)
	}
	return res
}

func (app *XrnApp) DeliverTx(txBytes []byte) (res abci.ResponseDeliverTx) {
	if app.pgIndexer != nil {
		app.pgIndexer.BeforeDeliverTx(txBytes)
	}
	res = app.BaseApp.DeliverTx(txBytes)
	if app.pgIndexer != nil {
		app.pgIndexer.AfterDeliverTx(txBytes, res)
	}
	return res
}

func (app *XrnApp) EndBlock(req abci.RequestEndBlock) (res abci.ResponseEndBlock) {
	res = app.BaseApp.EndBlock(req)
	if app.pgIndexer != nil {
		app.pgIndexer.OnEndBlock(req, res)
	}
	return res
}

func (app *XrnApp) Commit() (res abci.ResponseCommit) {
	res = app.BaseApp.Commit()
	if app.pgIndexer != nil {
		app.pgIndexer.OnCommit(res)
	}
	return res
}
