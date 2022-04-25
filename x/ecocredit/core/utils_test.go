package core

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"pgregory.net/rapid"
)

func TestUtils(t *testing.T) {
	t.Run("TestFormatClassID", rapid.MakeCheck(testFormatClassId))
	t.Run("TestInvalidClassID", rapid.MakeCheck(testInvalidClassID))
	t.Run("TestFormatBatchDenom", rapid.MakeCheck(testFormatBatchDenom))
	t.Run("TestInvalidBatchDenom", rapid.MakeCheck(testInvalidBatchDenom))
	t.Run("TestGetClassIdFromBatchDenom", rapid.MakeCheck(testGetClassIdFromBatchDenom))
}

func testFormatClassId(t *rapid.T) {
	creditType := genCreditType.Draw(t, "creditType").(*CreditType)
	classSeqNo := rapid.Uint64().Draw(t, "classSeqNo").(uint64)

	classId := FormatClassID(creditType.Abbreviation, classSeqNo)

	err := ValidateClassID(classId)
	require.NoError(t, err)
}

func testInvalidClassID(t *rapid.T) {
	classID := genInvalidClassId.Draw(t, "classID").(string)
	require.Error(t, ValidateClassID(classID))
}

func testFormatBatchDenom(t *rapid.T) {
	creditType := genCreditType.Draw(t, "creditType").(*CreditType)
	classSeqNo := rapid.Uint64().Draw(t, "classSeqNo").(uint64)
	batchSeqNo := rapid.Uint64().Draw(t, "batchSeqNo").(uint64)
	startDate := genTime.Draw(t, "startDate").(*time.Time)
	endDate := genTime.Draw(t, "endDate").(*time.Time)

	classId := FormatClassID(creditType.Abbreviation, classSeqNo)

	denom, err := FormatDenom(classId, batchSeqNo, startDate, endDate)
	require.NoError(t, err)
	t.Log(denom)

	err = ValidateDenom(denom)
	require.NoError(t, err)
}

func testInvalidBatchDenom(t *rapid.T) {
	batchDenom := genInvalidBatchDenom.Draw(t, "batchDenom").(string)
	require.Error(t, ValidateDenom(batchDenom))
}

func testGetClassIdFromBatchDenom(t *rapid.T) {
	creditType := genCreditType.Draw(t, "creditType").(*CreditType)
	classSeqNo := rapid.Uint64().Draw(t, "classSeqNo").(uint64)
	batchSeqNo := rapid.Uint64().Draw(t, "batchSeqNo").(uint64)
	startDate := genTime.Draw(t, "startDate").(*time.Time)
	endDate := genTime.Draw(t, "endDate").(*time.Time)

	classId := FormatClassID(creditType.Abbreviation, classSeqNo)

	denom, err := FormatDenom(classId, batchSeqNo, startDate, endDate)
	require.NoError(t, err)
	t.Log(denom)

	result := GetClassIdFromBatchDenom(denom)

	require.Equal(t, classId, result)
}

// genCreditType generates an empty credit type with a random valid abbreviation
var genCreditType = rapid.Custom(func(t *rapid.T) *CreditType {
	abbr := rapid.StringMatching(`[A-Z]{1,3}`).Draw(t, "abbr").(string)
	return &CreditType{
		Abbreviation: abbr,
	}
})

// genInvalidClassId generates strings that don't conform to the class id format
var genInvalidClassId = rapid.OneOf(
	rapid.StringMatching(`[a-zA-Z]*`),
	rapid.StringMatching(`[0-9]*`),
	rapid.StringMatching(`[A-Z]{4,}[0-9]*`),
)

// genInvalidBatchDenom generates strings that don't conform to the batch denom format
var genInvalidBatchDenom = rapid.OneOf(
	genInvalidClassId,
	rapid.StringMatching(`[A-Z]{1,3}[0-9]*-[a-zA-Z\-]*`),
)

// genTime generates time values up to the year ~9999 to avoid overflowing the
// denomination format.
var genTime = rapid.Custom(func(t *rapid.T) *time.Time {
	secs := rapid.Int64Range(0, 2e11).Draw(t, "secs").(int64)
	nanos := rapid.Int64Range(0, 1e15).Draw(t, "nanos").(int64)
	time := time.Unix(secs, nanos)
	return &time
})
