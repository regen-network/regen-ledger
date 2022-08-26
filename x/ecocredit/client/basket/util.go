package basketclient

import (
	"bytes"
	"encoding/json"
	"io/ioutil"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/x/ecocredit/basket"
)

func txFlags(cmd *cobra.Command) *cobra.Command {
	flags.AddTxFlagsToCmd(cmd)
	_ = cmd.MarkFlagRequired(flags.FlagFrom)
	return cmd
}

func parseBasketCredits(creditsFile string) ([]*basket.BasketCredit, error) {
	bz, err := ioutil.ReadFile(creditsFile)
	if err != nil {
		return nil, err
	}

	if err := types.CheckDuplicateKey(json.NewDecoder(bytes.NewReader(bz)), nil); err != nil {
		return nil, err
	}

	var credits []*basket.BasketCredit

	// using json package because array is not a proto message
	if err = json.Unmarshal(bz, &credits); err != nil {
		return nil, err
	}

	return credits, nil
}
