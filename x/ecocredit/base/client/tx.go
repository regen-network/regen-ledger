package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"

	sdkclient "github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	regentypes "github.com/regen-network/regen-ledger/types/v2"
	types "github.com/regen-network/regen-ledger/x/ecocredit/v3/base/types/v1"
)

const (
	FlagAddIssuers             string = "add-issuers"
	FlagReason                 string = "reason"
	FlagRemoveIssuers          string = "remove-issuers"
	FlagReferenceID            string = "reference-id"
	FlagRetirementJurisdiction string = "retirement-jurisdiction"
	FlagRetirementReason       string = "retirement-reason"
	FlagClassFee               string = "class-fee"
)

// TxCreateClassCmd returns a transaction command that creates a credit class.
func TxCreateClassCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-class [issuers] [credit-type-abbrev] [metadata] [flags]",
		Short: "Creates a new credit class with transaction author (--from) as admin",
		Long: fmt.Sprintf(`Creates a new credit class with transaction author (--from) as admin.

The transaction author must have permission to create a new credit class by
being on the list of %s. This is a governance parameter, which
can be queried via the command line.

They must also pay a fee (separate from the transaction fee) to create a credit
class. The list of accepted fees is defined by the %s parameter.

Parameters:

- issuers:    	       comma separated (no spaces) list of issuer account addresses
- credit-type-abbrev:  the abbreviation of a credit type
- metadata:            arbitrary data attached to the credit class info

Flags:

- class-fee: the fee that the class creator will pay to create the credit class. It must be >= the
required credit_class_fee param. If the credit_class_fee param is empty, no fee is required. 
We explicitly include the class creation fee here so that the class creator acknowledges paying 
the fee and is not surprised to learn that the they paid a fee without consent.

`,
			types.KeyAllowedClassCreators,
			types.KeyCreditClassFee,
		),
		Example: `regen tx ecocredit create-class regen1elq7ys34gpkj3jyvqee0h6yk4h9wsfxmgqelsw C regen:13toVgf5UjYBz6J29x28pLQyjKz5FpcW3f4bT5uRKGxGREWGKjEdXYG.rdf --class-fee 20000000uregen
regen tx ecocredit create-class regen1elq7ys34gpkj3jyvqee0h6yk4h9wsfxmgqelsw,regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6 C regen:13toVgf5UjYBz6J29x28pLQyjKz5FpcW3f4bT5uRKGxGREWGKjEdXYG.rdf --class-fee 20000000uregen`,
		Args: cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := sdkclient.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			// Get the class admin from the --from flag
			admin := clientCtx.GetFromAddress()

			// parse the comma-separated list of issuers
			issuers := strings.Split(args[0], ",")
			for i := range issuers {
				issuers[i] = strings.TrimSpace(issuers[i])
			}

			msg := types.MsgCreateClass{
				Admin:            admin.String(),
				Issuers:          issuers,
				Metadata:         args[2],
				CreditTypeAbbrev: args[1],
			}

			// parse and normalize credit class fee
			feeString, err := cmd.Flags().GetString(FlagClassFee)
			if err != nil {
				return err
			}
			if feeString != "" {
				fee, err := sdk.ParseCoinNormalized(feeString)
				if err != nil {
					return fmt.Errorf("failed to parse class-fee: %w", err)
				}

				msg.Fee = &fee
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	cmd.Flags().String(FlagClassFee, "", "the fee that the class creator will pay to create the credit class (e.g. \"20regen\")")

	return txFlags(cmd)
}

// TxCreateProjectCmd returns a transaction command that creates a new project.
func TxCreateProjectCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-project [class-id] [jurisdiction] [metadata] [flags]",
		Short: "Create a new project within a credit class",
		Long: `Create a new project within a credit class.
		
Parameters:

- class-id:      the ID of the credit class
- jurisdiction:  the jurisdiction of the project
- metadata:      any arbitrary metadata to attach to the project

Optional Flags:

- reference-id: a reference ID for the project`,
		Example: `regen tx ecocredit create-project C01 "US-WA 98225" regen:13toVgf5UjYBz6J29x28pLQyjKz5FpcW3f4bT5uRKGxGREWGKjEdXYG.rdf
regen tx ecocredit create-project C01 "US-WA 98225" regen:13toVgf5UjYBz6J29x28pLQyjKz5FpcW3f4bT5uRKGxGREWGKjEdXYG.rdf  --reference-id VCS-001`,
		Args: cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := sdkclient.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			referenceID, err := cmd.Flags().GetString(FlagReferenceID)
			if err != nil {
				return err
			}

			msg := types.MsgCreateProject{
				Admin:        clientCtx.GetFromAddress().String(),
				ClassId:      args[0],
				Jurisdiction: args[1],
				Metadata:     args[2],
				ReferenceId:  referenceID,
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)

		},
	}

	cmd.Flags().String(FlagReferenceID, "", "a reference ID for the project")

	return txFlags(cmd)
}

