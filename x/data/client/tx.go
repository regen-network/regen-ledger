package client

import (
	"net/url"

	"github.com/cosmos/cosmos-sdk/client/tx"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/regen-network/regen-ledger/x/data"

	sdkclient "github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

// TxCmd returns a root CLI command handler for all x/data transaction commands.
func TxCmd(name string) *cobra.Command {
	cmd := &cobra.Command{
		Use:                        name,
		Short:                      "Data transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       sdkclient.ValidateCmd,
	}

	cmd.AddCommand(
		MsgAnchorDataCmd(),
		MsgSignDataCmd(),
		MsgDefineResolverCmd(),
		MsgRegisterResolverCmd(),
	)

	return cmd
}

// MsgAnchorDataCmd creates a CLI command for Msg/AnchorData.
func MsgAnchorDataCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "anchor [iri]",
		Short: "Anchors a piece of data to the blockchain based on its secure " +
			"hash, effectively providing a tamper resistant timestamp.",
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := sdkclient.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			iri := args[0]
			if len(iri) == 0 {
				return sdkerrors.ErrInvalidRequest.Wrap("iri cannot be empty")
			}

			signer := clientCtx.GetFromAddress()
			content, err := data.ParseIRI(iri)
			if err != nil {
				return sdkerrors.ErrInvalidRequest.Wrapf("invalid iri: %s", err.Error())
			}

			msg := data.MsgAnchorData{
				Sender: signer.String(),
				Hash:   content,
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// MsgSignDataCmd creates a CLI command for Msg/SignData.
func MsgSignDataCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "sign [iri]",
		Short:   `Sign a piece of on-chain data.`,
		Long:    `Sign a piece of on-chain data, attesting to its validity. The data MUST be of graph type (rdf file extension).`,
		Example: "regen tx data sign regen:13toVgf5aZqSVSeJQv562xkkeoe3rr3bJWa29PHVKVf77VAkVMcDvVd.rdf",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := sdkclient.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			iri := args[0]
			if len(iri) == 0 {
				return sdkerrors.ErrInvalidRequest.Wrap("iri is required")
			}

			signer := clientCtx.GetFromAddress()
			content, err := data.ParseIRI(iri)
			if err != nil {
				return sdkerrors.ErrInvalidRequest.Wrapf("invalid iri: %s", err.Error())
			}
			graph := content.GetGraph()
			if graph == nil {
				return sdkerrors.ErrInvalidRequest.Wrap("can only sign graph data types")
			}

			msg := data.MsgSignData{
				Signers: []string{signer.String()},
				Hash:    graph,
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
		Short: `Registers data content hashes.`,
		Long: `RegisterResolver registers data content hashes.
Parameters:
  resolver_url: resolver_url is a resolver URL which should refer to an HTTP service which will respond to 
			  a GET request with the IRI of a ContentHash and return the content if it exists or a 404.
Flags:
  --from: from flag is the address of the resolver manager
		`,
		Example: "regen tx data define-resolver http://foo.bar --from manager",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := sdkclient.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			resolverUrl := args[0]
			if _, err := url.ParseRequestURI(resolverUrl); err != nil {
				return err
			}

			msg := data.MsgDefineResolver{
				Manager:     clientCtx.GetFromAddress().String(),
				ResolverUrl: resolverUrl,
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
		Use:     "register-resolver",
		Short:   ``,
		Long:    ``,
		Example: "",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {

			return nil
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
