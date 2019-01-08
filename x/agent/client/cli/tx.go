package cli

import (
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/utils"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtxb "github.com/cosmos/cosmos-sdk/x/auth/client/txbuilder"
	"github.com/spf13/cobra"
	"gitlab.com/regen-network/regen-ledger/x/agent"
	"strconv"
)

func addrsFromBech32Array(arr []string) []sdk.AccAddress {
	n := len(arr)
	res := make([]sdk.AccAddress, n)
	for i := 0; i < n; i++ {
		str := arr[i]
		acc, err := sdk.AccAddressFromBech32(str)
		if err != nil {
			panic(err)
		}
		res[i] = acc
	}
	return res
}

func AgentsFromArray(arr []string) []agent.AgentID {
	n := len(arr)
	res := make([]agent.AgentID, n)
	for i := 0; i < n; i++ {
		str := arr[i]
		id, err := strconv.ParseUint(str, 10, 64)
		if err != nil {
			panic(err)
		}
		res[i] = agent.AgentID(id)
	}
	return res
}

func GetCmdCreateAgent(cdc *codec.Codec) *cobra.Command {
	var threshold int
	var addrs []string
	var agents []string

	cmd := &cobra.Command{
		Use:   "create",
		Short: "create an agent",
		//Args:  cobra.MinimumNArgs(1),
		PreRun: func(cmd *cobra.Command, args []string) {

		},
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

			info := agent.AgentInfo{
				AuthPolicy:        agent.MultiSig,
				MultisigThreshold: threshold,
				Addresses:         addrsFromBech32Array(addrs),
				Agents:            AgentsFromArray(agents),
			}

			msg := agent.NewMsgCreateAgent(info, account)
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			cliCtx.PrintResponse = true

			return utils.CompleteAndBroadcastTxCli(txBldr, cliCtx, []sdk.Msg{msg})
		},
	}

	cmd.Flags().IntVar(&threshold, "threshold", 0, "Multisig threshold")
	cmd.Flags().StringArrayVar(&addrs, "addrs", []string{}, "Address")
	cmd.Flags().StringArrayVar(&agents, "agents", []string{}, "Agents")

	return cmd
}
