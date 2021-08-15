package client

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/cosmos/cosmos-sdk/client"
	sdkclient "github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"

	"github.com/regen-network/regen-ledger/x/ecocredit"
)

// TxCmd returns a root CLI command handler for all x/ecocredit transaction commands.
func TxCmd(name string) *cobra.Command {
	cmd := &cobra.Command{
		SuggestionsMinimumDistance: 2,
		DisableFlagParsing:         true,

		Use:   name,
		Short: "Ecocredit module transactions",
		RunE:  client.ValidateCmd,
	}
	cmd.AddCommand(
		TxCreateClassCmd(),
		TxGenBatchJSONCmd(),
		TxCreateBatchCmd(),
		TxSendCmd(),
		TxRetireCmd(),
		TxCancelCmd(),
	)
	return cmd
}

func txflags(cmd *cobra.Command) *cobra.Command {
	flags.AddTxFlagsToCmd(cmd)
	cmd.MarkFlagRequired(flags.FlagFrom)
	return cmd
}

func TxCreateClassCmd() *cobra.Command {
	return txflags(&cobra.Command{
		Use:   "create-class [designer] [issuer[,issuer]*] [credit type] [metadata]",
		Short: "Creates a new credit class",
		Long: `Creates a new credit class.

Parameters:
  designer:  	    address of the account which designed the credit class
  issuer:    	    comma separated (no spaces) list of issuer account addresses. Example: "addr1,addr2"
  credit type:    the credit class type (e.g. carbon, biodiversity, etc)
  metadata:  	    base64 encoded metadata - arbitrary data attached to the credit class info`,
		Args: cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			issuers := strings.Split(args[1], ",")
			for i := range issuers {
				issuers[i] = strings.TrimSpace(issuers[i])
			}
			if args[2] == "" {
				return sdkerrors.ErrInvalidRequest.Wrap("credit type is required")
			}
			creditType := args[2]
			if args[3] == "" {
				return errors.New("base64_metadata is required")
			}
			b, err := base64.StdEncoding.DecodeString(args[3])
			if err != nil {
				return sdkerrors.ErrInvalidRequest.Wrap("metadata is malformed, proper base64 string is required")
			}

			clientCtx, err := sdkclient.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			msg := ecocredit.MsgCreateClass{
				Designer: args[0], Issuers: issuers, Metadata: b, CreditType: creditType,
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	})
}

const (
	FlagClassId         string = "class-id"
	FlagIssuances       string = "issuances"
	FlagStartDate       string = "start-date"
	FlagEndDate         string = "end-date"
	FlagProjectLocation string = "project-location"
	FlagMetadata        string = "metadata"
)

func TxGenBatchJSONCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "gen-batch-json --class-id [class_id] --issuances [issuances] --start-date [start_date] --end-date [end_date] --project-location [project_location] --metadata [metadata]",
		Short: "Generates JSON to represent a new credit batch for use with create-batch command",
		Long: `Generates JSON to represent a new credit batch for use with create-batch command.

Required Flags:
  class_id:   id of the credit class
  issuances:  the amount of issuances to generate
  start-date: The beginning of the period during which this credit batch was
              quantified and verified. Format: yyyy-mm-dd.
  end-date:   The end of the period during which this credit batch was
              quantified and verified. Format: yyyy-mm-dd.
  project-location: the location of the credit batch (see documentation for proper project-location formats).
  metadata:   base64 encoded issuance metadata`,
		Args: cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			classId, err := cmd.Flags().GetString(FlagClassId)
			if err != nil {
				return err
			}

			templateIssuance := &ecocredit.MsgCreateBatch_BatchIssuance{
				Recipient:          "recipient-address",
				TradableAmount:     "tradable-amount",
				RetiredAmount:      "retired-amount",
				RetirementLocation: "retirement-location",
			}

			numIssuances, err := cmd.Flags().GetUint32(FlagIssuances)
			issuances := make([]*ecocredit.MsgCreateBatch_BatchIssuance, numIssuances)
			for i := range issuances {
				issuances[i] = templateIssuance
			}

			startDateStr, err := cmd.Flags().GetString(FlagStartDate)
			if err != nil {
				return err
			}
			startDate, err := ParseDate("start_date", startDateStr)
			if err != nil {
				return err
			}

			endDateStr, err := cmd.Flags().GetString(FlagEndDate)
			if err != nil {
				return err
			}
			endDate, err := ParseDate("end_date", endDateStr)
			if err != nil {
				return err
			}

			projectLocation, err := cmd.Flags().GetString(FlagProjectLocation)
			if err != nil {
				return err
			}

			metadataStr, err := cmd.Flags().GetString(FlagMetadata)
			if err != nil {
				return err
			}
			b, err := base64.StdEncoding.DecodeString(metadataStr)
			if err != nil {
				return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "metadata is malformed, proper base64 string is required")
			}

			msg := &ecocredit.MsgCreateBatch{
				ClassId:         classId,
				Issuance:        issuances,
				Metadata:        b,
				StartDate:       &startDate,
				EndDate:         &endDate,
				ProjectLocation: projectLocation,
			}

			// Marshal and output JSON of message
			msgJson, err := json.MarshalIndent(msg, "", "    ")
			fmt.Print(string(msgJson))

			return nil
		},
	}
	cmd.Flags().String(FlagClassId, "", "credit class")
	cmd.MarkFlagRequired(FlagClassId)
	cmd.Flags().Uint32(FlagIssuances, 0, "The number of template issuances to generate")
	cmd.MarkFlagRequired(FlagIssuances)
	cmd.Flags().String(FlagStartDate, "", "The beginning of the period during which this credit batch was quantified and verified. Format: yyyy-mm-dd.")
	cmd.MarkFlagRequired(FlagStartDate)
	cmd.Flags().String(FlagEndDate, "", "The end of the period during which this credit batch was quantified and verified. Format: yyyy-mm-dd.")
	cmd.MarkFlagRequired(FlagEndDate)
	cmd.Flags().String(FlagProjectLocation, "", "The location of the project that is backing the credits in this batch")
	cmd.MarkFlagRequired(FlagProjectLocation)
	cmd.Flags().String(FlagMetadata, "", "base64 encoded issuance metadata")
	return cmd
}

