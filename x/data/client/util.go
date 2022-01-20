package client

import (
	"github.com/spf13/cobra"

	sdkclient "github.com/cosmos/cosmos-sdk/client"
	"github.com/gogo/protobuf/proto"

	"github.com/regen-network/regen-ledger/x/data"
)

func print(cctx sdkclient.Context, res proto.Message, err error) error {
	if err != nil {
		return err
	}
	return cctx.PrintProto(res)
}

func mkQueryClient(cmd *cobra.Command) (data.QueryClient, sdkclient.Context, error) {
	ctx, err := sdkclient.GetClientQueryContext(cmd)
	if err != nil {
		return nil, sdkclient.Context{}, err
	}
	return data.NewQueryClient(ctx), ctx, err
}
