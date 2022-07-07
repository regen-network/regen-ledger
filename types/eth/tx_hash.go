package eth

import (
	"fmt"
	"regexp"
)

var RegexTxHash = `0x[0-9a-fA-F]{64}`

func IsValidTxHash(txHash string) bool {
	re := regexp.MustCompile(fmt.Sprintf(`^%s$`, RegexTxHash))
	return re.MatchString(txHash)
}
