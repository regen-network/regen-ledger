package client

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"strconv"
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
		txCreateBatch(),
		txSend(),
		txRetire(),
		txSetPrecision(),
	)
	return cmd
}

func txCreateClass() *cobra.Command {
	cmd := &cobra.Command{
		Use: "create_class [designer] [issuer[,issuer]*] [base64_metadata]",
		Short: `Creates a new credit class.
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
				return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "base64_metadata is malformed, proper base64 string is required")
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
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func txCreateBatch() *cobra.Command {
	cmd := &cobra.Command{
		Use: "create_batch [issuer] [class_id] [metadata] [issuance]",
		Short: `Issues new credit batach.
Parameters:
  issuer:    issuer address
  class_id:  credit class
  metadata:  base64 encoded issuance metadata
  issuance:  JSON encode issuance list,
             eg: '[{"recipient": "a1", "tradeableUnits": "10", "retiredUnits": "2"}]'`,
		Args: cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			b, err := base64.StdEncoding.DecodeString(args[2])
			if err != nil {
				return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "base64_metadata is malformed, proper base64 string is required")
			}
			var issuance = []*ecocredit.MsgCreateBatchRequest_BatchIssuance{}
			if err = json.Unmarshal([]byte(args[3]), &issuance); err != nil {
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
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func txSend() *cobra.Command {
	cmd := &cobra.Command{
		Use: "send [recipient] [credits]",
		Short: `Sends credits from the transaction author (--from) to the recipient.
Parameters:
  recipient: recipient address
  credits:   JSON encoded credit list
             eg: '[{"batchDenom": "100/2", "tradeableUnits": "5", "retiredUnits": "0"}]'`,
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			var credits = []*ecocredit.MsgSendRequest_SendUnits{}
			if err := json.Unmarshal([]byte(args[1]), &credits); err != nil {
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
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func txRetire() *cobra.Command {
	cmd := &cobra.Command{
		Use: "retire [holder] [credits]",
		Short: `Retires specified amount of a holder credits. Transaction author (--from) must
must be an authorized issuer.
  recipient: recipient address
  credits:   JSON encoded credit list
             eg: '[{"batchDenom": "100/2", "units": "5"}]'`,
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			var credits = []*ecocredit.MsgRetireRequest_RetireUnits{}
			if err := json.Unmarshal([]byte(args[1]), &credits); err != nil {
				return err
			}
			c, err := newMsgSrvClient(cmd)
			if err != nil {
				return err
			}
			msg := ecocredit.MsgRetireRequest{
				Issuer: c.Cctx.GetFromAddress().String(),
				Holder: args[0], Credits: credits,
			}
			_, err = c.client.Retire(cmd.Context(), &msg)
			return c.send(err)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func txSetPrecision() *cobra.Command {
	cmd := &cobra.Command{
		Use: "set_precision [batch_denom] [decimals]",
		Short: `SetPrecision allows an issuer to increase the decimal precision of a credit batch. It is an experimental feature. The number of decimal places allowed for a credit batch is determined by the original number of decimal places used with calling CreatBatch. SetPrecision allows the number of allowed decimal places to be increased, effectively making the supply more granular without actually changing any balances. It allows asset issuers to be able to issue an asset without needing to think about how many subdivisions are needed upfront. While it may not be relevant for credits which likely have a fairly stable market value, I wanted to experiment a bit and this serves as a proof of concept for a broader bank redesign where say for instance a coin like the ATOM or XRN could be issued in its own units rather than micro or nano-units. Instead an operation like SetPrecision would allow trading in micro, nano or pico in the future based on market demand. Arbitrary, unbounded precision is not desirable because this can lead to spam attacks (like sending 0.000000000000000000000000000001 coins). This is effectively fixed precision so under the hood it is still basically an integer, but the fixed precision can be increased so its more adaptable long term than just an integer.
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
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
