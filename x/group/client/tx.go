package client

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/version"
	authclient "github.com/cosmos/cosmos-sdk/x/auth/client"
	"github.com/regen-network/regen-ledger/x/group"
	"github.com/spf13/cobra"
)

const flagMembers = "members"

// NewTxCmd returns a root CLI command handler for all x/bank transaction commands.
func NewTxCmd() *cobra.Command {
	txCmd := &cobra.Command{
		Use:                        group.ModuleName,
		Short:                      "Group transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	txCmd.AddCommand(
		MsgCreateGroupCmd(),
		MsgUpdateGroupAdminCmd(),
		MsgUpdateGroupCommentCmd(),
		MsgUpdateGroupMembersCmd(),
		MsgCreateGroupAccountCmd(),
		MsgUpdateGroupAccountAdminCmd(),
		MsgUpdateGroupAccountDecisionPolicyCmd(),
		MsgUpdateGroupAccountCommentCmd(),
		MsgCreateProposalCmd(),
		MsgVoteCmd(),
		MsgExecCmd(),
	)

	return txCmd
}

// MsgCreateGroupCmd creates a CLI command for Msg/CreateGroup.
func MsgCreateGroupCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "create-group [admin] [comment]",
		Short: "Create a group which is an aggregation " +
			"of member accounts with associated weights and " +
			"an administrator account. Note, the '--from' flag is " +
			"ignored as it is implied from [admin].",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Create a group which is an aggregation of member accounts with associated weights and
an administrator account. Note, the '--from' flag is ignored as it is implied from [admin].
Members accounts can be given through a members JSON file that contains an array of members.

Example:
$ %s tx group create-group [admin] [comment] --members="path/to/members.json"

Where members.json contains:

