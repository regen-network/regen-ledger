package client

import (
	"fmt"
	"os"
	"strings"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"

	"github.com/regen-network/regen-ledger/x/intertx"
	intertxv1 "github.com/regen-network/regen-ledger/x/intertx/types/v1"
)

const (
	// FlagConnectionID is the connection end identifier on the controller chain
	FlagConnectionID = "connection-id"
	// FlagVersion is the controller chain channel version
	FlagVersion = "version"
)

// GetTxCmd creates and returns the intertx tx command
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        intertx.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", intertx.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		getRegisterAccountCmd(),
		getSubmitTxCmd(),
	)

	return cmd
}

func getRegisterAccountCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "register --connection-id <connection-id> --version <version>",
		Short: "register an interchain account",
		Long:  "register an interchain account for the chain corresponding to the connection-id.",
		Example: `
		regen tx intertx register --connection-id connection-0
		regen tx intertx register --connection-id connection-0 --version '{"version":"ics27-1","tx_type":"sdk_multi_msg","encoding":"proto3","host_connection_id":"connection-0","controller_connection_id":"connection-0","address":"regen14zs2x38lmkw4eqvl3lpml5l8crzaxn6mpvh79z"}'
		`,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := intertxv1.MsgRegisterAccount{
				Owner:        clientCtx.GetFromAddress().String(),
				ConnectionId: viper.GetString(FlagConnectionID),
				Version:      viper.GetString(FlagVersion),
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	cmd.Flags().String(FlagConnectionID, "", "the connection end identifier on the controller chain")
	cmd.Flags().String(FlagVersion, "", "the application version string")
	_ = cmd.MarkFlagRequired(FlagConnectionID)

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func getSubmitTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "submit-tx [connection-id] [msg.json]",
		Short: "submit a transaction to be sent to an ICA host chain",
		Long: strings.TrimSpace(`submit a transaction to be sent to the destination chain specified by the connection id.

The message in the JSON file MUST be an allowed message by the host chain's ica host 
module in order for it to succeed on the host chain.`),
		Example: `regen tx intertx submit-tx connection-0 msg.json

Example JSON:
{
  "@type": "/cosmos.bank.v1beta1.MsgSend",
  "from_address": "regen1yqr0pf38v9j7ah79wmkacau5mdspsc7l0sjeva",
  "to_address": "regen1df675r9vnf7pdedn4sf26svdsem3ugavgxmy46",
  "amount": [
    {
      "denom": "uregen",
      "amount": "1000000"
    }
  ]
}`,
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			cdc := codec.NewProtoCodec(clientCtx.InterfaceRegistry)

			var txMsg sdk.Msg
			if err := cdc.UnmarshalInterfaceJSON([]byte(args[1]), &txMsg); err != nil {

				// check for file path if JSON input is not provided
				contents, err := os.ReadFile(args[1])
				if err != nil {
					return errors.Wrap(err, "neither JSON input nor path to .json file for sdk msg were provided")
				}

				if err := cdc.UnmarshalInterfaceJSON(contents, &txMsg); err != nil {
					return errors.Wrap(err, "error unmarshalling sdk msg file")
				}
			}

			msg := intertxv1.NewMsgSubmitTx(clientCtx.GetFromAddress().String(), args[0], txMsg)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
