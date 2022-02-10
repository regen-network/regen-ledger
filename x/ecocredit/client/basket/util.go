package basket

import (
	"fmt"
	"io/ioutil"

	"github.com/cosmos/cosmos-sdk/client"

	keeper "github.com/regen-network/regen-ledger/x/ecocredit/basket"
)

func parseBasketCredits(clientCtx client.Context, creditsFile string) ([]*keeper.BasketCredit, error) {
	credits := keeper.BasketCredits{}

	if creditsFile == "" {
		return nil, fmt.Errorf("credits file path is empty")
	}

	contents, err := ioutil.ReadFile(creditsFile)
	if err != nil {
		return nil, err
	}

	if err := clientCtx.Codec.UnmarshalJSON(contents, &credits); err != nil {
		return nil, err
	}

	return credits.Credits, nil
}
