package client

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"

	"github.com/regen-network/regen-ledger/x/data"
)

// QueryCmd returns the parent command for all x/data query commands.
func QueryCmd(name string) *cobra.Command {
	cmd := &cobra.Command{
		Args:                       cobra.ExactArgs(1),
		Use:                        name,
		Short:                      "Query commands for the data module",
		Long:                       "Query commands for the data module.",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		QueryAnchorByIRICmd(),
		QueryAnchorByHashCmd(),
		QueryAttestationsByAttestorCmd(),
		QueryAttestationsByIRICmd(),
		QueryAttestationsByHashCmd(),
		QueryResolverCmd(),
		QueryResolversByIRICmd(),
		QueryResolversByHashCmd(),
		QueryResolversByURLCmd(),
		ConvertIRIToHashCmd(),
		ConvertHashToIRICmd(),
	)

	return cmd
}

// QueryAnchorByIRICmd creates a CLI command for Query/AnchorByIRI.
func QueryAnchorByIRICmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "anchor-by-iri [iri]",
		Short: "Query a data anchor by the IRI of the data",
		Long:  "Query a data anchor by the IRI of the data.",
		Example: formatExample(`
  regen q data anchor-by-iri regen:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf
		`),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			c, ctx, err := mkQueryClient(cmd)
			if err != nil {
				return err
			}

			res, err := c.AnchorByIRI(cmd.Context(), &data.QueryAnchorByIRIRequest{
				Iri: args[0],
			})

			return printQueryResponse(ctx, res, err)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// QueryAnchorByHashCmd creates a CLI command for Query/AnchorByHash.
func QueryAnchorByHashCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "anchor-by-hash [hash-json]",
		Short: "Query a data anchor by the ContentHash of the data",
		Long:  "Query a data anchor by the ContentHash of the data.",
		Example: formatExample(`
  regen q data anchor-by-hash hash.json

  where hash.json contains:
  {
    "graph": {
      "hash": "YWJjZGVmZ2hpamtsbW5vcHFyc3R1dnd4eXoxMjM0NTY=",
      "digest_algorithm": "DIGEST_ALGORITHM_BLAKE2B_256",
      "canonicalization_algorithm": "GRAPH_CANONICALIZATION_ALGORITHM_URDNA2015",
      "merkle_tree": "GRAPH_MERKLE_TREE_NONE_UNSPECIFIED"
    }
  }
		`),
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

			res, err := c.AnchorByHash(cmd.Context(), &data.QueryAnchorByHashRequest{
				ContentHash: contentHash,
			})

			return printQueryResponse(ctx, res, err)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// QueryAttestationsByAttestorCmd creates a CLI command for Query/AttestationsByAttestor.
func QueryAttestationsByAttestorCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "attestations-by-attestor [attestor]",
		Short: "Query data attestations by an attestor",
		Long:  "Query data attestations by an attestor with optional pagination flags.",
		Example: formatExample(`
  regen q data attestations-by-attestor regen16md38uw5z9v4du2dtq4qgake8ewyf36u6qgfza
  regen q data attestations-by-attestor regen16md38uw5z9v4du2dtq4qgake8ewyf36u6qgfza --limit 10 --count-total
		`),
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

			return printQueryResponse(ctx, res, err)
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
		Short: "Query data attestations by the IRI of the data",
		Long:  "Query data attestations by the IRI of the data with optional pagination flags.",
		Example: formatExample(`
  regen q data attestations-by-iri regen:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf
  regen q data attestations-by-iri regen:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf --limit 10 --count-total
		`),
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

			return printQueryResponse(ctx, res, err)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "attestations-by-iri")

	return cmd
}

// QueryAttestationsByHashCmd creates a CLI command for Query/AttestationsByHash.
func QueryAttestationsByHashCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "attestations-by-hash [hash-json]",
		Short: "Query data attestations by the ContentHash of the data",
		Long:  "Query data attestations by the ContentHash of the data with optional pagination flags.",
		Example: formatExample(`
  regen q data attestations-by-hash hash.json
  regen q data attestations-by-hash hash.json --limit 10 --count-total

  where hash.json contains:
  {
    "graph": {
      "hash": "YWJjZGVmZ2hpamtsbW5vcHFyc3R1dnd4eXoxMjM0NTY=",
      "digest_algorithm": "DIGEST_ALGORITHM_BLAKE2B_256",
      "canonicalization_algorithm": "GRAPH_CANONICALIZATION_ALGORITHM_URDNA2015",
      "merkle_tree": "GRAPH_MERKLE_TREE_NONE_UNSPECIFIED"
    }
  }
		`),
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

			pagination, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			res, err := c.AttestationsByHash(cmd.Context(), &data.QueryAttestationsByHashRequest{
				ContentHash: contentHash,
				Pagination:  pagination,
			})

			return printQueryResponse(ctx, res, err)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "attestations-by-hash")

	return cmd
}