[
	{
		"address": "addr1",
		"power": "1",
		"comment": "some comment"
	},
	{
		"address": "addr2",
		"power": "1",
		"comment": "some comment"
	}
]
`,
				version.AppName,
			),
		),
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			err := cmd.Flags().Set(flags.FlagFrom, args[0])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			members, err := parseMembersFlag(cmd.Flags())
			if err != nil {
				return err
			}

			msg := &group.MsgCreateGroupRequest{
				Admin:   clientCtx.GetFromAddress().String(),
				Members: members,
				Comment: args[1],
			}
			if err = msg.ValidateBasic(); err != nil {
				return fmt.Errorf("message validation failed: %w", err)
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().String(flagMembers, "", "Members file path")
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// MsgUpdateGroupMembersCmd creates a CLI command for Msg/UpdateGroupMembers.
func MsgUpdateGroupMembersCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-group-members [admin] [group-id]",
		Short: "Update a group's members. Set a member's weight to \"0\" to delete it.",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Update a group's members

Example:
$ %s tx group update-group-members [admin] [group-id] --members="path/to/members.json"

Where members.json contains:

[
	{
		"address": "addr1",
		"power": "1",
		"comment": "some new comment"
	},
	{
		"address": "addr2",
		"power": "0",
		"comment": "some comment"
	}
]

Set a member's weight to "0" to delete it.
`,
				version.AppName,
			),
		),
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			err := cmd.Flags().Set(flags.FlagFrom, args[0])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			members, err := parseMembersFlag(cmd.Flags())
			if err != nil {
				return err
			}

			groupID, err := strconv.ParseUint(args[1], 10, 64)
			if err != nil {
				return err
			}

			msg := &group.MsgUpdateGroupMembersRequest{
				Admin:         clientCtx.GetFromAddress().String(),
				MemberUpdates: members,
				GroupId:       group.ID(groupID),
			}
			if err = msg.ValidateBasic(); err != nil {
				return fmt.Errorf("message validation failed: %w", err)
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().String(flagMembers, "", "Members file path")
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// MsgUpdateGroupAdminCmd creates a CLI command for Msg/UpdateGroupAdmin.
func MsgUpdateGroupAdminCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-group-admin [admin] [group-id] [new-admin]",
		Short: "Update a group's admin",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			err := cmd.Flags().Set(flags.FlagFrom, args[0])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			groupID, err := strconv.ParseUint(args[1], 10, 64)
			if err != nil {
				return err
			}

			msg := &group.MsgUpdateGroupAdminRequest{
				Admin:    clientCtx.GetFromAddress().String(),
				NewAdmin: args[2],
				GroupId:  group.ID(groupID),
			}
			if err = msg.ValidateBasic(); err != nil {
				return fmt.Errorf("message validation failed: %w", err)
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// MsgUpdateGroupCommentCmd creates a CLI command for Msg/UpdateGroupComment.
func MsgUpdateGroupCommentCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-group-admin [admin] [group-id] [comment]",
		Short: "Update a group's admin",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			err := cmd.Flags().Set(flags.FlagFrom, args[0])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			groupID, err := strconv.ParseUint(args[1], 10, 64)
			if err != nil {
				return err
			}

			msg := &group.MsgUpdateGroupCommentRequest{
				Admin:   clientCtx.GetFromAddress().String(),
				Comment: args[2],
				GroupId: group.ID(groupID),
			}
			if err = msg.ValidateBasic(); err != nil {
				return fmt.Errorf("message validation failed: %w", err)
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// MsgCreateGroupAccountCmd creates a CLI command for Msg/CreateGroupAccount.
func MsgCreateGroupAccountCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "create-group-account [admin] [group-id] [comment] [decision-policy]",
		Short: "Create a group account which is an account " +
			"associated with a group and a decision policy. " +
			"Note, the '--from' flag is " +
			"ignored as it is implied from [admin].",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Create a group account which is an account associated with a group and a decision policy.
Note, the '--from' flag is ignored as it is implied from [admin].

Example:
$ %s tx group create-group-account [admin] [group-id] [comment] \
'{"@type":"/regen.group.v1alpha1.ThresholdDecisionPolicy", "threshold":"1", "timeout":"1s"}'

Where decision-policy.json contains:

{}
`,
				version.AppName,
			),
		),
		Args: cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			err := cmd.Flags().Set(flags.FlagFrom, args[0])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			groupID, err := strconv.ParseUint(args[1], 10, 64)
			if err != nil {
				return err
			}

			var policy group.DecisionPolicy
			if err := clientCtx.JSONMarshaler.UnmarshalInterfaceJSON([]byte(args[3]), &policy); err != nil {
				return err
			}

			msg, err := group.NewMsgCreateGroupAccountRequest(
				clientCtx.GetFromAddress(),
				group.ID(groupID),
				args[2],
				policy,
			)
			if err != nil {
				return err
			}
			if err = msg.ValidateBasic(); err != nil {
				return fmt.Errorf("message validation failed: %w", err)
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// MsgUpdateGroupAccountAdminCmd creates a CLI command for Msg/UpdateGroupAccountAdmin.
func MsgUpdateGroupAccountAdminCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-group-account-admin [admin] [group-account] [new-admin]",
		Short: "Update a group account admin",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			err := cmd.Flags().Set(flags.FlagFrom, args[0])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := &group.MsgUpdateGroupAccountAdminRequest{
				Admin:        clientCtx.GetFromAddress().String(),
				GroupAccount: args[1],
				NewAdmin:     args[2],
			}
			if err = msg.ValidateBasic(); err != nil {
				return fmt.Errorf("message validation failed: %w", err)
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// MsgUpdateGroupAccountDecisionPolicyCmd creates a CLI command for Msg/UpdateGroupAccountDecisionPolicy.
func MsgUpdateGroupAccountDecisionPolicyCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-group-account-policy [admin] [group-account] [decision-policy]",
		Short: "Update a group account decision policy",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			err := cmd.Flags().Set(flags.FlagFrom, args[0])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			var policy group.DecisionPolicy
			if err := clientCtx.JSONMarshaler.UnmarshalInterfaceJSON([]byte(args[3]), &policy); err != nil {
				return err
			}

			msg, err := group.NewMsgUpdateGroupAccountDecisionPolicyRequest(
				clientCtx.GetFromAddress(),
				args[1],
				policy,
			)
			if err != nil {
				return err
			}

			if err = msg.ValidateBasic(); err != nil {
				return fmt.Errorf("message validation failed: %w", err)
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// MsgUpdateGroupAccountCommentCmd creates a CLI command for Msg/MsgUpdateGroupAccountComment.
func MsgUpdateGroupAccountCommentCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-group-account-comment [admin] [group-account] [new-comment]",
		Short: "Update a group account comment",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			err := cmd.Flags().Set(flags.FlagFrom, args[0])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := &group.MsgUpdateGroupAccountCommentRequest{
				Admin:        clientCtx.GetFromAddress().String(),
				GroupAccount: args[1],
				Comment:      args[2],
			}
			if err = msg.ValidateBasic(); err != nil {
				return fmt.Errorf("message validation failed: %w", err)
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// MsgCreateProposalCmd creates a CLI command for Msg/CreateProposal.
func MsgCreateProposalCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-proposal [group-account] [proposer[,proposer]*] [msg_tx_json_file] [comment]",
		Short: "Submit a new proposal",
		Long: `Submit a new proposal.

Parameters:
			group-account: address of the group account
			proposer: comma separated (no spaces) list of proposer account addresses. Example: "addr1,addr2" 
			comment: comment for the proposal
			msg_tx_json_file: path to json file with messages that will be executed if the proposal is accepted.
`,
		Args: cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			proposers := strings.Split(args[1], ",")
			for i := range proposers {
				proposers[i] = strings.TrimSpace(proposers[i])
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			theTx, err := authclient.ReadTxFromFile(clientCtx, args[0])
			if err != nil {
				return err
			}
			msgs := theTx.GetMsgs()

			msg, err := group.NewMsgCreateProposalRequest(
				args[0],
				proposers,
				msgs,
				args[3],
			)
			if err != nil {
				return err
			}

			if err = msg.ValidateBasic(); err != nil {
				return fmt.Errorf("message validation failed: %w", err)
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// MsgVoteCmd creates a CLI command for Msg/Vote.
func MsgVoteCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "vote [proposal-id] [voter[,voter]*] [choice] [comment]",
		Short: "Vote on a proposal",
		Long: `Vote on a proposal.

Parameters:
			proposal-id: unique ID of the proposal
			voter: comma separated (no spaces) list of voter account addresses. Example: "addr1,addr2" 
			choice: choice of the voter(s)
				0: no-op
				1: no
				2: yes
				3: abstain
				4: veto
			comment: comment for the vote
`,
		Args: cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			voters := strings.Split(args[1], ",")
			for i := range voters {
				voters[i] = strings.TrimSpace(voters[i])
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			proposalID, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			choice, err := group.ChoiceFromString(args[2])
			if err != nil {
				return err
			}

			msg := &group.MsgVoteRequest{
				ProposalId: group.ProposalID(proposalID),
				Voters:     voters,
				Choice:     choice,
				Comment:    args[3],
			}
			if err != nil {
				return err
			}

			if err = msg.ValidateBasic(); err != nil {
				return fmt.Errorf("message validation failed: %w", err)
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// MsgMsgExecCmd creates a CLI command for Msg/MsgExec.
func MsgMsgExecCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "exec [proposal-id]",
		Short: "Execute a proposal",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			voters := strings.Split(args[1], ",")
			for i := range voters {
				voters[i] = strings.TrimSpace(voters[i])
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			proposalID, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			msg := &group.MsgExecRequest{
				ProposalId: group.ProposalID(proposalID),
				Signer:     clientCtx.GetFromAddress().String(),
			}
			if err != nil {
				return err
			}

			if err = msg.ValidateBasic(); err != nil {
				return fmt.Errorf("message validation failed: %w", err)
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
