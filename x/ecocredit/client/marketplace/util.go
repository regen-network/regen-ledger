package marketplace

import (
	"bytes"
	"encoding/json"
	"io/ioutil"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client/flags"

	"github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/x/ecocredit/marketplace"
)

func txFlags(cmd *cobra.Command) *cobra.Command {
	flags.AddTxFlagsToCmd(cmd)
	_ = cmd.MarkFlagRequired(flags.FlagFrom)
	return cmd
}

func parseSellOrders(jsonFile string) ([]*marketplace.MsgSell_Order, error) {
	bz, err := ioutil.ReadFile(jsonFile)
	if err != nil {
		return nil, err
	}

	if err := types.CheckDuplicateKey(json.NewDecoder(bytes.NewReader(bz)), nil); err != nil {
		return nil, err
	}

	var orders []*marketplace.MsgSell_Order

	// using json package because array is not a proto message
	err = json.Unmarshal(bz, &orders)
	if err != nil {
		return nil, err
	}

	return orders, nil
}

func parseSellUpdates(jsonFile string) ([]*marketplace.MsgUpdateSellOrders_Update, error) {
	bz, err := ioutil.ReadFile(jsonFile)
	if err != nil {
		return nil, err
	}

	if err := types.CheckDuplicateKey(json.NewDecoder(bytes.NewReader(bz)), nil); err != nil {
		return nil, err
	}

	var updates []*marketplace.MsgUpdateSellOrders_Update

	// using json package because array is not a proto message
	err = json.Unmarshal(bz, &updates)
	if err != nil {
		return nil, err
	}

	return updates, nil
}

func parseBuyOrders(jsonFile string) ([]*marketplace.MsgBuyDirect_Order, error) {
	bz, err := ioutil.ReadFile(jsonFile)
	if err != nil {
		return nil, err
	}

	if err := types.CheckDuplicateKey(json.NewDecoder(bytes.NewReader(bz)), nil); err != nil {
		return nil, err
	}

	var orders []*marketplace.MsgBuyDirect_Order

	// using json package because array is not a proto message
	err = json.Unmarshal(bz, &orders)
	if err != nil {
		return nil, err
	}

	return orders, nil
}
