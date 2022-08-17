package basket_test

import (
	"encoding/json"
	"testing"

	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"

	"github.com/regen-network/regen-ledger/x/ecocredit/basket"
)

type updateBasketFeeSuite struct {
	*baseSuite
	err error
}

func TestUpdateBasketFee(t *testing.T) {
	gocuke.NewRunner(t, &updateBasketFeeSuite{}).Path("./features/msg_update_basket_fee.feature").Run()
}

func (s *updateBasketFeeSuite) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
}

func (s *updateBasketFeeSuite) AliceAttemptsToUpdateBasketFeeWithProperties(a gocuke.DocString) {
	var msg *basket.MsgUpdateBasketFee

	err := json.Unmarshal([]byte(a.Content), &msg)
	require.NoError(s.t, err)

	_, s.err = s.k.UpdateBasketFee(s.ctx, msg)
}

func (s *updateBasketFeeSuite) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *updateBasketFeeSuite) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

func (s *updateBasketFeeSuite) ExpectErrorContains(a string) {
	require.ErrorContains(s.t, s.err, a)
}
