package client

import (
	sdkclient "github.com/cosmos/cosmos-sdk/client"
	"github.com/gogo/protobuf/proto"
	"github.com/spf13/cobra"

	"github.com/regen-network/regen-ledger/x/ecocredit"
)

// prints a query client response
func print(cctx sdkclient.Context, res proto.Message, err error) error {
	if err != nil {
		return err
	}
	return cctx.PrintProto(res)
}

func mkQueryClient(cmd *cobra.Command) (ecocredit.QueryClient, sdkclient.Context, error) {
	ctx, err := sdkclient.GetClientQueryContext(cmd)
	if err != nil {
		return nil, sdkclient.Context{}, err
	}
	return ecocredit.NewQueryClient(ctx), ctx, err
}
