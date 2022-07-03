package eth

import "regexp"

func IsValidEthereumTxHash(txHash string) bool {
	re := regexp.MustCompile("^0x[0-9a-fA-F]{64}$")
	return re.MatchString(txHash)
}
