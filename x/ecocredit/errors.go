package ecocredit

import sdkerrors "github.com/cosmos/cosmos-sdk/errors"

var (
	ErrParseFailure      = sdkerrors.Register(ModuleName, 2, "parse error")
	ErrInsufficientFunds = sdkerrors.Register(ModuleName, 3, "insufficient credit balance")
	ErrMaxLimit          = sdkerrors.Register(ModuleName, 4, "limit exceeded")
	ErrInvalidSellOrder  = sdkerrors.Register(ModuleName, 5, "invalid sell order")
	ErrInvalidBuyOrder   = sdkerrors.Register(ModuleName, 6, "invalid buy order")
	ErrNotFound          = sdkerrors.Register(ModuleName, 7, "not found")
	ErrInvalidInteger    = sdkerrors.Register(ModuleName, 8, "invalid integer")
)
