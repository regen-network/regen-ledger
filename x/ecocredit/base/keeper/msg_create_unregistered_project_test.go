package keeper

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"

	basev1beta1 "github.com/cosmos/cosmos-sdk/api/cosmos/base/v1beta1"
	sdk "github.com/cosmos/cosmos-sdk/types"

	api "github.com/regen-network/regen-ledger/api/v2/regen/ecocredit/v1"
	v1 "github.com/regen-network/regen-ledger/x/ecocredit/v3/base/types/v1"
)

type createUnregisteredProjectSuite struct {
	*baseSuite
	err     error
	res     *v1.MsgCreateUnregisteredProjectResponse
	balance sdk.Coins
}

func TestCreateUnregisteredProject(t *testing.T) {
	gocuke.NewRunner(t, &createUnregisteredProjectSuite{}).
		Path("./features/msg_create_unregistered_project.feature").
		Run()
}

func (s *createUnregisteredProjectSuite) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
	s.bankKeeper.EXPECT().
		GetBalance(gomock.Any(), s.addr, "regen").
		DoAndReturn(func(_ sdk.Context, addr sdk.AccAddress, denom string) sdk.Coin {
			if addr.Equals(s.addr) {
				ok, denom := s.balance.Find(denom)
				if ok {
					return denom
				}
			}
			return sdk.Coin{}
		}).
		AnyTimes()

	s.bankKeeper.EXPECT().
		SendCoinsFromAccountToModule(gomock.Any(), s.addr, gomock.Any(), gomock.Any()).
		DoAndReturn(func(_ sdk.Context, fromAddr sdk.AccAddress, _ string, coins sdk.Coins) error {
			if fromAddr.Equals(s.addr) {
				s.balance = s.balance.Sub(coins...)
				return nil
			}
			return fmt.Errorf("unexpected from address: %s", fromAddr)
		}).
		AnyTimes()
}

func (s *createUnregisteredProjectSuite) ProjectCreationFee(fee string) {
	coin, err := sdk.ParseCoinNormalized(fee)
	require.NoError(s.t, err)
	require.NoError(s.t, s.stateStore.ProjectFeeTable().Save(s.ctx, &api.ProjectFee{
		Fee: &basev1beta1.Coin{
			Denom:  coin.Denom,
			Amount: coin.Amount.String(),
		},
	}))
}

func (s *createUnregisteredProjectSuite) IHaveBalance(balance string) {
	coins, err := sdk.ParseCoinsNormalized(balance)
	require.NoError(s.t, err)
	s.balance = coins
}

func (s *createUnregisteredProjectSuite) ICreateAProjectWithJurisdictionAndFee(jurisdiction, fee string) {
	coin, err := sdk.ParseCoinNormalized(fee)
	require.NoError(s.t, err)
	s.res, s.err = s.k.CreateUnregisteredProject(s.ctx, &v1.MsgCreateUnregisteredProject{
		Admin:        s.addr.String(),
		Jurisdiction: jurisdiction,
		Fee:          &coin,
	})
}

func (s *createUnregisteredProjectSuite) ExpectTheProjectIsCreatedSuccessfully() {
	require.NoError(s.t, s.err)
	require.NotNil(s.t, s.res)
	require.NotEmpty(s.t, s.res.ProjectId)
}

func (s *createUnregisteredProjectSuite) ExpectBalance(a string) {
	coins, err := sdk.ParseCoinsNormalized(a)
	require.NoError(s.t, err)
	require.True(s.t, s.balance.IsEqual(coins), "expected balance %s, got %s", coins, s.balance)
}

func (s *createUnregisteredProjectSuite) ExpectErrorContains(x string) {
	if x == "" {
		require.NoError(s.t, s.err)
	} else {
		require.ErrorContains(s.t, s.err, x)
	}
}
