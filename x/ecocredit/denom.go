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
func FormatClassID(classSeqNo uint64) (string, error) {
	if classSeqNo > 999 {
		return "", fmt.Errorf("class sequence number exceeds limit of 999: got %d", classSeqNo)
	}

	return fmt.Sprintf("C%03d", classSeqNo), nil
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
// - <batch seq no> is the sequence number of the batch, padded to two digits
//
// e.g C001-20190101-20200101-01
//
// NB: This might differ from the actual denomination used.
func FormatDenom(classId string, batchSeqNo uint64, startDate *time.Time, endDate *time.Time) (string, error) {
	if batchSeqNo > 99 {
		return "", fmt.Errorf("batch sequence number exceeds limit of 99: got %d", batchSeqNo)
	}

	return fmt.Sprintf(
		"%s-%s-%s-%02d",

		// Class ID as three digits
		classId,

		// Start Date as YYYYMMDD
		startDate.Format("20060102"),

		// End Date as YYYYMMDD
		endDate.Format("20060102"),

		// Batch sequence number as two digits
		batchSeqNo,
	), nil
}

var (
	ReClassID        = `[A-Z]{1,3}[0-9]{3}`
	reFullClassID    = regexp.MustCompile(fmt.Sprintf(`^%s$`, ReClassID))
	ReBatchDenom     = fmt.Sprintf(`%s-[0-9]{8}-[0-9]{8}-[0-9]{2}`, ReClassID)
	reFullBatchDenom = regexp.MustCompile(fmt.Sprintf(`^%s$`, ReBatchDenom))
)

// Validate a class ID coforms to the format described in FormatClassID. The
// return is nil if the ID is valid.
func ValidateClassID(classId string) error {
	matches := reFullClassID.FindStringSubmatch(classId)
	if matches == nil {
		return fmt.Errorf("class ID didn't match the format: expected A000, got %s", classId)
	}
	return nil
}

// Validate a batch denomination conforms to the format described in
// FormatDenom. The return is nil if the denom is valid.
func ValidateDenom(denom string) error {
	matches := reFullBatchDenom.FindStringSubmatch(denom)
	if matches == nil {
		return fmt.Errorf("denomination didn't match the format: expected A000-00000000-00000000-00, got %s", denom)
	}
	return nil
}
