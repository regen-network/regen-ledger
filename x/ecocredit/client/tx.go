package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/spf13/cobra"
	"sigs.k8s.io/yaml"

	sdkclient "github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/regen-network/regen-ledger/types"
	basketcli "github.com/regen-network/regen-ledger/x/ecocredit/client/basket"
	marketplacecli "github.com/regen-network/regen-ledger/x/ecocredit/client/marketplace"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

// TxCmd returns a root CLI command handler for all x/ecocredit transaction commands.
func TxCmd(name string) *cobra.Command {
	cmd := &cobra.Command{
		SuggestionsMinimumDistance: 2,
		DisableFlagParsing:         true,

		Use:   name,
		Short: "Ecocredit module transactions",
		RunE:  sdkclient.ValidateCmd,
	}
	cmd.AddCommand(
		TxCreateClassCmd(),
		TxGenBatchJSONCmd(),
		TxCreateBatchCmd(),
		TxSendCmd(),
		TxRetireCmd(),
		TxCancelCmd(),
		TxUpdateClassMetadataCmd(),
		TxUpdateClassIssuersCmd(),
		TxUpdateClassAdminCmd(),
		TxCreateProject(),
		TxUpdateProjectAdminCmd(),
		TxUpdateProjectMetadataCmd(),
		basketcli.TxCreateBasket(),
		basketcli.TxPutInBasket(),
		basketcli.TxTakeFromBasket(),
		marketplacecli.TxSellCmd(),
		marketplacecli.TxUpdateSellOrdersCmd(),
		marketplacecli.TxBuyDirect(),
		marketplacecli.TxBuyDirectBatch(),
	)
	return cmd
}

func txFlags(cmd *cobra.Command) *cobra.Command {
	flags.AddTxFlagsToCmd(cmd)
	cmd.MarkFlagRequired(flags.FlagFrom)
	return cmd
}

