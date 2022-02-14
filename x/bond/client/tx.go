package client

import (
	"fmt"
	sdkclient "github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/spf13/cobra"
	"io/ioutil"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/regen-network/regen-ledger/v2/x/bond"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd(moduleName string) *cobra.Command {
	cmd := &cobra.Command{
		Use:                        moduleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", bond.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(TxIssueBondCmd())
	cmd.AddCommand(TxSellBondCmd())

	return cmd
}

// TxIssueBondCmd returns a transaction command that creates issues a bond.
func TxIssueBondCmd() *cobra.Command {
	return txflags(&cobra.Command{
		Use:   "issue-bond [msg-issue-bond-json-file]",
		Short: "Issues a new bond (--from) as admin",
		Long: fmt.Sprintf(
			`Creates a new bond with transaction author (--from) as holder.

Parameters:
  Parameters:
  msg-issue-bond-json-file: Path to a file containing a JSON object representing MsgIssueBond. The JSON has format:
			{
				"emission_denom": "emission-denom-01",
				"name": "Bond 01",
				"face_value": "15",
				"face_currency": "EUR",
				"issuance_date": "2022-02-21T00:00:00Z",
				"maturity_date": "2023-02-21T00:00:00Z",
				"coupon_rate": "8",
				"coupon_frequency": "quarterly",
				"project": "P01",
				"metadata": "Rm9v"
			}`),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := sdkclient.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			// Check input file is provided
			if args[0] == "" {
				return sdkerrors.ErrInvalidRequest.Wrap("Bond name is required")
			}

			msg, err := parseMsgIssueBond(clientCtx, args[0])
			if err != nil {
				return err
			}

			// Set holder as an author of the transaction
			msg.Holder = clientCtx.GetFromAddress().String()

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	})
}

// TxSellBondCmd returns a transaction command that sells a bond.
func TxSellBondCmd() *cobra.Command {
	return txflags(&cobra.Command{
		Use:   "sell-bond [bondId*] [amount*] [newOwner*]",
		Short: "Sell a bond that the (--from) is the owner",
		Long: fmt.Sprintf(
			`Changes the holder (--from) of the part of the bond to buyer.

They must also pay the fee associated with selling a bond, defined
by the %s parameter, so should make sure they have enough funds to cover that.

Parameters:
  bondId:    	       	bond identifier
  amount:    			the amount of the bond which will be selled to the new owner
  metadata:  	       	base64 encoded metadata`),
		Args: cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := sdkclient.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			if args[0] == "" {
				return sdkerrors.ErrInvalidRequest.Wrap("Bond ID is required")
			}
			bondId := args[0]

			if args[1] == "" {
				return sdkerrors.ErrInvalidRequest.Wrap("Bond amount is required")
			}
			amount := args[1]

			if args[2] == "" {
				return sdkerrors.ErrInvalidRequest.Wrap("Bond buyer is required")
			}
			buyer := args[2]

			// Get the class owner from the --from flag
			holder := clientCtx.GetFromAddress().String()

			msg := bond.MsgSellBond{
				BondId: bondId,
				Buyer:  buyer,
				Amount: amount,
				Holder: holder,
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	})
}

func txflags(cmd *cobra.Command) *cobra.Command {
	flags.AddTxFlagsToCmd(cmd)
	cmd.MarkFlagRequired(flags.FlagFrom)
	return cmd
}

func parseMsgIssueBond(clientCtx sdkclient.Context, batchFile string) (*bond.MsgIssueBond, error) {
	contents, err := ioutil.ReadFile(batchFile)
	if err != nil {
		return nil, err
	}

	var msg bond.MsgIssueBond
	err = clientCtx.Codec.UnmarshalJSON(contents, &msg)
	if err != nil {
		return nil, err
	}

	return &msg, nil
}
