package geo

import (
	"encoding/json"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/gogo/protobuf/grpc"
	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/regen-network/regen-ledger/x/geo/types"
	"github.com/spf13/cobra"
	abci "github.com/tendermint/tendermint/abci/types"
)

const (
	ModuleName = "geo"
)

var (
	_ module.AppModule      = AppModule{}
	_ module.AppModuleBasic = AppModuleBasic{}
)

// AppModuleBasic defines the basic application module used by the geo module.
type AppModuleBasic struct {
	cdc codec.Marshaler
}

var _ module.AppModuleBasic = AppModuleBasic{}

// module name
func (AppModuleBasic) Name() string {
	return ModuleName
}

// RegisterLegacyAminoCodec registers the geo module's types for the given codec.
func (AppModuleBasic) RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	types.RegisterLegacyAminoCodec(cdc)
}

// RegisterInterfaces registers the module's interface types
func (b AppModuleBasic) RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	types.RegisterInterfaces(registry)
}

// default genesis state
func (AppModuleBasic) DefaultGenesis(cdc codec.JSONMarshaler) json.RawMessage {
	return nil
	// TODO
	// return cdc.MustMarshalJSON(types.DefaultGenesisState())
}

// module validate genesis
func (AppModuleBasic) ValidateGenesis(cdc codec.JSONMarshaler, config client.TxEncodingConfig, bz json.RawMessage) error {
	return nil
	// TODO
	// var data types.GenesisState
	// err := cdc.UnmarshalJSON(bz, &data)
	// if err != nil {
	// 	return fmt.Errorf("failed to unmarshal %s genesis state: %w", types.ModuleName, err)
	// }

	// return types.ValidateGenesis(data)
}

// RegisterRESTRoutes registers the REST routes for the slashing module.
func (AppModuleBasic) RegisterRESTRoutes(clientCtx client.Context, rtr *mux.Router) {
	// TODO
}

// RegisterGRPCRoutes registers the gRPC Gateway routes for the slashig module.
func (AppModuleBasic) RegisterGRPCRoutes(clientCtx client.Context, mux *runtime.ServeMux) {
	// TODO
}

// GetTxCmd returns the root tx command for the slashing module.
func (AppModuleBasic) GetTxCmd() *cobra.Command {
	// TODO
	return nil
}

// GetQueryCmd returns no root query command for the slashing module.
func (AppModuleBasic) GetQueryCmd() *cobra.Command {
	// TODO
	return nil
}

//____________________________________________________________________________

// AppModule implements an application module for the geo module.
type AppModule struct {
	AppModuleBasic
	keeper Keeper
}

// NewAppModule creates a new AppModule object
func NewAppModule(cdc codec.Marshaler, keeper Keeper) AppModule {
	return AppModule{
		AppModuleBasic: AppModuleBasic{cdc: cdc},
		keeper:         keeper,
	}
}

// module name
func (AppModule) Name() string {
	return ModuleName
}

// register invariants
func (am AppModule) RegisterInvariants(ir sdk.InvariantRegistry) {
}

// module message route name
func (am AppModule) Route() sdk.Route {
	return sdk.NewRoute(types.RouterKey, NewHandler(am.keeper))
}

// module handler
func (am AppModule) NewHandler() sdk.Handler {
	return NewHandler(am.keeper)
}

// module querier route name
func (AppModule) QuerierRoute() string {
	// TODO
	return ""
}

// LegacyQuerierHandler returns the slashing module sdk.Querier.
func (am AppModule) LegacyQuerierHandler(legacyQuerierCdc *codec.LegacyAmino) sdk.Querier {
	// TODO
	return nil
}

// RegisterQueryService registers a GRPC query service to respond to the
// module-specific GRPC queries.
func (am AppModule) RegisterQueryService(server grpc.Server) {
	// TODO
	// types.RegisterQueryServer(server, am.keeper)
}

// module init-genesis
func (am AppModule) InitGenesis(ctx sdk.Context, cdc codec.JSONMarshaler, data json.RawMessage) []abci.ValidatorUpdate {
	return nil
	// TODO
	// var genesisState GenesisState
	// cdc.MustUnmarshalJSON(data, &genesisState)
	// InitGenesis(ctx, am.keeper, genesisState)
	// return []abci.ValidatorUpdate{}
}

// module export genesis
func (am AppModule) ExportGenesis(ctx sdk.Context, cdc codec.JSONMarshaler) json.RawMessage {
	// gs := ExportGenesis(ctx, am.keeper)
	// return cdc.MustMarshalJSON(gs)
	return nil
}

// module begin-block
func (AppModule) BeginBlock(_ sdk.Context, _ abci.RequestBeginBlock) {
}

// module end-block
func (am AppModule) EndBlock(ctx sdk.Context, _ abci.RequestEndBlock) []abci.ValidatorUpdate {
	return []abci.ValidatorUpdate{}
}
