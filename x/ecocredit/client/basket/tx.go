package basketclient

import (
	"strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"

	"github.com/regen-network/regen-ledger/x/ecocredit/basket"
)

func TxPutInBasket() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "put-in-basket [basket_denom] [credits_json_file]",
		Short: "add credits to the basket",
		Long: strings.TrimSpace(`add credits to the basket.

Parameters:
		basket_denom: basket identifier

Flags:
		from: account address of the owner

Example:
		$regen tx ecocredit put-in-basket [basket_denom] [credits_json_file]
		
		Where credits_json_file contains:
		
		[
			{
				"batch_denom": "C01-20210101-20220101-001",
				"amount": "10"
			},
			{
				"batch_denom": "C01-20210101-20220101-001",
				"amount": "10.5"
			}
		]
		`),
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			credits, err := parseBasketCredits(args[1])
			if err != nil {
				return err
			}

			msg := basket.MsgPut{
				Owner:       clientCtx.FromAddress.String(),
				BasketDenom: args[0],
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
