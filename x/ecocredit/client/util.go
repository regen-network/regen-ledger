package client

import (
	sdkclient "github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/gogo/protobuf/proto"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/regen-network/regen-ledger/client"
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

type msgSrvClient struct {
	Cctx   *sdkclient.Context
	conn   *client.ServiceMsgClientConn
	client ecocredit.MsgClient
	flags  *pflag.FlagSet
}

func newMsgSrvClient(cmd *cobra.Command) (msgSrvClient, error) {
	f := cmd.Flags()
	clientCtx, err := sdkclient.GetClientTxContext(cmd)
	if err != nil {
		return msgSrvClient{}, err
	}
	conn := &client.ServiceMsgClientConn{}
	return msgSrvClient{
		&clientCtx, conn, ecocredit.NewMsgClient(conn), f,
	}, nil
}

// executes a MsgService transaction
func (c msgSrvClient) send(err error) error {
	if err != nil {
		return err
	}
	return tx.GenerateOrBroadcastTxCLI(*c.Cctx, c.flags, c.conn.Msgs...)
}
