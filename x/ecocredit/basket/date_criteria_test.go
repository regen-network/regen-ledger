package basket

import (
	"testing"

	"github.com/gogo/protobuf/types"
	"github.com/stretchr/testify/require"
)

func TestDateCriteriaToApi(t *testing.T) {
	require := require.New(t)

	var dc *DateCriteria
	require.Nil(dc.ToApi(), "handles nil")

	tstamp := &types.Timestamp{Seconds: 10}
	dc = &DateCriteria{MinStartDate: tstamp}
	tstampStd, err := types.TimestampFromProto(tstamp)
	require.NoError(err)
	require.Equal(tstampStd, dc.ToApi().GetMinStartDate().AsTime(), "handles min start date")

	dur := &types.Duration{Seconds: 50}
	dc = &DateCriteria{StartDateWindow: dur}
	durStd, err := types.DurationFromProto(dur)
	require.NoError(err)
	dw := dc.ToApi().GetStartDateWindow()
	require.NotNil(dw)
	require.Equal(durStd, dw.AsDuration(), "handles window date")
}
