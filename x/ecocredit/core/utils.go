package core

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/regen-network/regen-ledger/x/ecocredit"
)

var errBadReq = sdkerrors.ErrInvalidRequest

const (
	// BridgePolygon is currently the only allowed target when calling
	// Msg/Bridge and the only allowed source (provided within OriginTx)
	// when calling Msg/BridgeReceive. This value is not required as the
	// source within basic OriginTx validation, allowing for manual bridge
	// operations to be performed from other sources with Msg/CreateBatch
	// and Msg/MintBatchCredits.
	// TODO: remove after we introduce governance gated chains
	// https://github.com/regen-network/regen-ledger/issues/1119
	BridgePolygon = "polygon"

	// MaxMetadataLength defines the max length of the metadata bytes field
	// for the credit-class & credit-batch.
	MaxMetadataLength = 256

	// MaxNoteLength defines the max length for note fields.
	MaxNoteLength = 512
)

var (
	RegexClassId      = `[A-Z]{1,3}[0-9]{2,}`
	RegexProjectId    = fmt.Sprintf(`%s-[A-Z0-9]{2,}`, RegexClassId)
	RegexBatchDenom   = fmt.Sprintf(`%s-[0-9]{8}-[0-9]{8}-[0-9]{3,}`, RegexProjectId)
	RegexJurisdiction = `([A-Z]{2})(?:-([A-Z0-9]{1,3})(?: ([a-zA-Z0-9 \-]{1,64}))?)?`

	regexClassId      = regexp.MustCompile(fmt.Sprintf(`^%s$`, RegexClassId))
	regexProjectId    = regexp.MustCompile(fmt.Sprintf(`^%s$`, RegexProjectId))
	regexBatchDenom   = regexp.MustCompile(fmt.Sprintf(`^%s$`, RegexBatchDenom))
	regexJurisdiction = regexp.MustCompile(fmt.Sprintf(`^%s$`, RegexJurisdiction))
)

// FormatClassId formats the unique identifier for a new credit class, based
// on the credit type abbreviation and the credit class sequence number. This
// format may evolve over time, but will maintain backwards compatibility.
//
// The current version has the format:
// <credit-type-abbrev><class-sequence>
//
// <credit-type-abbrev> is the unique identifier of the credit type
// <class-sequence> is the sequence number of the credit class, padded to at
// least three digits
//
// e.g. C01
func FormatClassId(creditTypeAbbreviation string, classSeqNo uint64) string {
	return fmt.Sprintf("%s%02d", creditTypeAbbreviation, classSeqNo)
}

// FormatProjectId formats the unique identifier for a new project, based on
// the credit class id and project sequence number. This format may evolve over
// time, but will maintain backwards compatibility.
//
// The current version has the format:
// <class-id>-<project-sequence>
//
// <class-id> is the unique identifier of the credit class (see FormatClassId)
// <project-sequence> is the sequence number of the project, padded to at least
// three digits
//
// e.g. C01-001
func FormatProjectId(classId string, projectSeqNo uint64) string {
	return fmt.Sprintf("%s-%03d", classId, projectSeqNo)
}

// FormatBatchDenom formats the unique denomination for a credit batch. This
// format may evolve over time, but will maintain backwards compatibility.
//
// The current version has the format:
// <project-id>-<start_date>-<end_date>-<batch_sequence>
//
// <project-id> is the unique identifier of the project (see FormatProjectId)
// <start-date> is the start date of the credit batch with the format YYYYMMDD
// <end-date> is the end date of the credit batch with the format YYYYMMDD
// <batch-sequence> is the sequence number of the credit batch, padded to at least
// three digits
//
// e.g. C01-001-20190101-20200101-001
func FormatBatchDenom(projectId string, batchSeqNo uint64, startDate, endDate *time.Time) (string, error) {
	return fmt.Sprintf(
		"%s-%s-%s-%03d",

		// Project ID string
		projectId,

		// Start Date as YYYYMMDD
		startDate.UTC().Format("20060102"),

		// End Date as YYYYMMDD
		endDate.UTC().Format("20060102"),

		// Batch sequence number padded to at least three digits
		batchSeqNo,
	), nil
}