// TxCreateClassCmd returns a transaction command that creates a credit class.
func TxCreateClassCmd() *cobra.Command {
	return txFlags(&cobra.Command{
		Use:   "create-class [issuer[,issuer]*] [credit type abbreviation] [metadata] [fee]",
		Short: "Creates a new credit class with transaction author (--from) as admin",
		Long: fmt.Sprintf(
			`Creates a new credit class with transaction author (--from) as admin.

The transaction author must have permission to create a new credit class by
being a member of the %s allowlist. This is a governance parameter, so can be
queried via the command line.

They must also pay the fee associated with creating a new credit class, defined
by the %s parameter, so should make sure they have enough funds to cover that.

Parameters:
  issuer:    	               comma separated (no spaces) list of issuer account addresses. Example: "addr1,addr2"
  credit type abbreviation:    the abbreviation of a credit type (e.g. "C", "BIO")
  metadata:  	               arbitrary data attached to the credit class info
  fee:                         fee to pay for the creation of the credit class (e.g. 10uatom, 10uregen)`,
			core.KeyAllowedClassCreators,
			core.KeyCreditClassFee,
		),
		Example: `
regen tx ecocredit create-class regen1el...xmgqelsw,regen2tl...xsgqdlhy C "metadata" 10uregen
		`,
		Args: cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := sdkclient.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			// Get the class admin from the --from flag
			admin := clientCtx.GetFromAddress()

			// Parse the comma-separated list of issuers
			issuers := strings.Split(args[0], ",")
			for i := range issuers {
				issuers[i] = strings.TrimSpace(issuers[i])
			}

			// Check credit type abbreviation is provided
			if args[1] == "" {
				return sdkerrors.ErrInvalidRequest.Wrap("credit type abbreviation is required")
			}
			creditTypeAbbrev := args[1]

			// Check that metadata is provided
			if args[2] == "" {
				return errors.New("metadata is required")
			}

			fee, err := sdk.ParseCoinNormalized(args[3])
			if err != nil {
				return err
			}

			msg := core.MsgCreateClass{
				Admin:            admin.String(),
				Issuers:          issuers,
				Metadata:         args[2],
				CreditTypeAbbrev: creditTypeAbbrev,
				Fee:              &fee,
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	})
}

const (
	FlagClassId             string = "class-id"
	FlagProjectId           string = "project-id"
	FlagIssuances           string = "issuances"
	FlagStartDate           string = "start-date"
	FlagEndDate             string = "end-date"
	FlagProjectJurisdiction string = "project-jurisdiction"
	FlagMetadata            string = "metadata"
	FlagAddIssuers          string = "add-issuers"
	FlagRemoveIssuers       string = "remove-issuers"
	FlagReferenceId         string = "reference-id"
	FlagIssuer              string = "issuer"
)

// TxGenBatchJSONCmd returns a transaction command that generates JSON to
// represent a new credit batch.
func TxGenBatchJSONCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: `gen-batch-json --project-id [project_id] --issuances [issuances] --start-date [start_date] --end-date [end_date] 
		--metadata [metadata] --issuer [issuer]`,
		Short: "Generates JSON to represent a new credit batch for use with create-batch command",
		Long: `Generates JSON to represent a new credit batch for use with create-batch command.

Required Flags:
  project_id: id of the project
  issuances:  the amount of issuances to generate
  start-date: The beginning of the period during which this credit batch was
              quantified and verified. Format: yyyy-mm-dd.
  end-date:   The end of the period during which this credit batch was
              quantified and verified. Format: yyyy-mm-dd.
  metadata:   issuance metadata
  issuer:     account address of the batch issuer
			  `,
		Args: cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			projectId, err := cmd.Flags().GetString(FlagProjectId)
			if err != nil {
				return err
			}

			templateIssuance := &core.BatchIssuance{
				Recipient:              "recipient-address",
				TradableAmount:         "tradable-amount",
				RetiredAmount:          "retired-amount",
				RetirementJurisdiction: "retirement-jurisdiction",
			}

			numIssuances, err := cmd.Flags().GetUint32(FlagIssuances)
			if err != nil {
				return err
			}

			issuances := make([]*core.BatchIssuance, numIssuances)
			for i := range issuances {
				issuances[i] = templateIssuance
			}

			startDateStr, err := cmd.Flags().GetString(FlagStartDate)
			if err != nil {
				return err
			}
			startDate, err := types.ParseDate("start_date", startDateStr)
			if err != nil {
				return err
			}

			endDateStr, err := cmd.Flags().GetString(FlagEndDate)
			if err != nil {
				return err
			}
			endDate, err := types.ParseDate("end_date", endDateStr)
			if err != nil {
				return err
			}

			metadata, err := cmd.Flags().GetString(FlagMetadata)
			if err != nil {
				return err
			}

			issuer, err := cmd.Flags().GetString(FlagIssuer)
			if err != nil {
				return err
			}

			msg := &core.MsgCreateBatch{
				ProjectId: projectId,
				Issuance:  issuances,
				Metadata:  metadata,
				StartDate: &startDate,
				EndDate:   &endDate,
				Issuer:    issuer,
			}

			// Marshal and output JSON of message
			ctx := sdkclient.GetClientContextFromCmd(cmd)
			msgJson, err := ctx.Codec.MarshalJSON(msg)
			if err != nil {
				return err
			}

			var formattedJson bytes.Buffer
			json.Indent(&formattedJson, msgJson, "", "    ")
			fmt.Println(formattedJson.String())

			return nil
		},
	}
	cmd.Flags().String(FlagProjectId, "", "project id")
	cmd.MarkFlagRequired(FlagProjectId)
	cmd.Flags().Uint32(FlagIssuances, 0, "The number of template issuances to generate")
	cmd.MarkFlagRequired(FlagIssuances)
	cmd.Flags().String(FlagStartDate, "", "The beginning of the period during which this credit batch was quantified and verified. Format: yyyy-mm-dd.")
	cmd.MarkFlagRequired(FlagStartDate)
	cmd.Flags().String(FlagEndDate, "", "The end of the period during which this credit batch was quantified and verified. Format: yyyy-mm-dd.")
	cmd.MarkFlagRequired(FlagEndDate)
	cmd.Flags().String(FlagMetadata, "", "arbitrary string attached to the credit batch")
	cmd.Flags().String(FlagIssuer, "", "account address of the batch issuer")
	cmd.MarkFlagRequired(FlagIssuer)
	return cmd
}