func TxCreateBatchCmd() *cobra.Command {
	var (
		startDate = time.Unix(10000, 10000).UTC()
		endDate   = time.Unix(10000, 10050).UTC()
	)
	createExampleBatchJSON, err := json.MarshalIndent(
		ecocredit.MsgCreateBatch{
			// Leave issuer empty, because we'll use --from flag
			Issuer:  "",
			ClassId: "1BX53GF",
			Issuance: []*ecocredit.MsgCreateBatch_BatchIssuance{
				{
					Recipient:          "regen1elq7ys34gpkj3jyvqee0h6yk4h9wsfxmgqelsw",
					TradableAmount:     "1000",
					RetiredAmount:      "15",
					RetirementLocation: "ST-UVW XY Z12",
				},
			},
			Metadata:        []byte{0x1, 0x2},
			StartDate:       &startDate,
			EndDate:         &endDate,
			ProjectLocation: "AB-CDE FG1 345",
		},
		"                              ",
		"    ",
	)
	if err != nil {
		panic("Couldn't marshal MsgCreateBatch to JSON")
	}
	return txflags(&cobra.Command{
		Use:   "create-batch [msg-create-batch-json-file]",
		Short: "Issues a new credit batch",
		Long: fmt.Sprintf(`Issues a new credit batch.

Parameters:
  msg-create-batch-json-file: Path to a file containing a JSON object
			      representing MsgCreateBatch. The JSON has format:
                              %s`, createExampleBatchJSON),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := sdkclient.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			// Parse the JSON file representing the request
			msg, err := parseMsgCreateBatch(clientCtx, args[0])
			if err != nil {
				return sdkerrors.ErrInvalidRequest.Wrapf("parsing batch JSON:\n%s", err.Error())
			}

			// Get the batch issuer from the --from flag
			issuer, err := cmd.Flags().GetString(flags.FlagFrom)
			if err != nil {
				return sdkerrors.ErrInvalidRequest.Wrap(err.Error())
			}
			msg.Issuer = issuer

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	})
}

func TxSendCmd() *cobra.Command {
	return txflags(&cobra.Command{
		Use:   "send [recipient] [credits]",
		Short: "Sends credits from the transaction author (--from) to the recipient",
		Long: `Sends credits from the transaction author (--from) to the recipient.

Parameters:
  recipient: recipient address
  credits:   YAML encoded credit list. Note: numerical values must be written in strings.
             eg: '[{batch_denom: "100/2", tradable_amount: "5", retired_amount: "0", retirement_location: "YY-ZZ 12345"}]'
             Note: "retirement_location" is only required when "retired_amount" is positive.`,
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			var credits = []*ecocredit.MsgSend_SendCredits{}
			if err := yaml.Unmarshal([]byte(args[1]), &credits); err != nil {
				return err
			}
			clientCtx, err := sdkclient.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			msg := ecocredit.MsgSend{
				Sender:    clientCtx.GetFromAddress().String(),
				Recipient: args[0], Credits: credits,
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	})
}

func TxRetireCmd() *cobra.Command {
	return txflags(&cobra.Command{
		Use:   "retire [credits] [retirement_location]",
		Short: "Retires a specified amount of credits from the account of the transaction author (--from)",
		Long: `Retires a specified amount of credits from the account of the transaction author (--from)

Parameters:
  credits:             YAML encoded credit list. Note: numerical values must be written in strings.
                       eg: '[{batch_denom: "100/2", amount: "5"}]'
  retirement_location: A string representing the location of the buyer or
                       beneficiary of retired credits. It has the form
                       <country-code>[-<region-code>[ <postal-code>]], where
                       country-code and region-code are taken from ISO 3166, and
                       postal-code being up to 64 alphanumeric characters.
                       eg: 'AA-BB 12345'`,
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			var credits = []*ecocredit.MsgRetire_RetireCredits{}
			if err := yaml.Unmarshal([]byte(args[0]), &credits); err != nil {
				return err
			}
			clientCtx, err := sdkclient.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			msg := ecocredit.MsgRetire{
				Holder:   clientCtx.GetFromAddress().String(),
				Credits:  credits,
				Location: args[1],
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	})
}

func TxCancelCmd() *cobra.Command {
	return txflags(&cobra.Command{
		Use:   "cancel [credits]",
		Short: "Cancels a specified amount of credits from the account of the transaction author (--from)",
		Long: `Cancels a specified amount of credits from the account of the transaction author (--from)

Parameters:
  credits:  comma-separated list of credits in the form [<amount> <batch-denom>]
            eg: '10 C01-20200101-20210101-001, 0.1 C01-20200101-20210101-001'`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			credits, err := parseCancelCreditsList(args[0])
			if err != nil {
				return err
			}
			clientCtx, err := sdkclient.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			msg := ecocredit.MsgCancel{
				Holder:  clientCtx.GetFromAddress().String(),
				Credits: credits,
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	})
}
