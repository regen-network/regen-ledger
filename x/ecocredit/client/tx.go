package client

import (
	"encoding/base64"
	"errors"
	"strconv"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	sdkclient "github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"

	"github.com/regen-network/regen-ledger/x/ecocredit"
)

// TxCmd returns a root CLI command handler for all x/data transaction commands.
func TxCmd(name string) *cobra.Command {
	cmd := &cobra.Command{
		SuggestionsMinimumDistance: 2,
		DisableFlagParsing:         true,

		Use:   name,
		Short: "Ecocredit module transactions",
		RunE:  client.ValidateCmd,
	}
	cmd.AddCommand(
		txflags(txCreateClass()),
		txflags(txCreateBatch()),
		txflags(txSend()),
		txflags(txRetire()),
		txflags(txSetPrecision()),
	)
	return cmd
}

func txflags(cmd *cobra.Command) *cobra.Command {
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func txCreateClass() *cobra.Command {
	return &cobra.Command{
		Use:   "create-class [designer] [issuer[,issuer]*] [metadata]",
		Short: "Creates a new credit class",
		Long: `Creates a new credit class.

Parameters:
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
				return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "metadata is malformed, proper base64 string is required")
			}

			clientCtx, err := sdkclient.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			msg := ecocredit.MsgCreateClassRequest{
				Designer: args[0], Issuers: issuers, Metadata: b,
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}
}

func txCreateBatch() *cobra.Command {
	return &cobra.Command{
		Use:   "create-batch [issuer] [class_id] [metadata] [issuance]",
		Short: "Issues a new credit batch",
		Long: `Issues a new credit batch.

Parameters:
  issuer:    issuer address
  class_id:  credit class
  metadata:  base64 encoded issuance metadata
  issuance:  YAML encode issuance list. Note: numerical values must be written in strings.
             eg: '[{recipient: "xrn:sdgkjhs2345u79ghisodg", tradable_units: "10", retired_units: "2", retirement_location: "YY-ZZ 12345"}]'
             Note: "tradable_units" and "retired_units" default to 0.
             Note: "retirement_location" is only required when "retired_units" is positive.`,
		Args: cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			b, err := base64.StdEncoding.DecodeString(args[2])
			if err != nil {
				return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "metadata is malformed, proper base64 string is required")
			}
			var issuance = []*ecocredit.MsgCreateBatchRequest_BatchIssuance{}
			if err = yaml.Unmarshal([]byte(args[3]), &issuance); err != nil {
				return err
			}

			clientCtx, err := sdkclient.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			msg := ecocredit.MsgCreateBatchRequest{
				Issuer: args[0], ClassId: args[1], Metadata: b, Issuance: issuance,
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}
}

func txSend() *cobra.Command {
	return &cobra.Command{
		Use:   "send [recipient] [credits]",
		Short: "Sends credits from the transaction author (--from) to the recipient",
		Long: `Sends credits from the transaction author (--from) to the recipient.

Parameters:
  recipient: recipient address
  credits:   YAML encoded credit list. Note: numerical values must be written in strings.
             eg: '[{batch_denom: "100/2", tradable_units: "5", retired_units: "0", retirement_location: "YY-ZZ 12345"}]'
             Note: "retirement_location" is only required when "retired_units" is positive.`,
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			var credits = []*ecocredit.MsgSendRequest_SendUnits{}
			if err := yaml.Unmarshal([]byte(args[1]), &credits); err != nil {
				return err
			}
			clientCtx, err := sdkclient.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			msg := ecocredit.MsgSendRequest{
				Sender:    clientCtx.GetFromAddress().String(),
				Recipient: args[0], Credits: credits,
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}
}

func txRetire() *cobra.Command {
	return &cobra.Command{
		Use:   "retire [credits] [retirement_location]",
		Short: "Retires a specified amounts of credits from the account of the transaction author (--from)",
		Long: `Retires a specified amounts of credits from the account of the transaction author (--from)

Parameters:
  credits:             YAML encoded credit list. Note: numerical values must be written in strings.
                       eg: '[{batch_denom: "100/2", units: "5"}]'
  retirement_location: A string representing the location of the buyer or
                       beneficiary of retired credits. It has the form
                       <country-code>[-<region-code>[ <postal-code>]], where
                       country-code and region-code are taken from ISO 3166, and
                       postal-code being up to 64 alphanumeric characters.
                       eg: 'AA-BB 12345'`,
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			var credits = []*ecocredit.MsgRetireRequest_RetireUnits{}
			if err := yaml.Unmarshal([]byte(args[0]), &credits); err != nil {
				return err
			}
			clientCtx, err := sdkclient.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			msg := ecocredit.MsgRetireRequest{
				Holder:   clientCtx.GetFromAddress().String(),
				Credits:  credits,
				Location: args[1],
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}
}

func txSetPrecision() *cobra.Command {
	return &cobra.Command{
		Use:   "set_precision [batch_denom] [decimals]",
		Short: "Allows an issuer to increase the decimal precision of a credit batch",
		Long: `Allows an issuer to increase the decimal precision of a credit batch. It is an experimental feature.

Parameters:
  batch_denom: credit batch ID
  decimals:    maximum number of decimals of precision`,
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			decimals, err := strconv.ParseUint(args[1], 10, 32)
			if err == nil {
				return err
			}
			clientCtx, err := sdkclient.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			msg := ecocredit.MsgSetPrecisionRequest{
				Issuer:     clientCtx.GetFromAddress().String(),
				BatchDenom: args[0], MaxDecimalPlaces: uint32(decimals),
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}
}
