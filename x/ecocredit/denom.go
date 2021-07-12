package ecocredit

import (
	"fmt"
	"regexp"
	"time"
)

// Calculate the denomination to use for a batch, based on the batch
// information. This format may evolve over time, but will maintain backwards
// compatibility.
//
// The initial version has format:
// <credit type abbreviation><class id>-<start date>-<end date>-<batch id>
// where:
// - <credit type> is the abbreviation for the credit type
// - <class id> is the Class ID, padded to three digits
// - <start date> is the start date of the batch in form YYYYMMDD
// - <end date> is the end date of the batch in form YYYYMMDD
// - <batch no> is the sequence number of the batch, padded to two digits
//
// e.g C001-20190101-20200101-01
//
// NB: This might differ from the actual denomination used.
func FormatDenom(classId uint64, batchId uint64, startDate *time.Time, endDate *time.Time) (string, error) {
	if classId > 999 {
		return "", fmt.Errorf("class id exceeds limit of 999: got %d", classId)
	}

	if batchId > 99 {
		return "", fmt.Errorf("batch id exceeds limit of 99: got %d", batchId)
	}

	return fmt.Sprintf(
		"C%03d-%s-%s-%02d",

		// Class ID as three digits
		classId,

		// Start Date as YYYYMMDD
		startDate.Format("20060102"),

		// End Date as YYYYMMDD
		endDate.Format("20060102"),

		// Batch ID as two digits
		batchId,
	), nil
}

var (
	ReBatchDenom     = `[A-Z]{1,3}[0-9]{3}-[0-9]{8}-[0-9]{8}-[0-9]{2}`
	reFullBatchDenom = regexp.MustCompile(fmt.Sprintf(`^%s$`, ReBatchDenom))
)

// Validate a batch denomination conforms to the format described in
// FormatDenom. The return is nil if the denom is valid.
func ValidateDenom(denom string) error {
	matches := reFullBatchDenom.FindStringSubmatch(denom)
	if matches == nil {
		return fmt.Errorf("Denomination didn't match the format: expected A000-00000000-00000000-00, got %s", denom)
	}
	return nil
}
