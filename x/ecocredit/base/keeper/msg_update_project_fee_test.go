package keeper

import (
	"testing"

	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"

	basev1beta1 "github.com/cosmos/cosmos-sdk/api/cosmos/base/v1beta1"
	sdk "github.com/cosmos/cosmos-sdk/types"

	api "github.com/regen-network/regen-ledger/api/v2/regen/ecocredit/v1"
	v1 "github.com/regen-network/regen-ledger/x/ecocredit/v3/base/types/v1"
)

type updateProjectFeeSuite struct {
	*baseSuite
	err   error
	addrs map[string]sdk.AccAddress
}

func TestUpdateProjectFee(t *testing.T) {
	gocuke.NewRunner(t, &updateProjectFeeSuite{}).
		Path("./features/msg_update_project_fee.feature").
		Run()
}

func (s *updateProjectFeeSuite) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
	s.addrs = map[string]sdk.AccAddress{
		"gov": s.k.authority,
		"bob": s.addr,
	}
}

func (s *updateProjectFeeSuite) CurrentFee(amount string) {
	if amount != "" {
		require.NoError(s.t, s.stateStore.ProjectFeeTable().Save(s.ctx, &api.ProjectFee{
			Fee: &basev1beta1.Coin{
				Denom:  "regen",
				Amount: amount,
			},
		}))
	}
}

func (s *updateProjectFeeSuite) UpdatesTheFeeTo(auth, fee string) {
	if fee == "" {
		_, s.err = s.k.UpdateProjectFee(s.ctx, &v1.MsgUpdateProjectFee{
			Authority: s.addrs[auth].String(),
		})
	} else {
		amount, ok := sdk.NewIntFromString(fee)
		require.True(s.t, ok, "invalid fee amount")
		_, s.err = s.k.UpdateProjectFee(s.ctx, &v1.MsgUpdateProjectFee{
			Authority: s.addrs[auth].String(),
			Fee:       &sdk.Coin{Denom: "regen", Amount: amount},
		})
	}
}

func (s *updateProjectFeeSuite) ExpectErrorContains(a string) {
	if a == "" {
		require.NoError(s.t, s.err)
	} else {
		require.ErrorContains(s.t, s.err, a)
	}
}

func (s *updateProjectFeeSuite) ExpectProjectFeeIs(a string) {
	actual, err := s.stateStore.ProjectFeeTable().Get(s.ctx)
	require.NoError(s.t, err)
	require.NotNil(s.t, actual)
	if a == "" {
		if actual.Fee == nil || actual.Fee.Amount == "" || actual.Fee.Amount == "0" {
			return
		}
		s.t.Fatalf("expected no fee, got %s", actual.Fee.Amount)
	} else {
		require.Equal(s.t, a, actual.Fee.Amount)
	}
}
