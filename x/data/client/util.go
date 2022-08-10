package client

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/gogo/protobuf/proto"
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"

	"github.com/regen-network/regen-ledger/x/data"
)

func formatExample(str string) string {
	str = strings.TrimPrefix(str, "\n")
	str = strings.TrimRight(str, "\t")
	return strings.TrimSuffix(str, "\n")
}

func printQueryResponse(ctx client.Context, res proto.Message, err error) error {
	if err != nil {
		return err
	}
	return ctx.PrintProto(res)
}

func mkQueryClient(cmd *cobra.Command) (data.QueryClient, client.Context, error) {
	ctx, err := client.GetClientQueryContext(cmd)
	if err != nil {
		return nil, client.Context{}, err
	}
	return data.NewQueryClient(ctx), ctx, err
}

func parseContentHash(clientCtx client.Context, filePath string) (*data.ContentHash, error) {
	contentHash := data.ContentHash{}

	if filePath == "" {
		return nil, fmt.Errorf("file path is empty")
	}

	bz, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	if err := clientCtx.Codec.UnmarshalJSON(bz, &contentHash); err != nil {
		return nil, err
	}

	return &contentHash, nil
}
