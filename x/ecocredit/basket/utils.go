package basket

import (
	"fmt"
	"regexp"

	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

var (
	RegexBasketDenom = `[a-zA-Z][a-zA-Z0-9/:._-]{2,127}`
	regexBasketDenom = regexp.MustCompile(fmt.Sprintf(`^%s$`, RegexBasketDenom))
)

const basketDenomPrefix = "eco"

// FormatBasketDenom formats denom and display denom:
// - denom: eco.<m.Exponent><m.CreditTypeAbbrev>.<m.Name>
// - display denom: eco.<m.CreditTypeAbbrev>.<m.Name>
func FormatBasketDenom(name, creditTypeAbbrev string, exponent uint32) (string, string, error) {
	denomPrefix, err := core.ExponentToPrefix(exponent)
	if err != nil {
		return "", "", err
	}

	denomTail := creditTypeAbbrev + "." + name
	denom := basketDenomPrefix + "." + denomPrefix + denomTail
	displayDenomName := basketDenomPrefix + "." + denomTail

	return denom, displayDenomName, nil
}
