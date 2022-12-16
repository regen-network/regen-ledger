package client

import (
	"bytes"
	"encoding/json"
	"os"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client/flags"

	regentypes "github.com/regen-network/regen-ledger/types/v2"
	types "github.com/regen-network/regen-ledger/x/ecocredit/v3/basket/types/v1"
)

func txFlags(cmd *cobra.Command) *cobra.Command {
	flags.AddTxFlagsToCmd(cmd)
	_ = cmd.MarkFlagRequired(flags.FlagFrom)
	return cmd
}

func parseBasketCredits(creditsFile string) ([]*types.BasketCredit, error) {
	bz, err := os.ReadFile(creditsFile)
	if err != nil {
		return nil, err
	}

	if err := regentypes.CheckDuplicateKey(json.NewDecoder(bytes.NewReader(bz)), nil); err != nil {
		return nil, err
	}

	var credits []*types.BasketCredit

	// using json package because array is not a proto message
	if err = json.Unmarshal(bz, &credits); err != nil {
		return nil, err
	}

	return credits, nil
}
