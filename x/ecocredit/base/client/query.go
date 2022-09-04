package client

import (
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"

	types "github.com/regen-network/regen-ledger/x/ecocredit/base/types/v1"
)

// QueryClassesCmd returns a query command that lists all credit classes.
func QueryClassesCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "classes",
		Short: "List all credit classes",
		Long:  "List all credit classes with optional pagination flags.",
		Example: `regen q ecocredit classes
regen q ecocredit classes --limit 10 --offset 10`,
		Args: cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			c, ctx, err := mkQueryClient(cmd)
			if err != nil {
				return err
			}

			pagination, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			res, err := c.Classes(cmd.Context(), &types.QueryClassesRequest{
				Pagination: pagination,
			})
			return printQueryResponse(ctx, res, err)
		},
	}
	flags.AddPaginationFlagsToCmd(cmd, "classes")
	return qflags(cmd)
}

// QueryClassCmd returns a query command that retrieves information for a
// given credit class.
func QueryClassCmd() *cobra.Command {
	return qflags(&cobra.Command{
		Use:     "class [class-id]",
		Short:   "Retrieve credit class information",
		Long:    "Retrieve credit class information.",
		Example: "regen q ecocredit class C01",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			c, ctx, err := mkQueryClient(cmd)
			if err != nil {
				return err
			}
			res, err := c.Class(cmd.Context(), &types.QueryClassRequest{
				ClassId: args[0],
			})
			return printQueryResponse(ctx, res, err)
		},
	})
}

// QueryClassIssuersCmd returns a query command that retrieves addresses of the
// credit class issuers.
func QueryClassIssuersCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "class-issuers [class-id]",
		Short: "Retrieve issuer addresses for a credit class",
		Long:  "Retrieve issuer addresses for a credit class with optional pagination flags.",
		Example: `regen q ecocredit class-issuers C01
regen q ecocredit class-issuers C01 --limit 10 --offset 10`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			c, ctx, err := mkQueryClient(cmd)
			if err != nil {
				return err
			}

			pagination, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			res, err := c.ClassIssuers(cmd.Context(), &types.QueryClassIssuersRequest{
				ClassId:    args[0],
				Pagination: pagination,
			})
			if err != nil {
				return err
			}

			return printQueryResponse(ctx, res, err)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, "class-issuers")
	return qflags(cmd)
}

// QueryProjectsCmd returns a query command that retrieves all projects.
func QueryProjectsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "projects",
		Short: "List all projects",
		Long:  "List all projects with optional pagination flags.",
		Example: `regen q ecocredit projects
regen q ecocredit projects --limit 10 --offset 10`,
		Args: cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			c, ctx, err := mkQueryClient(cmd)
			if err != nil {
				return err
			}

			pagination, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			res, err := c.Projects(cmd.Context(), &types.QueryProjectsRequest{
				Pagination: pagination,
			})

			return printQueryResponse(ctx, res, err)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, "projects")

	return qflags(cmd)
}

// QueryProjectsByClassCmd returns a query command that retrieves projects by credit class.
func QueryProjectsByClassCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "projects-by-class [class-id]",
		Short: "List projects by credit class",
		Long:  "List projects by credit class with optional pagination flags.",
		Example: `regen q ecocredit projects-by-class C01
regen q ecocredit projects-by-class C01 --limit 10 --offset 10`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			c, ctx, err := mkQueryClient(cmd)
			if err != nil {
				return err
			}

			pagination, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			res, err := c.ProjectsByClass(cmd.Context(), &types.QueryProjectsByClassRequest{
				ClassId:    args[0],
				Pagination: pagination,
			})
			return printQueryResponse(ctx, res, err)
		},
	}
	flags.AddPaginationFlagsToCmd(cmd, "projects-by-class")
	return qflags(cmd)
}

// QueryProjectCmd returns a query command that retrieves project information.
func QueryProjectCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "project [project-id]",
		Short:   "Retrieve project information",
		Long:    "Retrieve project information.",
		Example: `regen q ecocredit project C01-001`,
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			c, ctx, err := mkQueryClient(cmd)
			if err != nil {
				return err
			}

			res, err := c.Project(cmd.Context(), &types.QueryProjectRequest{
				ProjectId: args[0],
			})
			return printQueryResponse(ctx, res, err)
		},
	}

	return qflags(cmd)
}

