package app

import (
	"encoding/json"
	"fmt"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/cosmos/cosmos-sdk/x/upgrade"
	"github.com/regen-network/regen-ledger/index/postgresql"
	"github.com/regen-network/regen-ledger/x/consortium"
	"github.com/regen-network/regen-ledger/x/data"
	"github.com/regen-network/regen-ledger/x/esp"
	"github.com/regen-network/regen-ledger/x/geo"
	"github.com/regen-network/regen-ledger/x/group"
	"github.com/regen-network/regen-ledger/x/proposal"
	"github.com/regen-network/regen-ledger/x/schema"
	"github.com/spf13/viper"
	"github.com/tendermint/tendermint/libs/cli"
	"github.com/tendermint/tendermint/libs/log"
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
	tmtypes "github.com/tendermint/tendermint/types"
)

const (
	appName = "xrn"
)

type xrnApp struct {
	*bam.BaseApp
	cdc *codec.Codec

	keyMain          *sdk.KVStoreKey
	keyAccount       *sdk.KVStoreKey
	keyFeeCollection *sdk.KVStoreKey
	//schemaStoreKey  *sdk.KVStoreKey
	dataStoreKey       *sdk.KVStoreKey
	schemaStoreKey     *sdk.KVStoreKey
	espStoreKey        *sdk.KVStoreKey
	geoStoreKey        *sdk.KVStoreKey
	agentStoreKey      *sdk.KVStoreKey
	proposalStoreKey   *sdk.KVStoreKey
	upgradeStoreKey    *sdk.KVStoreKey
	consortiumStoreKey *sdk.KVStoreKey
	keyParams          *sdk.KVStoreKey
	tkeyParams         *sdk.TransientStoreKey

	accountKeeper       auth.AccountKeeper
	bankKeeper          bank.Keeper
	feeCollectionKeeper auth.FeeCollectionKeeper
	dataKeeper          data.Keeper
	schemaKeeper        schema.Keeper
	espKeeper           esp.Keeper
	geoKeeper           geo.Keeper
	agentKeeper         group.Keeper
	proposalKeeper      proposal.Keeper
	upgradeKeeper       upgrade.Keeper
	consortiumKeeper    consortium.Keeper
	paramsKeeper        params.Keeper

	pgIndexer postgresql.Indexer
}

