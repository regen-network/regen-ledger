package client

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/regen-network/regen-ledger/x/group"
	"github.com/spf13/cobra"
)

const FlagMembers = "members"

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
		// MsgCreateGroupAccountCmd(),
		// MsgUpdateGroupAccountAdminCmd(),
		// MsgUpdateGroupAccountDecisionPolicyCmd(),
		// MsgUpdateGroupAccountCommentCmd(),
		// MsgCreateProposalCmd(),
		// MsgVoteCmd(),
		// MsgExecCmd(),
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

			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err = client.ReadTxCommandFlags(clientCtx, cmd.Flags())
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

	cmd.Flags().String(FlagMembers, "", "Members file path")
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

			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err = client.ReadTxCommandFlags(clientCtx, cmd.Flags())
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

	cmd.Flags().String(FlagMembers, "", "Members file path")
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

			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err = client.ReadTxCommandFlags(clientCtx, cmd.Flags())
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

			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err = client.ReadTxCommandFlags(clientCtx, cmd.Flags())
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
			fmt.Sprintf(`Create a group account which is an account
associated with a group and a decision policy.
Note, the '--from' flag is ignored as it is implied from [admin].
Example:
$ %s tx group create-group-account [admin] [group-id] [comment] '{"@type":"/regen.group.v","key":"OauFcTKbN5Lx3fJL689cikXBqe+hcp6Y+x0rYUdR9Jk="}'

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

			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err = client.ReadTxCommandFlags(clientCtx, cmd.Flags())
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

			msg := &group.MsgCreateGroupAccountRequest{
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

	cmd.Flags().String(FlagMembers, "", "Members file path")
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
