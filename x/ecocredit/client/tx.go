package client

import (
	"encoding/base64"
	"errors"
	"strconv"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"

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

			c, err := newMsgSrvClient(cmd)
			if err != nil {
				return err
			}
			msg := ecocredit.MsgCreateClassRequest{
				Designer: args[0], Issuers: issuers, Metadata: b,
			}
			_, err = c.client.CreateClass(cmd.Context(), &msg)
			return c.send(err)
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
             eg: '[{recipient: "xrn:sdgkjhs2345u79ghisodg", tradable_units: "10", retired_units: "2"}]'
             Note: "tradable_units" and "retired_units" default to 0.`,
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

			msg := ecocredit.MsgCreateBatchRequest{
				Issuer: args[0], ClassId: args[1], Metadata: b, Issuance: issuance,
			}
			c, err := newMsgSrvClient(cmd)
			if err != nil {
				return err
			}
			_, err = c.client.CreateBatch(cmd.Context(), &msg)
			return c.send(err)
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
             eg: '[{batch_denom: "100/2", tradable_units: "5", retired_units: "0"}]'`,
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			var credits = []*ecocredit.MsgSendRequest_SendUnits{}
			if err := yaml.Unmarshal([]byte(args[1]), &credits); err != nil {
				return err
			}
			c, err := newMsgSrvClient(cmd)
			if err != nil {
				return err
			}
			msg := ecocredit.MsgSendRequest{
				Sender:    c.Cctx.GetFromAddress().String(),
				Recipient: args[0], Credits: credits,
			}
			_, err = c.client.Send(cmd.Context(), &msg)
			return c.send(err)
		},
	}
}

func txRetire() *cobra.Command {
	return &cobra.Command{
		Use:   "retire [credits]",
		Short: "Retires a specified amounts of credits from the account of the transaction author (--from)",
		Long: `Retires a specified amounts of credits from the account of the transaction author (--from)

Parameters:
  credits:  YAML encoded credit list. Note: numerical values must be written in strings.
            eg: '[{batch_denom: "100/2", units: "5"}]'`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			var credits = []*ecocredit.MsgRetireRequest_RetireUnits{}
			if err := yaml.Unmarshal([]byte(args[0]), &credits); err != nil {
				return err
			}
			c, err := newMsgSrvClient(cmd)
			if err != nil {
				return err
			}
			msg := ecocredit.MsgRetireRequest{
				Holder:  c.Cctx.GetFromAddress().String(),
				Credits: credits,
			}
			_, err = c.client.Retire(cmd.Context(), &msg)
			return c.send(err)
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
			c, err := newMsgSrvClient(cmd)
			if err != nil {
				return err
			}
			msg := ecocredit.MsgSetPrecisionRequest{
				Issuer:     c.Cctx.GetFromAddress().String(),
				BatchDenom: args[0], MaxDecimalPlaces: uint32(decimals),
			}
			_, err = c.client.SetPrecision(cmd.Context(), &msg)
			return c.send(err)
		},
	}
}
