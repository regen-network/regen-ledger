package basket

import (
	"strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"

	keeper "github.com/regen-network/regen-ledger/x/ecocredit/basket"
)

func TxAddToBasket() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-to-basket [owner] [basket_denom] [credits_json_file]",
		Short: "add credits to the basket",
		Long: strings.TrimSpace(`add credits to the basket.

Parameters:
		owner: account address of the owner of credits being put into the basket.
		Note, the '--from' flag is ignored as it is implied from [owner]
		basket_denom: basket denom is the basket denom to add credits to.

Example:
		$regen tx ecocredit add-to-basket [owner] [basket_denom] [credits_json_file]
		
		Where credits_json_file contains:
		
		{
			"credits": [
				{
					"batch_denom": "C01-20210101-20220101-001",
					"amount": "10",
				},
				{
					"batch_denom": "C01-20210101-20220101-001",
					"amount": "10.5",
				}
			]
		}
		`),
		Args: cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.Flags().Set(flags.FlagFrom, args[0])
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			credits, err := parseBasketCredits(clientCtx, args[2])
			if err != nil {
				return err
			}

			msg := keeper.MsgPut{
				Owner:       clientCtx.FromAddress.String(),
				BasketDenom: args[1],
				Credits:     credits,
			}

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
