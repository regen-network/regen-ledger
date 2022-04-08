package core

import (
	"math/rand"
	"testing"

	"github.com/regen-network/regen-ledger/types/testutil"
	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

var (
	s = rand.NewSource(1)
	r = rand.New(s)
)

func TestMsgCreateClass(t *testing.T) {
	t.Parallel()

	addr1 := testutil.GenAddress()
	addr2 := testutil.GenAddress()

	validFee := &sdk.Coin{
		Denom:  "regen",
		Amount: sdk.NewInt(10),
	}

	tests := map[string]struct {
		src    MsgCreateClass
		expErr bool
	}{
		"valid msg": {
			src: MsgCreateClass{
				Admin:            addr1,
				Issuers:          []string{addr1, addr2},
				CreditTypeAbbrev: "C",
				Metadata:         "hello",
				Fee:              validFee,
			},
			expErr: false,
		},
		"valid msg without metadata": {
			src: MsgCreateClass{
				Admin:            addr1,
				CreditTypeAbbrev: "C",
				Issuers:          []string{addr1, addr2},
				Fee:              validFee,
			},
			expErr: false,
		},
		"invalid without admin": {
			src:    MsgCreateClass{},
			expErr: true,
		},
		"invalid without issuers": {
			src: MsgCreateClass{
				Admin:            addr1,
				CreditTypeAbbrev: "C",
			},
			expErr: true,
		},
		"invalid with wrong issuers": {
			src: MsgCreateClass{
				Admin:            addr1,
				CreditTypeAbbrev: "C",
				Issuers:          []string{"xyz", "xyz1"},
			},
			expErr: true,
		},
		"invalid with wrong admin": {
			src: MsgCreateClass{
				Admin:            "wrongAdmin",
				CreditTypeAbbrev: "C",
				Issuers:          []string{addr1, addr2},
			},
			expErr: true,
		},
		"invalid with no credit type": {
			src: MsgCreateClass{
				Admin:   addr1,
				Issuers: []string{addr1, addr2},
			},
			expErr: true,
		},
		"invalid metadata maxlength is exceeded": {
			src: MsgCreateClass{
				Admin:            addr1,
				CreditTypeAbbrev: "C",
				Issuers:          []string{addr1, addr2},
				Metadata:         simtypes.RandStringOfLength(r, 288),
			},
			expErr: true,
		},
		"invalid bad fee denom": {
			src: MsgCreateClass{
				Admin:            addr1,
				CreditTypeAbbrev: "C",
				Issuers:          []string{addr1},
				Metadata:         "foo",
				Fee:              &sdk.Coin{Denom: "k,vm.zkx,cvzxk", Amount: sdk.NewInt(10)},
			},
			expErr: true,
		},
		"invalid bad fee amount": {
			src: MsgCreateClass{
				Admin:            addr1,
				CreditTypeAbbrev: "C",
				Issuers:          []string{addr1},
				Metadata:         "foo",
				Fee:              &sdk.Coin{Denom: "regen", Amount: sdk.NewInt(0)},
			},
			expErr: true,
		},
		"invalid duplicate issuer": {
			src: MsgCreateClass{
				Admin:            addr1.String(),
				Issuers:          []string{addr1.String(), addr2.String(), addr1.String()},
				CreditTypeAbbrev: "C",
				Metadata:         "hello",
				Fee:              validFee,
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
