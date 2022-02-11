package basket

import (
	"strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"

	keeper "github.com/regen-network/regen-ledger/x/ecocredit/basket"
)

const (
	FlagRetirementLocation = "retirement-location"
	FlagRetireOnTake       = "retire-on-take"
)

func TxTakeFromBasket() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "take-from-basket [owner] [basket_denom] [amount]",
		Short: "Takes credits from a basket",
		Long: strings.TrimSpace(`takes credits from a basket starting from the oldest credits first.

Parameters:
		owner: account address of the owner of the basket.
		Note, the '--from' flag is ignored as it is implied from [owner]
		basket_denom: basket denom is the basket denom to take credits from.
		amount: amount is a positive integer number of basket tokens to convert into credits.
Flags:
		retirement-location: retirement location is the optional retirement location for the credits
		                    which will be used only if --retire-on-take flag is true.
		retire-on-take: retire on take is a boolean that dictates whether the ecocredits
		                received in exchange for the basket tokens will be received as
		                retired or tradable credits.
		
		`),
		Args: cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.Flags().Set(flags.FlagFrom, args[0])
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			retirementLocation, err := cmd.Flags().GetString(FlagRetirementLocation)
			if err != nil {
				return err
			}

			retireOnTake, err := cmd.Flags().GetBool(FlagRetireOnTake)
			if err != nil {
				return err
			}

			msg := keeper.MsgTake{
				Owner:              clientCtx.FromAddress.String(),
				BasketDenom:        args[1],
				Amount:             args[2],
				RetirementLocation: retirementLocation,
				RetireOnTake:       retireOnTake,
			}

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	cmd.Flags().String(FlagRetirementLocation, "", "location for the credits which will be used only if --retire-on-take flag is true")
	cmd.Flags().Bool(FlagRetireOnTake, false, "dictates whether the ecocredits received in exchange for the basket tokens will be received as retired or tradable credits")

	return cmd
}
