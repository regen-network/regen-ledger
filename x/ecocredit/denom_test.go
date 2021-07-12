package ecocredit

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"pgregory.net/rapid"
)

func TestDenom(t *testing.T) {
	t.Run("TestValidateFormat", rapid.MakeCheck(testValidateFormat))
}

// Property: ValidateDenom(FormatDenom(a, b, c, d)) == nil
func testValidateFormat(t *rapid.T) {
	classId := rapid.Uint64().Draw(t, "classId").(uint64)
	batchId := rapid.Uint64().Draw(t, "batchId").(uint64)
	startDate := genTime.Draw(t, "startDate").(*time.Time)
	endDate := genTime.Draw(t, "endDate").(*time.Time)

	denom, err := FormatDenom(classId, batchId, startDate, endDate)
	if classId > 999 || batchId > 99 {
		require.Error(t, err)
	} else {
		require.NoError(t, err)

		err = ValidateDenom(denom)
		require.Nil(t, err)
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
