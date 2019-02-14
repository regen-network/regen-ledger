package client

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/spf13/cobra"
	"github.com/tendermint/go-amino"
	agentcmd "gitlab.com/regen-network/regen-ledger/x/agent/client/cli"
)

// ModuleClient exports all client functionality from this module
type ModuleClient struct {
	storeKey string
	cdc      *amino.Codec
}

func NewModuleClient(storeKey string, cdc *amino.Codec) ModuleClient {
	return ModuleClient{storeKey, cdc}
}

// GetQueryCmd returns the cli query commands for this module
func (mc ModuleClient) GetQueryCmd() *cobra.Command {
	agentQueryCmd := &cobra.Command{
		Use:   "agent",
		Short: "Querying commands for the agent module",
	}

	//dataQueryCmd.AddCommand(client.GetCommands(
	//	datacmd.GetCmdGetData(mc.storeKey, mc.cdc),
	//)...)

	return agentQueryCmd
}

// GetTxCmd returns the transaction commands for this module
func (mc ModuleClient) GetTxCmd() *cobra.Command {
	agentTxCmd := &cobra.Command{
		Use:   "agent",
		Short: "Agent transactions subcommands",
	}

	agentTxCmd.AddCommand(client.PostCommands(
		agentcmd.GetCmdCreateAgent(mc.cdc),
	)...)

	return agentTxCmd
}

