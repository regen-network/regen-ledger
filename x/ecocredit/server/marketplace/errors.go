package marketplace

import "github.com/cosmos/cosmos-sdk/types/errors"

const marketCodespace = "marketplace"

var (
	ErrBidTooLow = errors.Register(marketCodespace, 2, "bid price too low")
)
