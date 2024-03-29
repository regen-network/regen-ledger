package client

import (
	"bytes"
	"encoding/json"
	"os"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client/flags"

	regentypes "github.com/regen-network/regen-ledger/types/v2"
	types "github.com/regen-network/regen-ledger/x/ecocredit/v3/marketplace/types/v1"
)

func txFlags(cmd *cobra.Command) *cobra.Command {
	flags.AddTxFlagsToCmd(cmd)
	_ = cmd.MarkFlagRequired(flags.FlagFrom)
	return cmd
}

func parseSellOrders(jsonFile string) ([]*types.MsgSell_Order, error) {
	bz, err := os.ReadFile(jsonFile)
	if err != nil {
		return nil, err
	}

	if err := regentypes.CheckDuplicateKey(json.NewDecoder(bytes.NewReader(bz)), nil); err != nil {
		return nil, err
	}

	var orders []*types.MsgSell_Order

	// using json package because array is not a proto message
	err = json.Unmarshal(bz, &orders)
	if err != nil {
		return nil, err
	}

	return orders, nil
}

func parseSellUpdates(jsonFile string) ([]*types.MsgUpdateSellOrders_Update, error) {
	bz, err := os.ReadFile(jsonFile)
	if err != nil {
		return nil, err
	}

	if err := regentypes.CheckDuplicateKey(json.NewDecoder(bytes.NewReader(bz)), nil); err != nil {
		return nil, err
	}

	var updates []*types.MsgUpdateSellOrders_Update

	// using json package because array is not a proto message
	err = json.Unmarshal(bz, &updates)
	if err != nil {
		return nil, err
	}

	return updates, nil
}

func parseBuyOrders(jsonFile string) ([]*types.MsgBuyDirect_Order, error) {
	bz, err := os.ReadFile(jsonFile)
	if err != nil {
		return nil, err
	}

	if err := regentypes.CheckDuplicateKey(json.NewDecoder(bytes.NewReader(bz)), nil); err != nil {
		return nil, err
	}

	var orders []*types.MsgBuyDirect_Order

	// using json package because array is not a proto message
	err = json.Unmarshal(bz, &orders)
	if err != nil {
		return nil, err
	}

	return orders, nil
}