// TxCreateBatchCmd returns a transaction command that creates a credit batch.
func TxCreateBatchCmd() *cobra.Command {

	return txFlags(&cobra.Command{
		Use:   "create-batch [msg-create-batch-json-file]",
		Short: "Issues a new credit batch",
		Long: `Issues a new credit batch.

Parameters:
  msg-create-batch-json-file: Path to a file containing a JSON object
                              representing MsgCreateBatch. The JSON has format:
                              {
                                "project_id": "C01-001",
                                "issuance": [
                                  {
                                    "recipient":           "regen1elq7ys34gpkj3jyvqee0h6yk4h9wsfxmgqelsw",
                                    "tradable_amount":     "1000",
                                    "retired_amount":      "15",
                                    "retirement_jurisdiction": "ST-UVW XY Z12",
                                  },
                                ],
                                "metadata":         "metadata",
                                "start_date":       "1990-01-01",
                                "end_date":         "1995-10-31",
                                "project_jurisdiction": "AB-CDE FG1 345",
                                "issuer": "regen1elq7ys34gpkj3jyvqee0h6yk4h9wsfxmgqelsw"
                              }
                              `,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := sdkclient.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			contents, err := ioutil.ReadFile(args[0])
			if err != nil {
				return err
			}

			if err := checkDuplicateKey(json.NewDecoder(bytes.NewReader(contents)), nil); err != nil {
				return err
			}

			// Parse the JSON file representing the request
			msg, err := parseMsgCreateBatch(clientCtx, args[0])
			if err != nil {
				return sdkerrors.ErrInvalidRequest.Wrapf("parsing batch JSON:\n%s", err.Error())
			}

			// Get the batch issuer from the --from flag
			issuer := clientCtx.GetFromAddress()
			msg.Issuer = issuer.String()

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	})
}

