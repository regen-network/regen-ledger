package v1

import (
	"testing"

	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type msgUpdateProjectFee struct {
	gocuke.TestingT
	msg *MsgUpdateProjectFee
	err error
}

func TestMsgUpdateProjectFee(t *testing.T) {
	gocuke.NewRunner(t, &msgUpdateProjectFee{}).Path("./features/msg_update_project_fee.feature").Run()
}

func (s *msgUpdateProjectFee) Before() {
	s.msg = &MsgUpdateProjectFee{}
}

func (s *msgUpdateProjectFee) Authority(a string) {
	s.msg.Authority = a
}

func (s *msgUpdateProjectFee) Fee(a string) {
	coin, err := sdk.ParseCoinNormalized(a)
	require.NoError(s, err)
	s.msg.Fee = &coin
}

func (s *msgUpdateProjectFee) TheMessageIsValidated() {
	s.err = s.msg.ValidateBasic()
}

func (s *msgUpdateProjectFee) ExpectErrorContains(a string) {
	if a == "" {
		require.NoError(s, s.err)
	} else {
		require.ErrorContains(s, s.err, a)
	}
}

func (s *msgUpdateProjectFee) ExpectNoError() {
	require.NoError(s, s.err)
}

func (s *msgUpdateProjectFee) ExpectGetsignersReturns(a string) {
	require.Equal(s, a, s.msg.GetSigners()[0].String())
}

func (s *msgUpdateProjectFee) NilFee() {
	s.msg.Fee = nil
}