// QueryBatchesCmd returns a query command that retrieves credit batches for a
// given project.
func QueryBatchesCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "batches",
		Short: "List all credit batches",
		Long:  "List all credit batches with optional pagination flags.",
		Example: `regen q ecocredit batches
regen q ecocredit batches --limit 10 --offset 10`,
		Args: cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			c, ctx, err := mkQueryClient(cmd)
			if err != nil {
				return err
			}

			pagination, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			res, err := c.Batches(cmd.Context(), &types.QueryBatchesRequest{
				Pagination: pagination,
			})

			return printQueryResponse(ctx, res, err)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, "batches")

	return qflags(cmd)
}

// QueryBatchesByIssuerCmd returns a query command that retrieves credit batches by issuer.
func QueryBatchesByIssuerCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "batches-by-issuer [issuer]",
		Short: "List all credit batches by issuer",
		Long:  "List all credit batches by issuer with optional pagination flags.",
		Example: `regen q ecocredit batches-by-issuer regen1r9pl9gvr56kmclgkpjg3ynh4rm5am66f2a6y38
regen q ecocredit batches-by-issuer regen1r9pl9gvr56kmclgkpjg3ynh4rm5am66f2a6y38 --limit 10 --offset 10`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			c, ctx, err := mkQueryClient(cmd)
			if err != nil {
				return err
			}

			pagination, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			res, err := c.BatchesByIssuer(cmd.Context(), &types.QueryBatchesByIssuerRequest{
				Issuer:     args[0],
				Pagination: pagination,
			})

			return printQueryResponse(ctx, res, err)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, "batches-by-issuer")

	return qflags(cmd)
}

// QueryBatchesByClassCmd returns a query command that retrieves credit batches for a
// given credit class.
func QueryBatchesByClassCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "batches-by-class [class-id]",
		Short: "List all credit batches by credit class",
		Long:  "List all credit batches by credit class with pagination flags.",
		Example: `regen q ecocredit batches-by-class C01
regen q ecocredit batches-by-class C01 --limit 10 --offset 10`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			c, ctx, err := mkQueryClient(cmd)
			if err != nil {
				return err
			}

			pagination, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			res, err := c.BatchesByClass(cmd.Context(), &types.QueryBatchesByClassRequest{
				ClassId:    args[0],
				Pagination: pagination,
			})

			return printQueryResponse(ctx, res, err)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, "batches-by-class")

	return qflags(cmd)
}

// QueryBatchesByProjectCmd returns a query command that retrieves credit batches for a
// given project.
func QueryBatchesByProjectCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "batches-by-project [project-id]",
		Short: "List all credit batches by project",
		Long:  "List all credit batches by project with optional pagination flags.",
		Example: `regen q ecocredit batches-by-project C01-001
regen q ecocredit batches-by-project C01-001 --limit 10 --offset 10`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			c, ctx, err := mkQueryClient(cmd)
			if err != nil {
				return err
			}

			pagination, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			res, err := c.BatchesByProject(cmd.Context(), &types.QueryBatchesByProjectRequest{
				ProjectId:  args[0],
				Pagination: pagination,
			})

			return printQueryResponse(ctx, res, err)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, "batches-by-project")

	return qflags(cmd)
}

// QueryBatchCmd returns a query command that retrieves information for a
// given credit batch.
func QueryBatchCmd() *cobra.Command {
	return qflags(&cobra.Command{
		Use:     "batch [batch-denom]",
		Short:   "Retrieve credit batch information",
		Long:    "Retrieve credit batch information.",
		Example: "regen q ecocredit batch C01-001-20200101-20210101-001",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			c, ctx, err := mkQueryClient(cmd)
			if err != nil {
				return err
			}

			res, err := c.Batch(cmd.Context(), &types.QueryBatchRequest{
				BatchDenom: args[0],
			})
			return printQueryResponse(ctx, res, err)
		},
	})
}

// QueryBatchBalanceCmd returns a query command that retrieves the tradable and
// retired balances for a given credit batch and account address.
func QueryBatchBalanceCmd() *cobra.Command {
	return qflags(&cobra.Command{
		Use:     "batch-balance [batch-denom] [account]",
		Short:   "Retrieve the batch balance of an account",
		Long:    "Retrieve the batch balance of an account.",
		Example: "regen q ecocredit batch-balance C01-001-20200101-20210101-001 regen1r9pl9gvr56kmclgkpjg3ynh4rm5am66f2a6y38",
		Args:    cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			c, ctx, err := mkQueryClient(cmd)
			if err != nil {
				return err
			}
			res, err := c.Balance(cmd.Context(), &types.QueryBalanceRequest{
				BatchDenom: args[0], Address: args[1],
			})
			return printQueryResponse(ctx, res, err)
		},
	})
}

