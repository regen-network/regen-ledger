package cli

import (
	"github.com/regen-network/regen-ledger/x/proposal"
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/utils"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtxb "github.com/cosmos/cosmos-sdk/x/auth/client/txbuilder"
)

type ActionCreator func(cmd *cobra.Command, args []string) (proposal.ProposalAction, error)

func GetCmdPropose(cdc *codec.Codec, actionCreator ActionCreator) *cobra.Command {
	var exec bool

	cmd := &cobra.Command{
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc).WithAccountDecoder(cdc)

			txBldr := authtxb.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

			if err := cliCtx.EnsureAccountExists(); err != nil {
				return err
			}

			account := cliCtx.GetFromAddress()

			action, err := actionCreator(cmd, args)

			if err != nil {
				return err
			}

			msg := proposal.MsgCreateProposal{
				Proposer: account,
				Action:   action,
				Exec:     exec,
			}
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			cliCtx.PrintResponse = true

			return utils.CompleteAndBroadcastTxCLI(txBldr, cliCtx, []sdk.Msg{msg})
		},
	}
	cmd.Flags().BoolVar(&exec, "exec", false, "try to execute the proposal immediately")
	return cmd
}

func getRunVote(cdc *codec.Codec, approve bool) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		cliCtx := context.NewCLIContext().WithCodec(cdc).WithAccountDecoder(cdc)

		txBldr := authtxb.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

		if err := cliCtx.EnsureAccountExists(); err != nil {
			return err
		}

		account := cliCtx.GetFromAddress()

		id := proposal.MustDecodeProposalIDBech32(args[0])

		msg := proposal.MsgVote{
			ProposalId: id,
			Voter:      account,
			Vote:       approve,
		}
		err := msg.ValidateBasic()
		if err != nil {
			return err
		}

		cliCtx.PrintResponse = true

		return utils.CompleteAndBroadcastTxCLI(txBldr, cliCtx, []sdk.Msg{msg})
	}
}

func GetCmdApprove(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "approve [ID]",
		Short: "vote to approve a proposal",
		Args:  cobra.ExactArgs(1),
		RunE:  getRunVote(cdc, true),
	}
}

func GetCmdUnapprove(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "unapprove [ID]",
		Short: "vote to un-approve a proposal that you have previously approved",
		Args:  cobra.ExactArgs(1),
		RunE:  getRunVote(cdc, false),
	}
}

func GetCmdTryExec(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "try-exec [ID]",
		Short: "try to execute the proposal (will fail if not enough signers have approved it)",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc).WithAccountDecoder(cdc)

			txBldr := authtxb.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

			if err := cliCtx.EnsureAccountExists(); err != nil {
				return err
			}

			account := cliCtx.GetFromAddress()

			id := proposal.MustDecodeProposalIDBech32(args[0])

			msg := proposal.MsgTryExecuteProposal{
				ProposalId: id,
				Signer:     account,
			}
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}

			cliCtx.PrintResponse = true

			return utils.CompleteAndBroadcastTxCLI(txBldr, cliCtx, []sdk.Msg{msg})
		},
	}
}

func GetCmdWithdraw(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "withdraw [ID]",
		Short: "withdraw a proposer that you previously proposed",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc).WithAccountDecoder(cdc)

			txBldr := authtxb.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

			if err := cliCtx.EnsureAccountExists(); err != nil {
				return err
			}

			account := cliCtx.GetFromAddress()

			id := proposal.MustDecodeProposalIDBech32(args[0])

			msg := proposal.MsgWithdrawProposal{
				ProposalId: id,
				Proposer:   account,
			}
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}

			cliCtx.PrintResponse = true

			return utils.CompleteAndBroadcastTxCLI(txBldr, cliCtx, []sdk.Msg{msg})
		},
	}
}
