package basket

import (
	"encoding/json"
	"testing"

	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"

	v1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/basket/v1"
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

func (s *updateBasketFeeSuite) AliceAttemptsToUpdateBasketFeesWithProperties(a gocuke.DocString) {
	var msg *basket.MsgUpdateBasketFees

	err := json.Unmarshal([]byte(a.Content), &msg)
	require.NoError(s.t, err)

	_, s.err = s.k.UpdateBasketFees(s.ctx, msg)
}

func (s *updateBasketFeeSuite) ExpectBasketFeesWithProperties(a gocuke.DocString) {
	var expected *v1.BasketFees
	err := json.Unmarshal([]byte(a.Content), &expected)
	require.NoError(s.t, err)

	actual, err := s.stateStore.BasketFeesTable().Get(s.ctx)
	require.NoError(s.t, err)

	for i, fee := range expected.Fees {
		require.Equal(s.t, fee.Amount, actual.Fees[i].Amount)
		require.Equal(s.t, fee.Denom, actual.Fees[i].Denom)
	}
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
