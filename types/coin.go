package types

var (
	// CoinDenomRegex is used instead of DefaultCoinDenomRegex
	// to support basket denomination and DID characters
	// TODO: remove after updating to cosmos-sdk v0.46 #857
	CoinDenomRegex = `[a-zA-Z][a-zA-Z0-9/:._-]{2,127}`
)