// TxGenBatchJSONCmd returns a transaction command that generates JSON to
// represent a new credit batch.
func TxGenBatchJSONCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   `gen-batch-json [issuer] [project-id] [issuance-count] [metadata] [start-date] [end-date]`,
		Short: "Generates JSON to represent a new credit batch for use with create-batch command",
		Long: `Generates JSON to represent a new credit batch for use with create-batch command.

Parameters:

- issuer:          the account address of the credit batch issuer
- project-id:      the ID of the project
- issuance-count:  the number of issuance items to generate
- metadata:        any arbitrary metadata to attach to the credit batch
- start-date:      the beginning of the period during which this credit batch was
                   quantified and verified (format: yyyy-mm-dd)
- end-date:        the end of the period during which this credit batch was
                   quantified and verified (format: yyyy-mm-dd)`,
		Example: "regen tx ecocredit gen-batch-json regen1elq7ys34gpkj3jyvqee0h6yk4h9wsfxmgqelsw C01-001 2 regen:13toVgf5UjYBz6J29x28pLQyjKz5FpcW3f4bT5uRKGxGREWGKjEdXYG.rdf 2020-01-01 2021-01-01",
		Args:    cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) error {
			issuanceCount, err := strconv.ParseUint(args[2], 10, 8)
			if err != nil {
				return err
			}

			issuance := make([]*types.BatchIssuance, issuanceCount)
			for i := range issuance {
				issuance[i] = &types.BatchIssuance{}
			}

			startDate, err := regentypes.ParseDate("start date", args[4])
			if err != nil {
				return err
			}

			endDate, err := regentypes.ParseDate("end date", args[5])
			if err != nil {
				return err
			}

			msg := &types.MsgCreateBatch{
				Issuer:    args[0],
				ProjectId: args[1],
				Issuance:  issuance,
				Metadata:  args[3],
				StartDate: &startDate,
				EndDate:   &endDate,
			}

			// Marshal and output JSON of message
			ctx := sdkclient.GetClientContextFromCmd(cmd)
			msgJSON, err := ctx.Codec.MarshalJSON(msg)
			if err != nil {
				return err
			}

			var formattedJSON bytes.Buffer
			err = json.Indent(&formattedJSON, msgJSON, "", "  ")
			if err != nil {
				return err
			}
			fmt.Println(formattedJSON.String())

			return nil
		},
	}

	return cmd
}

// TxCreateBatchCmd returns a transaction command that creates a credit batch.
func TxCreateBatchCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-batch [batch-json]",
		Short: "Issues a new credit batch",
		Long: `Issues a new credit batch.

Parameters:

- batch-json:  path to JSON file containing credit batch information`,
		Example: `regen tx ecocredit create-batch batch.json

Example JSON:

{
  "project_id": "C01-001",
  "issuer": "regen1elq7ys34gpkj3jyvqee0h6yk4h9wsfxmgqelsw",
  "issuance": [
    {
      "recipient": "regen1elq7ys34gpkj3jyvqee0h6yk4h9wsfxmgqelsw",
      "tradable_amount": "1000"
    },
    {
      "recipient": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
      "retired_amount": "1000",
      "retirement_jurisdiction": "US-OR",
      "retirement_reason": "offsetting electricity consumption"
    }
  ],
  "metadata": "regen:13toVgf5UjYBz6J29x28pLQyjKz5FpcW3f4bT5uRKGxGREWGKjEdXYG.rdf",
  "start_date": "2020-01-01T00:00:00Z",
  "end_date": "2021-01-01T00:00:00Z"
}`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := sdkclient.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			// parse the JSON file representing the request
			msg, err := parseMsgCreateBatch(clientCtx, args[0])
			if err != nil {
				return sdkerrors.ErrInvalidRequest.Wrapf("failed to parse json: %s", err)
			}

			// Get the batch issuer from the --from flag
			msg.Issuer = clientCtx.GetFromAddress().String()

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	return txFlags(cmd)
}

