package client

import (
	"fmt"

	sdkclient "github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

// TxCmd returns a root CLI command handler for all x/data transaction commands.
func TxCmd(name string) *cobra.Command {
	cmd := &cobra.Command{
		Use:                        name,
		Short:                      "Data transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       sdkclient.ValidateCmd,
	}

	cmd.AddCommand(
		MsgAnchorDataCmd(),
		MsgSignDataCmd(),
	)

	return cmd
}

// MsgAnchorDataCmd creates a CLI command for Msg/AnchorData.
func MsgAnchorDataCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "anchor [sender] [cid]",
		Short: "Anchors a piece of data to the blockchain based on its secure " +
			"hash, effectively providing a tamper resistant timestamp.",
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			return fmt.Errorf("not implemented")
			//err := cmd.Flags().Set(flags.FlagFrom, args[0])
			//if err != nil {
			//	return err
			//}
			//
			//clientCtx, err := sdkclient.GetClientTxContext(cmd)
			//if err != nil {
			//	return err
			//}
			//
			//cid, err := gocid.Decode(args[1])
			//if err != nil {
			//	return err
			//}
			//
			//msg := data.MsgAnchorDataRequest{
			//	Sender: clientCtx.GetFromAddress().String(),
			//	Cid:    cid.Bytes(),
			//}
			//svcMsgClientConn := &client.ServiceMsgClientConn{}
			//msgClient := data.NewMsgClient(svcMsgClientConn)
			//_, err = msgClient.AnchorData(cmd.Context(), &msg)
			//if err != nil {
			//	return err
			//}
			//
			//return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), svcMsgClientConn.Msgs...)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// MsgSignDataCmd creates a CLI command for Msg/SignData.
func MsgSignDataCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sign [signer] [cid]",
		Short: `Sign an arbitrary piece of data on the blockchain.`,
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			return fmt.Errorf("not implemented")
			//err := cmd.Flags().Set(flags.FlagFrom, args[0])
			//if err != nil {
			//	return err
			//}
			//
			//clientCtx, err := sdkclient.GetClientTxContext(cmd)
			//if err != nil {
			//	return err
			//}
			//
			//cid, err := gocid.Decode(args[1])
			//if err != nil {
			//	return err
			//}
			//
			//msg := data.MsgSignDataRequest{
			//	Signers: []string{clientCtx.GetFromAddress().String()},
			//	Cid:     cid.Bytes(),
			//}
			//svcMsgClientConn := &client.ServiceMsgClientConn{}
			//msgClient := data.NewMsgClient(svcMsgClientConn)
			//_, err = msgClient.SignData(cmd.Context(), &msg)
			//if err != nil {
			//	return err
			//}
			//
			//return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), svcMsgClientConn.Msgs...)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