// QueryBatchSupplyCmd returns a query command that retrieves the tradable and
// retired supply of credits for a given credit batch.
func QueryBatchSupplyCmd() *cobra.Command {
	return qflags(&cobra.Command{
		Use:     "batch-supply [batch-denom]",
		Short:   "Retrieve the supply of a credit batch",
		Long:    "Retrieve the supply of a credit batch.",
		Example: "regen q ecocredit batch-supply C01-001-20200101-20210101-001",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			c, ctx, err := mkQueryClient(cmd)
			if err != nil {
				return err
			}
			res, err := c.Supply(cmd.Context(), &types.QuerySupplyRequest{
				BatchDenom: args[0],
			})
			return printQueryResponse(ctx, res, err)
		},
	})
}

// QueryCreditTypesCmd returns a query command that retrieves the list of
// approved credit types.
func QueryCreditTypesCmd() *cobra.Command {
	return qflags(&cobra.Command{
		Use:     "credit-types",
		Short:   "List all credit types",
		Long:    "List all credit types.",
		Example: "regen q ecocredit types",
		Args:    cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			c, ctx, err := mkQueryClient(cmd)
			if err != nil {
				return err
			}
			res, err := c.CreditTypes(cmd.Context(), &types.QueryCreditTypesRequest{})
			return printQueryResponse(ctx, res, err)
		},
	})
}

// QueryParamsCmd returns ecocredit module parameters.
func QueryParamsCmd() *cobra.Command {
	return qflags(&cobra.Command{
		Use:     "params",
		Short:   "List the current ecocredit module parameters",
		Long:    "List the current ecocredit module parameters.",
		Example: "regen q ecocredit params",
		RunE: func(cmd *cobra.Command, args []string) error {
			c, ctx, err := mkQueryClient(cmd)
			if err != nil {
				return err
			}
			res, err := c.Params(cmd.Context(), &types.QueryParamsRequest{})
			return printQueryResponse(ctx, res, err)
		},
	})
}

// QueryProjectsByReferenceIDCmd returns command that retrieves list of projects by reference id with pagination.
func QueryProjectsByReferenceIDCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "projects-by-reference-id [reference-id]",
		Short: "List all projects by reference ID",
		Long:  "List all projects by reference ID with optional pagination flags.",
		Example: `regen q ecocredit projects-by-reference-id VCS-001
regen q ecocredit projects-by-reference-id VCS-001 --limit 10 --offset 10`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			c, ctx, err := mkQueryClient(cmd)
			if err != nil {
				return err
			}

			pagination, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			res, err := c.ProjectsByReferenceId(cmd.Context(), &types.QueryProjectsByReferenceIdRequest{
				ReferenceId: args[0],
				Pagination:  pagination,
			})
			return printQueryResponse(ctx, res, err)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, "projects-by-reference-id")

	return qflags(cmd)
}

// QueryProjectsByAdminCmd returns command that retrieves list of projects by admin with pagination.
func QueryProjectsByAdminCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "projects-by-admin [admin]",
		Short: "List projects by admin",
		Long:  "List projects by admin with optional pagination flags.",
		Example: `regen q ecocredit projects-by-admin regen1elq7ys34gpkj3jyvqee0h6yk4h9wsfxmgqelsw
regen q ecocredit projects-by-admin regen1elq7ys34gpkj3jyvqee0h6yk4h9wsfxmgqelsw --limit 10 --offset 10`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			c, ctx, err := mkQueryClient(cmd)
			if err != nil {
				return err
			}

			pagination, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			res, err := c.ProjectsByAdmin(cmd.Context(), &types.QueryProjectsByAdminRequest{
				Admin:      args[0],
				Pagination: pagination,
			})
			return printQueryResponse(ctx, res, err)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, "projects-by-admin")

	return qflags(cmd)
}

