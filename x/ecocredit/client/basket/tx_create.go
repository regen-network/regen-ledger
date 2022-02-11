package basket

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gogo/protobuf/types"
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/regen-network/regen-ledger/x/ecocredit/basket"
)

const (
	FlagDisplayName       = "display-name"
	FlagExponent          = "exponent"
	FlagDisableAutoRetire = "disable-auto-retire"
	FlagCreditTypeName    = "credit-type-name"
	FlagAllowedClasses    = "allowed-classes"
	FlagMinimumStartDate  = "minimum-start-date"
	FlagStartDateWindow   = "start-date-window"
	FlagBasketFee         = "basket-fee"
)

func TxCreateBasket() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-basket [name]",
		Short: "Creates a bank denom that wraps credits",
		Long: strings.TrimSpace(`Creates a bank denom that wraps credits

Parameters:
		name: the name used to create a bank denom for this basket token.

Flags:
		display_name: the name used to create a bank denom display name.
		exponent: the exponent used for converting credits to basket tokens and for bank
			denom metadata. The exponent also limits the precision of credit amounts
			when putting credits into a basket. An exponent of 6 will mean that 10^6 units
			of a basket token will be issued for 1.0 credits and that this should be
			displayed as one unit in user interfaces. It also means that the maximum
			precision of credit amounts is 6 decimal places so that the need to round is
			eliminated. The exponent must be >= the precision of the credit type at the
			time the basket is created.
		disable_auto_retire: disables the auto-retirement of credits upon taking credits
			from the basket. The credits will be auto-retired if disable_auto_retire is
			false unless the credits were previously put into the basket by the address
			picking them from the basket, in which case they will remain tradable.
		credit_type_name: filters against credits from this credit type name (e.g. "carbon").
		allowed_classes: comma separated (no spaces) list of credit classes allowed to be put in
			the basket (e.g. "C01,C02").
		min_start_date: the earliest start date for batches of credits allowed into the basket.
			Note: either min_start_date or start-date-window is required and not both.
		start-date-window: the duration of time measured into the past which sets a cutoff for
			batch start dates when adding new credits to the basket.
			Note: either min_start_date or start-date-window is required and not both.
		basket_fee: the fee that the curator will pay to create the basket. It must be >= the
			required Params.basket_creation_fee. We include the fee explicitly here so that the
			curator explicitly acknowledges paying this fee and is not surprised to learn that the
			paid a big fee and didn't know beforehand.
		`),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			displayName, err := cmd.Flags().GetString(FlagDisplayName)
			if err != nil {
				return err
			}

			exponentString, err := cmd.Flags().GetString(FlagExponent)
			if err != nil {
				return err
			}
			exponent, err := strconv.ParseUint(exponentString, 10, 32)
			if err != nil {
				return err
			}

			disableAutoRetire, err := cmd.Flags().GetBool(FlagDisableAutoRetire)
			if err != nil {
				return err
			}

			creditTypeName, err := cmd.Flags().GetString(FlagCreditTypeName)
			if err != nil {
				return err
			}

			allowedClassesString, err := cmd.Flags().GetString(FlagAllowedClasses)
			if err != nil {
				return err
			}
			allowedClasses := strings.Split(allowedClassesString, ",")
			for i := range allowedClasses {
				allowedClasses[i] = strings.TrimSpace(allowedClasses[i])
			}

			minStartDateString, err := cmd.Flags().GetString(FlagMinimumStartDate)
			if err != nil {
				return err
			}
			startDateWindowString, err := cmd.Flags().GetString(FlagMinimumStartDate)
			if err != nil {
				return err
			}

			if minStartDateString == "" && startDateWindowString == "" {
				return fmt.Errorf("either min-start-date or start-date-window required")
			} else if minStartDateString != "" && startDateWindowString != "" {
				return fmt.Errorf("both min-start-date and start-date-window cannot be set")
			}

			dateCriteria := basket.DateCriteria{}

			if minStartDateString != "" {
				minStartDateTime, err := time.Parse("2006-01-02", minStartDateString)
				if err != nil {
					return fmt.Errorf("failed to parse min_start_date: %w", err)
				}
				minStartDate, err := types.TimestampProto(minStartDateTime)
				if err != nil {
					return fmt.Errorf("failed to parse min_start_date: %w", err)
				}
				dateCriteria.Sum = &basket.DateCriteria_MinStartDate{
					MinStartDate: minStartDate,
				}
			}

			if startDateWindowString != "" {
				startDateWindowInt, err := cmd.Flags().GetInt64(FlagMinimumStartDate)
				if err != nil {
					return err
				}
				startDateWindowDuration := time.Duration(startDateWindowInt)
				if err != nil {
					return fmt.Errorf("failed to parse start-date-window: %w", err)
				}
				startDateWindow := types.DurationProto(startDateWindowDuration)
				if err != nil {
					return fmt.Errorf("failed to parse start-date-window: %w", err)
				}
				dateCriteria.Sum = &basket.DateCriteria_StartDateWindow{
					StartDateWindow: startDateWindow,
				}
			}

			feeString, err := cmd.Flags().GetString(FlagBasketFee)
			if err != nil {
				return err
			}
			fee, err := sdk.ParseCoinsNormalized(feeString)
			if err != nil {
				return fmt.Errorf("failed to parse basket_fee: %w", err)
			}

			msg := basket.MsgCreate{
				Curator:           clientCtx.FromAddress.String(),
				Name:              args[0],
				DisplayName:       displayName,
				Exponent:          uint32(exponent),
				DisableAutoRetire: disableAutoRetire,
				CreditTypeName:    creditTypeName,
				AllowedClasses:    allowedClasses,
				DateCriteria:      &dateCriteria,
				Fee:               fee,
			}

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	// command flags
	cmd.Flags().String(FlagDisplayName, "", "the name used to create a bank denom display name")
	cmd.Flags().String(FlagExponent, "", "the exponent used for converting credits to basket tokens")
	cmd.Flags().Bool(FlagDisableAutoRetire, false, "dictates whether credits will be auto-retired upon taking")
	cmd.Flags().String(FlagCreditTypeName, "", "filters against credits from this credit type name (e.g. \"carbon\")")
	cmd.Flags().String(FlagAllowedClasses, "", "comma separated (no spaces) list of credit classes allowed to be put in the basket (e.g. \"C01,C02\")")
	cmd.Flags().String(FlagMinimumStartDate, "", "the earliest start date for batches of credits allowed into the basket (e.g. \"2012-01-01\")")
	cmd.Flags().String(FlagStartDateWindow, "", "sets a cutoff for batch start dates when adding new credits to the basket (e.g. \"1325404800\")")
	cmd.Flags().String(FlagBasketFee, "", "the fee that the curator will pay to create the basket (e.g. \"20regen\")")

	// required flags
	cmd.MarkFlagRequired(FlagDisplayName)
	cmd.MarkFlagRequired(FlagExponent)
	cmd.MarkFlagRequired(FlagCreditTypeName)
	cmd.MarkFlagRequired(FlagAllowedClasses)
	cmd.MarkFlagRequired(FlagBasketFee)

	return cmd
}
