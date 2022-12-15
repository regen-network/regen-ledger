package client

import (
	"fmt"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/regen-network/regen-ledger/x/data/v2"
)

// TxCmd returns a root CLI command handler for all x/data transaction commands.
func TxCmd(name string) *cobra.Command {
	cmd := &cobra.Command{
		Use:                        name,
		Short:                      "Data transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		MsgAnchorCmd(),
		MsgAttestCmd(),
		MsgDefineResolverCmd(),
		MsgRegisterResolverCmd(),
	)

	return cmd
}

// MsgAnchorCmd creates a CLI command for Msg/Anchor.
func MsgAnchorCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "anchor [iri]",
		Short: "Anchors a piece of data to the blockchain based on its secure " +
			"hash, effectively providing a tamper resistant timestamp.",
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			iri := args[0]
			if len(iri) == 0 {
				return sdkerrors.ErrInvalidRequest.Wrap("iri cannot be empty")
			}

			attestor := clientCtx.GetFromAddress()
			contentHash, err := data.ParseIRI(iri)
			if err != nil {
				return sdkerrors.ErrInvalidRequest.Wrapf("invalid iri: %s", err.Error())
			}

			msg := data.MsgAnchor{
				Sender:      attestor.String(),
				ContentHash: contentHash,
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// MsgAttestCmd creates a CLI command for Msg/Attest.
func MsgAttestCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "attest [iri]",
		Short: `Attest to the veracity of anchored data.`,
		Long: `Attest to the veracity of anchored data. The data MUST be of graph type (rdf file extension).
Attest to the veracity of more than one entry using a comma-separated (no spaces) list of IRIs.`,
		Example: "regen tx data attest regen:13toVgf5aZqSVSeJQv562xkkeoe3rr3bJWa29PHVKVf77VAkVMcDvVd.rdf",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			attestor := clientCtx.GetFromAddress()

			if len(args[0]) == 0 {
				return sdkerrors.ErrInvalidRequest.Wrap("at least one iri is required")
			}

			var contentHashes []*data.ContentHash_Graph

			iris := strings.Split(args[0], ",")

			for i := range iris {
				content, err := data.ParseIRI(iris[i])
				if err != nil {
					return sdkerrors.ErrInvalidRequest.Wrapf("invalid iri: %s", err.Error())
				}
				graph := content.GetGraph()
				if graph == nil {
					return sdkerrors.ErrInvalidRequest.Wrap("can only attest to graph data types")
				}
				contentHashes = append(contentHashes, graph)
			}

			msg := data.MsgAttest{
				Attestor:      attestor.String(),
				ContentHashes: contentHashes,
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// MsgDefineResolverCmd creates a CLI command for Msg/DefineResolver.
func MsgDefineResolverCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "define-resolver [resolver_url]",
		Short: `Defines a resolver URL and assigns it a new integer ID that can be used in calls to RegisterResolver.`,
		Long: `DefineResolver defines a resolver URL and assigns it a new integer ID that can be used in calls to RegisterResolver.
Parameters:
  resolver_url: resolver_url is a resolver URL which should refer to an HTTP service which will respond to 
			  a GET request with the IRI of a ContentHash and return the content if it exists or a 404.
Flags:
  --from: from flag is the address of the resolver manager
		`,
		Example: "regen tx data define-resolver https://foo.bar --from manager",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			resolverURL := args[0]
			if _, err := url.ParseRequestURI(resolverURL); err != nil {
				return err
			}

			msg := data.MsgDefineResolver{
				Manager:     clientCtx.GetFromAddress().String(),
				ResolverUrl: resolverURL,
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// MsgRegisterResolverCmd creates a CLI command for Msg/RegisterResolver.
func MsgRegisterResolverCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "register-resolver [resolver_id] [content_hashes_json]",
		Short: `Registers data content hashes.`,
		Long: `RegisterResolver registers data content hashes.
Parameters:
    resolver_id: resolver id is the ID of a resolver
	content_hashes_json: contains list of content hashes which the resolver claims to serve
Flags:
	--from: manager is the address of the resolver manager
		`,
		Example: `
			regen tx data register-resolver 1 content.json

			where content.json contains
			{
				"content_hashes": [
					{
						"graph": {
							"hash": "YWJjZGVmZ2hpamtsbW5vcHFyc3R1dnd4eXoxMjM0NTY=",
							"digest_algorithm": "DIGEST_ALGORITHM_BLAKE2B_256",
							"canonicalization_algorithm": "GRAPH_CANONICALIZATION_ALGORITHM_URDNA2015",
							"merkle_tree": "GRAPH_MERKLE_TREE_NONE_UNSPECIFIED"
						}
					}
				]
			}
			`,
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			resolverID, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return fmt.Errorf("invalid resolver id")
			}

			contentHashes, err := parseContentHashes(clientCtx, args[1])
			if err != nil {
				return err
			}

			msg := data.MsgRegisterResolver{
				Manager:       clientCtx.GetFromAddress().String(),
				ResolverId:    resolverID,
				ContentHashes: contentHashes,
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func parseContentHashes(clientCtx client.Context, filePath string) ([]*data.ContentHash, error) {
	contentHashes := data.ContentHashes{}

	if filePath == "" {
		return nil, fmt.Errorf("file path is empty")
	}

	bz, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	if err := clientCtx.Codec.UnmarshalJSON(bz, &contentHashes); err != nil {
		return nil, err
	}

	return contentHashes.ContentHashes, nil
}
