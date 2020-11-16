package client

import (
	"encoding/base64"
	"errors"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/spf13/cobra"

	"github.com/regen-network/regen-ledger/x/ecocredit"
)

// TxCmd returns a root CLI command handler for all x/data transaction commands.
func TxCmd() *cobra.Command {
	cmd := &cobra.Command{
		SuggestionsMinimumDistance: 2,
		DisableFlagParsing:         true,

		Use:   ecocredit.ModuleName,
		Short: "Ecocredit module transactions",
		RunE:  client.ValidateCmd,
	}

	cmd.AddCommand(
		txCreateClass(),
	)

	return cmd
}

func txCreateClass() *cobra.Command {
	cmd := &cobra.Command{
		Use: "create_class [designer] [issuer[,issuer]*] [base64_metadata]",
		Short: `Creates a new credit class.
  designer:  address of the account which designed the credit class
  issuer:    comma separated (no spaces) list of issuer account addresses. Example: "addr1,addr2"
  metadata:  base64 encoded metadata - arbitrary data attached to the credit class info`,
		Args: cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			issuers := strings.Split(args[1], ",")
			for i := range issuers {
				issuers[i] = strings.TrimSpace(issuers[i])
			}
			if args[2] == "" {
				return errors.New("base64_metadata is required")
			}
			b, err := base64.StdEncoding.DecodeString(args[2])
			if err != nil {
				return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "base64_metadata is malformed, proper base64 string is required")
			}

			c, ctx, err := mkMsgClient(cmd)
			if err != nil {
				return err
			}
			_, err = c.CreateClass(cmd.Context(), &ecocredit.MsgCreateClassRequest{
				Designer: args[0], Issuers: issuers, Metadata: b,
			})
			return mkTx(ctx, cmd, err)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
