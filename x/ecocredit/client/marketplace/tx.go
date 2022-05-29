package marketplace

import (
	"encoding/json"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"sigs.k8s.io/yaml"

	"github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/regen-network/regen-ledger/x/ecocredit/marketplace"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	FlagRetirementJurisdiction = "retirement-jurisdiction"
)

// TxSellCmd returns a transaction command that creates sell orders.
func TxSellCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sell [orders]",
		Short: "Creates new sell orders with transaction author (--from) as owner",
		Long: `Creates new sell orders with transaction author (--from) as owner.

Parameters:
  orders:  YAML encoded order list. Note: numerical values must be written in strings.
           eg: '[{batch_denom: "C01-20210101-20210201-001", quantity: "5", ask_price: "100regen", disable_auto_retire: false}]'
           eg: '[{batch_denom: "C01-20210101-20210201-001", quantity: "5", ask_price: "100regen", disable_auto_retire: false, expiration: "2024-01-01"}]'`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			owner := clientCtx.GetFromAddress()

			// declare orders array with ask price as string
			var strOrders []struct {
				BatchDenom        string `json:"batch_denom"`
				Quantity          string `json:"quantity"`
				AskPrice          string `json:"ask_price"`
				DisableAutoRetire bool   `json:"disable_auto_retire"`
				Expiration        string `json:"expiration"`
			}

			// unmarshal YAML encoded orders with ask price as string
			if err := yaml.Unmarshal([]byte(args[0]), &strOrders); err != nil {
				return err
			}

			orders := make([]*marketplace.MsgSell_Order, len(strOrders))

			// loop through orders with ask price as string
			for i, o := range strOrders {

				askPrice, err := sdk.ParseCoinNormalized(o.AskPrice)
				if err != nil {
					return err
				}

				// set order with ask price as sdk.Coin
				orders[i] = &marketplace.MsgSell_Order{
					BatchDenom:        o.BatchDenom,
					AskPrice:          &askPrice,
					Quantity:          o.Quantity,
					DisableAutoRetire: o.DisableAutoRetire,
				}

				if o.Expiration != "" {
					tm, err := types.ParseDate("expiration", o.Expiration)
					if err != nil {
						return err
					}
					orders[i].Expiration = &tm
				}
			}

			// create sell message
			msg := marketplace.MsgSell{
				Owner:  owner.String(),
				Orders: orders,
			}

			// generate and broadcast transaction
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// TxUpdateSellOrdersCmd returns a transaction command that creates sell orders.
func TxUpdateSellOrdersCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-sell-orders [updates]",
		Short: "Updates existing sell orders with transaction author (--from) as owner",
		Long: `Updates existing sell orders with transaction author (--from) as owner.

Parameters:
  updates:  YAML encoded update list. Note: numerical values must be written in strings.
           eg: '[{sell_order_id: "1", new_quantity: "5", new_ask_price: "200regen", disable_auto_retire: false}]'
           eg: '[{sell_order_id: "1", new_quantity: "5", new_ask_price: "200regen", disable_auto_retire: false, new_expiration: "2026-01-01"}]'`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			// get the order owner from the --from flag
			owner := clientCtx.GetFromAddress()

			// declare updates array with ask price as string
			var strUpdates []struct {
				SellOrderId       string `json:"sell_order_id"`
				NewQuantity       string `json:"new_quantity"`
				NewAskPrice       string `json:"new_ask_price"`
				DisableAutoRetire bool   `json:"disable_auto_retire"`
				NewExpiration     string `json:"new_expiration"`
			}

			// unmarshal YAML encoded updates with new ask price as string
			if err := yaml.Unmarshal([]byte(args[0]), &strUpdates); err != nil {
				return err
			}

			// declare updates array with new ask price as sdk.Coin
			updates := make([]*marketplace.MsgUpdateSellOrders_Update, len(strUpdates))

			// loop through updates with new ask price as string
			for i, u := range strUpdates {

				// parse sell order id
				sellOrderId, err := strconv.ParseUint(u.SellOrderId, 10, 64)
				if err != nil {
					return ecocredit.ErrInvalidSellOrder.Wrap(err.Error())
				}

				// parse and normalize new ask price as sdk.Coin
				askPrice, err := sdk.ParseCoinNormalized(u.NewAskPrice)
				if err != nil {
					return err
				}

				// set update with new ask price as sdk.Coin
				updates[i] = &marketplace.MsgUpdateSellOrders_Update{
					SellOrderId:       sellOrderId,
					NewAskPrice:       &askPrice,
					NewQuantity:       u.NewQuantity,
					DisableAutoRetire: u.DisableAutoRetire,
				}

				if u.NewExpiration != "" {
					tm, err := types.ParseDate("expiration", u.NewExpiration)
					if err != nil {
						return err
					}
					updates[i].NewExpiration = &tm
				}
			}

			// create update sell orders message
			msg := marketplace.MsgUpdateSellOrders{
				Owner:   owner.String(),
				Updates: updates,
			}

			// generate and broadcast transaction
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// TxBuyDirect returns a transaction command for a single direct buy order.
func TxBuyDirect() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "buy-direct [sell_order_id] [quantity] [bid_price] [disable_auto_retire] [flags]",
		Short: "Buy ecocredits from a specific sell order",
		Long: "Purchase ecocredits from a specific sell order. DisableAutoRetire can be set to false to retire the credits immediately upon purchase." +
			"When set to true, credits will be received in a tradable state, IF AND ONLY IF the sell order also has auto retire disabled. " +
			"NOTE: The bid price is the price paid PER credit. The total cost will be quantity * bid_price.",
		Example: "regen tx ecocredit buy-direct 194 300 40regen true --retirement-jurisdiction=US-NY",
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
			var retireJurisdiction string
			retireJurisdiction, err = cmd.Flags().GetString(FlagRetirementJurisdiction)
			if err != nil {
				return err
			}

			msg := marketplace.MsgBuyDirect{
				Buyer: clientCtx.GetFromAddress().String(),
				Orders: []*marketplace.MsgBuyDirect_Order{
					{
						SellOrderId:            sellOrderId,
						Quantity:               qtyStr,
						BidPrice:               &bidPrice,
						DisableAutoRetire:      disableAutoRetire,
						RetirementJurisdiction: retireJurisdiction,
					},
				},
			}
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}
	cmd.Flags().String(FlagRetirementJurisdiction, "", "the jurisdiction to use for retirement when auto retire is true.")
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

// TxBuyDirectBatch returns a transaction command for a batch direct buy order using a json file.
func TxBuyDirectBatch() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "buy-direct-batch [name_of_file.json]",
		Short: "Buy ecocredits from multiple sell orders",
		Long: "Batch purchase ecocredits using a json file. DisableAutoRetire can be set to false to " +
			"retire the credits immediately upon purchase. When set to true, credits will be received in a tradable state, " +
			"IF AND ONLY IF the sell order also has auto retire disabled. NOTE: The bid price is the price paid PER credit. " +
			"The total cost will be quantity * bid_price.",
		Example: strings.TrimSpace(`regen tx ecocredit buy-direct-batch batch.json
		where batch.json has the following form:
		[
			{
			   "sell_order_id": 52,
			   "quantity": "32.5",
			   "bid_price": {"denom": "uregen", "amount": "32000000"},
			   "disable_auto_retire": false,
			   "retirement_jurisdiction": "US-NY"
			},
		]`),
		Args: cobra.ExactArgs(1),
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

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
