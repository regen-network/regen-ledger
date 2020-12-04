package group

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrEmpty        = sdkerrors.Register(ModuleName, 202, "value is empty")
	ErrDuplicate    = sdkerrors.Register(ModuleName, 203, "duplicate value")
	ErrMaxLimit     = sdkerrors.Register(ModuleName, 204, "limit exceeded")
	ErrType         = sdkerrors.Register(ModuleName, 205, "invalid type")
	ErrInvalid      = sdkerrors.Register(ModuleName, 206, "invalid value")
	ErrUnauthorized = sdkerrors.Register(ModuleName, 207, "unauthorized")
	ErrModified     = sdkerrors.Register(ModuleName, 208, "modified")
	ErrExpired      = sdkerrors.Register(ModuleName, 209, "expired")
)