func NewXrnApp(logger log.Logger, db dbm.DB, postgresUrl string) *xrnApp {
	config := sdk.GetConfig()
	config.Seal()

	// First define the top level codec that will be shared by the different modules
	cdc := MakeCodec()

	// BaseApp handles interactions with Tendermint through the ABCI protocol
	txDecoder := auth.DefaultTxDecoder(cdc)
	bApp := bam.NewBaseApp(appName, logger, db, txDecoder)

	// Enable this for low-level debugging
	// bApp.SetCommitMultiStoreTracer(os.Stdout)

	// Here you initialize your application with the store keys it requires
	var app = &xrnApp{
		BaseApp: bApp,
		cdc:     cdc,

		keyMain:          sdk.NewKVStoreKey("main"),
		keyAccount:       sdk.NewKVStoreKey("acc"),
		keyFeeCollection: sdk.NewKVStoreKey("fee_collection"),
		//schemaStoreKey: sdk.NewKVStoreKey("schema"),
		dataStoreKey:       sdk.NewKVStoreKey("data"),
		schemaStoreKey:     sdk.NewKVStoreKey("schema"),
		espStoreKey:        sdk.NewKVStoreKey("esp"),
		geoStoreKey:        sdk.NewKVStoreKey("geo"),
		agentStoreKey:      sdk.NewKVStoreKey("group"),
		proposalStoreKey:   sdk.NewKVStoreKey("proposal"),
		upgradeStoreKey:    sdk.NewKVStoreKey("upgrade"),
		consortiumStoreKey: sdk.NewKVStoreKey("consortium"),
		keyParams:          sdk.NewKVStoreKey(params.StoreKey),
		tkeyParams:         sdk.NewTransientStoreKey(params.TStoreKey),
	}

	if len(postgresUrl) != 0 {
		pgIndexer, err := postgresql.NewIndexer(postgresUrl, txDecoder)
		if err == nil {
			pgIndexer.AddMigrations("geo", geo.PostgresMigrations)
			app.pgIndexer = pgIndexer
			logger.Info("Started PostgreSQL Indexer")
		} else {
			logger.Error("Error Starting PostgreSQL Indexer", err)
		}
	}

	app.paramsKeeper = params.NewKeeper(app.cdc, app.keyParams, app.tkeyParams)

	// The AccountKeeper handles address -> account lookups
	app.accountKeeper = auth.NewAccountKeeper(
		app.cdc,
		app.keyAccount,
		app.paramsKeeper.Subspace(auth.DefaultParamspace),
		auth.ProtoBaseAccount,
	)

	// The BankKeeper allows you perform sdk.Coins interactions
	app.bankKeeper = bank.NewBaseKeeper(app.accountKeeper,
		app.paramsKeeper.Subspace(bank.DefaultParamspace),
		bank.DefaultCodespace,
	)

	// The FeeCollectionKeeper collects transaction fees and renders them to the fee distribution module
	app.feeCollectionKeeper = auth.NewFeeCollectionKeeper(cdc, app.keyFeeCollection)

	app.schemaKeeper = schema.NewKeeper(app.schemaStoreKey, cdc)

	app.dataKeeper = data.NewKeeper(app.dataStoreKey, app.schemaKeeper, cdc)

	app.agentKeeper = group.NewKeeper(app.agentStoreKey, cdc)

	app.geoKeeper = geo.NewKeeper(app.geoStoreKey, cdc, app.pgIndexer)

	app.espKeeper = esp.NewKeeper(app.espStoreKey, app.agentKeeper, app.geoKeeper, cdc)

	app.upgradeKeeper = upgrade.NewKeeper(app.upgradeStoreKey, cdc)
	app.upgradeKeeper.SetDoShutdowner(app.shutdownOnUpgrade)
	app.upgradeKeeper.SetUpgradeHandler("test2", func(ctx sdk.Context, plan upgrade.Plan) {
		if app.pgIndexer != nil {
			app.pgIndexer.ApplyMigrations()
		}
	})

	app.consortiumKeeper = consortium.NewKeeper(app.consortiumStoreKey, cdc, app.agentKeeper, app.upgradeKeeper)

	proposalRouter := proposal.NewRouter().
		AddRoute("esp", app.espKeeper).
		AddRoute("consortium", app.consortiumKeeper)

	app.proposalKeeper = proposal.NewKeeper(app.proposalStoreKey, proposalRouter, cdc)

	// The AnteHandler handles signature verification and transaction pre-processing
	app.SetAnteHandler(auth.NewAnteHandler(app.accountKeeper, app.feeCollectionKeeper))

	// The app.Router is the main transaction router where each module registers its routes
	// Register the bank and data routes here
	app.Router().
		AddRoute("bank", bank.NewHandler(app.bankKeeper)).
		AddRoute("data", data.NewHandler(app.dataKeeper)).
		AddRoute("geo", geo.NewHandler(app.geoKeeper)).
		AddRoute("group", group.NewHandler(app.agentKeeper)).
		AddRoute("proposal", proposal.NewHandler(app.proposalKeeper)).
		AddRoute("schema", schema.NewHandler(app.schemaKeeper))

	// The app.QueryRouter is the main query router where each module registers its routes
	app.QueryRouter().
		AddRoute("acc", auth.NewQuerier(app.accountKeeper)).
		AddRoute("data", data.NewQuerier(app.dataKeeper)).
		AddRoute("group", group.NewQuerier(app.agentKeeper)).
		AddRoute("proposal", proposal.NewQuerier(app.proposalKeeper))

	// The initChainer handles translating the genesis.json file into initial state for the network
	app.SetInitChainer(app.initChainer)
	app.SetBeginBlocker(app.BeginBlocker)
	app.SetEndBlocker(app.EndBlocker)

	app.MountStores(
		app.keyMain,
		app.keyAccount,
		app.keyFeeCollection,
		app.schemaStoreKey,
		app.dataStoreKey,
		app.espStoreKey,
		app.geoStoreKey,
		app.agentStoreKey,
		app.proposalStoreKey,
		app.upgradeStoreKey,
		app.consortiumStoreKey,
		app.keyParams,
		app.tkeyParams,
	)

	err := app.LoadLatestVersion(app.keyMain)
	if err != nil {
		cmn.Exit(err.Error())
	}

	return app
}

// GenesisState represents chain state at the start of the chain. Any initial state (account balances) are stored here.
type GenesisState struct {
	Accounts []*auth.BaseAccount `json:"accounts"`
	Groups   []group.Group       `json:"groups"`
	AuthData auth.GenesisState   `json:"auth"`
	BankData bank.GenesisState   `json:"bank"`
}

