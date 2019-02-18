package cli

import (
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/utils"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtxb "github.com/cosmos/cosmos-sdk/x/auth/client/txbuilder"
	"github.com/spf13/cobra"
	utils2 "gitlab.com/regen-network/regen-ledger/utils"
	"gitlab.com/regen-network/regen-ledger/x/group"
	"strings"
)

func membersFromArray(arr []string) []group.Member {
	n := len(arr)
	res := make([]group.Member, n)
	for i := 0; i < n; i++ {
		strs := strings.Split(arr[i], "=")
		if len(strs) <= 0 {
			panic("empty array")
		}
		acc, err := sdk.AccAddressFromBech32(strs[0])
		if err != nil {
			panic(err)
		}
		mem := group.Member{Address: acc}
		if len(strs) == 2 {
			mem.Weight.SetString(strs[1], 10)
		} else {
			mem.Weight.SetInt64(1)
		}
		res[i] = mem
	}
	return res
}

func GetCmdCreateGroup(cdc *codec.Codec) *cobra.Command {
	var threshold int64
	var members []string

	cmd := &cobra.Command{
		Use:   "create",
		Short: "create an group",
		//Args:  cobra.MinimumNArgs(1),
		PreRun: func(cmd *cobra.Command, args []string) {

		},
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc).WithAccountDecoder(cdc)

			txBldr := authtxb.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

			if err := cliCtx.EnsureAccountExists(); err != nil {
				return err
			}

			account := cliCtx.GetFromAddress()

			info := group.Group{
				Members: membersFromArray(members),
			}
			info.DecisionThreshold.SetInt64(threshold)

			msg := group.NewMsgCreateGroup(info, account)
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}

			cliCtx.PrintResponse = true
			cliCtx.ResponseHandler = utils2.PrintCLIResponse_StringData

			return utils.CompleteAndBroadcastTxCLI(txBldr, cliCtx, []sdk.Msg{msg})
		},
	}

	cmd.Flags().Int64Var(&threshold, "decision-threshold", 0, "Decision threshold")
	cmd.Flags().StringArrayVar(&members, "members", []string{}, "Members")

	return cmd
}
