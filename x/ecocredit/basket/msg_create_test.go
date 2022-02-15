package basket

import (
	"testing"
	"time"

	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/types"
	gogotypes "github.com/gogo/protobuf/types"
	"github.com/stretchr/testify/require"
	"github.com/thanhpk/randstr"
)

func errorMatches(t *testing.T, err error, expect string) {
	if expect == "" {
		require.NoError(t, err)
	} else {
		require.Error(t, err)
		require.Contains(t, err.Error(), expect)
	}
}

func TestMsgCreateValidateBasic(t *testing.T) {
	_, _, addr1 := testdata.KeyTestPubAddr()
	a := addr1.String()
	name := randstr.String((nameMaxLen + nameMinLen) / 2)
	creditAbbr := randstr.String(10)
	descr := "my project description"
	start := &DateCriteria{&DateCriteria_MinStartDate{gogotypes.TimestampNow()}}

	classes := []string{"eco_class1"}

	tcs := []struct {
		id  string
		msg MsgCreate
		err string
	}{
		{"curator-1",
			MsgCreate{Curator: "wrong"},
			"malformed curator address"},
		{"name-1",
			MsgCreate{Curator: a, Name: ""},
			"name must be between"},
		{"name-long",
			MsgCreate{Curator: a, Name: randstr.String(nameMaxLen + 1)},
			"name must be between"},
		{"name-short",
			MsgCreate{Curator: a, Name: randstr.String(nameMinLen - 1)},
			"name must be between"},
		{"name-no-alpahnum",
			MsgCreate{Curator: a, Name: randstr.String(nameMinLen) + "*"},
			"name must be between"},
		{"exponent-1",
			MsgCreate{Curator: a, Name: name, Exponent: exponentMax + 1},
			"exponent must not be bigger than"},
		{"credity_type-1",
			MsgCreate{Curator: a, Name: name, Exponent: exponentMax},
			"credit_type_name must be defined"},
		{"credity_type-2",
			MsgCreate{Curator: a, Name: name, Exponent: exponentMax, CreditTypeAbbrev: randstr.String(creditTypeAbbrMaxLen + 1)},
			"credit_type_name must not be longer"},
		{"date_criteria-1",
			MsgCreate{Curator: a, Name: name, Exponent: exponentMax, CreditTypeAbbrev: creditAbbr, DateCriteria: &DateCriteria{}},
			"unsupported date_criteria value"},
		{"description",
			MsgCreate{Curator: a, Name: name, Exponent: exponentMax, CreditTypeAbbrev: creditAbbr, DateCriteria: start, Description: randstr.String(descrMaxLen + 1)},
			"description can't be longer"},
		{"allowed_classes-1",
			MsgCreate{Curator: a, Name: name, Exponent: exponentMax, CreditTypeAbbrev: creditAbbr, DateCriteria: start},
			"allowed_classes is required"},
		{"allowed_classes-2",
			MsgCreate{Curator: a, Name: name, Exponent: exponentMax, CreditTypeAbbrev: creditAbbr, DateCriteria: start, AllowedClasses: []string{"class1", ""}},
			"allowed_classes[1] must be defined"},
		{"fee-1",
			MsgCreate{Curator: a, Name: name, Exponent: exponentMax, CreditTypeAbbrev: creditAbbr, DateCriteria: start, AllowedClasses: classes, Fee: sdk.Coins{sdk.Coin{Denom: "1a"}}},
			"invalid denom"},
		{"fee-2", MsgCreate{Curator: a, Name: name, Exponent: exponentMax, CreditTypeAbbrev: creditAbbr, DateCriteria: start, AllowedClasses: classes, Fee: sdk.Coins{sdk.Coin{"aa", sdk.NewInt(-1)}}},
			"invalid denom"},

		{"good-1-fees-not-required",
			MsgCreate{Curator: a, Name: name, Exponent: 0, CreditTypeAbbrev: creditAbbr, DateCriteria: start, AllowedClasses: classes, Description: descr}, ""},
		{"good-date-criteria-not-required",
			MsgCreate{Curator: a, Name: name, Exponent: 6, CreditTypeAbbrev: creditAbbr, DateCriteria: nil, AllowedClasses: classes, Fee: sdk.Coins{sdk.NewInt64Coin("regen", 1)}}, ""},
	}

	for _, tc := range tcs {
		t.Run(tc.id, func(t *testing.T) {
			err := tc.msg.ValidateBasic()
			errorMatches(t, err, tc.err)
		})
	}
}

func TestMsgCreateValidateDateCriteria(t *testing.T) {
	tcs := []struct {
		id  string
		d   DateCriteria
		err string
	}{
		{"nil-min_start_date",
			DateCriteria{&DateCriteria_MinStartDate{nil}},
			"unsupported date_criteria value"},
		{"bad-min_start_date",
			DateCriteria{&DateCriteria_MinStartDate{
				&gogotypes.Timestamp{Seconds: time.Date(1400, 1, 1, 0, 0, 0, 0, time.UTC).Unix()}}},
			"date_criteria.min_start_date must be after"},
		{"nil-start_date_window",
			DateCriteria{&DateCriteria_StartDateWindow{}},
			"unsupported date_criteria value"},
		{"nil-start_date_window",
			DateCriteria{&DateCriteria_StartDateWindow{
				&gogotypes.Duration{Seconds: 3600}}},
			"date_criteria.start_date_window must be at least"},

		{"good-min_start_date",
			DateCriteria{&DateCriteria_MinStartDate{gogotypes.TimestampNow()}},
			""},
		{"good-start_date_window",
			DateCriteria{&DateCriteria_StartDateWindow{
				&gogotypes.Duration{Seconds: 3600 * 24 * 2}}},
			""},
	}
	for _, tc := range tcs {
		t.Run(tc.id, func(t *testing.T) {
			err := validateDateCriteria(&tc.d)
			errorMatches(t, err, tc.err)
		})
	}
}

func TestMsgCreateGetSigners(t *testing.T) {
	_, _, addr1 := testdata.KeyTestPubAddr()
	m := MsgCreate{Curator: addr1.String(), Name: "name", Exponent: 2}
	require.Equal(t, []sdk.AccAddress{addr1}, m.GetSigners())
}

func TestMsgCreateSignBytes(t *testing.T) {
	_, _, addr1 := testdata.KeyTestPubAddr()
	m := MsgCreate{Curator: addr1.String(), Name: "name", Exponent: 2}
	bz := m.GetSignBytes()
	require.NotEmpty(t, bz)
}
