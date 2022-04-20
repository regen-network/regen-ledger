package core

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"
	"github.com/cosmos/cosmos-sdk/x/gov/client/rest"
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/gov/client/cli"
	"github.com/cosmos/cosmos-sdk/x/gov/types"

	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

var CreditTypeProposalHandler = govclient.NewProposalHandler(TxCreditTypeProposalCmd, func(context client.Context) rest.ProposalRESTHandler {
	return rest.ProposalRESTHandler{
		SubRoute: "",
		Handler:  nil,
	}
})

func TxCreditTypeProposalCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "credit-type-proposal [path_to_file.json] [flags]",
		Args:  cobra.ExactArgs(1),
		Short: "Submit a proposal for a new credit type",
		Long: strings.TrimSpace(`Submit a proposal to add a new credit type. 
The json file MUST take the following form:
{
	"title": "some title",
	"description": "some description",
	"credit_type": {
					"abbreviation": "C",
					"name": "carbon",
					"unit": "metric ton C02",
					"precision": 6
	}
}
The credit type abbreviation MUST be unique, else the proposal will fail upon execution. Units are measurement units 
(i.e. metric ton C02). Precision is how many decimal places are allowed in the credits.`),
		Example: `regen tx gov submit-proposal credit-type-proposal my_file.json`,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			proposalFile, err := ioutil.ReadFile(args[0])
			if err != nil {
				return err
			}

			var proposal core.CreditTypeProposal
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
			var content types.Content = &proposal
			msg, err := types.NewMsgSubmitProposal(content, deposit, clientCtx.GetFromAddress())
			if err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	cmd.Flags().String(cli.FlagDeposit, "", "deposit of proposal")
	return cmd
}
