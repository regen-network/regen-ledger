package ecocredit

import (
	"regexp"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var reProjectID = regexp.MustCompile(`^[A-Za-z0-9]{2,16}$`)

// Validate a project ID conforms to the format described in reProjectID. The
// return is nil if the ID is valid.
func validateProjectID(projectID string) error {
	matches := reProjectID.FindStringSubmatch(projectID)
	if matches == nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "Invalid projectID: %s.", projectID)
	}

	return nil
}