// QueryResolverCmd creates a CLI command for Query/Resolver.
func QueryResolverCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "resolver [id]",
		Short: "Query a resolver by its unique identifier",
		Long:  "Query a resolver by its unique identifier.",
		Example: formatExample(`
  regen q data resolver 1
		`),
		Args: cobra.ExactArgs(1),
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

			return printQueryResponse(ctx, res, err)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// QueryResolversByIRICmd creates a CLI command for Query/ResolversByIRI.
func QueryResolversByIRICmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "resolvers-by-iri [iri]",
		Short: "Query resolvers with registered data by the IRI of the data",
		Long:  "Query resolvers with registered data by the IRI of the data with optional pagination flags.",
		Example: formatExample(`
  regen q data resolvers-by-iri regen:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf
  regen q data resolvers-by-iri regen:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf --limit 10 --count-total
		`),
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

			return printQueryResponse(ctx, res, err)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "resolvers-by-iri")

	return cmd
}

// QueryResolversByHashCmd creates a CLI command for Query/ResolversByHash.
func QueryResolversByHashCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "resolvers-by-hash [hash-json]",
		Short: "Query resolvers with registered data by the ContentHash of the data",
		Long:  "Query resolvers with registered data by the ContentHash of the data with optional pagination flags.",
		Example: formatExample(`
  regen q data resolvers-by-hash hash.json
  regen q data resolvers-by-hash hash.json --limit 10 --count-total

  where hash.json contains:
  {
    "graph": {
      "hash": "YWJjZGVmZ2hpamtsbW5vcHFyc3R1dnd4eXoxMjM0NTY=",
      "digest_algorithm": "DIGEST_ALGORITHM_BLAKE2B_256",
      "canonicalization_algorithm": "GRAPH_CANONICALIZATION_ALGORITHM_URDNA2015",
      "merkle_tree": "GRAPH_MERKLE_TREE_NONE_UNSPECIFIED"
    }
  }
		`),
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

			pagination, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			res, err := c.ResolversByHash(cmd.Context(), &data.QueryResolversByHashRequest{
				ContentHash: contentHash,
				Pagination:  pagination,
			})

			return printQueryResponse(ctx, res, err)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "resolvers-by-hash")

	return cmd
}

// QueryResolversByURLCmd creates a CLI command for Query/ResolversByURL.
func QueryResolversByURLCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "resolvers-by-url [url]",
		Short: "Query resolvers by URL",
		Long:  "Query resolvers by URL with optional pagination flags.",
		Example: formatExample(`
  regen q data resolvers-by-url https://foo.bar
  regen q data resolvers-by-url https://foo.bar --limit 10 --count-total
		`),
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

			return printQueryResponse(ctx, res, err)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "resolvers-by-url")

	return cmd
}

// ConvertIRIToHashCmd creates a CLI command for Query/ConvertIRIToHash.
func ConvertIRIToHashCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "convert-iri-to-hash [iri]",
		Short: "Convert an IRI to a ContentHash",
		Long:  "Convert an IRI to a ContentHash.",
		Example: formatExample(`
  regen q data convert-iri-to-hash regen:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf
		`),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			c, ctx, err := mkQueryClient(cmd)
			if err != nil {
				return err
			}

			res, err := c.ConvertIRIToHash(cmd.Context(), &data.ConvertIRIToHashRequest{
				Iri: args[0],
			})

			return printQueryResponse(ctx, res, err)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// ConvertHashToIRICmd creates a CLI command for Query/ConvertHashToIRI.
func ConvertHashToIRICmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "convert-hash-to-iri [hash-json]",
		Short: "Convert a ContentHash to an IRI",
		Long:  "Convert a ContentHash to an IRI.",
		Example: formatExample(`
  regen q data convert-hash-to-iri hash.json

  where hash.json contains:
  {
    "graph": {
      "hash": "YWJjZGVmZ2hpamtsbW5vcHFyc3R1dnd4eXoxMjM0NTY=",
      "digest_algorithm": "DIGEST_ALGORITHM_BLAKE2B_256",
      "canonicalization_algorithm": "GRAPH_CANONICALIZATION_ALGORITHM_URDNA2015",
      "merkle_tree": "GRAPH_MERKLE_TREE_NONE_UNSPECIFIED"
    }
  }
		`),
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

			return printQueryResponse(ctx, res, err)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
