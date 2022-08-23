package core

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"pgregory.net/rapid"
)

func TestUtils(t *testing.T) {
	t.Run("TestFormatClassID", rapid.MakeCheck(testFormatClassID))
	t.Run("TestInvalidClassID", rapid.MakeCheck(testInvalidClassID))
	t.Run("TestFormatProjectID", rapid.MakeCheck(testFormatProjectID))
	t.Run("TestInvalidProjectID", rapid.MakeCheck(testInvalidProjectID))
	t.Run("TestFormatBatchDenom", rapid.MakeCheck(testFormatBatchDenom))
	t.Run("TestInvalidBatchDenom", rapid.MakeCheck(testInvalidBatchDenom))
	t.Run("TestGetClassIDFromProjectID", rapid.MakeCheck(testGetClassIDFromProjectID))
	t.Run("TestGetClassIDFromBatchDenom", rapid.MakeCheck(testGetClassIDFromBatchDenom))
	t.Run("TestGetProjectIDFromBatchDenom", rapid.MakeCheck(testGetProjectIDFromBatchDenom))
	t.Run("GetCreditTypeAbbrevFromClassID", rapid.MakeCheck(testGetCreditTypeAbbrevFromClassID))
}

func testFormatClassID(t *rapid.T) {
	creditType := genCreditType.Draw(t, "creditType").(*CreditType)
	classSeq := rapid.Uint64().Draw(t, "classSeq").(uint64)

	classID := FormatClassID(creditType.Abbreviation, classSeq)

	t.Log(classID)

	err := ValidateClassID(classID)
	require.NoError(t, err)
}

func testInvalidClassID(t *rapid.T) {
	classID := genInvalidClassID.Draw(t, "classID").(string)
	require.Error(t, ValidateClassID(classID))
}

func testFormatProjectID(t *rapid.T) {
	creditType := genCreditType.Draw(t, "creditType").(*CreditType)
	classSeq := rapid.Uint64().Draw(t, "classSeq").(uint64)
	projectSeq := rapid.Uint64().Draw(t, "projectSeq").(uint64)

	classID := FormatClassID(creditType.Abbreviation, classSeq)
	projectID := FormatProjectID(classID, projectSeq)

	t.Log(projectID)

	err := ValidateProjectID(projectID)
	require.NoError(t, err)
}

func testInvalidProjectID(t *rapid.T) {
	projectID := genInvalidProjectID.Draw(t, "projectID").(string)
	require.Error(t, ValidateProjectID(projectID))
}

func testFormatBatchDenom(t *rapid.T) {
	creditType := genCreditType.Draw(t, "creditType").(*CreditType)
	classSeq := rapid.Uint64().Draw(t, "classSeq").(uint64)
	projectSeq := rapid.Uint64().Draw(t, "projectSeq").(uint64)
	batchSeq := rapid.Uint64().Draw(t, "batchSeq").(uint64)
	startDate := genTime.Draw(t, "startDate").(*time.Time)
	endDate := genTime.Draw(t, "endDate").(*time.Time)

	classID := FormatClassID(creditType.Abbreviation, classSeq)
	projectID := FormatProjectID(classID, projectSeq)
	denom, err := FormatBatchDenom(projectID, batchSeq, startDate, endDate)
	require.NoError(t, err)

	t.Log(denom)

	err = ValidateBatchDenom(denom)
	require.NoError(t, err)
}

func testInvalidBatchDenom(t *rapid.T) {
	batchDenom := genInvalidBatchDenom.Draw(t, "batchDenom").(string)
	require.Error(t, ValidateBatchDenom(batchDenom))
}

func testGetClassIDFromProjectID(t *rapid.T) {
	creditType := genCreditType.Draw(t, "creditType").(*CreditType)
	classSeq := rapid.Uint64().Draw(t, "classSeq").(uint64)
	projectSeq := rapid.Uint64().Draw(t, "projectSeq").(uint64)

	classID := FormatClassID(creditType.Abbreviation, classSeq)
	projectID := FormatProjectID(classID, projectSeq)

	result := GetClassIDFromProjectID(projectID)
	require.Equal(t, classID, result)
}

func testGetClassIDFromBatchDenom(t *rapid.T) {
	creditType := genCreditType.Draw(t, "creditType").(*CreditType)
	classSeq := rapid.Uint64().Draw(t, "classSeq").(uint64)
	projectSeq := rapid.Uint64().Draw(t, "projectSeq").(uint64)
	batchSeq := rapid.Uint64().Draw(t, "batchSeq").(uint64)
	startDate := genTime.Draw(t, "startDate").(*time.Time)
	endDate := genTime.Draw(t, "endDate").(*time.Time)

	classID := FormatClassID(creditType.Abbreviation, classSeq)
	projectID := FormatProjectID(classID, projectSeq)
	denom, err := FormatBatchDenom(projectID, batchSeq, startDate, endDate)
	require.NoError(t, err)

	result := GetClassIDFromBatchDenom(denom)
	require.Equal(t, classID, result)
}

func testGetProjectIDFromBatchDenom(t *rapid.T) {
	creditType := genCreditType.Draw(t, "creditType").(*CreditType)
	classSeq := rapid.Uint64().Draw(t, "classSeq").(uint64)
	projectSeq := rapid.Uint64().Draw(t, "projectSeq").(uint64)
	batchSeq := rapid.Uint64().Draw(t, "batchSeq").(uint64)
	startDate := genTime.Draw(t, "startDate").(*time.Time)
	endDate := genTime.Draw(t, "endDate").(*time.Time)

	classID := FormatClassID(creditType.Abbreviation, classSeq)
	projectID := FormatProjectID(classID, projectSeq)
	denom, err := FormatBatchDenom(projectID, batchSeq, startDate, endDate)
	require.NoError(t, err)

	result := GetProjectIDFromBatchDenom(denom)
	require.Equal(t, projectID, result)
}

func testGetCreditTypeAbbrevFromClassID(t *rapid.T) {
	creditType := genCreditType.Draw(t, "creditType").(*CreditType)
	classSeq := rapid.Uint64().Draw(t, "classSeq").(uint64)

	classID := FormatClassID(creditType.Abbreviation, classSeq)
	result := GetCreditTypeAbbrevFromClassID(classID)
	require.Equal(t, creditType.Abbreviation, result)
}

// genCreditType generates an empty credit type with a random valid abbreviation
var genCreditType = rapid.Custom(func(t *rapid.T) *CreditType {
	abbr := rapid.StringMatching(`[A-Z]{1,3}`).Draw(t, "abbr").(string)
	return &CreditType{
		Abbreviation: abbr,
	}
})

// genInvalidClassID generates strings that don't conform to the class id format
var genInvalidClassID = rapid.OneOf(
	rapid.StringMatching(`[a-zA-Z]*`),
	rapid.StringMatching(`[0-9]*`),
	rapid.StringMatching(`[A-Z]{4,}[0-9]*`),
)

// genInvalidProjectID generates strings that don't conform to the project id format
var genInvalidProjectID = rapid.OneOf(
	rapid.StringMatching(`[a-zA-Z]*`),
	rapid.StringMatching(`[0-9]*`),
	rapid.StringMatching(`[A-Z]{4,}[0-9]*`),
)

// genInvalidBatchDenom generates strings that don't conform to the batch denom format
var genInvalidBatchDenom = rapid.OneOf(
	genInvalidClassID,
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
