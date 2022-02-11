package basketclient

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/regen-network/regen-ledger/x/ecocredit/basket"
)

func parseBasketCredits(creditsFile string) ([]*basket.BasketCredit, error) {
	credits := []*basket.BasketCredit{}

	if creditsFile == "" {
		return nil, fmt.Errorf("credits file path is empty")
	}

	bz, err := ioutil.ReadFile(creditsFile)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(bz, &credits); err != nil {
		return nil, err
	}

	return credits, nil
}
