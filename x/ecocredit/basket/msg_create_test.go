package basket

import (
	"fmt"
	"testing"

	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/types"

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

	tcs := []struct {
		msg MsgCreate
		err string
	}{
		{MsgCreate{Curator: "wrong"}, "malformed curator address"},
		{MsgCreate{Curator: a, Name: ""}, "name must not be empty"},
		{MsgCreate{Curator: a, Name: randstr.String(101)}, "name must not"},
		{MsgCreate{Curator: a, Name: randstr.String(60), Exponent: 33}, "exponent must"},

		{MsgCreate{Curator: a, Name: randstr.String(60), Exponent: 0}, ""},
		{MsgCreate{Curator: a, Name: randstr.String(1), Exponent: 32}, ""},
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
