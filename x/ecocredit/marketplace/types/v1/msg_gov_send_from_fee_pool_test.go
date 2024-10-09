package v1

import (
	"testing"

	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type govSendFromFeePool struct {
	gocuke.TestingT
	msg *MsgGovSendFromFeePool
	err error
}

func TestGovSendFromFeePool(t *testing.T) {
	gocuke.NewRunner(t, &govSendFromFeePool{}).
		Path("./features/msg_gov_send_from_fee_pool.feature").
		Run()
}

func (s *govSendFromFeePool) Before() {
	s.msg = &MsgGovSendFromFeePool{}
}

func (s *govSendFromFeePool) Authority(a string) {
	s.msg.Authority = a
}

func (s *govSendFromFeePool) Recipient(a string) {
	s.msg.Recipient = a
}

func (s *govSendFromFeePool) Amount(a string) {
	coins, err := sdk.ParseCoinsNormalized(a)
	require.NoError(s, err)
	s.msg.Coins = coins
}

func (s *govSendFromFeePool) TheMessageIsValidated() {
	s.err = s.msg.ValidateBasic()
}

func (s *govSendFromFeePool) ExpectErrorContains(a string) {
	if a != "" {
		require.ErrorContains(s, s.err, a)
	} else {
		require.NoError(s, s.err)
	}
}
