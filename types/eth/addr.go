package eth

import (
	"fmt"
	"regexp"
)

var RegexAddress = `0x[0-9a-fA-F]{40}`

func IsValidAddress(addr string) bool {
	re := regexp.MustCompile(fmt.Sprintf(`^%s$`, RegexAddress))
	return re.MatchString(addr)
}
