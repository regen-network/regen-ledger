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

// FormatProjectID formats the ID to use for a new project, based on the credit class id and
// sequence number. This format may evolve over time, but will maintain backwards compatibility.
//
// The current version has the format:
// <credit_class_id>-<project_sequence>
//
// e.g. C01-001
func FormatProjectID(classId string, projectSeqNo uint64) string {
	return fmt.Sprintf("%s-%03d", classId, projectSeqNo)
}
