package keeper

import (
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"

	api "github.com/regen-network/regen-ledger/api/v2/regen/ecocredit/basket/v1"
	"github.com/regen-network/regen-ledger/types/v2/testutil"
	types "github.com/regen-network/regen-ledger/x/ecocredit/v3/basket/types/v1"
)

type updateDateCriteriaSuite struct {
	*baseSuite
	authority   sdk.AccAddress
	basketDenom string
	res         *types.MsgUpdateDateCriteriaResponse
	err         error
}

func TestUpdateDateCriteria(t *testing.T) {
	gocuke.NewRunner(t, &updateDateCriteriaSuite{}).Path("./features/msg_update_date_criteria.feature").Run()
}

func (s *updateDateCriteriaSuite) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
	s.authority = s.addrs[0]
	s.basketDenom = "eco.uC.NCT"
}

func (s *updateDateCriteriaSuite) TheAuthorityAddress(a string) {
	addr, err := sdk.AccAddressFromBech32(a)
	require.NoError(s.t, err)
	s.authority = addr
}

func (s *updateDateCriteriaSuite) ABasketWithDenom(a string) {
	err := s.stateStore.BasketTable().Insert(s.ctx, &api.Basket{
		BasketDenom: a,
	})
	require.NoError(s.t, err)
}

func (s *updateDateCriteriaSuite) AliceAttemptsToUpdateDateCriteriaWithMessage(a gocuke.DocString) {
	var msg types.MsgUpdateDateCriteria
	err := jsonpb.UnmarshalString(a.Content, &msg)
	require.NoError(s.t, err)

	s.res, s.err = s.k.UpdateDateCriteria(s.ctx, &msg)
}

func (s *updateDateCriteriaSuite) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *updateDateCriteriaSuite) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

func (s *updateDateCriteriaSuite) ExpectNoDateCriteria() {
	dc, err := s.k.Basket(s.ctx, &types.QueryBasketRequest{
		BasketDenom: s.basketDenom,
	})
	require.NoError(s.t, err)

	require.Equal(s.t, (*types.DateCriteria)(nil), dc.BasketInfo.DateCriteria)
}

func (s *updateDateCriteriaSuite) ExpectDateCriteria(a gocuke.DocString) {
	var expected types.DateCriteria
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	dc, err := s.k.Basket(s.ctx, &types.QueryBasketRequest{
		BasketDenom: s.basketDenom,
	})
	require.NoError(s.t, err)

	require.Equal(s.t, &expected, dc.BasketInfo.DateCriteria)
}

func (s *updateDateCriteriaSuite) ExpectEventWithProperties(a gocuke.DocString) {
	var event types.EventUpdateDateCriteria
	err := jsonpb.UnmarshalString(a.Content, &event)
	require.NoError(s.t, err)

	sdkEvent, found := testutil.GetEvent(&event, s.sdkCtx.EventManager().Events())
	require.True(s.t, found)

	err = testutil.MatchEvent(&event, sdkEvent)
	require.NoError(s.t, err)
}
