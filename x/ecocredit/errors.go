package ecocredit

import sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

var (
	ErrParseFailure        = sdkerrors.Register(ModuleName, 2, "parse error")
	ErrInsufficientCredits = sdkerrors.Register(ModuleName, 3, "insufficient credit balance")
	ErrMaxLimit            = sdkerrors.Register(ModuleName, 4, "limit exceeded")
	ErrInvalidSellOrder    = sdkerrors.Register(ModuleName, 5, "invalid sell order")
	ErrInvalidBuyOrder     = sdkerrors.Register(ModuleName, 6, "invalid buy order")
)
