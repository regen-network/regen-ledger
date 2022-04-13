package core

import (
	"fmt"
	"strconv"

	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/gov/client/cli"
	"github.com/cosmos/cosmos-sdk/x/gov/types"

	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

var CreditTypeProposalHandler = govclient.NewProposalHandler(TxCreditTypeProposalCmd, nil)

func TxCreditTypeProposalCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "credit-type-proposal [proposal-title] [proposal-description] [credit_type_abbreviation] [credit_type_name] [units] [precision] [flags]",
		Args:  cobra.ExactArgs(6),
		Short: "Submit a proposal for a new credit type",
		Long: "Submit a proposal to add a new credit type. The credit type abbreviation and name MUST be unique, else " +
			"the proposal will fail upon execution. Units are measurements units (i.e. metric tonne). Precision is how " +
			"many decimal places are allowed in the credits.",
		Example: `regen tx gov submit-proposal credit-type-proposal "Add Biodiversity Type" "A biodiversity type would be great..." BIO biodiversity "sq. meters" 3`,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			title, desc, abbrev, name, units, precisionStr := args[0], args[1], args[2], args[3], args[4], args[5]
			precisionU64, err := strconv.ParseUint(precisionStr, 10, 32)
			if err != nil {
				return fmt.Errorf("invalid precision %s: %w", precisionStr, err)
			}
			proposal := core.CreditTypeProposal{
				Title:       title,
				Description: desc,
				CreditType: &core.CreditType{
					Abbreviation: abbrev,
					Name:         name,
					Unit:         units,
					Precision:    uint32(precisionU64),
				},
			}
			if err := proposal.ValidateBasic(); err != nil {
				return fmt.Errorf("invalid proposal: %w", err)
			}

			from := clientCtx.GetFromAddress()

			depositStr, err := cmd.Flags().GetString(cli.FlagDeposit)
			if err != nil {
				return err
			}
			deposit, err := sdk.ParseCoinsNormalized(depositStr)
			if err != nil {
				return err
			}
			var content types.Content = &proposal
			msg, err := types.NewMsgSubmitProposal(content, deposit, from)
			if err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	cmd.Flags().String(cli.FlagDeposit, "", "deposit of proposal")
	return cmd
}
