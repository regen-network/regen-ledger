package cli

import (
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/utils"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtxb "github.com/cosmos/cosmos-sdk/x/auth/client/txbuilder"
	"github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/x/claim"
	"github.com/spf13/cobra"
)

// GetCmdSignClaim returns the tx claim sign command.
func GetCmdSignClaim(cdc *codec.Codec) *cobra.Command {
	var evidence []string
	cmd := &cobra.Command{
		Use:   "sign <content-address> [--evidence <evidence-addresses>] --from <signer>",
		Short: "sign a claim on the blockchain",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc).WithAccountDecoder(cdc)

			txBldr := authtxb.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

			if err := cliCtx.EnsureAccountExists(); err != nil {
				return err
			}

			account := cliCtx.GetFromAddress()

			contentAddr, err := types.DecodeBech32DataAddress(args[0])
			if err != nil {
				return err
			}

			evidenceAddrs := make([]types.DataAddress, len(evidence))
			for i, bech := range evidence {
				evidenceAddrs[i], err = types.DecodeBech32DataAddress(bech)
				if err != nil {
					return err
				}
			}

			msg := claim.MsgSignClaim{
				Content:  contentAddr,
				Signers:  []sdk.AccAddress{account},
				Evidence: evidenceAddrs,
			}
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			cliCtx.PrintResponse = true

			return utils.CompleteAndBroadcastTxCLI(txBldr, cliCtx, []sdk.Msg{msg})
		},
	}
	cmd.Flags().StringSliceVar(&evidence, "evidence", nil, "A comma-separated list of data addresses representing claim evidence")
	return cmd
}
