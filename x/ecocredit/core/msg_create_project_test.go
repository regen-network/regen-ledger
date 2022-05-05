package core

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/regen-network/regen-ledger/types/testutil"
)

func TestMsgCreateProject(t *testing.T) {
	t.Parallel()
	issuer := testutil.GenAddress()

	testCases := []struct {
		name   string
		src    MsgCreateProject
		expErr bool
	}{
		{
			"valid msg without project id",
			MsgCreateProject{
				Issuer:       issuer,
				ClassId:      "A00",
				Metadata:     "hello",
				Jurisdiction: "AB-CDE FG1 345",
			},
			false,
		},
		{
			"invalid issuer",
			MsgCreateProject{
				Issuer:       "invalid address",
				ClassId:      "A00",
				Metadata:     "hello",
				Jurisdiction: "AB-CDE FG1 345",
			},
			true,
		},
		{
			"invalid class id",
			MsgCreateProject{
				Issuer:       issuer,
				ClassId:      "ABCD",
				Metadata:     "hello",
				Jurisdiction: "AB-CDE FG1 345",
			},
			true,
		},
		{
			"invalid project jurisdiction",
			MsgCreateProject{
				Issuer:       issuer,
				ClassId:      "A01",
				Metadata:     "hello",
				Jurisdiction: "abcd",
			},
			true,
		},
		{
			"invalid: metadata is too large",
			MsgCreateProject{
				Issuer:       issuer,
				ClassId:      "A01",
				Metadata:     strings.Repeat("x", 288),
				Jurisdiction: "AB-CDE FG1 345",
			},
			true,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			err := test.src.ValidateBasic()
			if test.expErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
