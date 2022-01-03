package client

import (
	"fmt"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"

	"github.com/regen-network/regen-ledger/x/data"
)

// QueryCmd returns the parent command for all x/data CLI query commands
func QueryCmd(name string) *cobra.Command {
	queryByIRICmd := QueryByIRICmd()

	cmd := &cobra.Command{
		Args:  cobra.ExactArgs(1),
		Use:   fmt.Sprintf("%s [iri]", name),
		Short: "Querying commands for the data module",
		Long: strings.TrimSpace(`Querying commands for the data module.

If an IRI is passed as first argument, then this command will query timestamp,
signers and content (if available) for the given IRI. Otherwise, this command
will run the given subcommand.

Example (the two following commands are equivalent):
$ regen query data regen:113gdjFKcVCt13Za6vN7TtbgMM6LMSjRnu89BMCxeuHdkJ1hWUmy.rdf
$ regen query data by-iri regen:113gdjFKcVCt13Za6vN7TtbgMM6LMSjRnu89BMCxeuHdkJ1hWUmy.rdf`),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE: func(cmd *cobra.Command, args []string) error {
			// If 1st arg is NOT an IRI, parse subcommands as usual.
			if !strings.Contains(args[0], "regen:") {
				return client.ValidateCmd(cmd, args)
			}

			// Or else, we call QueryByIRICmd.
			return queryByIRICmd.RunE(cmd, args)
		},
	}

	cmd.AddCommand(
		queryByIRICmd,
		QueryBySignerCmd(),
		QuerySignersCmd(),
	)

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// QueryByIRICmd creates a CLI command for Query/Data.
func QueryByIRICmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "by-iri [iri]",
		Short: "Query for anchored data based on IRI",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			c, ctx, err := mkQueryClient(cmd)
			if err != nil {
				return err
			}

			res, err := c.ByIRI(cmd.Context(), &data.QueryByIRIRequest{
				Iri: args[0],
			})

			return print(ctx, res, err)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// QueryBySignerCmd creates a CLI command for Query/Data.
func QueryBySignerCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "by-signer [signer]",
		Short: "Query for anchored data based on signer",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			c, ctx, err := mkQueryClient(cmd)
			if err != nil {
				return err
			}

			pagination, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			res, err := c.BySigner(cmd.Context(), &data.QueryBySignerRequest{
				Signer:     args[0],
				Pagination: pagination,
			})

			return print(ctx, res, err)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// QuerySignersCmd creates a CLI command for Query/Data.
func QuerySignersCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "signers [iri]",
		Short: "Query for signers based on IRI",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			c, ctx, err := mkQueryClient(cmd)
			if err != nil {
				return err
			}

			pagination, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			res, err := c.Signers(cmd.Context(), &data.QuerySignersRequest{
				Iri:        args[0],
				Pagination: pagination,
			})

			return print(ctx, res, err)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