// TxSendCmd returns a transaction command that sends credits from one account
// to another.
func TxSendCmd() *cobra.Command {
	return txFlags(&cobra.Command{
		Use:   "send [recipient] [credits]",
		Short: "Sends credits from the transaction author (--from) to the recipient",
		Long: `Sends credits from the transaction author (--from) to the recipient.

Parameters:
  recipient: recipient address
  credits:   YAML encoded credit list. Note: numerical values must be written in strings.
             eg: '[{batch_denom: "C01-001-20210101-20210201-001", tradable_amount: "5", retired_amount: "0", retirement_jurisdiction: "YY-ZZ 12345"}]'
             Note: "retirement_jurisdiction" is only required when "retired_amount" is positive.`,
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			var credits = []*core.MsgSend_SendCredits{}
			if err := yaml.Unmarshal([]byte(args[1]), &credits); err != nil {
				return err
			}
			clientCtx, err := sdkclient.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			msg := core.MsgSend{
				Sender:    clientCtx.GetFromAddress().String(),
				Recipient: args[0], Credits: credits,
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	})
}

// TxRetireCmd returns a transaction command that retires credits.
func TxRetireCmd() *cobra.Command {
	return txFlags(&cobra.Command{
		Use:   "retire [credits] [retirement_jurisdiction]",
		Short: "Retires a specified amount of credits from the account of the transaction author (--from)",
		Long: `Retires a specified amount of credits from the account of the transaction author (--from)

Parameters:
  credits:             YAML encoded credit list. Note: numerical values must be written in strings.
                       eg: '[{batch_denom: "C01-20210101-20210201-001", amount: "5"}]'
  retirement_jurisdiction: A string representing the jurisdiction of the buyer or
                       beneficiary of retired credits. It has the form
                       <country-code>[-<region-code>[ <postal-code>]], where
                       country-code and region-code are taken from ISO 3166, and
                       postal-code being up to 64 alphanumeric characters.
                       eg: 'AA-BB 12345'`,
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			var credits []*core.Credits
			if err := yaml.Unmarshal([]byte(args[0]), &credits); err != nil {
				return err
			}
			clientCtx, err := sdkclient.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			msg := core.MsgRetire{
				Owner:        clientCtx.GetFromAddress().String(),
				Credits:      credits,
				Jurisdiction: args[1],
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	})
}

// TxCancelCmd returns a transaction command that cancels credits.
func TxCancelCmd() *cobra.Command {
	return txFlags(&cobra.Command{
		Use:   "cancel [credits] [reason]",
		Short: "Cancels a specified amount of credits from the account of the transaction author (--from)",
		Long: `Cancels a specified amount of credits from the account of the transaction author (--from)

Parameters:
  credits:  comma-separated list of credits in the form [<amount> <batch-denom>]
            eg: '10 C01-001-20200101-20210101-001, 0.1 C01-001-20200101-20210101-001'
  reason:   reason is any arbitrary string that specifies the reason for cancelling credits.
`,
		Example: `
regen tx ecocredit cancel '10 C01-001-20200101-20210101-001,0.1 C01-001-20200101-20210101-001' "bridging assets to another chain"
		`,
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			credits, err := parseCancelCreditsList(args[0])
			if err != nil {
				return err
			}
			clientCtx, err := sdkclient.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			msg := core.MsgCancel{
				Owner:   clientCtx.GetFromAddress().String(),
				Credits: credits,
				Reason:  args[1],
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	})
}

func TxUpdateClassMetadataCmd() *cobra.Command {
	return txFlags(&cobra.Command{
		Use:   "update-class-metadata [class-id] [metadata]",
		Short: "Updates the metadata for a specific credit class",
		Long: `Updates the metadata for a specific credit class. the '--from' flag must equal the credit class admin.

Parameters:
  class-id:  the class id that corresponds with the credit class you want to update
  metadata:  credit class metadata`,
		Args: cobra.ExactArgs(2),
		Example: `
regen tx ecocredit update-class-metadata C01 "some metadata"
		`,
		RunE: func(cmd *cobra.Command, args []string) error {

			if args[0] == "" {
				return errors.New("class-id is required")
			}
			classID := args[0]

			// Check that metadata is provided and decode it
			if args[1] == "" {
				return errors.New("metadata is required")
			}

			clientCtx, err := sdkclient.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := core.MsgUpdateClassMetadata{
				Admin:       clientCtx.GetFromAddress().String(),
				ClassId:     classID,
				NewMetadata: args[1],
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	})
}

func TxUpdateClassAdminCmd() *cobra.Command {
	return txFlags(&cobra.Command{
		Use:   "update-class-admin [class-id] [admin]",
		Short: "Updates the admin for a specific credit class",
		Long: `Updates the admin for a specific credit class. the '--from' flag must equal the current credit class admin.
               WARNING: Updating the admin replaces the current admin. Be sure the address entered is correct.

Parameters:
  class-id:  the class id that corresponds with the credit class you want to update
  new-admin: the address to overwrite the current admin address`,
		Example: `
regen tx ecocredit update-class-admin C01 regen1elq7ys34gpkj3jyvqee0h6yk4h9wsfxmgqelsw 
  `,
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {

			if args[0] == "" {
				return errors.New("class-id is required")
			}
			classID := args[0]

			// check for the address
			newAdmin := args[1]
			if newAdmin == "" {
				return errors.New("new admin address is required")
			}

			clientCtx, err := sdkclient.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := core.MsgUpdateClassAdmin{
				Admin:    clientCtx.GetFromAddress().String(),
				ClassId:  classID,
				NewAdmin: newAdmin,
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	})
}

func TxUpdateClassIssuersCmd() *cobra.Command {
	cmd := txFlags(&cobra.Command{
		Use:   "update-class-issuers [class-id]",
		Short: "Update the list of issuers for a specific credit class",
		Long: `Update the list of issuers for a specific credit class. the '--from' flag must equal the current credit class admin.

Parameters:
  class-id:  the class id that corresponds with the credit class you want to update
Flags:  
  add-issuers:    the new list of issuers to add to the class issuers list
  remove-issuers: the new list of issuers to remove from the class issuers list`,
		Example: `
regen tx ecocredit update-class-issuers C01 --add-issuers addr1,addr2,addr3
regen tx ecocredit update-class-issuers C01 --add-issuers addr1,addr2 --remove-issuers addr3,addr4
	`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {

			if args[0] == "" {
				return errors.New("class-id is required")
			}

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

			msg := core.MsgUpdateClassIssuers{
				Admin:         clientCtx.GetFromAddress().String(),
				ClassId:       args[0],
				AddIssuers:    addIssuers,
				RemoveIssuers: removeIssuers,
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	})
	cmd.Flags().StringSlice(FlagAddIssuers, []string{}, "comma separated (no spaces) list of addresses")
	cmd.Flags().StringSlice(FlagRemoveIssuers, []string{}, "comma separated (no spaces) list of addresses")

	return cmd
}

// TxCreateProject returns a transaction command that creates a new project.
func TxCreateProject() *cobra.Command {
	cmd := txFlags(&cobra.Command{
		Use:   "create-project [class-id] [project-jurisdiction] [metadata]",
		Short: "Create a new project within a credit class",
		Long: `Create a new project within a credit class.
		
		Parameters:
		class-id: id of the class
		project-jurisdiction: the jurisdiction of the project (see documentation for proper project-jurisdiction formats).
		metadata: any arbitrary metadata attached to the project.
		Flags:
		reference-id: project reference id
		`,
		Example: `
regen tx ecocredit create-project C01 "AA-BB 12345" metadata
regen tx ecocredit create-project C01 "AA-BB 12345" metadata --reference-id R01
		`,
		Args: cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			if args[0] == "" {
				return errors.New("class-id is required")
			}

			classID := args[0]

			if args[1] == "" {
				return errors.New("project jurisdiction is required")
			}

			projectJurisdiction := args[1]
			if err := core.ValidateJurisdiction(projectJurisdiction); err != nil {
				return err
			}

			if args[2] == "" {
				return errors.New("metadata is required")
			}

			clientCtx, err := sdkclient.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := core.MsgCreateProject{
				Admin:        clientCtx.GetFromAddress().String(),
				ClassId:      classID,
				Jurisdiction: projectJurisdiction,
				Metadata:     args[2],
			}

			referenceId, err := cmd.Flags().GetString(FlagReferenceId)
			if err != nil {
				return err
			}
			referenceId = strings.TrimSpace(referenceId)
			msg.ReferenceId = referenceId

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)

		},
	})

	cmd.Flags().String(FlagReferenceId, "", "project reference id")

	return cmd
}

func TxUpdateProjectAdminCmd() *cobra.Command {
	return txFlags(&cobra.Command{
		Use:   "update-project-admin [project_id] [new_admin_address] [flags]",
		Short: "Update the project admin address",
		Long: "Update the project admin to the provided new_admin_address. Please double check the address as " +
			"this will forfeit control of the project. Passing an invalid address could cause loss of control of the project.",
		Example: "regen tx ecocredit update-project-admin VERRA1 regen1ynugxwpp4lfpy0epvfqwqkpuzkz62htnex3op",
		Args:    cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := sdkclient.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			projectId, newAdmin := args[0], args[1]
			msg := core.MsgUpdateProjectAdmin{
				Admin:     clientCtx.GetFromAddress().String(),
				NewAdmin:  newAdmin,
				ProjectId: projectId,
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	})
}

func TxUpdateProjectMetadataCmd() *cobra.Command {
	return txFlags(&cobra.Command{
		Use:     "update-project-metadata [project-id] [new_metadata]",
		Short:   "Update the project metadata",
		Long:    "Update the project metadata, overwriting the project's current metadata.",
		Example: `regen tx ecocredit update-project-metadata VERRA1 "some metadata"`,
		Args:    cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := sdkclient.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			projectId, newMetadata := args[0], args[1]
			msg := core.MsgUpdateProjectMetadata{
				Admin:       clientCtx.GetFromAddress().String(),
				NewMetadata: newMetadata,
				ProjectId:   projectId,
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	})
}
