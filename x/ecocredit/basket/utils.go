package basket

import (
	"fmt"
	"regexp"
)

var (
	RegexBasketDenom = `[a-zA-Z][a-zA-Z0-9/:._-]{2,127}`
	regexBasketDenom = regexp.MustCompile(fmt.Sprintf(`^%s$`, RegexBasketDenom))
)
