package ecocredit

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
	"time"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// FormatClassID formats the ID to use for a new credit class, based on the credit type and
// sequence number. This format may evolve over time, but will maintain backwards compatibility.
//
// The current version has the format:
// <credit type abbreviation><class seq no>
//
// e.g. C01
func FormatClassID(creditTypeAbbreviation string, classSeqNo uint64) string {
	return fmt.Sprintf("%s%02d", creditTypeAbbreviation, classSeqNo)
}

// FormatDenom formats the denomination to use for a credit batch. This format may evolve over
// time, but will maintain backwards compatibility.
//
// The current version has the format:
// <project_id>-<start_date>-<end_date>-<batch_sequence>
//
// where:
// - <project id> is the unique identifier of the project and has the format:
//     <credit_type_abbrev><class_id>-<project_sequence> (see FormatProjectId)
// - <start date> is the start date of the credit batch and has the format YYYYMMDD
// - <end date> is the end date of the credit batch and has the format YYYYMMDD
// - <batch seq no> is the sequence number of the credit batch, padded to at least
//   three digits
//
// e.g. C01-001-20190101-20200101-001
func FormatDenom(projectId string, batchSeqNo uint64, startDate, endDate *time.Time) (string, error) {
	return fmt.Sprintf(
		"%s-%s-%s-%03d",

		// Project ID string
		projectId,

		// Start Date as YYYYMMDD
		startDate.Format("20060102"),

		// End Date as YYYYMMDD
		endDate.Format("20060102"),

		// Batch sequence number padded to at least three digits
		batchSeqNo,
	), nil
}

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

// GetClassIdFromBatchDenom returns the classID in a batch denom
func GetClassIdFromBatchDenom(denom string) string {
	var s strings.Builder
	for _, r := range denom {
		if r != '-' {
			s.WriteRune(r)
			continue
		}
		break
	}
	return s.String()
}

// exponent prefix map https://en.wikipedia.org/wiki/Metric_prefix
var exponentPrefixMap = map[uint32]string{
	0:  "",
	1:  "d",
	2:  "c",
	3:  "m",
	6:  "u",
	9:  "n",
	12: "p",
	15: "f",
	18: "a",
	21: "z",
	24: "y",
}
var validExponents string

func init() {
	var exponents = make([]uint32, 0, len(exponentPrefixMap))
	for e := range exponentPrefixMap {
		exponents = append(exponents, e)
	}
	sort.Slice(exponents, func(i, j int) bool { return exponents[i] < exponents[j] })
	validExponents = fmt.Sprint(exponents)
}

// ExponentToPrefix returns a denom prefix for a given exponent.
// Returns error if the exponent is not supported.
func ExponentToPrefix(exponent uint32) (string, error) {
	e, ok := exponentPrefixMap[exponent]
	if !ok {
		return "", sdkerrors.ErrInvalidRequest.Wrapf("exponent must be one of %s", validExponents)
	}
	return e, nil
}