// TxSendCmd returns a transaction command that sends credits from a single batch from one account
// to another.
func TxSendCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "send [amount] [batch-denom] [recipient] [flags]",
		Short: "Sends credits from a single batch from the transaction author (--from) to the recipient",
		Long: `Sends credits from a single batch from the transaction author (--from) to the recipient.

By default, the credits will be sent as tradable. Use the --retirement-jurisdiction flag to retire the credits to the recipient address.

Parameters:

- amount:       the amount of credits to send
- batch-denom:  the denomination of the credit batch
- recipient:    the recipient account address`,
		Example: "regen tx ecocredit send 20 C01-001-20200101-20210101-001 regen18xvpj53vaupyfejpws5sktv5lnas5xj2phm3cf --from myKey",
		Args:    cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := sdkclient.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			retireTo, err := cmd.Flags().GetString(FlagRetirementJurisdiction)
			if err != nil {
				return err
			}

			tradableAmount := args[0]

			retiredAmount := "0"
			if len(retireTo) > 0 {
				tradableAmount = "0"
				retiredAmount = args[0]
			}

			reason, err := cmd.Flags().GetString(FlagRetirementReason)
			if err != nil {
				return err
			}

			credit := types.MsgSend_SendCredits{
				TradableAmount:         tradableAmount,
				BatchDenom:             args[1],
				RetiredAmount:          retiredAmount,
				RetirementJurisdiction: retireTo,
				RetirementReason:       reason,
			}

			msg := types.MsgSend{
				Sender:    clientCtx.GetFromAddress().String(),
				Recipient: args[2],
				Credits:   []*types.MsgSend_SendCredits{&credit},
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	cmd.Flags().String(FlagRetirementJurisdiction, "", "Jurisdiction to retire the credits to. If empty, credits are not retired. (default empty)")
	cmd.Flags().String(FlagRetirementReason, "", "the reason for retiring the credits (optional)")

	return txFlags(cmd)
}

// TxSendBulkCmd returns a transaction command that can send credits from multiple batches from one account
// to another
func TxSendBulkCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "send-bulk [recipient] [credits-json]",
		Short: "Sends credits from multiple batches from the transaction author (--from) to the recipient",
		Long: `Sends credits from multiple batches from the transaction author (--from) to the recipient.

Parameters:

- recipient:     the recipient account address
- credits-json:  path to JSON file containing credits to send`,
		Example: `regen tx ecocredit send-bulk regen18xvpj53vaupyfejpws5sktv5lnas5xj2phm3cf credits.json

Example JSON:

[
  {
    "batch_denom": "C01-001-20200101-20210101-001",
    "tradable_amount": "500"
  },
  {
    "batch_denom": "C01-001-20200101-20210101-002",
    "tradable_amount": "50",
    "retired_amount": "100",
    "retirement_jurisdiction": "YY-ZZ 12345",
    "retirement_reason": "offsetting electricity consumption"
  }
]`,
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := sdkclient.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			// parse the JSON file representing the credits
			credits, err := parseSendCredits(args[1])
			if err != nil {
				return sdkerrors.ErrInvalidRequest.Wrapf("failed to parse json: %s", err)
			}

			msg := types.MsgSend{
				Sender:    clientCtx.GetFromAddress().String(),
				Recipient: args[0],
				Credits:   credits,
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	return txFlags(cmd)
}

// TxRetireCmd returns a transaction command that retires credits.
func TxRetireCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "retire [credits] [retirement-jurisdiction]",
		Short: "Retires a specified amount of credits from the account of the transaction author (--from)",
		Long: `Retires a specified amount of credits from the account of the transaction author (--from)

Parameters:

- credits:                  path to JSON file containing credits to retire
- retirement-jurisdiction:  the jurisdiction in which the credit will be retired`,
		Example: `
regen tx ecocredit retire credits.json "US-WA 98225"

Example JSON:
[
  {
    "batch_denom": "C01-001-20200101-20210101-001",
    "amount": "5"
  },
  {
    "batch_denom": "C01-001-20200101-20210101-002",
    "amount": "10"
  }
]`,
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := sdkclient.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			// parse the JSON file representing the credits
			credits, err := parseCredits(args[0])
			if err != nil {
				return sdkerrors.ErrInvalidRequest.Wrapf("failed to parse json: %s", err)
			}

			reason, err := cmd.Flags().GetString(FlagReason)
			if err != nil {
				return err
			}

			msg := types.MsgRetire{
				Owner:        clientCtx.GetFromAddress().String(),
				Credits:      credits,
				Jurisdiction: args[1],
				Reason:       reason,
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	cmd.Flags().String(FlagReason, "", "the reason for retiring the credits (optional)")

	return txFlags(cmd)
}

