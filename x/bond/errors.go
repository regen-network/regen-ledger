package bond

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/bond module sentinel errors
var (
	ErrInsufficientFunds = sdkerrors.Register(ModuleName, 1, "insufficient bond balance to sell")
	ErrMaxLimit          = sdkerrors.Register(ModuleName, 2, "limit exceeded")
)
