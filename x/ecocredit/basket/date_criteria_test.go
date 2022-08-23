package basket

import (
	"testing"
	"time"

	"github.com/gogo/protobuf/types"
	"github.com/stretchr/testify/require"
)

func TestDateCriteriaToAPI(t *testing.T) {
	t.Parallel()

	require := require.New(t)

	var dc *DateCriteria
	require.Nil(dc.ToAPI(), "handles nil")

	tstamp := &types.Timestamp{Seconds: 10}
	dc = &DateCriteria{MinStartDate: tstamp}
	tstampStd, err := types.TimestampFromProto(tstamp)
	require.NoError(err)
	require.Equal(tstampStd, dc.ToAPI().GetMinStartDate().AsTime(), "handles min start date")

	dur := &types.Duration{Seconds: 50}
	dc = &DateCriteria{StartDateWindow: dur}
	durStd, err := types.DurationFromProto(dur)
	require.NoError(err)
	dw := dc.ToAPI().GetStartDateWindow()
	require.NotNil(dw)
	require.Equal(durStd, dw.AsDuration(), "handles window date")
}

func TestValidateDateCriteria(t *testing.T) {
	t.Parallel()

	tcs := []struct {
		id  string
		d   DateCriteria
		err string
	}{
		{
			"bad-min_start_date",
			DateCriteria{MinStartDate: &types.Timestamp{
				Seconds: time.Date(1400, 1, 1, 0, 0, 0, 0, time.UTC).Unix()},
			},
			"min_start_date must be after",
		},
		{
			"bad-start_date_window",
			DateCriteria{StartDateWindow: &types.Duration{Seconds: 3600}},
			"start_date_window must be at least 1 day",
		},
		{
			"both-min_start_date-start_date_window-set",
			DateCriteria{
				MinStartDate:    types.TimestampNow(),
				StartDateWindow: &types.Duration{Seconds: 3600 * 24 * 2},
			},
			"only one of min_start_date, start_date_window, or years_in_the_past must be set",
		},
		{
			"both-min_start_date-years_in_the_past-set",
			DateCriteria{
				MinStartDate:   types.TimestampNow(),
				YearsInThePast: 10,
			},
			"only one of min_start_date, start_date_window, or years_in_the_past must be set",
		},
		{
			"both-start_date_window-years_in_the_past-set",
			DateCriteria{
				StartDateWindow: &types.Duration{Seconds: 3600 * 24 * 2},
				YearsInThePast:  10,
			},
			"only one of min_start_date, start_date_window, or years_in_the_past must be set",
		},
		{
			"good-min_start_date",
			DateCriteria{MinStartDate: types.TimestampNow()},
			"",
		},
		{
			"good-start_date_window",
			DateCriteria{StartDateWindow: &types.Duration{Seconds: 3600 * 24 * 2}},
			"",
		},
		{
			"good-years_in_the_past",
			DateCriteria{YearsInThePast: 10},
			"",
		},
	}
	for _, tc := range tcs {
		t.Run(tc.id, func(t *testing.T) {
			t.Parallel()

			err := tc.d.Validate()
			errorMatches(t, err, tc.err)
		})
	}
}

func errorMatches(t *testing.T, err error, expect string) {
	if expect == "" {
		require.NoError(t, err)
	} else {
		require.Error(t, err)
		require.Contains(t, err.Error(), expect)
	}
}
