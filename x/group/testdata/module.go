package testdata

import (
	"encoding/json"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
	abci "github.com/tendermint/tendermint/abci/types"
)

const ModuleName = "testdata"

type AppModule struct {
	keeper Keeper
}

func NewAppModule(keeper Keeper) AppModule {
	return AppModule{
		keeper: keeper,
	}
}

func (a AppModule) Name() string {
	return ModuleName
}

func (a AppModule) RegisterCodec(cdc *codec.Codec) {
	RegisterCodec(cdc)
}

func (a AppModule) DefaultGenesis(cdc codec.JSONMarshaler) json.RawMessage {
	return nil
}

func (a AppModule) ValidateGenesis(cdc codec.JSONMarshaler, bz json.RawMessage) error {
	panic("implement me")
}

func (a AppModule) RegisterRESTRoutes(clientCtx client.Context, r *mux.Router) {
	panic("implement me")
}

func (a AppModule) GetTxCmd(clientCtx client.Context) *cobra.Command {
	panic("implement me")
}

func (a AppModule) GetQueryCmd(*codec.Codec) *cobra.Command {
	panic("implement me")
}

func (a AppModule) InitGenesis(ctx sdk.Context, cdc codec.JSONMarshaler, data json.RawMessage) []abci.ValidatorUpdate {
	return nil
}

func (a AppModule) ExportGenesis(ctx sdk.Context, cdc codec.JSONMarshaler) json.RawMessage {
	return nil
}

func (a AppModule) RegisterInvariants(sdk.InvariantRegistry) {
}

func (a AppModule) Route() string {
	return ModuleName
}

func (a AppModule) NewHandler() sdk.Handler {
	return NewHandler(a.keeper)
}

func (a AppModule) QuerierRoute() string {
	return ModuleName
}

func (a AppModule) NewQuerierHandler() sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, error) {
		panic("not implemented")
	}
}
func (a AppModule) BeginBlock(sdk.Context, abci.RequestBeginBlock) {}

func (a AppModule) EndBlock(sdk.Context, abci.RequestEndBlock) []abci.ValidatorUpdate {
	return nil
}
