package client

import (
	"bytes"
	"encoding/json"
	"io/ioutil"

	"github.com/gogo/protobuf/proto"
	"github.com/spf13/cobra"

	sdkclient "github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"

	"github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

func txFlags(cmd *cobra.Command) *cobra.Command {
	flags.AddTxFlagsToCmd(cmd)
	cmd.MarkFlagRequired(flags.FlagFrom)
	return cmd
}

func qflags(cmd *cobra.Command) *cobra.Command {
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

func printQueryResponse(clientCtx sdkclient.Context, res proto.Message, err error) error {
	if err != nil {
		return err
	}
	return clientCtx.PrintProto(res)
}

func mkQueryClient(cmd *cobra.Command) (core.QueryClient, sdkclient.Context, error) {
	ctx, err := sdkclient.GetClientQueryContext(cmd)
	if err != nil {
		return nil, sdkclient.Context{}, err
	}
	return core.NewQueryClient(ctx), ctx, err
}

func parseMsgCreateBatch(clientCtx sdkclient.Context, jsonFile string) (*core.MsgCreateBatch, error) {
	bz, err := ioutil.ReadFile(jsonFile)
	if err != nil {
		return nil, err
	}

	if err := types.CheckDuplicateKey(json.NewDecoder(bytes.NewReader(bz)), nil); err != nil {
		return nil, err
	}

	var msg core.MsgCreateBatch
	err = clientCtx.Codec.UnmarshalJSON(bz, &msg)
	if err != nil {
		return nil, err
	}

	return &msg, nil
}

func parseCredits(jsonFile string) ([]*core.Credits, error) {
	bz, err := ioutil.ReadFile(jsonFile)
	if err != nil {
		return nil, err
	}

	if err := types.CheckDuplicateKey(json.NewDecoder(bytes.NewReader(bz)), nil); err != nil {
		return nil, err
	}

	var credits []*core.Credits

	// using json package because array is not a proto message
	err = json.Unmarshal(bz, &credits)
	if err != nil {
		return nil, err
	}

	return credits, nil
}

func parseSendCredits(jsonFile string) ([]*core.MsgSend_SendCredits, error) {
	bz, err := ioutil.ReadFile(jsonFile)
	if err != nil {
		return nil, err
	}

	if err := types.CheckDuplicateKey(json.NewDecoder(bytes.NewReader(bz)), nil); err != nil {
		return nil, err
	}

	var sendCredits []*core.MsgSend_SendCredits

	// using json package because array is not a proto message
	err = json.Unmarshal(bz, &sendCredits)
	if err != nil {
		return nil, err
	}

	return sendCredits, nil
}
