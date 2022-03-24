package marketplace

import (
	"strconv"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"

	sdkclient "github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/regen-network/regen-ledger/x/ecocredit/client/utils"
	"github.com/regen-network/regen-ledger/x/ecocredit/marketplace"
)

// TxBuyCmd returns a transaction command that creates sell orders.
func TxBuyCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "buy [orders]",
		Short: "Creates new buy orders with transaction author (--from) as buyer",
		Long: `Creates new buy orders with transaction author (--from) as buyer.

Parameters:
  orders:  YAML encoded order list. Note: numerical values must be written in strings.
           eg: '[{sell_order_id: "1", quantity: "5", bid_price: "100regen", disable_auto_retire: false}]'
           eg: '[{sell_order_id: "1", quantity: "5", bid_price: "100regen", disable_auto_retire: false, expiration: "2024-01-01"}]'`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := sdkclient.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			// get the order buyer from the --from flag
			buyer := clientCtx.GetFromAddress()

			// declare orders array with bid price as string
			var strOrders []struct {
				SellOrderId       string `json:"sell_order_id"`
				Quantity          string `json:"quantity"`
				BidPrice          string `json:"bid_price"`
				DisableAutoRetire bool   `json:"disable_auto_retire"`
				Expiration        string `json:"expiration"`
			}

			// unmarshal YAML encoded orders with new bid price as string
			if err := yaml.Unmarshal([]byte(args[0]), &strOrders); err != nil {
				return err
			}

			// declare orders array with new bid price as sdk.Coin
			orders := make([]*marketplace.MsgBuy_Order, len(strOrders))

			// loop through orders with new bid price as string
			for i, order := range strOrders {

				// parse sell order id
				sellOrderId, err := strconv.ParseUint(order.SellOrderId, 10, 64)
				if err != nil {
					return ecocredit.ErrInvalidSellOrder.Wrap(err.Error())
				}

				// set sell order id as buy order selection
				selection := &marketplace.MsgBuy_Order_Selection{
					Sum: &marketplace.MsgBuy_Order_Selection_SellOrderId{SellOrderId: sellOrderId},
				}

				// parse and normalize new bid price as sdk.Coin
				bidPrice, err := sdk.ParseCoinNormalized(order.BidPrice)
				if err != nil {
					return err
				}

				// set order with new bid price as sdk.Coin
				orders[i] = &marketplace.MsgBuy_Order{
					Selection:         selection,
					BidPrice:          &bidPrice,
					Quantity:          order.Quantity,
					DisableAutoRetire: order.DisableAutoRetire,
				}

				// parse and set expiration
				if order.Expiration != "" {
					expiration, err := utils.ParseDate("expiration", order.Expiration)
					if err != nil {
						return err
					}
					orders[i].Expiration = &expiration
				}
			}

			// create buy message
			msg := marketplace.MsgBuy{
				Buyer:  buyer.String(),
				Orders: orders,
			}

			// generate and broadcast transaction
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

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
			clientCtx, err := sdkclient.GetClientTxContext(cmd)
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
					if err := utils.ParseAndSetDate(&orders[i].Expiration, "expiration", o.Expiration); err != nil {
						return err
					}
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
			clientCtx, err := sdkclient.GetClientTxContext(cmd)
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
					if err := utils.ParseAndSetDate(&updates[i].NewExpiration, "expiration", u.NewExpiration); err != nil {
						return err
					}
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

// TxAllowAskDenomCmd returns a transaction command which authorizes a new ask denom.
func TxAllowAskDenomCmd() *cobra.Command {
	// TODO: implement
	return &cobra.Command{}
}
