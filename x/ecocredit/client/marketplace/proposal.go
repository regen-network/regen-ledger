package marketplace

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"
	"github.com/cosmos/cosmos-sdk/x/gov/client/cli"
	govv1beta1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"

	"github.com/regen-network/regen-ledger/x/ecocredit/marketplace"
)

var AllowDenomProposalHandler = govclient.NewProposalHandler(TxAllowDenomProposal)

func TxAllowDenomProposal() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "allow-denom-proposal [path_to_file.json] [flags]",
		Args:  cobra.ExactArgs(1),
		Short: "Submit a proposal to add a denom to the list of allowed denoms",
		Long: strings.TrimSpace(`Submit a proposal to add a denom to the list of allowed denoms for use in the marketplace. 
The json file MUST take the following form:
{
    "title": "some title",
    "description": "some description",
    "denom": {
	"bank_denom": "uregen",
        "display_denom": "regen",
        "exponent": 6
    }
}
The bank denom is the underlying coin denom (i.e. ibc/CDC4587874B85BEA4FCEC3CEA5A1195139799A1FEE711A07D972537E18FD). 
Display denom is used for display purposes, and serves as the name of the coin denom (i.e. ATOM). Exponent is used to 
relate the bank_denom to the display_denom and is informational`),
		Example: `regen tx gov submit-proposal allow-denom-proposal my_file.json --deposit=100regen`,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			proposalFile, err := ioutil.ReadFile(args[0])
			if err != nil {
				return err
			}

			var proposal marketplace.AllowDenomProposal
			err = json.Unmarshal(proposalFile, &proposal)
			if err != nil {
				return err
			}
			if err := proposal.ValidateBasic(); err != nil {
				return fmt.Errorf("invalid proposal: %w", err)
			}

			depositStr, err := cmd.Flags().GetString(cli.FlagDeposit)
			if err != nil {
				return err
			}
			deposit, err := sdk.ParseCoinsNormalized(depositStr)
			if err != nil {
				return err
			}
			var content govv1beta1.Content = &proposal
			msg, err := govv1beta1.NewMsgSubmitProposal(content, deposit, clientCtx.GetFromAddress())
			if err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	cmd.Flags().String(cli.FlagDeposit, "", "deposit of proposal")
	return cmd
}
