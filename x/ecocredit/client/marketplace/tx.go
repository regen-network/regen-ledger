package marketplace

import (
	"encoding/json"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/regen-network/regen-ledger/x/ecocredit/marketplace"
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	FlagRetirementLocation = "retirement-location"
)

// TxBuyDirect returns a transaction command for a single direct buy order.
func TxBuyDirect() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "buy-direct [sell_order_id] [quantity] [bid_price] [disable_auto_retire]",
		Short: "Buy ecocredits from a specific sell order",
		Long: "Purchase ecocredits from a specific sell order. AutoRetire can be set to true to retire the credits immediately upon purchase." +
			"NOTE: The bid price is the price paid PER credit. The total cost will be quantity * bid_price.",
		Example: "regen tx ecocredit buy-direct 194 300 40regen true --retirement-location=US-NY",
		Args:    cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			sellOrderIdStr, qtyStr, bidPriceStr, autoRetireStr := args[0], args[1], args[2], args[3]

			sellOrderId, err := strconv.ParseUint(sellOrderIdStr, 10, 64)
			if err != nil {
				return err
			}

			bidPrice, err := sdk.ParseCoinNormalized(bidPriceStr)
			if err != nil {
				return err
			}

			disableAutoRetire, err := strconv.ParseBool(autoRetireStr)
			if err != nil {
				return err
			}
			var retireLocation string
			retireLocation, err = cmd.Flags().GetString(FlagRetirementLocation)
			if err != nil {
				return err
			}

			msg := marketplace.MsgBuyDirect{
				Buyer: clientCtx.GetFromAddress().String(),
				Orders: []*marketplace.MsgBuyDirect_Order{
					{
						SellOrderId:        sellOrderId,
						Quantity:           qtyStr,
						BidPrice:           &bidPrice,
						DisableAutoRetire:  disableAutoRetire,
						RetirementLocation: retireLocation,
					},
				},
			}
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	cmd.Flags().String(FlagRetirementLocation, "", "the location to use for retirement when auto retire is true.")
	return cmd
}

// TxBuyDirectBatch returns a transaction command for a batch direct buy order using a json file.
func TxBuyDirectBatch() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "buy-direct-batch [name_of_file.json]",
		Short: "Buy ecocredits from multiple sell orders",
		Long:  "Directly buy ecocredits from different sell orders using a json file.",
		Example: strings.TrimSpace(`regen tx ecocredit buy-direct-batch batch.json
		where batch.json has the following form:
		[
			{
			   "sell_order_id": 52,
			   "quantity": "32.5",
			   "bid_price": {"denom": "uregen", "amount": "32000000"},
			   "disable_auto_retire": false,
			   "retirement_location": "US-NY"
			},
		]`),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			batch, err := ioutil.ReadFile(args[0])
			if err != nil {
				return err
			}

			var orders []*marketplace.MsgBuyDirect_Order
			err = json.Unmarshal(batch, &orders)
			if err != nil {
				return err
			}

			msg := marketplace.MsgBuyDirect{
				Buyer:  clientCtx.GetFromAddress().String(),
				Orders: orders,
			}
			if err = msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
