package client

import (
	"github.com/spf13/cobra"

	sdkclient "github.com/cosmos/cosmos-sdk/client"

	baseclient "github.com/regen-network/regen-ledger/x/ecocredit/base/client"
	basketcli "github.com/regen-network/regen-ledger/x/ecocredit/basket/client"
	marketclient "github.com/regen-network/regen-ledger/x/ecocredit/marketplace/client"
)

const (
	FlagAddIssuers             string = "add-issuers"
	FlagRemoveIssuers          string = "remove-issuers"
	FlagReferenceID            string = "reference-id"
	FlagRetirementJurisdiction string = "retirement-jurisdiction"
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
		baseclient.TxCreateClassCmd(),
		baseclient.TxCreateProjectCmd(),
		baseclient.TxGenBatchJSONCmd(),
		baseclient.TxCreateBatchCmd(),
		baseclient.TxSendCmd(),
		baseclient.TxSendBulkCmd(),
		baseclient.TxRetireCmd(),
		baseclient.TxCancelCmd(),
		baseclient.TxUpdateClassMetadataCmd(),
		baseclient.TxUpdateClassIssuersCmd(),
		baseclient.TxUpdateClassAdminCmd(),
		baseclient.TxUpdateProjectAdminCmd(),
		baseclient.TxUpdateProjectMetadataCmd(),
		basketcli.TxCreateBasketCmd(),
		basketcli.TxPutInBasketCmd(),
		basketcli.TxTakeFromBasketCmd(),
		marketclient.TxSellCmd(),
		marketclient.TxUpdateSellOrdersCmd(),
		marketclient.TxBuyDirectCmd(),
		marketclient.TxBuyDirectBulkCmd(),
		marketclient.TxCancelSellOrderCmd(),
	)
	return cmd
}
