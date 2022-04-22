package ecocredit

import (
	"fmt"
	"regexp"
)

var (
	ReClassID        = `[A-Z]{1,3}[0-9]{2,}`
	reFullClassID    = regexp.MustCompile(fmt.Sprintf(`^%s$`, ReClassID))
	ReBatchDenom     = fmt.Sprintf(`%s-[0-9]{8}-[0-9]{8}-[0-9]{3,}`, ReClassID)
	reFullBatchDenom = regexp.MustCompile(fmt.Sprintf(`^%s$`, ReBatchDenom))
)

// ValidateClassID validates a class ID conforms to the format described in FormatClassID. The
// return is nil if the ID is valid.
func ValidateClassID(classId string) error {
	matches := reFullClassID.FindStringSubmatch(classId)
	if matches == nil {
		return ErrParseFailure.Wrapf("class ID didn't match the format: expected A00, got %s", classId)
	}
	return nil
}

// ValidateDenom validates a batch denomination conforms to the format described in
// FormatDenom. The return is nil if the denom is valid.
func ValidateDenom(denom string) error {
	matches := reFullBatchDenom.FindStringSubmatch(denom)
	if matches == nil {
		return ErrParseFailure.Wrap("invalid denom. Valid denom format is: A00-00000000-00000000-000")
	}
	return nil
}
