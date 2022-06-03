package client

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"

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
attestors and content (if available) for the given IRI. Otherwise, this command
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
		QueryByAttestorCmd(),
		QueryHashByIRICmd(),
		QueryIRIByHashCmd(),
		QueryAttestorsCmd(),
		QueryResolverCmd(),
		QueryResolversByIriCmd(),
		QueryResolversByUrlCmd(),
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

// QueryByAttestorCmd creates a CLI command for Query/ByAttestor.
func QueryByAttestorCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "by-attestor [attestor]",
		Short: "Query for anchored data based on an attestor",
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

			res, err := c.ByAttestor(cmd.Context(), &data.QueryByAttestorRequest{
				Attestor:   args[0],
				Pagination: pagination,
			})

			return print(ctx, res, err)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// QueryHashByIRICmd creates a CLI command for Query/HashByIRI.
func QueryHashByIRICmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "content-hash [iri]",
		Short: "Query for content hash based on IRI",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			c, ctx, err := mkQueryClient(cmd)
			if err != nil {
				return err
			}

			res, err := c.HashByIRI(cmd.Context(), &data.QueryHashByIRIRequest{
				Iri: args[0],
			})

			return print(ctx, res, err)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// QueryIRIByHashCmd creates a CLI command for Query/IRIByHash.
func QueryIRIByHashCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "iri [content-hash-json]",
		Short: "Query for IRI based on content hash",
		Long: `Query for IRI based on content hash provided in json format.

Parameters:
  content-hash-json: contains the content hash formatted as json`,
		Example: `regen q data iri content.json

where content.json contains:
{
  "graph": {
    "hash": "YWJjZGVmZ2hpamtsbW5vcHFyc3R1dnd4eXoxMjM0NTY=",
    "digest_algorithm": "DIGEST_ALGORITHM_BLAKE2B_256",
    "canonicalization_algorithm": "GRAPH_CANONICALIZATION_ALGORITHM_URDNA2015",
    "merkle_tree": "GRAPH_MERKLE_TREE_NONE_UNSPECIFIED"
  }
}`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			c, ctx, err := mkQueryClient(cmd)
			if err != nil {
				return err
			}

			contentHash, err := parseContentHash(ctx, args[0])
			if err != nil {
				return err
			}

			res, err := c.IRIByHash(cmd.Context(), &data.QueryIRIByHashRequest{
				ContentHash: contentHash,
			})

			return print(ctx, res, err)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// QueryAttestorsCmd creates a CLI command for Query/Attestors.
func QueryAttestorsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "attestors [iri]",
		Short: "Query for attestors based on IRI",
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

			res, err := c.AttestorsByIRI(cmd.Context(), &data.QueryAttestorsByIRIRequest{
				Iri:        args[0],
				Pagination: pagination,
			})

			return print(ctx, res, err)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// QueryResolverCmd creates a CLI command for Query/Resolver.
func QueryResolverCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "resolver [id]",
		Short:   "Query for a resolver based on resolver ID",
		Long:    "Query for a resolver based on resolver ID.",
		Example: "regen q data resolver 1",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			c, ctx, err := mkQueryClient(cmd)
			if err != nil {
				return err
			}

			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return fmt.Errorf("invalid resolver id: %s", err)
			}

			res, err := c.Resolver(cmd.Context(), &data.QueryResolverRequest{
				Id: id,
			})

			return print(ctx, res, err)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// QueryResolversByIriCmd creates a CLI command for Query/Resolvers.
func QueryResolversByIriCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "resolvers-by-iri [iri]",
		Short: "Query for resolvers based on the IRI of anchored data",
		Long:  "Query for resolvers based on the IRI of anchored data with pagination.",
		Example: `
regen q data resolvers-by-iri regen:113gdjFKcVCt13Za6vN7TtbgMM6LMSjRnu89BMCxeuHdkJ1hWUmy.rdf
regen q data resolvers-by-iri regen:113gdjFKcVCt13Za6vN7TtbgMM6LMSjRnu89BMCxeuHdkJ1hWUmy.rdf --pagination.limit 10
`,
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

			res, err := c.ResolversByIRI(cmd.Context(), &data.QueryResolversByIRIRequest{
				Iri:        args[0],
				Pagination: pagination,
			})

			return print(ctx, res, err)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "resolvers-by-iri")

	return cmd
}

// QueryResolversByUrlCmd creates a CLI command for Query/Resolvers.
func QueryResolversByUrlCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "resolvers-by-url [url]",
		Short: "Query for resolvers based on URL",
		Long:  "Query for resolvers based on the URL of the resolver.",
		Example: `
regen q data resolvers-by-url https://foo.bar
regen q data resolvers-by-url https://foo.bar --pagination.limit 10
`,
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

			res, err := c.ResolversByUrl(cmd.Context(), &data.QueryResolversByUrlRequest{
				Url:        args[0],
				Pagination: pagination,
			})

			return print(ctx, res, err)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "resolvers-by-url")

	return cmd
}

func parseContentHash(clientCtx client.Context, filePath string) (*data.ContentHash, error) {
	contentHash := data.ContentHash{}

	if filePath == "" {
		return nil, fmt.Errorf("file path is empty")
	}

	bz, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	if err := clientCtx.Codec.UnmarshalJSON(bz, &contentHash); err != nil {
		return nil, err
	}

	return &contentHash, nil
}
