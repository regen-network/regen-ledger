package marketplace

import "github.com/cosmos/cosmos-sdk/types/errors"

const marketCodespace = "marketplace"

var (
	ErrBidTooLow = errors.Register(marketCodespace, 0, "bid price too low")
)
