package marketplace

import (
	"strconv"

	"github.com/spf13/cobra"
	"sigs.k8s.io/yaml"

	sdkclient "github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/regen-network/regen-ledger/x/ecocredit/marketplace"
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
