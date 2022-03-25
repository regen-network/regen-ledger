package structvalid

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// ErrsToError creates an error if the errs slice is not empty. The content of the error is
// a formatted content of the slice.
func ErrsToError(errs []string) error {
	if len(errs) == 0 {
		return nil
	}
	return sdkerrors.ErrInvalidRequest.Wrapf("%v", errs)
}
