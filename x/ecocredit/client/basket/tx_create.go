package basket

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"strconv"
	"strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"

	keeper "github.com/regen-network/regen-ledger/x/ecocredit/basket"
)

const (
	FlagDisableAutoRetire = "disable-auto-retire"
)

func TxCreateBasket() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-basket [curator] [name] [display_name] [exponent] [credit_type_name] [allowed_classes] [min_start_date]",
		Short: "Creates a basket",
		Long: strings.TrimSpace(`Creates a basket
Parameters:
		curator: the account address of the curator of the basket.
			Note, the '--from' flag is ignored as it is implied from [curator]
		name: the name used to create a bank denom for this basket token.
		display_name: the name used to create a bank Metadata display name.
		exponent: the exponent that will be used for converting credits to basket tokens
			and for bank denom metadata. It also limits the precision of credit amounts
			when putting credits into a basket. An exponent of 6 will mean that 10^6 units
			of a basket token will be issued for 1.0 credits and that this should be
			displayed as one unit in user interfaces. It also means that the maximum
			precision of credit amounts is 6 decimal places so that the need to round is
			eliminated. The exponent must be >= the precision of the credit type at the
			time the basket is created.
		credit_type_name: filters against credits from this credit type name.
		allowed_classes: comma separated (no spaces) list of credit classes allowed to be put in
			the basket. Example: "C01,C02,B01"
		min_start_date: the earliest start date for batches of credits allowed into the basket.
		fee: the fee that the curator will pay to create the basket. It must be >= the required
			Params.basket_creation_fee. We include the fee explicitly here so that the curator
			explicitly acknowledges paying this fee and is not surprised to learn that the paid
			a big fee and didn't know beforehand.
Flags:
		disable_auto_retire: disables the auto-retirement of credits upon taking credits
			from the basket. The credits will be auto-retired if disable_auto_retire is
			false unless the credits were previously put into the basket by the address
			picking them from the basket, in which case they will remain tradable.
		`),
		Args: cobra.ExactArgs(7),
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.Flags().Set(flags.FlagFrom, args[0])
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			exponent, err := strconv.ParseUint(args[3], 10, 32)
			if err != nil {
				return err
			}

			disableAutoRetire, err := cmd.Flags().GetBool(FlagDisableAutoRetire)
			if err != nil {
				return err
			}

			classes := strings.Split(args[5], ",")
			for i := range classes {
				classes[i] = strings.TrimSpace(classes[i])
			}

			coins, err := sdk.ParseCoinsNormalized(args[6])
			if err != nil {
				return fmt.Errorf("failed to parse coins: %w", err)
			}

			msg := keeper.MsgCreate{
				Curator:           clientCtx.FromAddress.String(),
				Name:              args[1],
				DisplayName:       args[2],
				Exponent:          uint32(exponent),
				DisableAutoRetire: disableAutoRetire,
				CreditTypeName:    args[4],
				AllowedClasses:    classes,
				Fee:               coins,
			}

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	cmd.Flags().Bool(FlagDisableAutoRetire, false, "dictates whether credits will be auto-retired upon taking credits from the basket")

	return cmd
}
