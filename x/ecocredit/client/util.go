package client

import (
	"github.com/cosmos/cosmos-sdk/client"
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

func mkQueryClient(cmd *cobra.Command) (ecocredit.QueryClient, client.Context, error) {
	ctx := client.GetClientContextFromCmd(cmd)
	ctx, err := client.ReadQueryCommandFlags(ctx, cmd.Flags())
	return ecocredit.NewQueryClient(ctx), ctx, err
}
