package v1

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
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
	var ptrCoins []*sdk.Coin
	for _, coin := range coins {
		ptrCoins = append(ptrCoins, &coin)
	}
	s.msg.Amount = ptrCoins
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
