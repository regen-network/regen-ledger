package cli

import (
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/regen-network/regen-ledger/types"
	"github.com/spf13/cobra"
)

// GetCmdGetData returns a command
func GetCmdGetData(storeName string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "get [id]",
		Short: "get data by id",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			bech := args[0]

			addr, err := types.DecodeBech32DataAddress(bech)
			if err != nil {
				return err
			}

			_, err = cliCtx.QueryStore([]byte(addr), storeName)
			if err != nil {
				return err
			}

			panic("TODO")
			//graph, err := binary.DeserializeGraph(resolver, res)
			//
			//return cliCtx.PrintOutput(graph.String())
		},
	}
}
