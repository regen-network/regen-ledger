package util

// import (
// 	"encoding/base64"
// 	"encoding/json"
// 	"fmt"
// 	"github.com/tendermint/tendermint/rpc/core/types"
// )

// type KVPair struct {
// 	Key   string `json:"key,omitempty"`
// 	Value string `json:"value,omitempty"`
// }

// type CLIResponse struct {
// 	Code      uint32   `json:"code,omitempty"`
// 	Codespace string   `json:"codespace,omitempty"`
// 	Data      string   `json:"data,omitempty"`
// 	Log       string   `json:"log,omitempty"`
// 	Info      string   `json:"info,omitempty"`
// 	GasWanted int64    `json:"gas_wanted,omitempty"`
// 	GasUsed   int64    `json:"gas_used,omitempty"`
// 	Tags      []KVPair `json:"tags,omitempty"`
// }

// func PrintCLIResponse_Base64Data(res *core_types.ResultBroadcastTxCommit) (*core_types.ResultBroadcastTxCommit, error) {
// 	return PrintCLIResponse(res, func(data []byte) string {
// 		return base64.URLEncoding.EncodeToString(data)
// 	})
// }

// func PrintCLIResponse_StringData(res *core_types.ResultBroadcastTxCommit) (*core_types.ResultBroadcastTxCommit, error) {
// 	return PrintCLIResponse(res, func(data []byte) string { return string(data) })
// }

// func PrintCLIResponse(res *core_types.ResultBroadcastTxCommit, dataEncoder func(data []byte) string) (*core_types.ResultBroadcastTxCommit, error) {
// 	tags := make([]KVPair, len(res.DeliverTx.Tags))
// 	for i, tag := range res.DeliverTx.Tags {
// 		tags[i] = KVPair{Key: string(tag.Key), Value: string(tag.Value)}
// 	}
// 	cliRes := CLIResponse{
// 		Code:      res.DeliverTx.Code,
// 		Codespace: res.DeliverTx.Codespace,
// 		Log:       res.DeliverTx.Log,
// 		Info:      res.DeliverTx.Info,
// 		GasWanted: res.DeliverTx.GasWanted,
// 		GasUsed:   res.DeliverTx.GasUsed,
// 		Data:      dataEncoder(res.DeliverTx.Data),
// 		Tags:      tags,
// 	}
// 	b, err := json.MarshalIndent(cliRes, "", "  ")
// 	if err != nil {
// 		fmt.Print("error: ", err)
// 	} else {
// 		fmt.Printf("%s\n", b)
// 	}
// 	return res, nil
// }
