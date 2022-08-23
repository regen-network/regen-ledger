package marketplace

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/regen-network/regen-ledger/x/ecocredit/marketplace"
)

const (
	FlagRetirementJurisdiction = "retirement-jurisdiction"
)

// TxSellCmd returns a transaction command that creates sell orders.
func TxSellCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sell [orders-json]",
		Short: "Creates new sell orders with transaction author (--from) as seller",
		Long: `Creates new sell orders with transaction author (--from) as seller.

Parameters:

- orders-json:  path to JSON file containing orders to create

Example JSON:

[
  {
    "batch_denom": "C01-001-20200101-20210101-001",
    "quantity": "5",
    "ask_price": {
      "denom": "uregen",
      "amount" "100000000"
	},
    "disable_auto_retire": "true"
  },
  {
    "batch_denom": "C01-001-20200101-20210101-002",
    "quantity": "10",
    "ask_price": {
      "denom": "uregen",
      "amount" "100000000"
	},
    "disable_auto_retire": false,
    "expiration": "2024-01-01T00:00:00Z"
  }
]`,
		Args:    cobra.ExactArgs(1),
		Example: "regen tx ecocredit sell orders.json",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			orders, err := parseSellOrders(args[0])
			if err != nil {
				return sdkerrors.ErrInvalidRequest.Wrapf("failed to parse json: %s", err)
			}

			// create sell message
			msg := marketplace.MsgSell{
				Seller: clientCtx.GetFromAddress().String(),
				Orders: orders,
			}

			// generate and broadcast transaction
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	return txFlags(cmd)
}

// TxUpdateSellOrdersCmd returns a transaction command that creates sell orders.
func TxUpdateSellOrdersCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-sell-orders [updates-json]",
		Short: "Updates existing sell orders with transaction author (--from) as seller",
		Long: `Updates existing sell orders with transaction author (--from) as seller.

Parameters:

- updates-json:  path to JSON file containing orders to update

Example JSON:

[
  {
    "sell_order_id": 1,
    "new_quantity": "5",
    "new_ask_price": {
      "denom": "uregen",
      "amount" "100000000"
	},
    "disable_auto_retire": true
  },
  {
    "sell_order_id": 2,
    "new_quantity": "10",
    "new_ask_price": {
      "denom": "uregen",
      "amount" "100000000"
	},
    "disable_auto_retire": false,
    "new_expiration": "2024-01-01T00:00:00Z"
  }
]`,
		Args:    cobra.ExactArgs(1),
		Example: "regen tx ecocredit update-sell-orders updates.json",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			updates, err := parseSellUpdates(args[0])
			if err != nil {
				return sdkerrors.ErrInvalidRequest.Wrapf("failed to parse json: %s", err)
			}

			// create update sell orders message
			msg := marketplace.MsgUpdateSellOrders{
				Seller:  clientCtx.GetFromAddress().String(),
				Updates: updates,
			}

			// generate and broadcast transaction
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	return txFlags(cmd)
}

// TxBuyDirectCmd returns a transaction command for a single direct buy order.
func TxBuyDirectCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "buy-direct [sell-order-id] [quantity] [bid-price] [disable-auto-retire] [flags]",
		Short: "Buy ecocredits from a specific sell order",
		Long: `Purchase ecocredits from a specific sell order.

DisableAutoRetire can be set to false to retire the credits immediately
upon purchase. When set to true, credits will be received in a tradable
state, IF AND ONLY IF the sell order also has auto retire disabled.

NOTE: The bid price is the price paid PER credit. The total cost will be quantity * bid_price.`,
		Example: `regen tx ecocredit buy-direct 1 300 10000000uregen true --retirement-jurisdiction "US-WA 98225"`,
		Args:    cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			sellOrderID, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			bidPrice, err := sdk.ParseCoinNormalized(args[2])
			if err != nil {
				return err
			}

			disableAutoRetire, err := strconv.ParseBool(args[3])
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
						SellOrderId:            sellOrderID,
						Quantity:               args[1],
						BidPrice:               &bidPrice,
						DisableAutoRetire:      disableAutoRetire,
						RetirementJurisdiction: retireJurisdiction,
					},
				},
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	cmd.Flags().String(FlagRetirementJurisdiction, "", "the jurisdiction to use for retirement when auto retire is true.")

	return txFlags(cmd)
}

// TxBuyDirectBulkCmd returns a transaction command for a batch direct buy order using a json file.
func TxBuyDirectBulkCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "buy-direct-bulk [orders-json]",
		Short: "Buy ecocredits from multiple sell orders",
		Long: `Purchase ecocredits from multiple sell orders.

DisableAutoRetire can be set to false to retire the credits immediately
upon purchase. When set to true, credits will be received in a tradable
state, IF AND ONLY IF the sell order also has auto retire disabled.

NOTE: The bid price is the price paid PER credit. The total cost will be quantity * bid_price.`,
		Example: `regen tx ecocredit buy-direct-bulk orders.json

Example JSON:
[
  {
    "sell_order_id": 1,
    "quantity": "32.5",
    "bid_price": {
      "denom": "uregen",
      "amount": "32000000"
    },
    "disable_auto_retire": false,
    "retirement_jurisdiction": "US-NY"
  },
  {
    "sell_order_id": 2,
    "quantity": "32.5",
    "bid_price": {
      "denom": "uregen",
      "amount": "32000000"
    },
    "disable_auto_retire": false,
    "retirement_jurisdiction": "US-NY"
  }
]`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			orders, err := parseBuyOrders(args[0])
			if err != nil {
				return sdkerrors.ErrInvalidRequest.Wrapf("failed to parse json: %s", err)
			}

			msg := marketplace.MsgBuyDirect{
				Buyer:  clientCtx.GetFromAddress().String(),
				Orders: orders,
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	return txFlags(cmd)
}

// TxCancelSellOrderCmd returns a transaction command that cancels sell order.
func TxCancelSellOrderCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "cancel-sell-order [order-id]",
		Short:   "Cancel existing sell orders with transaction author (--from) as seller",
		Long:    "Cancel existing sell orders with transaction author (--from) as seller",
		Example: "regen tx ecocredit cancel-sell-order 1",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return fmt.Errorf("invalid order id: %s", err)
			}

			msg := marketplace.MsgCancelSellOrder{
				Seller:      clientCtx.GetFromAddress().String(),
				SellOrderId: id,
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	return txFlags(cmd)
}
