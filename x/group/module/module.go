package module

import (
	"encoding/json"
	"fmt"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/regen-network/regen-ledger/x/group"
	"github.com/regen-network/regen-ledger/x/group/server"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
	abci "github.com/tendermint/tendermint/abci/types"
)

type AppModuleBasic struct{}

var _ module.AppModule = AppModule{}
var _ module.AppModuleBasic = AppModuleBasic{}

func (a AppModuleBasic) Name() string {
	return group.ModuleName
}

func (AppModuleBasic) RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {}

func (a AppModuleBasic) DefaultGenesis(marshaler codec.JSONMarshaler) json.RawMessage {
	return marshaler.MustMarshalJSON(group.NewGenesisState())
}

func (a AppModuleBasic) ValidateGenesis(cdc codec.JSONMarshaler, config client.TxEncodingConfig, bz json.RawMessage) error {
	var data group.GenesisState
	if err := cdc.UnmarshalJSON(bz, &data); err != nil {
		return fmt.Errorf("failed to unmarshal %s genesis state: %w", group.ModuleName, err)
	}
	return data.Validate()
}

func (a AppModuleBasic) RegisterRESTRoutes(clientCtx client.Context, r *mux.Router) {
}

func (a AppModuleBasic) RegisterGRPCGatewayRoutes(client.Context, *runtime.ServeMux) {}

func (a AppModuleBasic) GetTxCmd() *cobra.Command {
	return nil
}

func (a AppModuleBasic) GetQueryCmd() *cobra.Command {
	//return cli.GetQueryCmd(StoreKey, cdc)
	return nil
}

// RegisterInterfaces registers module concrete types into protobuf Any.
func (AppModuleBasic) RegisterInterfaces(registry types.InterfaceRegistry) {
	group.RegisterTypes(registry)
}

type AppModule struct {
	AppModuleBasic
	storeKey   sdk.StoreKey
	paramSpace paramstypes.Subspace
	router     sdk.Router
}

func NewAppModule(storeKey sdk.StoreKey, paramSpace paramstypes.Subspace, router sdk.Router) AppModule {
	// set KeyTable if it has not already been set
	if !paramSpace.HasKeyTable() {
		paramSpace = paramSpace.WithKeyTable(paramstypes.NewKeyTable().RegisterParamSet(&group.Params{}))
	}
	return AppModule{
		AppModuleBasic: AppModuleBasic{},
		storeKey:       storeKey,
		paramSpace:     paramSpace,
		router:         router,
	}
}

func (a AppModule) InitGenesis(ctx sdk.Context, cdc codec.JSONMarshaler, data json.RawMessage) []abci.ValidatorUpdate {
	var genesisState group.GenesisState
	cdc.MustUnmarshalJSON(data, &genesisState)
	if err := genesisState.Validate(); err != nil {
		panic(fmt.Sprintf("failed to validate %s genesis state: %s", group.ModuleName, err))
	}
	a.paramSpace.SetParamSet(ctx, &genesisState.Params) // TODO: revisit if this makes sense
	return []abci.ValidatorUpdate{}
}

func (a AppModule) ExportGenesis(ctx sdk.Context, cdc codec.JSONMarshaler) json.RawMessage {
	var p group.Params
	a.paramSpace.GetParamSet(ctx, &p)
	return cdc.MustMarshalJSON(&group.GenesisState{
		Params: p,
	})
}

func (a AppModule) RegisterInvariants(sdk.InvariantRegistry) {
	// todo: anything to check?
	// todo: check that tally sums must never have less than block before ?
}

// Route returns the message routing key for the group module.
func (a AppModule) Route() sdk.Route {
	return sdk.Route{}
}

func (a AppModule) QuerierRoute() string {
	return ""
}

func (a AppModule) LegacyQuerierHandler(*codec.LegacyAmino) sdk.Querier {
	return nil
}

func (a AppModule) RegisterServices(configurator module.Configurator) {
	server.RegisterServices(a.storeKey, a.paramSpace, a.router, configurator)
}

func (a AppModule) BeginBlock(sdk.Context, abci.RequestBeginBlock) {}

func (a AppModule) EndBlock(sdk.Context, abci.RequestEndBlock) []abci.ValidatorUpdate {
	return []abci.ValidatorUpdate{}
}
