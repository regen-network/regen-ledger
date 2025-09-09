package module

import (
	"context"
	"encoding/json"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/cobra"

	abci "github.com/cometbft/cometbft/abci/types"

	govkeeper "github.com/cosmos/cosmos-sdk/x/gov/keeper"

	storetypes "cosmossdk.io/store/types"
	sdkclient "github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/regen-network/regen-ledger/orm/model/ormdb"
	"github.com/regen-network/regen-ledger/orm/types/ormjson"

	"github.com/regen-network/regen-ledger/x/ecocredit/v4"
	basesims "github.com/regen-network/regen-ledger/x/ecocredit/v4/base/simulation"
	basetypes "github.com/regen-network/regen-ledger/x/ecocredit/v4/base/types/v1"
	basetypesv1alpha1 "github.com/regen-network/regen-ledger/x/ecocredit/v4/base/types/v1alpha1"
	basketsims "github.com/regen-network/regen-ledger/x/ecocredit/v4/basket/simulation"
	baskettypes "github.com/regen-network/regen-ledger/x/ecocredit/v4/basket/types/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/v4/client"
	"github.com/regen-network/regen-ledger/x/ecocredit/v4/genesis"
	marketsims "github.com/regen-network/regen-ledger/x/ecocredit/v4/marketplace/simulation"
	markettypes "github.com/regen-network/regen-ledger/x/ecocredit/v4/marketplace/types/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/v4/server"
	"github.com/regen-network/regen-ledger/x/ecocredit/v4/simulation"
)

var (
	_ module.AppModule           = &Module{}
	_ module.AppModuleBasic      = Module{}
	_ module.AppModuleSimulation = Module{}
)

const (
	ConsensusVersion = 4 // ConsensusVersion is the module consensus version
)

// Module implements the AppModule interface.
type Module struct {
	key           storetypes.StoreKey
	authority     sdk.AccAddress
	Keeper        server.Keeper
	accountKeeper ecocredit.AccountKeeper
	bankKeeper    ecocredit.BankKeeper
	govKeeper     *govkeeper.Keeper

	// legacySubspace is used solely for migration of x/ecocredit managed parameters
	legacySubspace paramtypes.Subspace
}

// NewModule returns a new Module.
func NewModule(
	storeKey storetypes.StoreKey,
	authority sdk.AccAddress,
	accountKeeper ecocredit.AccountKeeper,
	bankKeeper ecocredit.BankKeeper,
	legacySubspace paramtypes.Subspace,
	govKeeper *govkeeper.Keeper,
) *Module {
	return &Module{
		key:            storeKey,
		legacySubspace: legacySubspace,
		bankKeeper:     bankKeeper,
		accountKeeper:  accountKeeper,
		authority:      authority,
		govKeeper:      govKeeper,
		Keeper:         server.NewServer(storeKey, accountKeeper, bankKeeper, authority),
	}
}

// IsOnePerModuleType implements the depinject.OnePerModuleType interface.
func (am Module) IsOnePerModuleType() {}

// IsAppModule implements the appmodule.AppModule interface.
func (am Module) IsAppModule() {}

/* -------------------- AppModule -------------------- */

// ConsensusVersion implements AppModule/ConsensusVersion.
func (Module) ConsensusVersion() uint64 { return ConsensusVersion }

// Name implements AppModule/Name.
func (am Module) Name() string {
	return ecocredit.ModuleName
}

// RegisterInvariants implements AppModule/RegisterInvariants.
func (am Module) RegisterInvariants(reg sdk.InvariantRegistry) {
	am.Keeper.RegisterInvariants(reg)
}

// RegisterInterfaces implements AppModule/RegisterInterfaces.
func (am Module) RegisterInterfaces(registry types.InterfaceRegistry) {
	baskettypes.RegisterTypes(registry)
	basetypes.RegisterTypes(registry)
	markettypes.RegisterTypes(registry)

	// legacy types to support querying historical events
	basetypesv1alpha1.RegisterTypes(registry)
}

// RegisterServices implements AppModule/RegisterServices.
func (am *Module) RegisterServices(cfg module.Configurator) {
	baseK := am.Keeper.GetBaseKeeper()
	basetypes.RegisterMsgServer(cfg.MsgServer(), baseK)
	basetypes.RegisterQueryServer(cfg.QueryServer(), baseK)

	basketK := am.Keeper.GetBasketKeeper()
	baskettypes.RegisterMsgServer(cfg.MsgServer(), basketK)
	baskettypes.RegisterQueryServer(cfg.QueryServer(), basketK)

	marketK := am.Keeper.GetMarketKeeper()
	markettypes.RegisterMsgServer(cfg.MsgServer(), marketK)
	markettypes.RegisterQueryServer(cfg.QueryServer(), marketK)

	migrator := server.NewMigrator(am.Keeper, am.legacySubspace)
	if err := cfg.RegisterMigration(ecocredit.ModuleName, 2, migrator.Migrate2to3); err != nil {
		panic(err)
	}

	if err := cfg.RegisterMigration(ecocredit.ModuleName, 3, migrator.Migrate3to4); err != nil {
		panic(err)
	}
}

