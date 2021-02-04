package client

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/regen-network/regen-ledger/x/group"
	"github.com/spf13/cobra"
)

// QueryCmd returns the cli query commands for the group module.
func QueryCmd(name string) *cobra.Command {
	queryCmd := &cobra.Command{
		Use:                        name,
		Short:                      "Querying commands for the group module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	queryCmd.AddCommand(
		QueryGroupInfoCmd(),
		QueryGroupAccountInfoCmd(),
		QueryGroupMembersCmd(),
		// QueryGroupsByAdminCmd(),
		// QueryGroupAccountsByGroupCmd(),
		// QueryGroupAccountsByAdminCmd(),
		// QueryProposalCmd(),
		// QueryProposalsByGroupAccountCmd(),
		// QueryVoteByProposalVoterCmd(),
		// QueryVotesByProposalCmd(),
		// QueryVotesByVoterCmd(),
	)

	return queryCmd
}

// QueryGroupInfoCmd creates a CLI command for Query/GroupInfo.
func QueryGroupInfoCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "group-info [id]",
		Short: "Query for group info by group id",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			groupID, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			queryClient := group.NewQueryClient(clientCtx)

			res, err := queryClient.GroupInfo(cmd.Context(), &group.QueryGroupInfoRequest{
				GroupId: group.ID(groupID),
			})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// QueryGroupAccountInfoCmd creates a CLI command for Query/GroupAccountInfo.
func QueryGroupAccountInfoCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "group-account-info [group-account]",
		Short: "Query for group account info by group account address",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := group.NewQueryClient(clientCtx)

			res, err := queryClient.GroupAccountInfo(cmd.Context(), &group.QueryGroupAccountInfoRequest{
				GroupAccount: args[0],
			})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// QueryGroupMembersCmd creates a CLI command for Query/GroupMembers.
func QueryGroupMembersCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "group-members [id]",
		Short: "Query for group members by group id with pagination flags",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			groupID, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := group.NewQueryClient(clientCtx)

			res, err := queryClient.GroupMembers(cmd.Context(), &group.QueryGroupMembersRequest{
				GroupId:    group.ID(groupID),
				Pagination: pageReq,
			})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// QueryGroupsByAdminCmd creates a CLI command for Query/GroupsByAdmin.
func QueryGroupsByAdminCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "groups-by-admin [admin]",
		Short: "Query for groups by admin account address with pagination flags",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := group.NewQueryClient(clientCtx)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			res, err := queryClient.GroupsByAdmin(cmd.Context(), &group.QueryGroupsByAdminRequest{
				Admin:      args[0],
				Pagination: pageReq,
			})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
