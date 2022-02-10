package basket

import (
	"fmt"
	"testing"

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
	name := randstr.String(nameMaxLen)
	dName := randstr.String((displayNameMaxLen + displayNameMinLen) / 2)
	creditName := randstr.String(10)
	start := gogotypes.TimestampNow()
	classes := []string{"eco_class1"}

	tcs := []struct {
		msg MsgCreate
		err string
	}{
		{MsgCreate{Curator: "wrong"}, "malformed curator address"},
		{MsgCreate{Curator: a, Name: ""}, "name must not be empty"},
		{MsgCreate{Curator: a, Name: randstr.String(nameMaxLen + 1)}, "name must not be empty and must not be longer than"},
		{MsgCreate{Curator: a, Name: name, DisplayName: ""}, "display_name must be between"},
		{MsgCreate{Curator: a, Name: name, DisplayName: randstr.String(displayNameMaxLen + 1)}, "display_name must be between"},
		{MsgCreate{Curator: a, Name: name, DisplayName: dName, Exponent: exponentMax + 1}, "exponent must not be bigger than"},
		{MsgCreate{Curator: a, Name: name, DisplayName: dName, Exponent: exponentMax},
			"credit_type_name must be defined"},
		{MsgCreate{Curator: a, Name: name, DisplayName: dName, Exponent: exponentMax, CreditTypeName: randstr.String(creditTNameMaxLen + 1)},
			"credit_type_name must not be longer"},
		{MsgCreate{Curator: a, Name: name, DisplayName: dName, Exponent: exponentMax, CreditTypeName: creditName, MinStartDate: nil},
			"min_start_date must be defined"},
		{MsgCreate{Curator: a, Name: name, DisplayName: dName, Exponent: exponentMax, CreditTypeName: creditName, MinStartDate: start},
			"allowed_classes is required"},
		{MsgCreate{Curator: a, Name: name, DisplayName: dName, Exponent: exponentMax, CreditTypeName: creditName, MinStartDate: start, AllowedClasses: []string{"class1", ""}},
			"allowed_classes[1] must be defined"},

		{MsgCreate{Curator: a, Name: name, DisplayName: dName, Exponent: 0, CreditTypeName: creditName, MinStartDate: start, AllowedClasses: classes}, ""},
		// {MsgCreate{Curator: a, Name: randstr.String(1), Exponent: 32}, ""},
	}

	for i, tc := range tcs {
		t.Run(fmt.Sprint("test-", i), func(t *testing.T) {
			err := tc.msg.ValidateBasic()
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