// TxCancelCmd returns a transaction command that cancels credits.
func TxCancelCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cancel [credits-json] [reason]",
		Short: "Cancels a specified amount of credits from the account of the transaction author (--from)",
		Long: `Cancels a specified amount of credits from the account of the transaction author (--from)

Parameters:

- credits-json:  path to JSON file containing credits to retire
- reason:        any arbitrary string that specifies the reason for cancelling credits`,
		Example: `regen tx ecocredit cancel credits.json "transferring credits to another registry"

Example JSON:

[
  {
    "batch_denom": "C01-001-20200101-20210101-001",
    "amount": "5"
  },
  {
    "batch_denom": "C01-001-20200101-20210101-002",
    "amount": "10"
  }
]`,
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := sdkclient.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			// parse the JSON file representing the credits
			credits, err := parseCredits(args[0])
			if err != nil {
				return sdkerrors.ErrInvalidRequest.Wrapf("failed to parse json: %s", err)
			}

			msg := types.MsgCancel{
				Owner:   clientCtx.GetFromAddress().String(),
				Credits: credits,
				Reason:  args[1],
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	return txFlags(cmd)
}

// TxUpdateClassMetadataCmd returns a transaction command that updates class metadata.
func TxUpdateClassMetadataCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-class-metadata [class-id] [new-metadata]",
		Short: "Updates the metadata for a specific credit class",
		Long: `Updates the metadata for a specific credit class.

The '--from' flag must equal the credit class admin.

Parameters:

- class-id:      the class id that corresponds with the credit class you want to update
- new-metadata:  any arbitrary metadata to attach to the credit class`,
		Args:    cobra.ExactArgs(2),
		Example: `regen tx ecocredit update-class-metadata C01 regen:13toVgf5UjYBz6J29x28pLQyjKz5FpcW3f4bT5uRKGxGREWGKjEdXYG.rdf`,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := sdkclient.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.MsgUpdateClassMetadata{
				Admin:       clientCtx.GetFromAddress().String(),
				ClassId:     args[0],
				NewMetadata: args[1],
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	return txFlags(cmd)
}

// TxUpdateClassAdminCmd returns a transaction command that updates class admin.
func TxUpdateClassAdminCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-class-admin [class-id] [new-admin]",
		Short: "Updates the admin for a specific credit class",
		Long: `Updates the admin for a specific credit class.

The '--from' flag must equal the current credit class admin.

WARNING: Updating the admin replaces the current admin. Be sure the new admin account address is entered correctly.

Parameters:

- class-id:   the ID of the credit class to update
- new-admin:  the new admin account address that will overwrite the current admin account address`,
		Example: "regen tx ecocredit update-class-admin C01 regen1elq7ys34gpkj3jyvqee0h6yk4h9wsfxmgqelsw",
		Args:    cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := sdkclient.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.MsgUpdateClassAdmin{
				Admin:    clientCtx.GetFromAddress().String(),
				ClassId:  args[0],
				NewAdmin: args[1],
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	return txFlags(cmd)
}

