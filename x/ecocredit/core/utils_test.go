package core

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"pgregory.net/rapid"
)

func TestUtils(t *testing.T) {
	t.Run("TestFormatClassId", rapid.MakeCheck(testFormatClassId))
	t.Run("TestInvalidClassId", rapid.MakeCheck(testInvalidClassId))
	t.Run("TestFormatProjectId", rapid.MakeCheck(testFormatProjectId))
	t.Run("TestInvalidProjectId", rapid.MakeCheck(testInvalidProjectId))
	t.Run("TestFormatBatchDenom", rapid.MakeCheck(testFormatBatchDenom))
	t.Run("TestInvalidBatchDenom", rapid.MakeCheck(testInvalidBatchDenom))
	t.Run("TestGetClassIdFromBatchDenom", rapid.MakeCheck(testGetClassIdFromBatchDenom))
	t.Run("GetCreditTypeAbbrevFromClassId", rapid.MakeCheck(testGetCreditTypeAbbrevFromClassId))
}

func testFormatClassId(t *rapid.T) {
	creditType := genCreditType.Draw(t, "creditType").(*CreditType)
	classSeq := rapid.Uint64().Draw(t, "classSeq").(uint64)

	classId := FormatClassId(creditType.Abbreviation, classSeq)

	t.Log(classId)

	err := ValidateClassId(classId)
	require.NoError(t, err)
}

func testInvalidClassId(t *rapid.T) {
	classId := genInvalidClassId.Draw(t, "classId").(string)
	require.Error(t, ValidateClassId(classId))
}

func testFormatProjectId(t *rapid.T) {
	creditType := genCreditType.Draw(t, "creditType").(*CreditType)
	classSeq := rapid.Uint64().Draw(t, "classSeq").(uint64)
	projectSeq := rapid.Uint64().Draw(t, "projectSeq").(uint64)

	classId := FormatClassId(creditType.Abbreviation, classSeq)
	projectId := FormatProjectId(classId, projectSeq)

	t.Log(projectId)

	err := ValidateProjectId(projectId)
	require.NoError(t, err)
}

func testInvalidProjectId(t *rapid.T) {
	projectId := genInvalidProjectId.Draw(t, "projectId").(string)
	require.Error(t, ValidateProjectId(projectId))
}

func testFormatBatchDenom(t *rapid.T) {
	creditType := genCreditType.Draw(t, "creditType").(*CreditType)
	classSeq := rapid.Uint64().Draw(t, "classSeq").(uint64)
	projectSeq := rapid.Uint64().Draw(t, "projectSeq").(uint64)
	batchSeq := rapid.Uint64().Draw(t, "batchSeq").(uint64)
	startDate := genTime.Draw(t, "startDate").(*time.Time)
	endDate := genTime.Draw(t, "endDate").(*time.Time)

	classId := FormatClassId(creditType.Abbreviation, classSeq)
	projectId := FormatProjectId(classId, projectSeq)
	denom, err := FormatBatchDenom(projectId, batchSeq, startDate, endDate)
	require.NoError(t, err)

	t.Log(denom)

	err = ValidateBatchDenom(denom)
	require.NoError(t, err)
}

func testInvalidBatchDenom(t *rapid.T) {
	batchDenom := genInvalidBatchDenom.Draw(t, "batchDenom").(string)
	require.Error(t, ValidateBatchDenom(batchDenom))
}

func testGetClassIdFromBatchDenom(t *rapid.T) {
	creditType := genCreditType.Draw(t, "creditType").(*CreditType)
	classSeq := rapid.Uint64().Draw(t, "classSeq").(uint64)
	projectSeq := rapid.Uint64().Draw(t, "projectSeq").(uint64)
	batchSeq := rapid.Uint64().Draw(t, "batchSeq").(uint64)
	startDate := genTime.Draw(t, "startDate").(*time.Time)
	endDate := genTime.Draw(t, "endDate").(*time.Time)

	classId := FormatClassId(creditType.Abbreviation, classSeq)
	projectId := FormatProjectId(classId, projectSeq)
	denom, err := FormatBatchDenom(projectId, batchSeq, startDate, endDate)
	require.NoError(t, err)

	result := GetClassIdFromBatchDenom(denom)
	require.Equal(t, classId, result)
}

func testGetCreditTypeAbbrevFromClassId(t *rapid.T) {
	creditType := genCreditType.Draw(t, "creditType").(*CreditType)
	classSeq := rapid.Uint64().Draw(t, "classSeq").(uint64)

	classId := FormatClassId(creditType.Abbreviation, classSeq)
	result := GetCreditTypeAbbrevFromClassId(classId)
	require.Equal(t, creditType.Abbreviation, result)
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

// genInvalidProjectId generates strings that don't conform to the project id format
var genInvalidProjectId = rapid.OneOf(
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
