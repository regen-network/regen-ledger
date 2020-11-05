package client

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/gogo/protobuf/proto"
	"github.com/spf13/cobra"

	"github.com/regen-network/regen-ledger/x/ecocredit"
)

func print(ctx client.Context, res proto.Message, err error) error {
	if err != nil {
		return err
	}
	return ctx.PrintOutput(res)
}

func mkTx(ctx client.Context, cmd *cobra.Command, err error) error {
	if err != nil {
		return err
	}
	return tx.GenerateOrBroadcastTxCLI(ctx, cmd.Flags())
}

func mkQueryClient(cmd *cobra.Command) (ecocredit.QueryClient, client.Context, error) {
	ctx := client.GetClientContextFromCmd(cmd)
	ctx, err := client.ReadQueryCommandFlags(ctx, cmd.Flags())
	return ecocredit.NewQueryClient(ctx), ctx, err
}

func mkMsgClient(cmd *cobra.Command) (ecocredit.MsgClient, client.Context, error) {
	ctx := client.GetClientContextFromCmd(cmd)
	ctx, err := client.ReadTxCommandFlags(ctx, cmd.Flags())
	return ecocredit.NewMsgClient(ctx), ctx, err
}