// ValidateCreditTypeAbbreviation validates a credit type abbreviation, ensuring it
// is only 1-3 uppercase letters. The return is nil if the abbreviation is valid.
func ValidateCreditTypeAbbreviation(abbr string) error {
	if abbr == "" {
		return ecocredit.ErrParseFailure.Wrap("credit type abbreviation cannot be empty")
	}
	reAbbr := regexp.MustCompile(`^[A-Z]{1,3}$`)
	if !reAbbr.Match([]byte(abbr)) {
		return ecocredit.ErrParseFailure.Wrapf("credit type abbreviation must be 1-3 uppercase latin letters: got %s", abbr)
	}
	return nil
}

// ValidateClassId validates a class ID conforms to the format described in
// FormatClassId. The return is nil if the ID is valid.
func ValidateClassId(classId string) error {
	if classId == "" {
		return ecocredit.ErrParseFailure.Wrap("class id cannot be empty")
	}
	matches := regexClassId.FindStringSubmatch(classId)
	if matches == nil {
		return ecocredit.ErrParseFailure.Wrapf("class ID didn't match the format: expected A00, got %s", classId)
	}
	return nil
}

// ValidateProjectId validates a project ID conforms to the format described
// in FormatProjectId. The return is nil if the ID is valid.
func ValidateProjectId(projectId string) error {
	if projectId == "" {
		return ecocredit.ErrParseFailure.Wrap("project id cannot be empty")
	}
	matches := regexProjectId.FindStringSubmatch(projectId)
	if matches == nil {
		return sdkerrors.Wrapf(ecocredit.ErrParseFailure, "invalid project id: %s", projectId)
	}
	return nil
}

// ValidateBatchDenom validates a batch denomination conforms to the format
// described in FormatBatchDenom. The return is nil if the denom is valid.
func ValidateBatchDenom(denom string) error {
	if denom == "" {
		return ecocredit.ErrParseFailure.Wrap("batch denom cannot be empty")
	}
	matches := regexBatchDenom.FindStringSubmatch(denom)
	if matches == nil {
		return ecocredit.ErrParseFailure.Wrap("invalid batch denom: expected format A00-000-00000000-00000000-000")
	}
	return nil
}

// ValidateJurisdiction checks that the country and region conform to ISO 3166 and
// the postal code is valid. This is a simple regex check and doesn't check that
// the country or subdivision codes actually exist. This is because the codes
// could change at short notice and we don't want to hardfork to keep up-to-date
// with that information. The return is nil if the jurisdiction is valid.
func ValidateJurisdiction(jurisdiction string) error {
	if jurisdiction == "" {
		return ecocredit.ErrParseFailure.Wrap("jurisdiction cannot be empty, expected format <country-code>[-<region-code>[ <postal-code>]]")
	}
	matches := regexJurisdiction.FindStringSubmatch(jurisdiction)
	if matches == nil {
		return ecocredit.ErrParseFailure.Wrapf("invalid jurisdiction: %s, expected format <country-code>[-<region-code>[ <postal-code>]]", jurisdiction)
	}

	return nil
}

// GetClassIdFromProjectId returns the credit class ID in a project ID.
func GetClassIdFromProjectId(projectId string) string {
	var s strings.Builder
	for _, r := range projectId {
		if r != '-' {
			s.WriteRune(r)
			continue
		}
		break
	}
	return s.String()
}

// GetClassIdFromBatchDenom returns the credit class ID in a batch denom.
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

// GetProjectIdFromBatchDenom returns the credit project ID in a batch denom.
func GetProjectIdFromBatchDenom(denom string) string {
	var s strings.Builder
	c := 0
	for _, r := range denom {
		if r == '-' {
			c++
		}
		if r != '-' || c != 2 {
			s.WriteRune(r)
			continue
		}
		break
	}
	return s.String()
}

// GetCreditTypeAbbrevFromClassId returns the credit type abbreviation in a credit class id
func GetCreditTypeAbbrevFromClassId(classId string) string {
	var s strings.Builder
	for _, r := range classId {
		if !unicode.IsNumber(r) {
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
		return "", ecocredit.ErrParseFailure.Wrapf("exponent must be one of %s", validExponents)
	}
	return e, nil
}
