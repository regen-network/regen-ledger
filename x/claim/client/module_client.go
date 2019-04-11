package client

import (
	"github.com/cosmos/cosmos-sdk/client"
	claimcmd "github.com/regen-network/regen-ledger/x/claim/client/cli"
	"github.com/spf13/cobra"
	"github.com/tendermint/go-amino"
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
	queryCmd := &cobra.Command{
		Use:   "claim",
		Short: "Querying commands for the claim module",
	}

	queryCmd.AddCommand(client.GetCommands(
		claimcmd.GetSignaturesQueryCmd(mc.storeKey, mc.cdc),
		claimcmd.GetEvidenceQueryCmd(mc.storeKey, mc.cdc),
	)...)

	return queryCmd
}

// GetTxCmd returns the transaction commands for this module
func (mc ModuleClient) GetTxCmd() *cobra.Command {
	txCmd := &cobra.Command{
		Use:   "claim",
		Short: "Claim transactions subcommands",
	}

	txCmd.AddCommand(client.PostCommands(
		claimcmd.GetCmdSignClaim(mc.cdc),
	)...)

	return txCmd
}

