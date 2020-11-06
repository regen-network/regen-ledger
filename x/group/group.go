package group

import (
	"encoding/json"
	"fmt"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/gorilla/mux"
	"github.com/regen-network/regen-ledger/orm"
	"github.com/spf13/cobra"
	abci "github.com/tendermint/tendermint/abci/types"
)

const (
	// ModuleName is the module name constant used in many places
	ModuleName = "group"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	DefaultParamspace = ModuleName
)

// AccountCondition returns a condition to build a group account address.
func AccountCondition(id uint64) Condition {
	return NewCondition("group", "account", orm.EncodeSequence(id))
}

type AppModuleBasic struct{}

var _ module.AppModule = AppModule{}
var _ module.AppModuleBasic = AppModuleBasic{}

func (a AppModuleBasic) Name() string {
	return ModuleName
}

func (AppModuleBasic) RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	RegisterLegacyAminoCodec(cdc)
}

func (a AppModuleBasic) DefaultGenesis(marshaler codec.JSONMarshaler) json.RawMessage {
	return marshaler.MustMarshalJSON(NewGenesisState())
}

func (a AppModuleBasic) ValidateGenesis(cdc codec.JSONMarshaler, config client.TxEncodingConfig, bz json.RawMessage) error {
	var data GenesisState
	if err := cdc.UnmarshalJSON(bz, &data); err != nil {
		return fmt.Errorf("failed to unmarshal %s genesis state: %w", ModuleName, err)
	}
	return data.Validate()
}

func (a AppModuleBasic) RegisterRESTRoutes(clientCtx client.Context, r *mux.Router) {
	//rest.RegisterRoutes(ctx, r, moduleCdc, RouterKey)
	// todo: what client functions do we want to support?
	return
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
func (AppModuleBasic) RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	RegisterInterfaces(registry)
}

type AppModule struct {
	AppModuleBasic
	keeper Keeper
}

func NewAppModule(keeper Keeper) AppModule {
	return AppModule{
		AppModuleBasic: AppModuleBasic{},
		keeper:         keeper,
	}
}

func (a AppModule) InitGenesis(ctx sdk.Context, cdc codec.JSONMarshaler, data json.RawMessage) []abci.ValidatorUpdate {
	var genesisState GenesisState
	cdc.MustUnmarshalJSON(data, &genesisState)
	if err := genesisState.Validate(); err != nil {
		panic(fmt.Sprintf("failed to validate %s genesis state: %s", ModuleName, err))
	}
	a.keeper.setParams(ctx, genesisState.Params) // TODO: revisit if this makes sense
	return []abci.ValidatorUpdate{}
}

func (a AppModule) ExportGenesis(ctx sdk.Context, cdc codec.JSONMarshaler) json.RawMessage {
	return cdc.MustMarshalJSON(ExportGenesis(ctx, a.keeper))
}

func (a AppModule) RegisterInvariants(sdk.InvariantRegistry) {
	// todo: anything to check?
	// todo: check that tally sums must never have less than block before ?
}

// Route returns the message routing key for the group module.
func (a AppModule) Route() sdk.Route {
	return sdk.NewRoute(RouterKey, NewHandler(a.keeper))
}

func (a AppModule) NewHandler() sdk.Handler {
	return NewHandler(a.keeper)
}

func (a AppModule) QuerierRoute() string {
	return QuerierRoute
}

func (a AppModule) LegacyQuerierHandler(*codec.LegacyAmino) sdk.Querier {
	return NewQuerier(a.keeper)
}

func (a AppModule) RegisterServices(configurator module.Configurator) {
}

func (a AppModule) BeginBlock(sdk.Context, abci.RequestBeginBlock) {}

func (a AppModule) EndBlock(sdk.Context, abci.RequestEndBlock) []abci.ValidatorUpdate {
	return []abci.ValidatorUpdate{}
}
