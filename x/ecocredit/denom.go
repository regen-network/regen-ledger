package ecocredit

import (
	"fmt"
	"regexp"
	"time"
)

// Calculate the ID to use for a new credit class, based on the credit type and
// sequence number. This format may evolve over time, but will maintain
// backwards compatibility.
//
// The initial version has format:
// <credit type abbreviation><class seq no>
func FormatClassID(creditType *CreditType, classSeqNo uint64) (string, error) {
	return fmt.Sprintf("%s%02d", creditType.Abbreviation, classSeqNo), nil
}

// Calculate the denomination to use for a batch, based on the batch
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
// e.g C01-20190101-20200101-001
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

// Validate a class ID conforms to the format described in FormatClassID. The
// return is nil if the ID is valid.
func ValidateClassID(classId string) error {
	matches := reFullClassID.FindStringSubmatch(classId)
	if matches == nil {
		return ErrParseFailure.Wrapf("class ID didn't match the format: expected A00, got %s", classId)
	}
	return nil
}

// Validate a batch denomination conforms to the format described in
// FormatDenom. The return is nil if the denom is valid.
func ValidateDenom(denom string) error {
	matches := reFullBatchDenom.FindStringSubmatch(denom)
	if matches == nil {
		return ErrParseFailure.Wrapf("denomination didn't match the format: expected A00-00000000-00000000-000, got %s", denom)
	}
	return nil
}
