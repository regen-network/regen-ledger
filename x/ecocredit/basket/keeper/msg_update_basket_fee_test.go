package keeper

import (
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"

	api "github.com/regen-network/regen-ledger/api/v2/regen/ecocredit/basket/v1"
	types "github.com/regen-network/regen-ledger/x/ecocredit/v3/basket/types/v1"
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
	var msg types.MsgUpdateBasketFee

	err := jsonpb.UnmarshalString(a.Content, &msg)
	require.NoError(s.t, err)

	_, s.err = s.k.UpdateBasketFee(s.ctx, &msg)
}

func (s *updateBasketFeeSuite) ExpectBasketFeeWithProperties(a gocuke.DocString) {
	var expected api.BasketFee
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	actual, err := s.stateStore.BasketFeeTable().Get(s.ctx)
	require.NoError(s.t, err)

	if expected.Fee.Denom != "" {
		require.Equal(s.t, expected.Fee.Amount, actual.Fee.Amount)
		require.Equal(s.t, expected.Fee.Denom, actual.Fee.Denom)
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
