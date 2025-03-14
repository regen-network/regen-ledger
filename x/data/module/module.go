package module

import (
	"context"
	"encoding/json"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/cobra"

	tmtypes "github.com/cometbft/cometbft/abci/types"

	sdkclient "github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/orm/model/ormdb"
	"github.com/cosmos/cosmos-sdk/orm/types/ormjson"
	storeTypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"

	"github.com/regen-network/regen-ledger/x/data/v3"
	"github.com/regen-network/regen-ledger/x/data/v3/client"
	"github.com/regen-network/regen-ledger/x/data/v3/genesis"
	"github.com/regen-network/regen-ledger/x/data/v3/server"
	"github.com/regen-network/regen-ledger/x/data/v3/simulation"
)

var (
	_ module.AppModule           = &Module{}
	_ module.AppModuleBasic      = Module{}
	_ module.AppModuleSimulation = Module{}
)

type Module struct {
	ak     data.AccountKeeper
	bk     data.BankKeeper
	sk     storeTypes.StoreKey
	keeper server.Keeper
}

func (a Module) InitGenesis(s sdk.Context, jsonCodec codec.JSONCodec, message json.RawMessage) []tmtypes.ValidatorUpdate {
	update, err := a.keeper.InitGenesis(s, jsonCodec, message)
	if err != nil {
		panic(err)
	}
	return update
}

func (a Module) ExportGenesis(s sdk.Context, jsonCodec codec.JSONCodec) json.RawMessage {
	jsn, err := a.keeper.ExportGenesis(s, jsonCodec)
	if err != nil {
		panic(err)
	}
	return jsn
}

func (a Module) RegisterInvariants(_ sdk.InvariantRegistry) {}

func (a *Module) RegisterServices(cfg module.Configurator) {
	impl := server.NewServer(a.sk, a.ak, a.bk)
	data.RegisterMsgServer(cfg.MsgServer(), impl)
	data.RegisterQueryServer(cfg.QueryServer(), impl)
	a.keeper = impl
}

var _ module.AppModuleBasic = Module{}
var _ module.AppModuleSimulation = &Module{}

func NewModule(sk storeTypes.StoreKey, ak data.AccountKeeper, bk data.BankKeeper) *Module {
	return &Module{
		ak: ak,
		bk: bk,
		sk: sk,
	}
}

func (a Module) Name() string {
	return data.ModuleName
}

func (a Module) RegisterInterfaces(registry types.InterfaceRegistry) {
	data.RegisterTypes(registry)
}

func (a Module) RegisterGRPCGatewayRoutes(clientCtx sdkclient.Context, mux *runtime.ServeMux) {
	err := data.RegisterQueryHandlerClient(context.Background(), mux, data.NewQueryClient(clientCtx))
	if err != nil {
		panic(err)
	}
}

func (a Module) DefaultGenesis(codec.JSONCodec) json.RawMessage {
	db, err := ormdb.NewModuleDB(&data.ModuleSchema, ormdb.ModuleDBOptions{})
	if err != nil {
		panic(err)
	}

	jsonTarget := ormjson.NewRawMessageTarget()
	err = db.DefaultJSON(jsonTarget)
	if err != nil {
		panic(err)
	}

	bz, err := jsonTarget.JSON()
	if err != nil {
		panic(err)
	}

	return bz
}

func (a Module) ValidateGenesis(_ codec.JSONCodec, _ sdkclient.TxEncodingConfig, bz json.RawMessage) error {
	db, err := ormdb.NewModuleDB(&data.ModuleSchema, ormdb.ModuleDBOptions{})
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

func (a Module) GetQueryCmd() *cobra.Command {
	return client.QueryCmd(a.Name())
}

func (a Module) GetTxCmd() *cobra.Command {
	return client.TxCmd(a.Name())
}

// ConsensusVersion implements AppModule/ConsensusVersion.
func (Module) ConsensusVersion() uint64 { return 1 }

// RegisterLegacyAminoCodec registers the data module's types on the given LegacyAmino codec.
func (a Module) RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	data.RegisterLegacyAminoCodec(cdc)
}

// AppModuleSimulation functions

// GenerateGenesisState creates a randomized GenesisState of the data module.
func (Module) GenerateGenesisState(simState *module.SimulationState) {
	simulation.RandomizedGenState(simState)
}

// RegisterStoreDecoder registers a decoder for data module's types
func (Module) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {
}

// WeightedOperations returns all the data module operations with their respective weights.
func (a Module) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	querier := a.keeper.QueryServer()

	return simulation.WeightedOperations(
		simState.AppParams, simState.Cdc,
		a.ak, a.bk,
		querier,
	)
}
