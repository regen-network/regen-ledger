package cli

//import (
//	"github.com/cosmos/cosmos-sdk/client/context"
//	"github.com/cosmos/cosmos-sdk/client/utils"
//	"github.com/cosmos/cosmos-sdk/codec"
//	sdk "github.com/cosmos/cosmos-sdk/types"
//	authtxb "github.com/cosmos/cosmos-sdk/x/auth/client/txbuilder"
//	"github.com/regen-network/regen-ledger/x/schema"
//	"github.com/spf13/cobra"
//)

//func GetCmdDefineProperty(cdc *codec.Codec) *cobra.Command {
//	cmd := &cobra.Command{
//		Use:   "define-property",
//		Short: "define a schema property on the blockchain",
//		Args:  cobra.ExactArgs(1),
//		RunE: func(cmd *cobra.Command, args []string) error {
//			cliCtx := context.NewCLIContext().WithCodec(cdc).WithAccountDecoder(cdc)
//
//			txBldr := authtxb.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))
//
//			if err := cliCtx.EnsureAccountExists(); err != nil {
//				return err
//			}
//
//			account := cliCtx.GetFromAddress()
//
//			if err != nil {
//				return err
//			}
//
//			msg := schema.PropertyDefinition{
//			}
//			err = msg.ValidateBasic()
//			if err != nil {
//				return err
//			}
//
//			cliCtx.PrintResponse = true
//
//			return utils.CompleteAndBroadcastTxCLI(txBldr, cliCtx, []sdk.Msg{msg})
//		},
//	}
//	return cmd
//}