// QueryCreditTypeCmd returns a query command that retrieves credit type
// information by abbreviation.
func QueryCreditTypeCmd() *cobra.Command {
	return qflags(&cobra.Command{
		Use:     "credit-type [abbreviation]",
		Short:   "Retrieve credit type information",
		Long:    "Retrieve credit type information.",
		Example: "regen q ecocredit credit-type C",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			c, ctx, err := mkQueryClient(cmd)
			if err != nil {
				return err
			}
			res, err := c.CreditType(cmd.Context(), &types.QueryCreditTypeRequest{
				Abbreviation: args[0],
			})
			return printQueryResponse(ctx, res, err)
		},
	})
}

// QueryClassFeeCmd returns a query command that retrieves the credit class fees.
func QueryClassFeeCmd() *cobra.Command {
	return qflags(&cobra.Command{
		Use:     "class-fee",
		Short:   "Retrieve the credit class creation fee",
		Long:    "Retrieve the credit class creation fee",
		Example: "regen q ecocredit class-fee",
		RunE: func(cmd *cobra.Command, args []string) error {
			c, ctx, err := mkQueryClient(cmd)
			if err != nil {
				return err
			}
			res, err := c.ClassFee(cmd.Context(), &types.QueryClassFeeRequest{})
			return printQueryResponse(ctx, res, err)
		},
	})
}

// QueryClassCreatorAllowlistCmd returns a query command that retrieves the
// class allowlist enabled setting.
func QueryClassCreatorAllowlistCmd() *cobra.Command {
	return qflags(&cobra.Command{
		Use:     "class-creator-allowlist",
		Short:   "Retrieve the class creator allowlist enabled setting",
		Long:    "Retrieve the class creator allowlist enabled setting",
		Example: "regen q ecocredit class-creator-allowlist",
		RunE: func(cmd *cobra.Command, args []string) error {
			c, ctx, err := mkQueryClient(cmd)
			if err != nil {
				return err
			}
			res, err := c.ClassCreatorAllowlist(cmd.Context(), &types.QueryClassCreatorAllowlistRequest{})
			return printQueryResponse(ctx, res, err)
		},
	})
}

// QueryAllowedClassCreatorsCmd returns a query command that retrives the list of allowed
// credit class creators with pagination.
func QueryAllowedClassCreatorsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "allowed-class-creators",
		Short: "Retrieve the allowed credit class creators",
		Long:  "Retrieve the list of allowed credit class creators with pagination",
		Example: `
		regen q ecocredit allowed-class-creators
		regen q ecocredit allowed-class-creators --limit 10`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, ctx, err := mkQueryClient(cmd)
			if err != nil {
				return err
			}

			pagination, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			res, err := c.AllowedClassCreators(cmd.Context(), &types.QueryAllowedClassCreatorsRequest{
				Pagination: pagination,
			})
			return printQueryResponse(ctx, res, err)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, "batches-by-project")

	return qflags(cmd)
}

// QueryAllBalances returns a query command that retrieves a list of all ecocredit balances
// with pagination.
func QueryAllBalances() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "all-balances",
		Short: "Retrieve all ecocredit balances",
		Long:  "Retrieve all ecocredit balances across all addresses and batch denoms with pagination",
		Example: `
		regen q ecocredit all-balances
		regen q ecocredit all-balances --limit 10
		`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, ctx, err := mkQueryClient(cmd)
			if err != nil {
				return err
			}
			pagination, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			res, err := c.AllBalances(cmd.Context(), &types.QueryAllBalancesRequest{
				Pagination: pagination,
			})
			return printQueryResponse(ctx, res, err)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, "all-balances")

	return qflags(cmd)
}

// QueryAllowedBridgeChains returns a query command that retrieves a list of chain that are allowed to be used
// in bridge operations.
func QueryAllowedBridgeChains() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "allowed-bridge-chains",
		Short:   "Retrieve the list of allowed bridge chains",
		Long:    "Retrieve the list of chains that are allowed to be used in bridge operations",
		Example: "regen q ecocredit allowed-bridge-chains",
		RunE: func(cmd *cobra.Command, args []string) error {
			c, ctx, err := mkQueryClient(cmd)
			if err != nil {
				return err
			}
			res, err := c.AllowedBridgeChains(cmd.Context(), &types.QueryAllowedBridgeChainsRequest{})
			return printQueryResponse(ctx, res, err)
		},
	}
	flags.AddPaginationFlagsToCmd(cmd, "allowed-bridge-chains")
	return qflags(cmd)
}
