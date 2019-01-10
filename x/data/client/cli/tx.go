package cli

import (
	"github.com/spf13/cobra"
	utils2 "gitlab.com/regen-network/regen-ledger/utils"
	"gitlab.com/regen-network/regen-ledger/x/data"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/utils"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtxb "github.com/cosmos/cosmos-sdk/x/auth/client/txbuilder"
)

func GetCmdStoreData(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "store <data>",
		Short: "store some data on the blockchain",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc).WithAccountDecoder(cdc)

			txBldr := authtxb.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

			if err := cliCtx.EnsureAccountExists(); err != nil {
				return err
			}

			account, err := cliCtx.GetFromAddress()
			if err != nil {
				return err
			}

			msg := data.NewMsgStoreData([]byte(args[0]), account)
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			cliCtx.PrintResponse = true
			cliCtx.ResponsePrinter = utils2.PrintCLIResponse_Base64Data

			return utils.CompleteAndBroadcastTxCli(txBldr, cliCtx, []sdk.Msg{msg})
		},
	}
}
