package basket

import "cosmossdk.io/errors"

var codespace = "ecocredit/basket"

var (
	ErrCantDisableRetire = errors.Register(codespace, 2, "can't disable retirement when taking from this basket")
)
