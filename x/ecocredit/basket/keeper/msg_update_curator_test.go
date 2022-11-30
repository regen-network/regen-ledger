package keeper

import (
	"encoding/json"
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"

	api "github.com/regen-network/regen-ledger/api/v2/regen/ecocredit/basket/v1"
	"github.com/regen-network/regen-ledger/types/v2/testutil"
	types "github.com/regen-network/regen-ledger/x/ecocredit/v3/basket/types/v1"
)

type updateBasketCurator struct {
	*baseSuite
	alice sdk.AccAddress
	bob   sdk.AccAddress
	res   *types.MsgUpdateCuratorResponse
	err   error
}

func TestUpdateBasketCurator(t *testing.T) {
	gocuke.NewRunner(t, &updateBasketCurator{}).Path("./features/msg_update_curator.feature").Run()
}

func (s *updateBasketCurator) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
	s.alice = s.addrs[0]
	s.bob = s.addrs[1]
}

func (s *updateBasketCurator) ABasketWithPropertiesAndCuratorAlice(a gocuke.DocString) {
	var basket types.Basket
	err := jsonpb.UnmarshalString(a.Content, &basket)
	require.NoError(s.t, err)

	err = s.k.stateStore.BasketTable().Save(s.ctx, &api.Basket{
		BasketDenom:      basket.BasketDenom,
		Name:             basket.Name,
		CreditTypeAbbrev: basket.CreditTypeAbbrev,
		Curator:          s.alice,
	})
	require.NoError(s.t, err)
}

func (s *updateBasketCurator) BobAttemptsToUpdateBasketCuratorWithBasketDenom(a string) {
	s.res, s.err = s.k.UpdateCurator(s.ctx, &types.MsgUpdateCurator{
		Denom:      a,
		Curator:    s.bob.String(),
		NewCurator: s.alice.String(),
	})
}

func (s *updateBasketCurator) ExpectBasketWithDenomAndCuratorBob(a string) {
	basket, err := s.k.Basket(s.ctx, &types.QueryBasketRequest{
		BasketDenom: a,
	})
	require.NoError(s.t, err)
	require.Equal(s.t, basket.BasketInfo.BasketDenom, a)
	require.Equal(s.t, basket.BasketInfo.Curator, s.bob.String())
}

func (s *updateBasketCurator) AliceAttemptsToUpdateBasketCuratorWithDenom(a string) {
	s.res, s.err = s.k.UpdateCurator(s.ctx, &types.MsgUpdateCurator{
		Denom:      a,
		Curator:    s.alice.String(),
		NewCurator: s.bob.String(),
	})
}

func (s *updateBasketCurator) AliceAttemptsToUpdateBasketCuratorWithDenomAndNewCuratorBob(a string) {
	s.res, s.err = s.k.UpdateCurator(s.ctx, &types.MsgUpdateCurator{
		Curator:    s.alice.String(),
		Denom:      a,
		NewCurator: s.bob.String(),
	})
}

func (s *updateBasketCurator) BobAttemptsToUpdateBasketCuratorWithDenom(a string) {
	s.res, s.err = s.k.UpdateCurator(s.ctx, &types.MsgUpdateCurator{
		Curator:    s.bob.String(),
		Denom:      a,
		NewCurator: s.alice.String(),
	})
}

func (s *updateBasketCurator) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *updateBasketCurator) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

func (s *updateBasketCurator) ExpectErrorContains(a string) {
	require.ErrorContains(s.t, s.err, a)
}

func (s *updateBasketCurator) ExpectEventWithProperties(a gocuke.DocString) {
	var event types.EventUpdateCurator
	err := json.Unmarshal([]byte(a.Content), &event)
	require.NoError(s.t, err)

	sdkEvent, found := testutil.GetEvent(&event, s.sdkCtx.EventManager().Events())
	require.True(s.t, found)

	err = testutil.MatchEvent(&event, sdkEvent)
	require.NoError(s.t, err)
}
