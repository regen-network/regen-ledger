package client

import (
	"github.com/cosmos/cosmos-sdk/client"
	geocmd "github.com/regen-network/regen-ledger/x/geo/client/cli"
	"github.com/spf13/cobra"
	"github.com/tendermint/go-amino"
)

// ModuleClient exports all client functionality from this module
type ModuleClient struct {
	//storeKey string
	cdc *amino.Codec
}

func NewModuleClient(cdc *amino.Codec) ModuleClient {
	return ModuleClient{cdc}
}

func (mc ModuleClient) GetQueryCmd() *cobra.Command {
	geoQueryCmd := &cobra.Command{
		Use:   "geo",
		Short: "Querying commands for the geo module",
	}

	//geoQueryCmd.AddCommand(client.GetCommands(
	//	//geocmd.GetCmdGetData(mc.storeKey, mc.cdc),
	//)...)

	return geoQueryCmd
}

func (mc ModuleClient) GetTxCmd() *cobra.Command {
	geoTxCmd := &cobra.Command{
		Use:   "geo",
		Short: "geo transactions subcommands",
	}

	cdc := mc.cdc

	geoTxCmd.AddCommand(client.PostCommands(
		geocmd.GetCmdStoreGeometry(cdc),
	)...)

	return geoTxCmd
}
