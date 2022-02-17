package basket

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"gotest.tools/v3/assert"
	"testing"

	"github.com/gogo/protobuf/types"
	"github.com/stretchr/testify/require"
)

func TestDateCriteriaToApi(t *testing.T) {
	require := require.New(t)

	var dc *DateCriteria
	require.Nil(dc.ToApi(), "handles nil")

	tstamp := &types.Timestamp{Seconds: 10}
	dc = &DateCriteria{&DateCriteria_MinStartDate{tstamp}}
	tstampStd, err := types.TimestampFromProto(tstamp)
	require.NoError(err)
	require.Equal(tstampStd, dc.ToApi().GetMinStartDate().AsTime(), "handles min start date")

	dur := &types.Duration{Seconds: 50}
	dc = &DateCriteria{&DateCriteria_StartDateWindow{dur}}
	durStd, err := types.DurationFromProto(dur)
	require.NoError(err)
	dw := dc.ToApi().GetStartDateWindow()
	require.NotNil(dw)
	require.Equal(durStd, dw.AsDuration(), "handles window date")
}

func TestAminoDate(t *testing.T) {
	cdc := codec.NewLegacyAmino()
	dur := types.Duration{Seconds: 50}
	bsk := MsgCreate{
		Name: "foo",
		DateCriteria: &DateCriteria{Sum: &DateCriteria_StartDateWindow{StartDateWindow: &dur}},
		Exponent: 32,
	}

	// test Un/MarshalJSON
	bz, err := cdc.MarshalJSON(bsk)
	assert.NilError(t, err)
	var bsk2 MsgCreate
	err = cdc.UnmarshalJSON(bz, &bsk2)
	assert.NilError(t, err)

	// test regular Un/Marshal
	bz, err = cdc.Marshal(bsk)
	assert.NilError(t, err)
	var bsk3 MsgCreate
	err = cdc.Unmarshal(bz, &bsk3)
	assert.NilError(t, err)


	assert.DeepEqual(t, bsk, bsk2)
	assert.DeepEqual(t, bsk2, bsk3)
}