func (app *xrnApp) initChainer(ctx sdk.Context, req abci.RequestInitChain) abci.ResponseInitChain {
	stateJSON := req.AppStateBytes

	genesisState := new(GenesisState)
	err := app.cdc.UnmarshalJSON(stateJSON, genesisState)
	if err != nil {
		panic(err)
	}

	for _, acc := range genesisState.Accounts {
		acc.AccountNumber = app.accountKeeper.GetNextAccountNumber(ctx)
		app.accountKeeper.SetAccount(ctx, acc)
	}

	for _, g := range genesisState.Groups {
		app.agentKeeper.CreateGroup(ctx, g)
	}

	app.consortiumKeeper.SetValidators(ctx, req.Validators)

	auth.InitGenesis(ctx, app.accountKeeper, app.feeCollectionKeeper, genesisState.AuthData)
	bank.InitGenesis(ctx, app.bankKeeper, genesisState.BankData)

	return abci.ResponseInitChain{}
}

func (app *xrnApp) shutdownOnUpgrade(ctx sdk.Context, plan upgrade.Plan) {
	if len(plan.Info) != 0 {
		home := viper.GetString(cli.HomeFlag)
		_ = ioutil.WriteFile(filepath.Join(home, "data", "upgrade-info"), []byte(plan.Info), 0644)
	}
	ctx.Logger().Error(fmt.Sprintf("UPGRADE \"%s\" NEEDED needed at height %d: %s", plan.Name, ctx.BlockHeight(), plan.Info))
	os.Exit(1)
}

func (app *xrnApp) BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock) abci.ResponseBeginBlock {
	app.upgradeKeeper.BeginBlocker(ctx, req)
	return abci.ResponseBeginBlock{}
}

func (app *xrnApp) EndBlocker(ctx sdk.Context, req abci.RequestEndBlock) abci.ResponseEndBlock {
	//validatorUpdates := app.consortiumKeeper.EndBlocker(ctx)
	//return abci.ResponseEndBlock{ValidatorUpdates: validatorUpdates}
	return abci.ResponseEndBlock{}
}

func (app *xrnApp) InitChain(req abci.RequestInitChain) (res abci.ResponseInitChain) {
	res = app.BaseApp.InitChain(req)
	if app.pgIndexer != nil {
		app.pgIndexer.OnInitChain(req, res)
	}
	return res
}

func (app *xrnApp) BeginBlock(req abci.RequestBeginBlock) (res abci.ResponseBeginBlock) {
	res = app.BaseApp.BeginBlock(req)
	if app.pgIndexer != nil {
		app.pgIndexer.OnBeginBlock(req, res)
	}
	return res
}

func (app *xrnApp) DeliverTx(txBytes []byte) (res abci.ResponseDeliverTx) {
	app.pgIndexer.BeforeDeliverTx(txBytes)
	res = app.BaseApp.DeliverTx(txBytes)
	if app.pgIndexer != nil {
		app.pgIndexer.AfterDeliverTx(txBytes, res)
	}
	return res
}

func (app *xrnApp) EndBlock(req abci.RequestEndBlock) (res abci.ResponseEndBlock) {
	res = app.BaseApp.EndBlock(req)
	if app.pgIndexer != nil {
		app.pgIndexer.OnEndBlock(req, res)
	}
	return res
}

func (app *xrnApp) Commit() (res abci.ResponseCommit) {
	res = app.BaseApp.Commit()
	if app.pgIndexer != nil {
		app.pgIndexer.OnCommit(res)
	}
	return res
}

// ExportAppStateAndValidators does the things
func (app *xrnApp) ExportAppStateAndValidators() (appState json.RawMessage, validators []tmtypes.GenesisValidator, err error) {
	ctx := app.NewContext(true, abci.Header{})
	accounts := []*auth.BaseAccount{}

	appendAccountsFn := func(acc auth.Account) bool {
		account := &auth.BaseAccount{
			Address: acc.GetAddress(),
			Coins:   acc.GetCoins(),
		}

		accounts = append(accounts, account)
		return false
	}

	app.accountKeeper.IterateAccounts(ctx, appendAccountsFn)

	genState := GenesisState{Accounts: accounts}
	appState, err = codec.MarshalJSONIndent(app.cdc, genState)
	if err != nil {
		return nil, nil, err
	}

	return appState, validators, err
}

// MakeCodec generates the necessary codecs for Amino
func MakeCodec() *codec.Codec {
	var cdc = codec.New()
	auth.RegisterCodec(cdc)
	bank.RegisterCodec(cdc)
	data.RegisterCodec(cdc)
	esp.RegisterCodec(cdc)
	geo.RegisterCodec(cdc)
	group.RegisterCodec(cdc)
	proposal.RegisterCodec(cdc)
	consortium.RegisterCodec(cdc)
	upgrade.RegisterCodec(cdc)
	sdk.RegisterCodec(cdc)
	codec.RegisterCrypto(cdc)
	return cdc
}
