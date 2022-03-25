package core

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

var (
	s = rand.NewSource(1)
	r = rand.New(s)
)

func TestMsgCreateClass(t *testing.T) {
	t.Parallel()

	_, _, addr1 := testdata.KeyTestPubAddr()
	_, _, addr2 := testdata.KeyTestPubAddr()
	tests := map[string]struct {
		src    MsgCreateClass
		expErr bool
	}{
		"valid msg": {
			src: MsgCreateClass{
				Admin:            addr1.String(),
				Issuers:          []string{addr1.String(), addr2.String()},
				CreditTypeAbbrev: "C",
				Metadata:         "hello",
			},
			expErr: false,
		},
		"valid msg without metadata": {
			src: MsgCreateClass{
				Admin:            addr1.String(),
				CreditTypeAbbrev: "C",
				Issuers:          []string{addr1.String(), addr2.String()},
			},
			expErr: false,
		},
		"invalid without admin": {
			src:    MsgCreateClass{},
			expErr: true,
		},
		"invalid without issuers": {
			src: MsgCreateClass{
				Admin:            addr1.String(),
				CreditTypeAbbrev: "C",
			},
			expErr: true,
		},
		"invalid with wrong issuers": {
			src: MsgCreateClass{
				Admin:            addr1.String(),
				CreditTypeAbbrev: "C",
				Issuers:          []string{"xyz", "xyz1"},
			},
			expErr: true,
		},
		"invalid with wrong admin": {
			src: MsgCreateClass{
				Admin:            "wrongAdmin",
				CreditTypeAbbrev: "C",
				Issuers:          []string{addr1.String(), addr2.String()},
			},
			expErr: true,
		},
		"invalid with no credit type": {
			src: MsgCreateClass{
				Admin:   addr1.String(),
				Issuers: []string{addr1.String(), addr2.String()},
			},
			expErr: true,
		},
		"invalid metadata maxlength is exceeded": {
			src: MsgCreateClass{
				Admin:            addr1.String(),
				CreditTypeAbbrev: "C",
				Issuers:          []string{addr1.String(), addr2.String()},
				Metadata:         simtypes.RandStringOfLength(r, 288),
			},
			expErr: true,
		},
	}

	for msg, test := range tests {
		t.Run(msg, func(t *testing.T) {
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
