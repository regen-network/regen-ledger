package keeper

import (
	"testing"

	"github.com/cosmos/gogoproto/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/regen-network/regen-ledger/types/v2/testutil"
	"github.com/regen-network/regen-ledger/x/ecocredit/v4"
	types "github.com/regen-network/regen-ledger/x/ecocredit/v4/base/types/v1"
)

type burnRegenSuite struct {
	*baseSuite
	msg *types.MsgBurnRegen
	err error
}

func TestBurnRegen(t *testing.T) {
	gocuke.NewRunner(t, &burnRegenSuite{}).Path("./features/msg_burn_regen.feature").Run()
}

func (s *burnRegenSuite) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
}

func (s *burnRegenSuite) TheMessage(a gocuke.DocString) {
	s.msg = &types.MsgBurnRegen{}
	err := jsonpb.UnmarshalString(a.Content, s.msg)
	require.NoError(s.t, err)
}
func (s *burnRegenSuite) ItIsExecuted() {
	_, s.err = s.k.BurnRegen(s.ctx, s.msg)
}

func (s *burnRegenSuite) ExpectAreSentFromToTheEcocreditModule(coinsStr, addrStr string) {
	coins, err := sdk.ParseCoinsNormalized(coinsStr)
	require.NoError(s.t, err)
	addr, err := sdk.AccAddressFromBech32(addrStr)
	require.NoError(s.t, err)
	s.bankKeeper.EXPECT().SendCoinsFromAccountToModule(s.ctx, addr, ecocredit.ModuleName, coins).Return(nil)
}

func (s *burnRegenSuite) ExpectAreBurnedByTheEcocreditModule(coinsStr string) {
	coins, err := sdk.ParseCoinsNormalized(coinsStr)
	require.NoError(s.t, err)
	s.bankKeeper.EXPECT().BurnCoins(s.ctx, ecocredit.ModuleName, coins).Return(nil)
}

func (s *burnRegenSuite) ExpectTheEventIsEmitted(a gocuke.DocString) {
	event := &types.EventBurnRegen{}
	err := jsonpb.UnmarshalString(a.Content, event)
	require.NoError(s.t, err)

	sdkEvent, found := testutil.GetEvent(event, s.sdkCtx.EventManager().Events())
	require.True(s.t, found)

	err = testutil.MatchEvent(event, sdkEvent)
	require.NoError(s.t, err)
}

func (s *burnRegenSuite) ExpectNoError() {
	require.NoError(s.t, s.err)
}
