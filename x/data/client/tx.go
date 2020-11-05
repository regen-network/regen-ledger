package client

import (
	"context"
	"encoding/base64"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	gocid "github.com/ipfs/go-cid"
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

// MsgAnchorDataCmd creates a CLI command for Msg/AnchorData.
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

			cid, err := gocid.Decode(args[0])
			if err != nil {
				return err
			}

			msg := data.MsgAnchorDataRequest{
				Sender: clientCtx.GetFromAddress().String(),
				Cid:    cid.Bytes(),
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

// MsgSignDataCmd creates a CLI command for Msg/SignData.
func MsgSignDataCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sign [cid]",
		Short: `Sign an arbitrary piece of data on the blockchain.`,
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadTxCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			cid, err := gocid.Decode(args[0])
			if err != nil {
				return err
			}

			msg := data.MsgSignDataRequest{
				Signers: []string{clientCtx.GetFromAddress().String()},
				Cid:     cid.Bytes(),
			}
			svcMsgClientConn := &util.ServiceMsgClientConn{}
			msgClient := data.NewMsgClient(svcMsgClientConn)
			_, err = msgClient.SignData(context.Background(), &msg)
			if err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), svcMsgClientConn.Msgs...)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// MsgStoreDataCmd creates a CLI command for Msg/StoreData.
func MsgStoreDataCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "store [cid] [content-as-base64]",
		Short: `Store a piece of data corresponding to a CID on the blockchain.`,
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadTxCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			cid, err := gocid.Decode(args[0])
			if err != nil {
				return err
			}

			content, err := base64.StdEncoding.DecodeString(args[1])
			if err != nil {
				return err
			}

			msg := data.MsgStoreDataRequest{
				Sender:  clientCtx.GetFromAddress().String(),
				Cid:     cid.Bytes(),
				Content: content,
			}
			svcMsgClientConn := &util.ServiceMsgClientConn{}
			msgClient := data.NewMsgClient(svcMsgClientConn)
			_, err = msgClient.StoreData(context.Background(), &msg)
			if err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), svcMsgClientConn.Msgs...)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
