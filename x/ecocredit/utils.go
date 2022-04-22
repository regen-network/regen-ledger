package ecocredit

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ReClassID        = `[A-Z]{1,3}[0-9]{2,}`
	reFullClassID    = regexp.MustCompile(fmt.Sprintf(`^%s$`, ReClassID))
	ReBatchDenom     = fmt.Sprintf(`%s-[0-9]{8}-[0-9]{8}-[0-9]{3,}`, ReClassID)
	reFullBatchDenom = regexp.MustCompile(fmt.Sprintf(`^%s$`, ReBatchDenom))
)

func ValidateClassID(classId string) error {
	matches := reFullClassID.FindStringSubmatch(classId)
	if matches == nil {
		return ErrParseFailure.Wrapf("class ID didn't match the format: expected A00, got %s", classId)
	}
	return nil
}

var reProjectID = regexp.MustCompile(`^[A-Za-z0-9]{2,16}$`)

func ValidateProjectID(projectID string) error {
	matches := reProjectID.FindStringSubmatch(projectID)
	if matches == nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid project id: %s.", projectID)
	}

	return nil
}

func ValidateDenom(denom string) error {
	matches := reFullBatchDenom.FindStringSubmatch(denom)
	if matches == nil {
		return ErrParseFailure.Wrap("invalid denom. Valid denom format is: A00-00000000-00000000-000")
	}
	return nil
}

var reJurisdiction = regexp.MustCompile(`^([A-Z]{2})(?:-([A-Z0-9]{1,3})(?: ([a-zA-Z0-9 \-]{1,64}))?)?$`)

func ValidateJurisdiction(jurisdiction string) error {
	matches := reJurisdiction.FindStringSubmatch(jurisdiction)
	if matches == nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "Invalid jurisdiction: %s.\nJurisdiction should have format <country-code>[-<region-code>[ <postal-code>]].\n", jurisdiction)
	}

	return nil
}

func FormatProjectID(classID string, projectSeqNo uint64) string {
	return fmt.Sprintf("%s%02d", classID, projectSeqNo)
}

func NormalizeCreditTypeName(name string) string {
	return fastRemoveWhitespace(strings.ToLower(name))
}

func fastRemoveWhitespace(str string) string {
	var b strings.Builder
	b.Grow(len(str))
	for _, ch := range str {
		if !unicode.IsSpace(ch) {
			b.WriteRune(ch)
		}
	}
	return b.String()
}
