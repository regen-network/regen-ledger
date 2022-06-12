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
	queryAnchorByIRICmd := QueryAnchorByIRICmd()

	cmd := &cobra.Command{
		Args:  cobra.ExactArgs(1),
		Use:   fmt.Sprintf("%s [iri]", name),
		Short: "Query commands for the data module",
		Long: strings.TrimSpace(`Query commands for the data module.

If an IRI is passed as first argument, then this command will query timestamp,
attestors and content (if available) for the given IRI. Otherwise, this command
will run the given subcommand.

Example (the two following commands are equivalent):
$ regen query data regen:113gdjFKcVCt13Za6vN7TtbgMM6LMSjRnu89BMCxeuHdkJ1hWUmy.rdf
$ regen query data anchor-by-iri regen:113gdjFKcVCt13Za6vN7TtbgMM6LMSjRnu89BMCxeuHdkJ1hWUmy.rdf`),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE: func(cmd *cobra.Command, args []string) error {
			// If 1st arg is NOT an IRI, parse subcommands as usual.
			if !strings.Contains(args[0], "regen:") {
				return client.ValidateCmd(cmd, args)
			}

			// Or else, we call QueryAnchorByIRICmd.
			return queryAnchorByIRICmd.RunE(cmd, args)
		},
	}

	cmd.AddCommand(
		queryAnchorByIRICmd,
		QueryAttestationsByAttestorCmd(),
		QueryAttestationsByIRICmd(),
		QueryResolverCmd(),
		QueryResolversByIRICmd(),
		QueryResolversByURLCmd(),
		ConvertIRIToHashCmd(),
		ConvertHashToIRICmd(),
	)

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// QueryAnchorByIRICmd creates a CLI command for Query/AnchorByIRI.
func QueryAnchorByIRICmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "anchor-by-iri [iri]",
		Short: "Query for anchored data by IRI",
		Long:  "Query for anchored data by IRI.",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			c, ctx, err := mkQueryClient(cmd)
			if err != nil {
				return err
			}

			res, err := c.AnchorByIRI(cmd.Context(), &data.QueryAnchorByIRIRequest{
				Iri: args[0],
			})

			return print(ctx, res, err)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// QueryAttestationsByAttestorCmd creates a CLI command for Query/AttestationsByAttestor.
func QueryAttestationsByAttestorCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "attestations-by-attestor [attestor]",
		Short: "Query for attestations by attestor",
		Long:  "Query for attestations by attestor with optional pagination flags.",
		Example: `
regen q data attestations-by-attestor regen16md38uw5z9v4du2dtq4qgake8ewyf36u6qgfza
regen q data attestations-by-attestor regen16md38uw5z9v4du2dtq4qgake8ewyf36u6qgfza --limit 1 --count-total
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

			res, err := c.AttestationsByAttestor(cmd.Context(), &data.QueryAttestationsByAttestorRequest{
				Attestor:   args[0],
				Pagination: pagination,
			})

			return print(ctx, res, err)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "attestations-by-attestor")

	return cmd
}

// QueryAttestationsByIRICmd creates a CLI command for Query/AttestationsByIRI.
func QueryAttestationsByIRICmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "attestations-by-iri [iri]",
		Short: "Query for attestations by IRI",
		Long:  "Query for attestations by IRI with optional pagination flags.",
		Example: `
regen q data attestations-by-iri regen:113gdjFKcVCt13Za6vN7TtbgMM6LMSjRnu89BMCxeuHdkJ1hWUmy.rdf
regen q data attestations-by-iri regen:113gdjFKcVCt13Za6vN7TtbgMM6LMSjRnu89BMCxeuHdkJ1hWUmy.rdf --limit 1 --count-total
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

			res, err := c.AttestationsByIRI(cmd.Context(), &data.QueryAttestationsByIRIRequest{
				Iri:        args[0],
				Pagination: pagination,
			})

			return print(ctx, res, err)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "attestations-by-iri")

	return cmd
}

// QueryResolverCmd creates a CLI command for Query/Resolver.
func QueryResolverCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "resolver [id]",
		Short:   "Query for a resolver by the resolver ID",
		Long:    "Query for a resolver by the resolver ID.",
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

// QueryResolversByIRICmd creates a CLI command for Query/ResolversByIRI.
func QueryResolversByIRICmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "resolvers-by-iri [iri]",
		Short: "Query for resolvers by the IRI of the anchored data",
		Long:  "Query for resolvers by the IRI of the anchored data with optional pagination flags.",
		Example: `
regen q data resolvers-by-iri regen:113gdjFKcVCt13Za6vN7TtbgMM6LMSjRnu89BMCxeuHdkJ1hWUmy.rdf
regen q data resolvers-by-iri regen:113gdjFKcVCt13Za6vN7TtbgMM6LMSjRnu89BMCxeuHdkJ1hWUmy.rdf --limit 1 --count-total
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

// QueryResolversByURLCmd creates a CLI command for Query/ResolversByURL.
func QueryResolversByURLCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "resolvers-by-url [url]",
		Short: "Query for resolvers by the URL of the resolver",
		Long:  "Query for resolvers by the URL of the resolver with optional pagination flags.",
		Example: `
regen q data resolvers-by-url https://foo.bar
regen q data resolvers-by-url https://foo.bar --limit 10 --count-total
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

			res, err := c.ResolversByURL(cmd.Context(), &data.QueryResolversByURLRequest{
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

// ConvertIRIToHashCmd creates a CLI command for Query/ConvertIRIToHash.
func ConvertIRIToHashCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "iri-to-hash [iri]",
		Short: "Convert an IRI to a ContentHash",
		Long:  "Convert an IRI to a ContentHash.",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			c, ctx, err := mkQueryClient(cmd)
			if err != nil {
				return err
			}

			res, err := c.ConvertIRIToHash(cmd.Context(), &data.ConvertIRIToHashRequest{
				Iri: args[0],
			})

			return print(ctx, res, err)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// ConvertHashToIRICmd creates a CLI command for Query/ConvertHashToIRI.
func ConvertHashToIRICmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "hash-to-iri [content-hash-json]",
		Short: "Convert a ContentHash to an IRI",
		Long: `Convert a ContentHash to an IRI.

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

			res, err := c.ConvertHashToIRI(cmd.Context(), &data.ConvertHashToIRIRequest{
				ContentHash: contentHash,
			})

			return print(ctx, res, err)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

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
