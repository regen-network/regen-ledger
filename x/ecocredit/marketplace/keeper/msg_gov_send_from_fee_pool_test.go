package keeper

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/golang/mock/gomock"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"

	types "github.com/regen-network/regen-ledger/x/ecocredit/v3/marketplace/types/v1"
)

type govSendFromFeePool struct {
	*baseSuite
	err             error
	msg             *types.MsgGovSendFromFeePool
	moduleBalances  map[string]sdk.Coins
	accountBalances map[string]sdk.Coins
}

func TestGovSendFromFeePool(t *testing.T) {
	gocuke.NewRunner(t, &govSendFromFeePool{}).Path("./features/msg_gov_send_from_fee_pool.feature").Run()
}

func (s *govSendFromFeePool) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t, 2)
	s.msg = &types.MsgGovSendFromFeePool{}
	s.moduleBalances = make(map[string]sdk.Coins)
	s.accountBalances = make(map[string]sdk.Coins)
	s.bankKeeper.EXPECT().
		SendCoinsFromModuleToAccount(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
		AnyTimes().
		DoAndReturn(func(_ sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amount sdk.Coins) error {
			newModBalance, neg := s.moduleBalances[senderModule].SafeSub(amount...)
			if neg {
				return sdkerrors.ErrInsufficientFunds
			}

			s.moduleBalances[senderModule] = newModBalance
			s.accountBalances[recipientAddr.String()] = s.accountBalances[recipientAddr.String()].Add(amount...)
			return nil
		})
}

func (s *govSendFromFeePool) AuthorityIsSetToTheKeeperAuthority() {
	s.msg.Authority = s.k.authority.String()
}

func (s *govSendFromFeePool) AuthorityIsSetTo(a string) {
	s.msg.Authority = a
}

func (s *govSendFromFeePool) Recipient(a string) {
	s.msg.Recipient = a
}

func (s *govSendFromFeePool) SendAmount(a string) {
	coins, err := sdk.ParseCoinsNormalized(a)
	require.NoError(s.t, err)
	s.msg.Amount = coins
}

func (s *govSendFromFeePool) FeePoolBalance(a string) {
	coins, err := sdk.ParseCoinsNormalized(a)
	require.NoError(s.t, err)
	s.moduleBalances[s.k.feePoolName] = coins
}

func (s *govSendFromFeePool) RecipientBalance(a string) {
	coins, err := sdk.ParseCoinsNormalized(a)
	require.NoError(s.t, err)
	s.accountBalances[s.msg.Recipient] = coins
}

func (s *govSendFromFeePool) FundsAreSent() {
	_, s.err = s.k.GovSendFromFeePool(s.ctx, s.msg)
}

func (s *govSendFromFeePool) ExpectErrorContains(a string) {
	if s.err != nil {
		require.ErrorContains(s.t, s.err, a)
	} else {
		require.NoError(s.t, s.err)
	}
}

func (s *govSendFromFeePool) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *govSendFromFeePool) ExpectFeePoolBalance(a string) {
	coins, err := sdk.ParseCoinsNormalized(a)
	require.NoError(s.t, err)
	require.True(s.t, coins.IsEqual(s.moduleBalances[s.k.feePoolName]))
}

func (s *govSendFromFeePool) ExpectRecipientBalance(a string) {
	coins, err := sdk.ParseCoinsNormalized(a)
	require.NoError(s.t, err)
	require.True(s.t, coins.IsEqual(s.accountBalances[s.msg.Recipient]))
}
