package ecocredit

import (
	"fmt"
	"regexp"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var reProjectID = regexp.MustCompile(`^[A-Za-z0-9]{2,16}$`)

// Validate a project ID conforms to the format described in reProjectID. The
// return is nil if the ID is valid.
func ValidateProjectID(projectID string) error {
	matches := reProjectID.FindStringSubmatch(projectID)
	if matches == nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "Invalid projectID: %s.", projectID)
	}

	return nil
}

// Calculate the ID to use for a new project, based on the credit type and
// sequence number.
//
// The initial version has format:
// <credit type abbreviation><class seq no>
func FormatProjectID(creditType CreditType, classSeqNo uint64) string {
	return fmt.Sprintf("%s%02d", creditType.Abbreviation, classSeqNo)
}
