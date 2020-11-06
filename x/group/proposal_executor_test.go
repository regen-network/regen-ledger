package group_test

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	testdatagroup "github.com/regen-network/regen-ledger/testutil/testdata/group"
	"github.com/regen-network/regen-ledger/x/group"
	testutil "github.com/regen-network/regen-ledger/x/group/testutil"
	"github.com/stretchr/testify/require"
)

func TestDoExecuteMsgs(t *testing.T) {
	specs := map[string]struct {
		srcAccount sdk.AccAddress
		srcMsgs    []sdk.Msg
		srcHandler sdk.Handler
		expErr     bool
	}{
		"all good": {
			srcAccount: []byte("my-group-acct-addrss"),
			srcMsgs:    []sdk.Msg{&testdatagroup.MsgAuthenticate{Signers: []sdk.AccAddress{[]byte("my-group-acct-addrss")}}},
			srcHandler: mockHandler(&sdk.Result{}, nil),
		},
		"not authz by group account": {
			srcAccount: []byte("my-group-acct-addrss"),
			srcMsgs:    []sdk.Msg{&testdatagroup.MsgAuthenticate{Signers: []sdk.AccAddress{[]byte("any--other---address")}}},
			srcHandler: alwaysPanicHandler(),
			expErr:     true,
		},
		"mixed group account msgs": {
			srcAccount: []byte("my-group-acct-addrss"),
			srcMsgs: []sdk.Msg{
				&testdatagroup.MsgAuthenticate{Signers: []sdk.AccAddress{[]byte("my-group-acct-addrss")}},
				&testdatagroup.MsgAuthenticate{Signers: []sdk.AccAddress{[]byte("any--other---address")}},
			},
			srcHandler: alwaysPanicHandler(),
			expErr:     true,
		},
		"no handler": {
			srcAccount: []byte("my-group-acct-addrss"),
			srcMsgs:    []sdk.Msg{NonRoutableMsg{}},
			srcHandler: alwaysPanicHandler(),
			expErr:     true,
		},
		"not panic on nil result": {
			srcAccount: []byte("my-group-acct-addrss"),
			srcMsgs:    []sdk.Msg{&testdatagroup.MsgAuthenticate{Signers: []sdk.AccAddress{[]byte("my-group-acct-addrss")}}},
			srcHandler: mockHandler(nil, nil),
		},
	}
	for msg, spec := range specs {
		t.Run(msg, func(t *testing.T) {
			router := baseapp.NewRouter().AddRoute(sdk.NewRoute(testdatagroup.ModuleName, spec.srcHandler))
			_, err := group.DoExecuteMsgs(testutil.NewContext(), router, spec.srcAccount, spec.srcMsgs)
			if spec.expErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func mockHandler(r *sdk.Result, t error) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (result *sdk.Result, err error) {
		return r, t
	}
}

func alwaysPanicHandler() sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (result *sdk.Result, err error) {
		panic("not supposed to be called")
	}
}

type NonRoutableMsg struct {
	sdk.Msg
}

func (m NonRoutableMsg) Route() string {
	return "not_routable"
}
func (m NonRoutableMsg) GetSigners() []sdk.AccAddress {
	return nil
}
