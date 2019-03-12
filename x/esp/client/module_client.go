package client

import (
	"github.com/cosmos/cosmos-sdk/client"
	espcmd "github.com/regen-network/regen-ledger/x/esp/client/cli"
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
	espQueryCmd := &cobra.Command{
		Use:   "esp",
		Short: "Querying commands for the ESP module",
	}

	//espQueryCmd.AddCommand(client.GetCommands(
	//	//espcmd.GetCmdGetData(mc.storeKey, mc.cdc),
	//)...)

	return espQueryCmd
}

func (mc ModuleClient) GetTxCmd() *cobra.Command {
	espTxCmd := &cobra.Command{
		Use:   "esp",
		Short: "ESP transactions subcommands",
	}

	cdc := mc.cdc

	espTxCmd.AddCommand(client.PostCommands(
		espcmd.GetCmdProposeVersion(cdc),
		espcmd.GetCmdReportResult(cdc),
	)...)

	return espTxCmd
}
