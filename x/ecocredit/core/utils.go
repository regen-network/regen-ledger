package core

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
	"time"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/regen-network/regen-ledger/x/ecocredit"
)

var errBadReq = sdkerrors.ErrInvalidRequest

// MaxMetadataLength defines the max length of the metadata bytes field
// for the credit-class & credit-batch.
// TODO: This could be used as params once x/params is upgraded to use protobuf
const MaxMetadataLength = 256

var reJurisdiction = regexp.MustCompile(`^([A-Z]{2})(?:-([A-Z0-9]{1,3})(?: ([a-zA-Z0-9 \-]{1,64}))?)?$`)

// ValidateJurisdiction checks that the country and region conform to ISO 3166 and
// the postal code is valid. This is a simple regex check and doesn't check that
// the country or subdivision codes actually exist. This is because the codes
// could change at short notice and we don't want to hardfork to keep up-to-date
// with that information.
func ValidateJurisdiction(jurisdiction string) error {
	matches := reJurisdiction.FindStringSubmatch(jurisdiction)
	if matches == nil {
		return sdkerrors.ErrInvalidRequest.Wrapf("Invalid jurisdiction: %s.\nJurisdiction should have format <country-code>[-<region-code>[ <postal-code>]].\n", jurisdiction)
	}

	return nil
}

// reProjectID defines regular expression to check if the string contains only alphanumeric characters
// and is between 2 ~ 16 characters long.
//
// e.g. P01, C01P01, 123
var reProjectID = regexp.MustCompile(`^[A-Za-z0-9]{2,16}$`)

// ValidateProjectID validates a project ID conforms to the format described in reProjectID. The
// return is nil if the ID is valid.
func ValidateProjectID(projectID string) error {
	matches := reProjectID.FindStringSubmatch(projectID)
	if matches == nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid project id: %s.", projectID)
	}

	return nil
}

// FormatProjectID formats the ID to use for a new project, based on the class id and
// the project sequence number.
//
// The initial version has format:
// <class id><project seq no>
func FormatProjectID(classID string, projectSeqNo uint64) string {
	return fmt.Sprintf("%s%02d", classID, projectSeqNo)
}

// FormatClassID formats the ID to use for a new credit class, based on the credit type and
// sequence number. This format may evolve over time, but will maintain
// backwards compatibility.
//
// The initial version has format:
// <credit type abbreviation><class seq no>
func FormatClassID(creditTypeAbbreviation string, classSeqNo uint64) string {
	return fmt.Sprintf("%s%02d", creditTypeAbbreviation, classSeqNo)
}

// FormatDenom formats the denomination to use for a batch, based on the batch
// information. This format may evolve over time, but will maintain backwards
// compatibility.
//
// The initial version has format:
// <class id>-<start date>-<end date>-<batch seq no>
// where:
// - <class id> is the string ID of the credit class
// - <start date> is the start date of the batch in form YYYYMMDD
// - <end date> is the end date of the batch in form YYYYMMDD
// - <batch seq no> is the sequence number of the batch, padded to at least
//   three digits
//
// e.g. C01-20190101-20200101-001
//
// NB: This might differ from the actual denomination used.
func FormatDenom(classId string, batchSeqNo uint64, startDate *time.Time, endDate *time.Time) (string, error) {
	return fmt.Sprintf(
		"%s-%s-%s-%03d",

		// Class ID string
		classId,

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
		return ecocredit.ErrParseFailure.Wrapf("class ID didn't match the format: expected A00, got %s", classId)
	}
	return nil
}

// ValidateDenom validates a batch denomination conforms to the format described in
// FormatDenom. The return is nil if the denom is valid.
func ValidateDenom(denom string) error {
	matches := reFullBatchDenom.FindStringSubmatch(denom)
	if matches == nil {
		return ecocredit.ErrParseFailure.Wrap("invalid denom. Valid denom format is: A00-00000000-00000000-000")
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
