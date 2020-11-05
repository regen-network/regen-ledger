package client

import (
	"context"
	"encoding/base64"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"

	"github.com/regen-network/regen-ledger/util"
	"github.com/regen-network/regen-ledger/x/data"
)

// TxCmd returns a root CLI command handler for all x/data transaction commands.
func TxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        data.ModuleName,
		Short:                      "Data transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		MsgAnchorDataCmd(),
	)

	return cmd
}

// MsgAnchorDataCmd created a CLI tx command for MsgAnchorData.
func MsgAnchorDataCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "anchor [cid]",
		Short: `Anchors a piece of data to the blockchain based on its secure
		hash, effectively providing a tamper resistant timestamp.`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadTxCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			cidBz, err := base64.StdEncoding.DecodeString(args[0])
			if err != nil {
				return err
			}

			msg := data.MsgAnchorDataRequest{
				Sender: clientCtx.GetFromAddress().String(),
				Cid:    cidBz,
			}
			svcMsgClientConn := &util.ServiceMsgClientConn{}
			msgClient := data.NewMsgClient(svcMsgClientConn)
			_, err = msgClient.AnchorData(context.Background(), &msg)
			if err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), svcMsgClientConn.Msgs...)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
