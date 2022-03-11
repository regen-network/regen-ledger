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
	name := randstr.String((nameMaxLen+nameMinLen)/2, "ABCDEFGHIJKL")
	creditAbbr := "FOO"
	descr := "my project description"
	start := &DateCriteria{MinStartDate: gogotypes.TimestampNow()}

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
			"name must start with an alphabetic character"},
		{"name-long",
			MsgCreate{Curator: a, Name: randstr.String(nameMaxLen + 1)},
			"name must start with an alphabetic character"},
		{"name-short",
			MsgCreate{Curator: a, Name: randstr.String(nameMinLen - 1)},
			"name must start with an alphabetic character"},
		{"name-no-alphanum",
			MsgCreate{Curator: a, Name: randstr.String(nameMinLen) + "*"},
			"name must start with an alphabetic character"},
		{"name-no-alpah-prefix",
			MsgCreate{Curator: a, Name: "1" + randstr.String(nameMinLen)},
			"name must start with an alphabetic character"},
		{"exponent-1",
			MsgCreate{Curator: a, Name: name, Exponent: 4},
			"exponent must be one of [0 1 2 3 6 9 12 15 18 21 24]"},
		{"exponent-2",
			MsgCreate{Curator: a, Name: name, Exponent: 17},
			"exponent must be one of [0 1 2 3 6 9 12 15 18 21 24]"},
		{"credit_type-1",
			MsgCreate{Curator: a, Name: name, Exponent: 3},
			"credit type abbreviation must be 1-3"},
		{"credit_type-2",
			MsgCreate{Curator: a, Name: name, Exponent: 3, CreditTypeAbbrev: randstr.String(creditTypeAbbrMaxLen + 1)},
			"credit type abbreviation must be 1-3"},
		{"description",
			MsgCreate{Curator: a, Name: name, Exponent: 3, CreditTypeAbbrev: creditAbbr, DateCriteria: start, Description: randstr.String(descrMaxLen + 1)},
			"description can't be longer"},
		{"allowed_classes-1",
			MsgCreate{Curator: a, Name: name, Exponent: 3, CreditTypeAbbrev: creditAbbr, DateCriteria: start},
			"allowed_classes is required"},
		{"allowed_classes-2",
			MsgCreate{Curator: a, Name: name, Exponent: 3, CreditTypeAbbrev: creditAbbr, DateCriteria: start, AllowedClasses: []string{"class1", ""}},
			"allowed_classes[1] must be defined"},
		{"fee-1",
			MsgCreate{Curator: a, Name: name, Exponent: 3, CreditTypeAbbrev: creditAbbr, DateCriteria: start, AllowedClasses: classes, Fee: sdk.Coins{sdk.Coin{Denom: "1a"}}},
			"invalid denom"},
		{"fee-2", MsgCreate{Curator: a, Name: name, Exponent: 3, CreditTypeAbbrev: creditAbbr, DateCriteria: start, AllowedClasses: classes, Fee: sdk.Coins{sdk.Coin{"aa", sdk.NewInt(-1)}}},
			"invalid denom"},
		{"good-1-fees-not-required",
			MsgCreate{Curator: a, Name: name, Exponent: 0, CreditTypeAbbrev: creditAbbr, DateCriteria: start, AllowedClasses: classes, Description: descr}, ""},
		{"good-date-criteria-not-required",
			MsgCreate{Curator: a, Name: name, Exponent: 18, CreditTypeAbbrev: creditAbbr, DateCriteria: nil, AllowedClasses: classes, Fee: sdk.Coins{sdk.NewInt64Coin("regen", 1)}}, ""},
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
		{
			"bad-min_start_date",
			DateCriteria{MinStartDate: &gogotypes.Timestamp{
				Seconds: time.Date(1400, 1, 1, 0, 0, 0, 0, time.UTC).Unix()},
			},
			"date_criteria.min_start_date must be after",
		},
		{
			"bad-start_date_window",
			DateCriteria{StartDateWindow: &gogotypes.Duration{Seconds: 3600}},
			"date_criteria.start_date_window must be at least 1 day",
		},
		{
			"both-min_start_date-start_date_window-set",
			DateCriteria{
				MinStartDate:    gogotypes.TimestampNow(),
				StartDateWindow: &gogotypes.Duration{Seconds: 3600 * 24 * 2},
			},
			"only one of date_criteria.min_start_date, date_criteria.start_date_window, or date_criteria.years_in_the_past must be set",
		},
		{
			"both-min_start_date-years_in_the_past-set",
			DateCriteria{
				MinStartDate:   gogotypes.TimestampNow(),
				YearsInThePast: 10,
			},
			"only one of date_criteria.min_start_date, date_criteria.start_date_window, or date_criteria.years_in_the_past must be set",
		},
		{
			"both-start_date_window-years_in_the_past-set",
			DateCriteria{
				StartDateWindow: &gogotypes.Duration{Seconds: 3600 * 24 * 2},
				YearsInThePast:  10,
			},
			"only one of date_criteria.min_start_date, date_criteria.start_date_window, or date_criteria.years_in_the_past must be set",
		},
		{
			"good-min_start_date",
			DateCriteria{MinStartDate: gogotypes.TimestampNow()},
			"",
		},
		{
			"good-start_date_window",
			DateCriteria{StartDateWindow: &gogotypes.Duration{Seconds: 3600 * 24 * 2}},
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

func TestBasketDenom(t *testing.T) {
	tcs := []struct {
		tname        string
		abbrev       string
		exponent     uint32
		denom        string
		displayDenom string
		err          bool
	}{
		{"wrong exponent",
			"X", 5, "", "", true},
		{"exponent-0",
			"X", 0, "eco.X.foo", "eco.X.foo", false},
		{"exponent-1`",
			"X", 1, "eco.dX.foo", "eco.X.foo", false},
		{"exponent-2",
			"X", 2, "eco.cX.foo", "eco.X.foo", false},
		{"exponent-6",
			"X", 6, "eco.uX.foo", "eco.X.foo", false},
	}
	require := require.New(t)
	for _, tc := range tcs {
		d, displayD, err := BasketDenom("foo", tc.abbrev, tc.exponent)
		if tc.err {
			require.Error(err, tc.tname)
		} else {
			require.NoError(err, tc.tname)
			require.Equal(tc.denom, d, tc.tname)
			require.Equal(tc.displayDenom, displayD, tc.tname)
		}
	}
}