// TxUpdateClassIssuersCmd returns a transaction command that updates class issuers.
func TxUpdateClassIssuersCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-class-issuers [class-id]",
		Short: "Update the list of issuers for a specific credit class",
		Long: `Update the list of issuers for a specific credit class.

The '--from' flag must equal the current credit class admin.

Parameters:

- class-id:  the ID of the credit class to update

Flags:

- add-issuers:     the new list of issuers to add to the class issuers list
- remove-issuers:  the new list of issuers to remove from the class issuers list`,
		Example: `regen tx ecocredit update-class-issuers C01 --add-issuers addr1,addr2,addr3
regen tx ecocredit update-class-issuers C01 --add-issuers addr1,addr2 --remove-issuers addr3,addr4`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {

			// parse add issuers
			addIssuers, err := cmd.Flags().GetStringSlice(FlagAddIssuers)
			if err != nil {
				return err
			}
			for i := range addIssuers {
				addIssuers[i] = strings.TrimSpace(addIssuers[i])
			}

			// parse remove issuers
			removeIssuers, err := cmd.Flags().GetStringSlice(FlagRemoveIssuers)
			if err != nil {
				return err
			}
			for i := range removeIssuers {
				removeIssuers[i] = strings.TrimSpace(removeIssuers[i])
			}

			clientCtx, err := sdkclient.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.MsgUpdateClassIssuers{
				Admin:         clientCtx.GetFromAddress().String(),
				ClassId:       args[0],
				AddIssuers:    addIssuers,
				RemoveIssuers: removeIssuers,
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	cmd.Flags().StringSlice(FlagAddIssuers, []string{}, "comma separated (no spaces) list of addresses")
	cmd.Flags().StringSlice(FlagRemoveIssuers, []string{}, "comma separated (no spaces) list of addresses")

	return txFlags(cmd)
}

// TxUpdateProjectAdminCmd returns a transaction command that updates project admin.
func TxUpdateProjectAdminCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-project-admin [project-id] [new-admin] [flags]",
		Short: "Update the project admin address",
		Long: `Update the project admin to the provided new_admin_address.

The '--from' flag must equal the current credit class admin.

WARNING: Updating the admin replaces the current admin. Be sure the new admin account address is entered correctly.

Parameters:

- project-id:  the ID of the project to update
- new-admin:   the new admin account address that will overwrite the current admin account address`,
		Example: "regen tx ecocredit update-project-admin C01-001 regen1ynugxwpp4lfpy0epvfqwqkpuzkz62htnex3op",
		Args:    cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := sdkclient.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.MsgUpdateProjectAdmin{
				Admin:     clientCtx.GetFromAddress().String(),
				NewAdmin:  args[1],
				ProjectId: args[0],
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	return txFlags(cmd)
}

// TxUpdateProjectMetadataCmd returns a transaction command that updates project metadata.
func TxUpdateProjectMetadataCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "update-project-metadata [project-id] [new-metadata]",
		Short:   "Update the project metadata",
		Long:    "Update the project metadata, overwriting the project's current metadata.",
		Example: "regen tx ecocredit update-project-metadata C01-001 regen:13toVgf5UjYBz6J29x28pLQyjKz5FpcW3f4bT5uRKGxGREWGKjEdXYG.rdf",
		Args:    cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := sdkclient.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.MsgUpdateProjectMetadata{
				Admin:       clientCtx.GetFromAddress().String(),
				NewMetadata: args[1],
				ProjectId:   args[0],
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	return txFlags(cmd)
}

// TxUpdateBatchMetadataCmd returns a transaction command that updates batch metadata.
func TxUpdateBatchMetadataCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-batch-metadata [batch-denom] [new-metadata]",
		Short: "Updates the metadata for a specific credit batch",
		Long: `Updates the metadata for a specific credit batch.

The '--from' flag must equal the credit batch issuer.

Parameters:

- batch-denom:   the batch denom of the credit batch to be updated
- new-metadata:  any arbitrary metadata to attach to the credit batch`,
		Args:    cobra.ExactArgs(2),
		Example: `regen tx ecocredit update-batch-metadata C01-001-20200101-20210101-001 regen:13toVgf5UjYBz6J29x28pLQyjKz5FpcW3f4bT5uRKGxGREWGKjEdXYG.rdf`,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := sdkclient.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.MsgUpdateBatchMetadata{
				Issuer:      clientCtx.GetFromAddress().String(),
				BatchDenom:  args[0],
				NewMetadata: args[1],
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	return txFlags(cmd)
}

// TxBridgeCmd returns a transaction command that bridges credits to another chain.
func TxBridgeCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "bridge [target] [recipient] [credits-json]",
		Short: "Bridge credits to another chain",
		Long: `Bridge credits to another chain.

The '--from' flag must equal the owner of the credits.

Parameters:

- target:       the target chain (e.g. "polygon")
- recipient:    the address of the recipient on the other chain
- credits-json: path to JSON file containing credits to bridge`,
		Example: `regen tx ecocredit bridge polygon 0x0000000000000000000000000000000000000001 credits.json

Example JSON:

[
  {
    "batch_denom": "C01-001-20200101-20210101-001",
    "amount": "5"
  },
  {
    "batch_denom": "C01-001-20200101-20210101-002",
    "amount": "10"
  }
]`,
		Args: cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := sdkclient.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			// parse the JSON file representing the credits
			credits, err := parseCredits(args[2])
			if err != nil {
				return sdkerrors.ErrInvalidRequest.Wrapf("failed to parse json: %s", err)
			}

			msg := types.MsgBridge{
				Owner:     clientCtx.GetFromAddress().String(),
				Target:    args[0],
				Recipient: args[1],
				Credits:   credits,
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	return txFlags(cmd)
}
