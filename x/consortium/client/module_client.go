package client

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/spf13/cobra"
	"github.com/tendermint/go-amino"
	consortiumcmd "gitlab.com/regen-network/regen-ledger/x/consortium/client/cli"
)

// ModuleClient exports all client functionality from this module
type ModuleClient struct {
	cdc *amino.Codec
}

func NewModuleClient(cdc *amino.Codec) ModuleClient {
	return ModuleClient{cdc}
}

// GetQueryCmd returns the cli query commands for this module
func (mc ModuleClient) GetQueryCmd() *cobra.Command {
	consortiumQueryCmd := &cobra.Command{
		Use:   "consortium",
		Short: "Querying commands for the consortium module",
	}

	return consortiumQueryCmd
}

// GetTxCmd returns the transaction commands for this module
func (mc ModuleClient) GetTxCmd() *cobra.Command {
	consortiumTxCmd := &cobra.Command{
		Use:   "consortium",
		Short: "Consortium transactions subcommands",
	}

	consortiumTxCmd.AddCommand(client.PostCommands(
		consortiumcmd.GetCmdProposeUpgrade(mc.cdc),
	)...)

	return consortiumTxCmd
}
