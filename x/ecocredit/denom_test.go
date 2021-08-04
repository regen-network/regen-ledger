package ecocredit

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"pgregory.net/rapid"
)

func TestDenom(t *testing.T) {
	t.Run("TestValidateFormatClassID", rapid.MakeCheck(testValidateFormatClassID))
	t.Run("TestValidateFormatDenom", rapid.MakeCheck(testValidateFormatDenom))
}

// Property: ValidateClassID(FormatClassID(a)) == nil IFF a <= 999
func testValidateFormatClassID(t *rapid.T) {
	classSeqNo := rapid.Uint64().Draw(t, "classSeqNo").(uint64)

	classId, err := FormatClassID(classSeqNo)
	if classSeqNo > 999 {
		require.Error(t, err)
	} else {
		require.NoError(t, err)

		err = ValidateClassID(classId)
		require.NoError(t, err)
	}
}

// Property: ValidateDenom(FormatDenom(a, b, c, d)) == nil IFF b <= 99
func testValidateFormatDenom(t *rapid.T) {
	classSeqNo := rapid.Uint64Range(0, 999).Draw(t, "classSeqNo").(uint64)
	batchSeqNo := rapid.Uint64().Draw(t, "batchSeqNo").(uint64)
	startDate := genTime.Draw(t, "startDate").(*time.Time)
	endDate := genTime.Draw(t, "endDate").(*time.Time)

	classId, err := FormatClassID(classSeqNo)
	require.NoError(t, err)

	denom, err := FormatDenom(classId, batchSeqNo, startDate, endDate)
	if batchSeqNo > 99 {
		require.Error(t, err)
	} else {
		require.NoError(t, err)

		err = ValidateDenom(denom)
		require.NoError(t, err)
	}
}

// genTime generates time values up to the year ~9999 to avoid overflowing the
// denomination format.
var genTime = rapid.Custom(func(t *rapid.T) *time.Time {
	secs := rapid.Int64Range(0, 2e11).Draw(t, "secs").(int64)
	nanos := rapid.Int64Range(0, 1e15).Draw(t, "nanos").(int64)
	time := time.Unix(secs, nanos)
	return &time
})