// RegisterGRPCGatewayRoutes implements AppModule/RegisterGRPCGatewayRoutes.
func (am Module) RegisterGRPCGatewayRoutes(clientCtx sdkclient.Context, mux *runtime.ServeMux) {
	ctx := context.Background()
	err := basetypes.RegisterQueryHandlerClient(ctx, mux, basetypes.NewQueryClient(clientCtx))
	if err != nil {
		panic(err)
	}
	err = baskettypes.RegisterQueryHandlerClient(ctx, mux, baskettypes.NewQueryClient(clientCtx))
	if err != nil {
		panic(err)
	}
	err = markettypes.RegisterQueryHandlerClient(ctx, mux, markettypes.NewQueryClient(clientCtx))
	if err != nil {
		panic(err)
	}
}

// RegisterLegacyAminoCodec implements AppModule/RegisterLegacyAminoCodec.
func (am Module) RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	basetypes.RegisterLegacyAminoCodec(cdc)
	baskettypes.RegisterLegacyAminoCodec(cdc)
	markettypes.RegisterLegacyAminoCodec(cdc)
}

// InitGenesis implements AppModule/InitGenesis.
func (am Module) InitGenesis(s sdk.Context, jsonCodec codec.JSONCodec, message json.RawMessage) []abci.ValidatorUpdate {
	update, err := am.Keeper.InitGenesis(s, jsonCodec, message)
	if err != nil {
		panic(err)
	}
	return update
}

// ExportGenesis implements AppModule/ExportGenesis.
func (am Module) ExportGenesis(s sdk.Context, jsonCodec codec.JSONCodec) json.RawMessage {
	genesis, err := am.Keeper.ExportGenesis(s, jsonCodec)
	if err != nil {
		panic(err)
	}
	return genesis
}

// DefaultGenesis implements AppModule/DefaultGenesis.
func (am Module) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
	db, err := ormdb.NewModuleDB(&ecocredit.ModuleSchema, ormdb.ModuleDBOptions{})
	if err != nil {
		panic(err)
	}

	jsonTarget := ormjson.NewRawMessageTarget()
	err = db.DefaultJSON(jsonTarget)
	if err != nil {
		panic(err)
	}

	creditTypes := genesis.DefaultCreditTypes()
	err = genesis.MergeCreditTypesIntoTarget(creditTypes, jsonTarget)
	if err != nil {
		panic(err)
	}

	classFee := genesis.DefaultClassFee()
	err = genesis.MergeClassFeeIntoTarget(cdc, classFee, jsonTarget)
	if err != nil {
		panic(err)
	}

	basketFee := genesis.DefaultBasketFee()
	err = genesis.MergeBasketFeeIntoTarget(cdc, basketFee, jsonTarget)
	if err != nil {
		panic(err)
	}

	allowedDenoms := genesis.DefaultAllowedDenoms()
	err = genesis.MergeAllowedDenomsIntoTarget(allowedDenoms, jsonTarget)
	if err != nil {
		panic(err)
	}

	bz, err := jsonTarget.JSON()
	if err != nil {
		panic(err)
	}

	return bz
}

// ValidateGenesis implements AppModule/ValidateGenesis.
func (am Module) ValidateGenesis(_ codec.JSONCodec, _ sdkclient.TxEncodingConfig, bz json.RawMessage) error {
	db, err := ormdb.NewModuleDB(&ecocredit.ModuleSchema, ormdb.ModuleDBOptions{})
	if err != nil {
		return err
	}

	jsonSource, err := ormjson.NewRawMessageSource(bz)
	if err != nil {
		return err
	}

	err = db.ValidateJSON(jsonSource)
	if err != nil {
		return err
	}

	return genesis.ValidateGenesis(bz)
}

// GetTxCmd implements AppModule/GetTxCmd.
func (am Module) GetTxCmd() *cobra.Command {
	return client.TxCmd(am.Name())
}

// GetQueryCmd implements AppModule/GetQueryCmd.
func (am Module) GetQueryCmd() *cobra.Command {
	return client.QueryCmd(am.Name())
}

// BeginBlock checks if there are any expired sell or buy orders and removes them from state.
func (am Module) BeginBlock(ctx sdk.Context) {
	err := BeginBlocker(ctx, am.Keeper)
	if err != nil {
		panic(err)
	}
}

/* -------------------- AppModuleSimulation -------------------- */

// GenerateGenesisState creates a randomized GenesisState of the ecocredit module.
func (Module) GenerateGenesisState(simState *module.SimulationState) {
	simulation.RandomizedGenState(simState)
}

// RegisterStoreDecoder implements AppModuleSimulation/RegisterStoreDecoder.
func (Module) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations implements AppModuleSimulation/WeightedOperations.
func (am Module) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	baseServer, basketServer, marketServer := am.Keeper.QueryServers()
	baseOps := basesims.WeightedOperations(
		simState.AppParams,
		am.accountKeeper,
		am.bankKeeper,
		*am.govKeeper,
		baseServer,
		basketServer,
		marketServer,
		am.authority,
	)

	basketOps := basketsims.WeightedOperations(
		simState.AppParams,
		am.accountKeeper,
		am.bankKeeper,
		*am.govKeeper,
		baseServer,
		basketServer,
		am.authority,
	)

	marketplaceOps := marketsims.WeightedOperations(simState.AppParams,
		am.accountKeeper,
		am.bankKeeper,
		baseServer,
		marketServer,
		*am.govKeeper,
		am.authority,
	)

	return append(append(baseOps, basketOps...), marketplaceOps...)
}
