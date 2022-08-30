package ecocredit

import "cosmossdk.io/errors"

var (
	ErrParseFailure        = errors.Register(ModuleName, 2, "parse error")
	ErrInsufficientCredits = errors.Register(ModuleName, 3, "insufficient credit balance")
	ErrMaxLimit            = errors.Register(ModuleName, 4, "limit exceeded")
	ErrInvalidSellOrder    = errors.Register(ModuleName, 5, "invalid sell order")
	ErrInvalidBuyOrder     = errors.Register(ModuleName, 6, "invalid buy order")
)
