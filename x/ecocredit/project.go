package ecocredit

import (
	"fmt"
	"regexp"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// reProjectID defines regular expression to check if the string contains only alphanumeric characters 
// and is between 2 ~ 16 characters long.
//
// e.g. P01, C01P01, 123
var reProjectID = regexp.MustCompile(`^[A-Za-z0-9]{2,16}$`)

// Validate a project ID conforms to the format described in reProjectID. The
// return is nil if the ID is valid.
func ValidateProjectID(projectID string) error {
	matches := reProjectID.FindStringSubmatch(projectID)
	if matches == nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid project id: %s.", projectID)
	}

	return nil
}

// Calculate the ID to use for a new project, based on the class id and
// the project sequence number.
//
// The initial version has format:
// <class id><project seq no>
func FormatProjectID(classID string, projectSeqNo uint64) string {
	return fmt.Sprintf("%s%02d", classID, projectSeqNo)
}
